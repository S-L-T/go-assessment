package port

import (
	"context"
	"github.com/google/uuid"
	"github.com/S-L-T/go-assessment/internal/core/domain"
)

type CompanyReader interface {
	Get(ctx context.Context, id uuid.UUID) (domain.Company, error)
}

type CompanyWriter interface {
	Insert(ctx context.Context, company domain.Company) (uuid.UUID, error)
	Update(ctx context.Context, company domain.Company) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type CompanyRepository interface {
	CompanyReader
	CompanyWriter
	Healthcheck
}
