package use_case

import (
	"context"
	"github.com/google/uuid"
	"github.com/S-L-T/go-assessment/helper"
	"github.com/S-L-T/go-assessment/internal/core/domain"
	"github.com/S-L-T/go-assessment/internal/core/port"
)

type Company struct {
	ctx  context.Context
	repo port.CompanyRepository
}

func NewCompanyUseCase(
	ctx context.Context,
	r port.CompanyRepository,
) port.CompanyUseCase {
	return &Company{
		ctx:  ctx,
		repo: r,
	}
}

func (c Company) Get(id uuid.UUID) (domain.Company, error) {
	company, err := c.repo.Get(c.ctx, id)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return domain.Company{}, err
	}

	return company, nil
}

func (c Company) Create(
	name string,
	description string,
	totalEmployees uint,
	isRegistered bool,
	companyType domain.CompanyType,
) (uuid.UUID, error) {
	company := domain.Company{
		ID:             uuid.New(),
		Name:           name,
		Description:    description,
		TotalEmployees: totalEmployees,
		IsRegistered:   isRegistered,
		Type:           companyType,
	}

	id, err := c.repo.Insert(c.ctx, company)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return uuid.UUID{}, err
	}

	return id, nil
}

func (c Company) Update(
	id uuid.UUID,
	name string,
	description string,
	totalEmployees uint,
	isRegistered bool,
	companyType domain.CompanyType,
) error {
	company := domain.Company{
		ID:             id,
		Name:           name,
		Description:    description,
		TotalEmployees: totalEmployees,
		IsRegistered:   isRegistered,
		Type:           companyType,
	}
	err := c.repo.Update(c.ctx, company)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return err
	}

	return nil
}

func (c Company) Delete(id uuid.UUID) error {
	err := c.repo.Delete(c.ctx, id)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		return err
	}

	return nil
}

func (c Company) IsAlive(ctx context.Context) bool {
	return c.repo.IsAlive(c.ctx)
}
