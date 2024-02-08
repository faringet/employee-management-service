package entityservice

import "time"

type Entity struct {
	ID                           int
	WorkspaceID                  int
	CompanyStatusID              int
	CustomerStatusID             int
	OrganizationSizeCategoriesID int
	Name                         *string
	BoldBiSiteName               *string
	Details                      *string
	ImportLock                   *bool
	CreatedAt                    *time.Time
	UpdatedAt                    *time.Time
	CreatedBy                    *int
	UpdatedBy                    *int
}
