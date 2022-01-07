package model

// Column names of the fields belonging to gorm.Model.
var ModelColumnNames = modelColumnNames{
	CreatedAt: "created_at",
	DeletedAt: "deleted_at",
	ID:        "id",
	UpdatedAt: "updated_at",
}

// Column names of the fields belonging to gorm.Model.
type modelColumnNames struct {
	// Column name of the CreatedAt field.
	CreatedAt string
	// Column name of the DeletedAt field.
	DeletedAt string
	// Column name of the ID field.
	ID string
	// Column name of the UpdatedAt field.
	UpdatedAt string
}
