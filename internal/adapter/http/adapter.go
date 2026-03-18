package http

import (
	"github.com/google/uuid"
	"github.com/S-L-T/go-assessment/internal/core/domain"
)

type ErrorRes struct {
	Error string `json:"error"`
}

type GetReqAdapter struct {
	ID string `json:"id" validate:"required"`
}

type GetResAdapter struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	TotalEmployees int    `json:"total_employees"`
	IsRegistered   bool   `json:"is_registered"`
	Type           int    `json:"type"`
}

func NewGetResAdapter(c domain.Company) GetResAdapter {
	return GetResAdapter{
		ID:             c.ID.String(),
		Name:           c.Name,
		Description:    c.Description,
		TotalEmployees: int(c.TotalEmployees),
		IsRegistered:   c.IsRegistered,
		Type:           int(c.Type),
	}
}

type PutReqAdapter struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description"`
	TotalEmployees uint   `json:"total_employees" validate:"required"`
	IsRegistered   bool   `json:"is_registered" validate:"required"`
	Type           uint8  `json:"type" validate:"required"`
}

type PutResAdapter struct {
	ID string `json:"id"`
}

func (a PutReqAdapter) ToDomain() (domain.Company, error) {
	return domain.Company{
		Name:           a.Name,
		Description:    a.Description,
		TotalEmployees: uint(a.TotalEmployees),
		IsRegistered:   a.IsRegistered,
		Type:           domain.CompanyType(a.Type),
	}, nil
}

type PatchReqAdapter struct {
	ID             string `json:"id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description"`
	TotalEmployees uint    `json:"total_employees" validate:"required"`
	IsRegistered   bool   `json:"is_registered" validate:"required"`
	Type           uint8    `json:"type" validate:"required"`
}

func (a PatchReqAdapter) ToDomain() (domain.Company, error) {
	parsedID, err := uuid.Parse(a.ID)
	if err != nil {
		return domain.Company{}, err
	}

	return domain.Company{
		ID:             parsedID,
		Name:           a.Name,
		Description:    a.Description,
		TotalEmployees: uint(a.TotalEmployees),
		IsRegistered:   a.IsRegistered,
		Type:           domain.CompanyType(a.Type),
	}, nil
}

type DeleteReqAdapter struct {
	ID string `json:"id" validate:"required"`
}
