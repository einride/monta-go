package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.einride.tech/monta"
)

func main() {
	cmd := newMontaCommand()
	cmd.AddCommand(newLoginCommand())
	cmd.AddCommand(newMeCommand())
	cmd.AddCommand(newListSitesCommand())
	cmd.AddCommand(newListChargePointsCommand())
	cmd.AddCommand(newListChargesCommand())
	cmd.AddCommand(newListWalletTransactionsCommand())
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const (
	authConfigFile = "monta/auth.json"
	allPagesValue  = -1
)

const (
	apiCommandAnnotation   = "monta_annotation_api_command"
	authCommandAnnotation  = "monta_annotation_auth_command"
	argumentFlagAnnotation = "monta_annotation_argument_flag"
	authFlagAnnotation     = "monta_annotation_auth_flag"
)

func newMontaCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monta",
		Short: "Monta Partner API CLI",
		Long: strings.TrimSpace(
			"" +
				"Monta Partner API CLI\n" +
				"\n" +
				"The Monta Partner API enables you to develop your own use cases around Monta-managed charge points.\n" +
				"\n" +
				"https://partner-api.monta.app/docs",
		),
	}
	cmd.SetHelpFunc(helpFunc)
	cmd.SetUsageFunc(usageFunc)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.PersistentFlags().String("client-id", "", "client ID to authenticate with")
	cmd.PersistentFlags().Lookup("client-id").Annotations = map[string][]string{
		authFlagAnnotation: {},
	}
	cmd.PersistentFlags().String("client-secret", "", "client secret to authenticate with")
	cmd.PersistentFlags().Lookup("client-secret").Annotations = map[string][]string{
		authFlagAnnotation: {},
	}
	return cmd
}

func newClientWithAuthentication(cmd *cobra.Command) (*monta.Client, error) {
	var options []monta.ClientOption
	if cmd.Flags().Changed("client-id") && cmd.Flags().Changed("client-secret") {
		clientID, err := cmd.Flags().GetString("client-id")
		if err != nil {
			return nil, err
		}
		clientSecret, err := cmd.Flags().GetString("client-secret")
		if err != nil {
			return nil, err
		}
		options = append(options, monta.WithClientIDAndSecret(clientID, clientSecret))
	}
	authFilepath, err := xdg.ConfigFile(authConfigFile)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(authFilepath); err == nil {
		data, err := os.ReadFile(authFilepath)
		if err != nil {
			return nil, err
		}
		var token monta.Token
		if err := json.Unmarshal(data, &token); err != nil {
			return nil, err
		}
		now := time.Now()
		if token.AccessTokenExpirationTime.Before(now) {
			if token.RefreshTokenExpirationTime.Before(now) {
				return nil, fmt.Errorf("cached tokens have expired, please re-authenticate with `monta login`")
			}
			refreshedToken, err := monta.NewClient().RefreshToken(cmd.Context(), &monta.RefreshTokenRequest{
				RefreshToken: token.RefreshToken,
			})
			if err != nil {
				return nil, err
			}
			refreshedTokenData, err := json.MarshalIndent(refreshedToken, "", "  ")
			if err != nil {
				return nil, err
			}
			if err := os.WriteFile(authFilepath, refreshedTokenData, 0o600); err != nil {
				return nil, err
			}
			token = *refreshedToken
		}
		options = append(options, monta.WithToken(&token))
	}
	return monta.NewClient(options...), nil
}

func newMeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "me",
		Short: "Get information about the current authenticated user",
		Annotations: map[string]string{
			apiCommandAnnotation: "GET /v1/auth/me",
		},
	}
	cmd.SetHelpFunc(helpFunc)
	cmd.SetUsageFunc(usageFunc)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		client, err := newClientWithAuthentication(cmd)
		if err != nil {
			return err
		}
		me, err := client.GetMe(cmd.Context())
		if err != nil {
			return err
		}
		data, err := json.MarshalIndent(me, "", "  ")
		if err != nil {
			return err
		}
		cmd.Println(string(data))
		return nil
	}
	return cmd
}

func newListSitesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sites",
		Short: "List your charge sites",
		Annotations: map[string]string{
			apiCommandAnnotation: "GET /v1/sites",
		},
	}
	cmd.SetHelpFunc(helpFunc)
	cmd.SetUsageFunc(usageFunc)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	page := cmd.Flags().Int("page", allPagesValue, "page number to retrieve")
	cmd.Flag("page").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	perPage := cmd.Flags().Int("per-page", 10, "number of items per page")
	cmd.Flag("per-page").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		client, err := newClientWithAuthentication(cmd)
		if err != nil {
			return err
		}
		var allPages bool
		if *page == allPagesValue {
			allPages = true
			*page = 0
		}
		for {
			response, err := client.ListSites(cmd.Context(), &monta.ListSitesRequest{
				PageFilters: monta.PageFilters{
					Page:    *page,
					PerPage: *perPage,
				},
			})
			if err != nil {
				return err
			}
			if !allPages {
				data, err := json.MarshalIndent(response, "", "  ")
				if err != nil {
					return err
				}
				cmd.Println(string(data))
				break
			}
			for _, site := range response.Sites {
				data, err := json.MarshalIndent(site, "", "  ")
				if err != nil {
					return err
				}
				cmd.Println(string(data))
			}
			if *page >= int(response.PageMeta.TotalPageCount) {
				break
			}
			*page = int(response.PageMeta.CurrentPage + 1)
		}
		return nil
	}
	return cmd
}

func newListChargePointsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "charge-points",
		Short: "List your charge points",
		Annotations: map[string]string{
			apiCommandAnnotation: "GET /v1/charge-points",
		},
	}
	cmd.SetHelpFunc(helpFunc)
	cmd.SetUsageFunc(usageFunc)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	page := cmd.Flags().Int("page", allPagesValue, "page number to retrieve")
	cmd.Flag("page").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	perPage := cmd.Flags().Int("per-page", 10, "number of items per page")
	cmd.Flag("per-page").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	siteID := cmd.Flags().Int64("site-id", 0, "site ID to filter by")
	cmd.Flag("site-id").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		client, err := newClientWithAuthentication(cmd)
		if err != nil {
			return err
		}
		var allPages bool
		if *page == allPagesValue {
			allPages = true
			*page = 0
		}
		for {
			request := &monta.ListChargePointsRequest{
				PageFilters: monta.PageFilters{
					Page:    *page,
					PerPage: *perPage,
				},
			}
			if cmd.Flags().Changed("site-id") {
				request.SiteID = siteID
			}
			response, err := client.ListChargePoints(cmd.Context(), request)
			if err != nil {
				return err
			}
			if !allPages {
				data, err := json.MarshalIndent(response, "", "  ")
				if err != nil {
					return err
				}
				cmd.Println(string(data))
				break
			}
			for _, chargePoint := range response.ChargePoints {
				data, err := json.MarshalIndent(chargePoint, "", "  ")
				if err != nil {
					return err
				}
				cmd.Println(string(data))
			}
			if *page >= int(response.PageMeta.TotalPageCount) {
				break
			}
			*page = int(response.PageMeta.CurrentPage + 1)
		}
		return nil
	}
	return cmd
}

func newListChargesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "charges",
		Short: "List your charging transactions",
		Annotations: map[string]string{
			apiCommandAnnotation: "GET /v1/charges",
		},
	}
	cmd.SetHelpFunc(helpFunc)
	cmd.SetUsageFunc(usageFunc)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	page := cmd.Flags().Int("page", allPagesValue, "page number to retrieve")
	cmd.Flag("page").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	perPage := cmd.Flags().Int("per-page", 10, "number of items per page")
	cmd.Flag("per-page").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	teamID := cmd.Flags().Int64("team-id", 0, "team ID to filter by")
	cmd.Flag("team-id").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	chargePointID := cmd.Flags().Int64("charge-point-id", 0, "charge point ID to filter by")
	cmd.Flag("charge-point-id").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		client, err := newClientWithAuthentication(cmd)
		if err != nil {
			return err
		}
		var allPages bool
		if *page == allPagesValue {
			allPages = true
			*page = 0
		}
		for {
			request := &monta.ListChargesRequest{
				PageFilters: monta.PageFilters{
					Page:    *page,
					PerPage: *perPage,
				},
			}
			if cmd.Flags().Changed("team-id") {
				request.TeamID = teamID
			}
			if cmd.Flags().Changed("charge-point-id") {
				request.ChargePointID = chargePointID
			}
			response, err := client.ListCharges(cmd.Context(), request)
			if err != nil {
				return err
			}
			if !allPages {
				data, err := json.MarshalIndent(response, "", "  ")
				if err != nil {
					return err
				}
				cmd.Println(string(data))
				break
			}
			for _, charge := range response.Charges {
				data, err := json.MarshalIndent(charge, "", "  ")
				if err != nil {
					return err
				}
				cmd.Println(string(data))
			}
			if *page >= int(response.PageMeta.TotalPageCount) {
				break
			}
			*page = int(response.PageMeta.CurrentPage + 1)
		}
		return nil
	}
	return cmd
}

func newListWalletTransactionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wallet-transactions",
		Short: "List your wallet transactions",
		Annotations: map[string]string{
			apiCommandAnnotation: "GET /v1/wallet-transactiosn",
		},
	}
	cmd.SetHelpFunc(helpFunc)
	cmd.SetUsageFunc(usageFunc)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	page := cmd.Flags().Int("page", allPagesValue, "page number to retrieve")
	cmd.Flag("page").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	perPage := cmd.Flags().Int("per-page", 10, "number of items per page")
	cmd.Flag("per-page").Annotations = map[string][]string{
		argumentFlagAnnotation: {},
	}
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		client, err := newClientWithAuthentication(cmd)
		if err != nil {
			return err
		}
		var allPages bool
		if *page == allPagesValue {
			allPages = true
			*page = 0
		}
		for {
			request := &monta.ListWalletTransactionsRequest{
				PageFilters: monta.PageFilters{
					Page:    *page,
					PerPage: *perPage,
				},
			}
			response, err := client.ListWalletTransactions(cmd.Context(), request)
			if err != nil {
				return err
			}
			if !allPages {
				data, err := json.MarshalIndent(response, "", "  ")
				if err != nil {
					return err
				}
				cmd.Println(string(data))
				break
			}
			for _, walletTransaction := range response.WalletTransactions {
				data, err := json.MarshalIndent(walletTransaction, "", "  ")
				if err != nil {
					return err
				}
				cmd.Println(string(data))
			}
			if *page >= int(response.PageMeta.TotalPageCount) {
				break
			}
			*page = int(response.PageMeta.CurrentPage + 1)
		}
		return nil
	}
	return cmd
}

func newLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with the Monta Partner API",
		Annotations: map[string]string{
			authCommandAnnotation: "POST /v1/auth/token",
		},
	}
	cmd.SetHelpFunc(helpFunc)
	cmd.SetUsageFunc(usageFunc)
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		authFilepath, err := xdg.ConfigFile(authConfigFile)
		if err != nil {
			return err
		}
		clientID, err := cmd.Flags().GetString("client-id")
		if err != nil {
			return err
		}
		clientSecret, err := cmd.Flags().GetString("client-secret")
		if err != nil {
			return err
		}
		client := monta.NewClient()
		token, err := client.CreateToken(cmd.Context(), &monta.CreateTokenRequest{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		})
		if err != nil {
			return err
		}
		tokenData, err := json.MarshalIndent(&token, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(authFilepath, tokenData, 0o600); err != nil {
			return err
		}
		cmd.Println("Successfully authenticated to Monta Partner API.")
		cmd.Println()
		cmd.Printf(
			"  (access token expires at %s in %s)\n",
			token.AccessTokenExpirationTime.Format(time.RFC3339),
			time.Until(token.AccessTokenExpirationTime).Round(time.Second),
		)
		cmd.Printf(
			"  (refresh token expires at %s in %s)\n",
			token.RefreshTokenExpirationTime.Format(time.RFC3339),
			time.Until(token.RefreshTokenExpirationTime).Round(time.Second),
		)
		return nil
	}
	return cmd
}

func helpFunc(cmd *cobra.Command, _ []string) {
	_ = usageFunc(cmd)
}

