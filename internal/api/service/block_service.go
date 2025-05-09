package service

import (
	"context"
	"fmt"

	"zond-api/internal/api/dto"
	blockRepo "zond-api/internal/api/repository/block"
)

type BlockService struct {
	repo blockRepo.BlockRepository
}

func NewBlockService(repo blockRepo.BlockRepository) *BlockService {
	return &BlockService{repo: repo}
}

func (s *BlockService) GetLatestBlocks(ctx context.Context, page, limit int) (dto.BlocksPaginatedResponse, error) {
	blocks, total, err := s.repo.GetPaginatedBlocks(ctx, page, limit)
	if err != nil {
		return dto.BlocksPaginatedResponse{}, err
	}

	var blockResponses []dto.BlockResponse
	for _, block := range blocks {
		blockResponses = append(blockResponses, dto.BlockResponse{
			BlockNumber:      block.BlockNumber,
			BlockHash:        fmt.Sprintf("0x%x", block.BlockHash),
			Timestamp:        block.Timestamp,
			MinerAddress:     fmt.Sprintf("0x%x", block.MinerAddress),
			Canonical:        block.Canonical,
			ParentHash:       fmt.Sprintf("0x%x", block.ParentHash),
			GasUsed:          block.GasUsed,
			GasLimit:         block.GasLimit,
			Size:             block.Size,
			TransactionCount: block.TransactionCount,
			ExtraData:        fmt.Sprintf("0x%x", block.ExtraData),
			BaseFeePerGas:    block.BaseFeePerGas,
			TransactionsRoot: fmt.Sprintf("0x%x", block.TransactionsRoot),
			StateRoot:        fmt.Sprintf("0x%x", block.StateRoot),
			ReceiptsRoot:     fmt.Sprintf("0x%x", block.ReceiptsRoot),
			LogsBloom:        fmt.Sprintf("0x%x", block.LogsBloom),
			ChainID:          block.ChainID,
			RetrievedFrom:    block.RetrievedFrom,
		})
	}

	return dto.BlocksPaginatedResponse{
		Blocks: blockResponses,
		Pagination: dto.PaginationInfo{
			Total: total,
			Page:  page,
			Limit: limit,
		},
	}, nil
}

func (s *BlockService) GetBlockByNumber(blockNumber int64) (*dto.BlockResponse, error) {
	block, err := s.repo.GetBlockByNumber(blockNumber)
	if err != nil {
		return nil, err
	}
	return &dto.BlockResponse{
		BlockNumber:      block.BlockNumber,
		BlockHash:        fmt.Sprintf("0x%x", block.BlockHash),
		Timestamp:        block.Timestamp,
		MinerAddress:     fmt.Sprintf("0x%x", block.MinerAddress),
		Canonical:        block.Canonical,
		ParentHash:       fmt.Sprintf("0x%x", block.ParentHash),
		GasUsed:          block.GasUsed,
		GasLimit:         block.GasLimit,
		Size:             block.Size,
		TransactionCount: block.TransactionCount,
		ExtraData:        fmt.Sprintf("0x%x", block.ExtraData),
		BaseFeePerGas:    block.BaseFeePerGas,
		TransactionsRoot: fmt.Sprintf("0x%x", block.TransactionsRoot),
		StateRoot:        fmt.Sprintf("0x%x", block.StateRoot),
		ReceiptsRoot:     fmt.Sprintf("0x%x", block.ReceiptsRoot),
		LogsBloom:        fmt.Sprintf("0x%x", block.LogsBloom),
		ChainID:          block.ChainID,
		RetrievedFrom:    block.RetrievedFrom,
	}, nil
}

func (s *BlockService) GetForkedBlocks(limit, offset int) (dto.BlocksPaginatedResponse, error) {
	blocks, err := s.repo.GetForkedBlocks(limit, offset)
	if err != nil {
		return dto.BlocksPaginatedResponse{}, err
	}

	var blockResponses []dto.BlockResponse
	for _, block := range blocks {
		blockResponses = append(blockResponses, dto.BlockResponse{
			BlockNumber:      block.BlockNumber,
			BlockHash:        fmt.Sprintf("0x%x", block.BlockHash),
			Timestamp:        block.Timestamp,
			MinerAddress:     fmt.Sprintf("0x%x", block.MinerAddress),
			Canonical:        block.Canonical,
			ParentHash:       fmt.Sprintf("0x%x", block.ParentHash),
			GasUsed:          block.GasUsed,
			GasLimit:         block.GasLimit,
			Size:             block.Size,
			TransactionCount: block.TransactionCount,
			ExtraData:        fmt.Sprintf("0x%x", block.ExtraData),
			BaseFeePerGas:    block.BaseFeePerGas,
			TransactionsRoot: fmt.Sprintf("0x%x", block.TransactionsRoot),
			StateRoot:        fmt.Sprintf("0x%x", block.StateRoot),
			ReceiptsRoot:     fmt.Sprintf("0x%x", block.ReceiptsRoot),
			LogsBloom:        fmt.Sprintf("0x%x", block.LogsBloom),
			ChainID:          block.ChainID,
			RetrievedFrom:    block.RetrievedFrom,
			ReorgDepth:       block.ReorgDepth,
		})
	}

	return dto.BlocksPaginatedResponse{
		Blocks: blockResponses,
		Pagination: dto.PaginationInfo{
			Total: len(blockResponses),
			Page:  offset/limit + 1,
			Limit: limit,
		},
	}, nil
}

func (s *BlockService) GetBlockByHash(ctx context.Context, hash string) (*dto.BlockResponse, error) {
	block, err := s.repo.GetBlockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	return &dto.BlockResponse{
		BlockNumber:      block.BlockNumber,
		BlockHash:        fmt.Sprintf("0x%x", block.BlockHash),
		Timestamp:        block.Timestamp,
		MinerAddress:     fmt.Sprintf("0x%x", block.MinerAddress),
		Canonical:        block.Canonical,
		ParentHash:       fmt.Sprintf("0x%x", block.ParentHash),
		GasUsed:          block.GasUsed,
		GasLimit:         block.GasLimit,
		Size:             block.Size,
		TransactionCount: block.TransactionCount,
		ExtraData:        fmt.Sprintf("0x%x", block.ExtraData),
		BaseFeePerGas:    block.BaseFeePerGas,
		TransactionsRoot: fmt.Sprintf("0x%x", block.TransactionsRoot),
		StateRoot:        fmt.Sprintf("0x%x", block.StateRoot),
		ReceiptsRoot:     fmt.Sprintf("0x%x", block.ReceiptsRoot),
		LogsBloom:        fmt.Sprintf("0x%x", block.LogsBloom),
		ChainID:          block.ChainID,
		RetrievedFrom:    block.RetrievedFrom,
	}, nil
}
