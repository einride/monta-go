package main

import (
	"context"

	"go.einride.tech/sage/sg"
)

type MontaCmd sg.Namespace

func (MontaCmd) Default(ctx context.Context) error {
	sg.Deps(ctx, MontaCmd.GoModTidy)
	return nil
}

func (MontaCmd) GoModTidy(ctx context.Context) error {
	sg.Logger(ctx).Println("tidying Go module files...")
	cmd := sg.Command(ctx, "go", "mod", "tidy", "-v")
	cmd.Dir = sg.FromGitRoot("cmd", "monta")
	return cmd.Run()
}

func DependabotFix(ctx context.Context) error {
	sg.SerialDeps(ctx, Proto.All, Terraform.All, Backend.Default)
	sg.Deps(ctx, ConvcoCheck, FormatMarkdown, FormatYAML, GoLicenses)
	return nil
}
