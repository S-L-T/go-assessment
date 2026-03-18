package domain

import "github.com/google/uuid"

type Company struct {
	ID             uuid.UUID
	Name           string
	Description    string
	TotalEmployees uint
	IsRegistered   bool
	Type           CompanyType
}
