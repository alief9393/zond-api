package account_abstraction

import "zond-api/internal/domain/model"

type AccountAbstractionRepository interface {
	GetBundleTransactions(limit, offset int) ([]model.BundleTransaction, int, error)
	GetAATransactions(limit, offset int) ([]model.AATransaction, int, error)
}
