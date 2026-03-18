package mysql

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/S-L-T/go-assessment/internal/core/domain"
)

type Adapter struct {
	ID             uuid.UUID
	Name           string
	Description    sql.NullString
	TotalEmployees uint
	IsRegistered   bool
	Type           uint8
}

func NewAdapter(c domain.Company) Adapter {
	d := sql.NullString{
		String: c.Description,
		Valid:  true,
	}

	return Adapter{
		ID:             c.ID,
		Name:           c.Name,
		Description:    d,
		TotalEmployees: c.TotalEmployees,
		IsRegistered:   c.IsRegistered,
		Type:           uint8(c.Type),
	}
}

func (a Adapter) ToDomain() (domain.Company, error) {
	d := ""
	if a.Description.Valid {
		d = a.Description.String
	}

	return domain.Company{
		ID:             a.ID,
		Name:           a.Name,
		Description:    d,
		TotalEmployees: a.TotalEmployees,
		IsRegistered:   a.IsRegistered,
		Type:           domain.CompanyType(a.Type),
	}, nil
}