func usageFunc(cmd *cobra.Command) error {
	out := cmd.ErrOrStderr()
	defer cmd.SetErr(out)
	tw := tabwriter.NewWriter(out, 2, 0, 2, ' ', 0)
	cmd.SetErr(tw)
	if cmd.Short != "" {
		cmd.PrintErrln()
		if cmd.Long != "" {
			cmd.PrintErrln(cmd.Long)
		} else {
			cmd.PrintErrln(cmd.Short)
		}
	}
	cmd.PrintErrln()
	cmd.PrintErrln("USAGE")
	cmd.PrintErrln(" ", cmd.Use, "<command>")
	if commands := getCommands(cmd, commandHasAnnotation(apiCommandAnnotation)); len(commands) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("API COMMANDS")
		printCommands(cmd, commands)
	}
	if commands := getCommands(cmd, commandHasAnnotation(authCommandAnnotation)); len(commands) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("AUTH COMMANDS")
		printCommands(cmd, commands)
	}
	if otherCommands := getCommands(cmd, func(cmd *cobra.Command) bool {
		_, isAPICommand := cmd.Annotations[apiCommandAnnotation]
		_, isAuthCommand := cmd.Annotations[authCommandAnnotation]
		return !isAPICommand && !isAuthCommand
	}); len(otherCommands) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("OTHER COMMANDS")
		printCommands(cmd, otherCommands)
	}
	if flags := getFlags(cmd, isArgumentFlag); len(flags) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("ARGUMENT FLAGS")
		printFlags(cmd, flags)
	}
	if flags := getFlags(cmd, isAuthFlag); len(flags) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("AUTH FLAGS")
		printFlags(cmd, flags)
	}
	if flags := getFlags(cmd, func(command *cobra.Command, flag *pflag.Flag) bool {
		return !isArgumentFlag(cmd, flag) && !isAuthFlag(cmd, flag)
	}); len(flags) > 0 {
		cmd.PrintErrln()
		cmd.PrintErrln("OTHER FLAGS")
		printFlags(cmd, flags)
	}
	return tw.Flush()
}

func printFlags(cmd *cobra.Command, flags []*pflag.Flag) {
	var hasShorthand bool
	for _, flag := range flags {
		if flag.Shorthand != "" {
			hasShorthand = true
			break
		}
	}
	for _, flag := range flags {
		if flag.Hidden {
			continue
		}
		var line strings.Builder
		_, _ = line.WriteString("  ")
		if hasShorthand {
			if flag.Shorthand == "" {
				_, _ = line.WriteString("  ")
			} else {
				_, _ = line.WriteString("-" + flag.Shorthand)
			}
			_, _ = line.WriteString("  ")
		}
		_, _ = line.WriteString("--" + flag.Name)
		_, _ = line.WriteString("\t")
		_, _ = line.WriteString(flag.Value.Type())
		_, _ = line.WriteString("\t")
		_, _ = line.WriteString(flag.Usage)
		if flag.DefValue != "" && flag.DefValue != "false" {
			_, _ = line.WriteString(" (" + flag.DefValue + ")")
		}
		cmd.PrintErrln(line.String())
	}
}

func commandHasAnnotation(annotation string) func(command *cobra.Command) bool {
	return func(command *cobra.Command) bool {
		_, ok := command.Annotations[annotation]
		return ok
	}
}

func getCommands(cmd *cobra.Command, fn func(*cobra.Command) bool) []*cobra.Command {
	cmds := cmd.Commands()
	result := make([]*cobra.Command, 0, len(cmds))
	for _, subCmd := range cmds {
		if fn(subCmd) {
			result = append(result, subCmd)
		}
	}
	return result
}

func getFlags(cmd *cobra.Command, fn func(*cobra.Command, *pflag.Flag) bool) []*pflag.Flag {
	flags := cmd.Flags()
	flags.SortFlags = false
	result := make([]*pflag.Flag, 0, flags.NFlag())
	flags.VisitAll(func(flag *pflag.Flag) {
		if fn(cmd, flag) {
			result = append(result, flag)
		}
	})
	return result
}

func isArgumentFlag(_ *cobra.Command, flag *pflag.Flag) bool {
	_, ok := flag.Annotations[argumentFlagAnnotation]
	return ok
}

func isAuthFlag(_ *cobra.Command, flag *pflag.Flag) bool {
	_, ok := flag.Annotations[authFlagAnnotation]
	return ok
}

func printCommands(cmd *cobra.Command, commands []*cobra.Command) {
	for _, command := range commands {
		cmd.PrintErrln("  " + command.Name() + "\t" + command.Short)
	}
}
