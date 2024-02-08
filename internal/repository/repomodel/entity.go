package repomodel

import "time"

type Entity struct {
	ID                           int        `db:"id"`
	Name                         *string    `db:"name"`
	WorkspaceID                  int        `db:"workspace_id"`
	CompanyStatusID              *int       `db:"company_status_id"`
	CustomerStatusID             *int       `db:"customer_status_id"`
	OrganizationSizeCategoriesID *int       `db:"organization_size_categories_id"`
	BoldBISiteName               *string    `db:"bold_bi_site_name"`
	Details                      *string    `db:"details"`
	ImportLock                   *bool      `db:"import_lock"`
	LogoID                       *int       `db:"logo_id"`
	IDOrgstructure               *int       `db:"id_orgstructure"`
	CreatedBy                    *int       `db:"created_by"`
	UpdatedBy                    *int       `db:"updated_by"`
	CreatedAt                    *time.Time `db:"created_at"`
	UpdatedAt                    *time.Time `db:"updated_at"`
}

type CustomerStatus struct {
	ID          int        `db:"id"`
	Name        *string    `db:"name"`
	Code        *string    `db:"code"`
	Description *string    `db:"description"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	CreatedBy   *int       `db:"created_by"`
	UpdatedBy   *int       `db:"updated_by"`
}
