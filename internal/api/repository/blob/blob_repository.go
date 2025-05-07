package blob

import "zond-api/internal/domain/model"

type BlobRepository interface {
	GetBlobs(limit, offset int) ([]model.Blob, int, error)
}
