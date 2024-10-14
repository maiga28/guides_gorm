package models

import (
	"database/sql"
	"time"

	"github.com/maiga28/guides_gorm/initializers"
)

// init loads environment variables and initializes the database
func init() {
	initializers.LocalEnvVariables()
	initializers.Database()
}

type Users struct {
	ID           uint           `gorm:"primaryKey"` // GORM manages primary keys automatically
	Name         string         // Regular string field
	Email        *string        // Nullable string
	Age          uint8          // Unsigned 8-bit integer
	Birthday     *time.Time     // Nullable time
	MemberNumber sql.NullString // Nullable string
	ActivatedAt  sql.NullTime   // Nullable time
	CreatedAt    time.Time      // Managed by GORM for creation time
	UpdatedAt    time.Time      // Managed by GORM for update time
}

// TableName overrides the default table name
func (Users) TableName() string {
	return "users"
}
