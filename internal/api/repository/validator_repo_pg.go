package repository

import (
	"context"
	"zond-api/internal/api/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ValidatorRepo interface {
	GetValidators(ctx context.Context) (*dto.ValidatorsResponse, error)
}

type ValidatorRepoPG struct {
	db *pgxpool.Pool
}

func NewValidatorRepoPG(db *pgxpool.Pool) ValidatorRepo {
	return &ValidatorRepoPG{db: db}
}

func (r *ValidatorRepoPG) GetValidators(ctx context.Context) (*dto.ValidatorsResponse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT public_key, validator_index, balance, status, effective_balance, 
		       activation_epoch, exit_epoch, chain_id, retrieved_from
		FROM validators
		ORDER BY validator_index ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var validators []dto.ValidatorResponse
	for rows.Next() {
		var validator dto.ValidatorResponse
		err := rows.Scan(
			&validator.PublicKey, &validator.Index, &validator.Balance, &validator.Status,
			&validator.EffectiveBalance, &validator.ActivationEpoch, &validator.ExitEpoch,
			&validator.ChainID, &validator.RetrievedFrom,
		)
		if err != nil {
			return nil, err
		}
		validators = append(validators, validator)
	}

	return &dto.ValidatorsResponse{Validators: validators}, nil
}
