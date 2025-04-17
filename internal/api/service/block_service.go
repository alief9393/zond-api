package service

import (
	"fmt"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository"
)

type BlockService struct {
	repo repository.BlockRepository
}

func NewBlockService(repo repository.BlockRepository) *BlockService {
	return &BlockService{repo: repo}
}

func (s *BlockService) GetLatestBlocks(limit, offset int) (dto.BlocksResponse, error) {
	blocks, err := s.repo.GetLatestBlocks(limit, offset)
	if err != nil {
		return dto.BlocksResponse{}, err
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
	return dto.BlocksResponse{Blocks: blockResponses}, nil
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

func (s *BlockService) GetForkedBlocks(limit, offset int) (dto.BlocksResponse, error) {
	blocks, err := s.repo.GetForkedBlocks(limit, offset)
	if err != nil {
		return dto.BlocksResponse{}, err
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
	return dto.BlocksResponse{Blocks: blockResponses}, nil
}
