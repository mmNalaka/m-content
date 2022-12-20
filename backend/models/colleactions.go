package models

import (
	"github.com/google/uuid"
)

// CollocationsCreateRequest is the request body for the collocations.create endpoint
type CollocationsCreateRequest struct {
	Name      string        `json:"name" validate:"required,max=255"`
	Slug      string        `json:"slug" validate:"required,max=255"`
	Notes     string        `json:"notes"`
	Singleton bool          `json:"singleton"`
	CreatedBy uuid.NullUUID `json:"created_by" validate:"required"`
}
