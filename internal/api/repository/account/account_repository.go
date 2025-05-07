package account

import (
	"context"
	"zond-api/internal/domain/model"
)

type AccountRepository interface {
	GetTopAccounts(ctx context.Context, page, limit int) ([]model.Account, int, error)
}
