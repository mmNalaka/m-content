// Code generated by sqlc. DO NOT EDIT.

package sqlc

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/tabbed/pqtype"
)

type Collection struct {
	ID         uuid.UUID             `json:"id"`
	Slug       string                `json:"slug"`
	Name       sql.NullString        `json:"name"`
	Notes      sql.NullString        `json:"notes"`
	Singleton  bool                  `json:"singleton"`
	Schema     json.RawMessage       `json:"schema"`
	Listrule   sql.NullString        `json:"listrule"`
	Viewrule   sql.NullString        `json:"viewrule"`
	Createrule sql.NullString        `json:"createrule"`
	Updaterule sql.NullString        `json:"updaterule"`
	Deleterule sql.NullString        `json:"deleterule"`
	Tags       pqtype.NullRawMessage `json:"tags"`
	Meta       pqtype.NullRawMessage `json:"meta"`
	CreatedBy  uuid.NullUUID         `json:"created_by"`
	UpdatedBy  uuid.NullUUID         `json:"updated_by"`
	CreatedAt  time.Time             `json:"created_at"`
	UpdatedAt  time.Time             `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID             `json:"id"`
	Email     string                `json:"email"`
	Password  string                `json:"password"`
	Name      string                `json:"name"`
	Avatar    sql.NullString        `json:"avatar"`
	Location  sql.NullString        `json:"location"`
	Title     sql.NullString        `json:"title"`
	Bio       sql.NullString        `json:"bio"`
	Status    string                `json:"status"`
	Tags      pqtype.NullRawMessage `json:"tags"`
	Meta      pqtype.NullRawMessage `json:"meta"`
	LastLogin sql.NullTime          `json:"last_login"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}
