package apm

import (
	"fmt"
)

// KeyTransaction represents information about a New Relic key transaction.
type KeyTransaction struct {
	ID              int                       `json:"id,omitempty"`
	Name            string                    `json:"name,omitempty"`
	TransactionName string                    `json:"transaction_name,omitempty"`
	HealthStatus    string                    `json:"health_status,omitempty"`
	LastReportedAt  string                    `json:"last_reported_at,omitempty"`
	Reporting       bool                      `json:"reporting"`
	Summary         ApplicationSummary        `json:"application_summary,omitempty"`
	EndUserSummary  ApplicationEndUserSummary `json:"end_user_summary,omitempty"`
	Links           KeyTransactionLinks       `json:"links,omitempty"`
}

// KeyTransactionLinks represents associations for a key transaction.
type KeyTransactionLinks struct {
	Application int `json:"application,omitempty"`
}

// ListKeyTransactionsParams represents a set of filters to be
// used when querying New Relic key transactions.
type ListKeyTransactionsParams struct {
	Name string `url:"filter[name],omitempty"`
	IDs  []int  `url:"filter[ids],omitempty,comma"`
}

// ListKeyTransactions returns all key transactions for an account.
func (apm *APM) ListKeyTransactions(params *ListKeyTransactionsParams) ([]*KeyTransaction, error) {
	results := []*KeyTransaction{}
	nextURL := "/key_transactions.json"

	for nextURL != "" {
		response := keyTransactionsResponse{}
		resp, err := apm.client.Get(nextURL, &params, &response)

		if err != nil {
			return nil, err
		}

		results = append(results, response.KeyTransactions...)

		paging := apm.pager.Parse(resp)
		nextURL = paging.Next
	}

	return results, nil
}

// GetKeyTransaction returns a specific key transaction by ID.
func (apm *APM) GetKeyTransaction(id int) (*KeyTransaction, error) {
	response := keyTransactionResponse{}
	u := fmt.Sprintf("/key_transactions/%d.json", id)

	_, err := apm.client.Get(u, nil, &response)

	if err != nil {
		return nil, err
	}

	return &response.KeyTransaction, nil
}

type keyTransactionsResponse struct {
	KeyTransactions []*KeyTransaction `json:"key_transactions,omitempty"`
}

type keyTransactionResponse struct {
	KeyTransaction KeyTransaction `json:"key_transaction,omitempty"`
}
