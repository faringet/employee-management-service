package handler

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/engagerocketco/go-common/ns"
	"github.com/engagerocketco/go-common/tests"
)

func TestGetCommunicationTemplateByIDWithEntityHandler(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*4)
	defer cancel()
	handler, pgRepo, conn, mockEntityService := NewTestServer(ctx, t)

	testCases := map[string]tests.TestCase{
		"get communication template by id with entity success": {
			Expected:   `{"id":1,"owner_entity_id":1,"header_logo_id":1,"reminder_days_id":1,"is_send_report":false,"name":"test_name","description":"test desc","time_send_report":"2024-01-25T11:09:25.103433Z","created_at":"2024-01-25T11:09:25.103433Z","updated_at":"2024-01-25T11:09:25.103433Z","created_by":1,"updated_by":1,"entity":{"id":1,"workspace_id":1,"company_status_id":1,"customer_status_id":1,"organization_size_categories_id":1,"name":"test_entity","bold_bi_site_name":"test_entity","details":"test_entity","import_lock":true,"created_at":"2022-10-25T12:30:00Z","updated_at":"2022-10-25T12:30:00Z","created_by":1,"updated_by":1}}`,
			StatusCode: http.StatusOK,
			Method:     "GET",
			Setup: func() func() {
				tx, err := tests.NewTx(conn)
				if err != nil {
					t.Fatalf("failed to open tx: %s", err)
				}

				str := "test_entity"
				id := 1
				date := time.Date(2022, 10, 25, 12, 30, 0, 0, time.UTC)
				boolPtr := true

				mockEntityService.Res = &ns.GetEntityResponse{
					ID:                           1,
					Name:                         &str,
					WorkspaceID:                  1,
					CompanyStatusID:              &id,
					CustomerStatusID:             &id,
					OrganizationSizeCategoriesID: &id,
					BoldBISiteName:               &str,
					Details:                      &str,
					ImportLock:                   &boolPtr,
					LogoID:                       &id,
					IDOrgstructure:               &id,
					CreatedBy:                    &id,
					UpdatedBy:                    &id,
					CreatedAt:                    &date,
					UpdatedAt:                    &date,
				}

				pgRepo.DB = tx

				addCommunicationTemplateForTest(ctx, t, tx)

				return func() {
					pgRepo.DB = conn
					if err := tx.Rollback(); err != nil {
						t.Fatalf("failed to rollback tx: %s", err)
					}
					mockEntityService.Res = nil
				}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/1/entity"
			},
		},
		"get communication template by id with entity not found": {
			Expected:   `{"error":"communication template records not found"}`,
			StatusCode: http.StatusNotFound,
			Method:     "GET",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/15/entity"
			},
		},
		"get communication template by id with entity error id must be an integer": {
			Expected:   `{"error":"request validation failed","details":[{"field":"id","message":"status id must be an integer"}]}`,
			StatusCode: http.StatusUnprocessableEntity,
			Method:     "GET",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/sd/entity"
			},
		},
		"get communication template by id with entity internal error": {
			Expected:   `{"error":"internal server error"}`,
			StatusCode: http.StatusInternalServerError,
			Method:     "GET",
			Setup: func() func() {
				tx, err := tests.NewTx(conn)
				if err != nil {
					t.Fatalf("failed to open tx: %s", err)
				}

				pgRepo.DB = tx

				if _, err := tx.ExecContext(ctx, "DROP TABLE templ_template_comms CASCADE"); err != nil {
					t.Fatalf("failed to drop table: %s", err)
				}

				return func() {
					pgRepo.DB = conn
					if err := tx.Rollback(); err != nil {
						t.Fatalf("failed to rollback tx: %s", err)
					}
				}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/1/entity"
			},
		},
		"get communication template by id with entity nats internal error": {
			Expected:   `{"error":"internal server error"}`,
			StatusCode: http.StatusInternalServerError,
			Method:     "GET",
			Setup: func() func() {
				tx, err := tests.NewTx(conn)
				if err != nil {
					t.Fatalf("failed to open tx: %s", err)
				}
				pgRepo.DB = tx

				addCommunicationTemplateForTest(ctx, t, tx)
				mockEntityService.Err = errors.New("some error")

				return func() {
					pgRepo.DB = conn
					if err := tx.Rollback(); err != nil {
						t.Fatalf("failed to rollback tx: %s", err)
					}
					mockEntityService.Err = nil
				}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/1/entity"
			},
		},
		"get communication template by id with entity nats not found internal error": {
			Expected:   `{"error":"communication template records not found"}`,
			StatusCode: http.StatusNotFound,
			Method:     "GET",
			Setup: func() func() {
				tx, err := tests.NewTx(conn)
				if err != nil {
					t.Fatalf("failed to open tx: %s", err)
				}
				pgRepo.DB = tx

				addCommunicationTemplateForTest(ctx, t, tx)
				mockEntityService.Err = ns.ErrEntityNotFound

				return func() {
					pgRepo.DB = conn
					if err := tx.Rollback(); err != nil {
						t.Fatalf("failed to rollback tx: %s", err)
					}
					mockEntityService.Err = nil
				}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/1/entity"
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(
			name,
			func(t *testing.T) {
				tests.RunTest(t, &testCase, handler, "")
			},
		)
	}
}
