package monta

import (
	"encoding/json"
	"time"
)

// WalletTransaction is a wallet transaction.
type WalletTransaction struct {
	// ID of the transaction.
	ID int64 `json:"id"`

	// FromAmount is the amount sent from the sender.
	FromAmount float64 `json:"fromAmount"`

	// FromCurrency is the currency sent by the sender.
	FromCurrency Currency `json:"fromCurrency"`

	// FromType is the type of sender.
	FromType ToFromType `json:"fromType"`

	// From is the raw JSON message representing sender of the transaction.
	From json.RawMessage `json:"from"`

	// FromTeam holds the parsed value of From when [FromType] is [TypeTeam].
	FromTeam *Team `json:"-"`

	// FromOperator holds the parsed value of From when [FromType] is [TypeOperator].
	FromOperator *Operator `json:"-"`

	// ToAmount is the amount received by the receiver.
	ToAmount float64 `json:"toAmount"`

	// Type of receiver: operator, team
	ToCurrency Currency `json:"toCurrency"`

	// Type of the sender.
	ToType ToFromType `json:"toType"`

	// From is the sender of the transaction.
	To json.RawMessage `json:"to"`

	// ToOperator is used when [ToType] is "operator".
	ToOperator *Operator `json:"-"`

	// ToTeam is used when [ToType] is "team".
	ToTeam *Team `json:"-"`

	// Exchange rate used for currency conversion.
	ExchangeRate float64 `json:"exchangeRate"`

	// Creation date of transaction.
	CreatedAt time.Time `json:"createdAt"`

	// Update date of transaction.
	UpdatedAt time.Time `json:"updatedAt"`

	// Reference type of this transaction.
	ReferenceType string `json:"referenceType"`

	// Reference id of this transaction.
	ReferenceID string `json:"referenceId"`

	// Transaction group of this transaction.
	Group TransactionGroup `json:"group"`

	// Transaction state of this transaction
	State WalletTransactionState `json:"state"`

	// A note that has been entered for this transaction.
	Note string `json:"note"`
}

// UnmarshalJSON implements [json.Unmarshaler].
func (w *WalletTransaction) UnmarshalJSON(data []byte) error {
	type jsonTransaction WalletTransaction
	var t jsonTransaction
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*w = WalletTransaction(t)
	switch w.FromType {
	case ToFromTypeTeam:
		var fromTeam Team
		if err := json.Unmarshal(w.From, &fromTeam); err != nil {
			return err
		}
		w.FromTeam = &fromTeam
	case ToFromTypeOperator:
		var fromOperator Operator
		if err := json.Unmarshal(w.From, &fromOperator); err != nil {
			return err
		}
		w.FromOperator = &fromOperator
	}
	switch w.ToType {
	case ToFromTypeTeam:
		var toTeam Team
		if err := json.Unmarshal(w.To, &toTeam); err != nil {
			return err
		}
		w.ToTeam = &toTeam
	case ToFromTypeOperator:
		var toOperator Operator
		if err := json.Unmarshal(w.To, &toOperator); err != nil {
			return err
		}
		w.ToOperator = &toOperator
	}
	return nil
}

// MarshalJSON implements [json.Marshaler].
func (w *WalletTransaction) MarshalJSON() ([]byte, error) {
	type jsonTransaction WalletTransaction
	t := jsonTransaction(*w)
	switch w.FromType {
	case ToFromTypeTeam:
		fromTeamData, err := json.Marshal(w.FromTeam)
		if err != nil {
			return nil, err
		}
		t.From = fromTeamData
	case ToFromTypeOperator:
		fromOperatorData, err := json.Marshal(w.FromOperator)
		if err != nil {
			return nil, err
		}
		t.From = fromOperatorData
	}
	switch w.ToType {
	case ToFromTypeTeam:
		toTeamData, err := json.Marshal(w.ToTeam)
		if err != nil {
			return nil, err
		}
		t.To = toTeamData
	case ToFromTypeOperator:
		toOperatorData, err := json.Marshal(w.ToOperator)
		if err != nil {
			return nil, err
		}
		t.To = toOperatorData
	}
	return json.Marshal(&t)
}

// ToFromType denotes the type of transaction sender or receiver.
type ToFromType string

// Known [ToFromType] values.
const (
	ToFromTypeOperator = "operator"
	ToFromTypeTeam     = "team"
)
