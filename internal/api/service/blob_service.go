package service

import (
	"context"
	"zond-api/internal/api/dto"
	"zond-api/internal/api/repository/blob"
)

type BlobService interface {
	GetBlobs(ctx context.Context, page, limit int) ([]dto.BlobResponse, int, error)
}

type blobService struct {
	repo blob.BlobRepository
}

func NewBlobService(repo blob.BlobRepository) BlobService {
	return &blobService{repo: repo}
}

func (s *blobService) GetBlobs(ctx context.Context, page, limit int) ([]dto.BlobResponse, int, error) {
	offset := (page - 1) * limit
	blobs, total, err := s.repo.GetBlobs(limit, offset)
	if err != nil {
		return nil, 0, err
	}

	var result []dto.BlobResponse
	for _, b := range blobs {
		result = append(result, dto.BlobResponse{
			VersionedHash: b.VersionedHash,
			TxHash:        b.TxHash,
			BlockNumber:   b.BlockNumber,
			Timestamp:     b.Timestamp,
			BlobSender:    b.BlobSender,
			GasPrice:      b.GasPrice,
			Size:          b.Size,
			RetrievedFrom: b.RetrievedFrom,
		})
	}

	return result, total, nil
}
