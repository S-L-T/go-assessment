package port

import (
	"github.com/google/uuid"
	"github.com/S-L-T/go-assessment/internal/core/domain"
)

type CompanyUseCase interface {
	Get(id uuid.UUID) (domain.Company, error)
	Create(
		name string,
		description string,
		totalEmployees uint,
		isRegistered bool,
		companyType domain.CompanyType,
	) (uuid.UUID, error)
	Update(
		id uuid.UUID,
		name string,
		description string,
		totalEmployees uint,
		isRegistered bool,
		companyType domain.CompanyType,
	) error
	Delete(id uuid.UUID) error
	Healthcheck
}
