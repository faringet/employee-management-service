package handler

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/engagerocketco/go-common/tests"
)

func addCommunicationTemplateForTest(ctx context.Context, t *testing.T, tx *tests.Tx) {
	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE file RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("failed to truncate table: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE workspace RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("failed to truncate table: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE company_status RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("failed to truncate table: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE customer_status RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("failed to truncate table: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE organization_size_categories RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("failed to truncate table: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE entity RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("failed to truncate table: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "TRUNCATE TABLE templ_template_comms RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("failed to truncate table: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "INSERT INTO file (name, path) VALUES ($1, $2)",
		"test_name", "test_path"); err != nil {
		t.Fatalf("failed to insert row: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "INSERT INTO workspace (name, proj_m, industry, sign_in_method, demo) VALUES ($1, $2, $3, $4, $5)",
		"test_name", false, "test_industry", 1, true); err != nil {
		t.Fatalf("failed to insert row: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "INSERT INTO company_status (name, code, description) VALUES ($1, $2, $3)",
		"test_status", 123, "test_desc"); err != nil {
		t.Fatalf("failed to insert row: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "INSERT INTO customer_status (name, code, description) VALUES ($1, $2, $3)",
		"test_status", 123, "test_desc"); err != nil {
		t.Fatalf("failed to insert row: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "INSERT INTO organization_size_categories (name, code, description, number_of_employees) VALUES ($1, $2, $3, $4)",
		"test_name", 123, "test desc", 1); err != nil {
		t.Fatalf("failed to insert row: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "INSERT INTO entity (name, workspace_id, company_status_id, customer_status_id, organization_size_categories_id, bold_bi_site_name, import_lock, created_by, updated_by, logo_id, id_orgstructure) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		"test_name", 1, 1, 1, 1, "test", false, 1, 1, 1, 1); err != nil {
		t.Fatalf("failed to insert row: %s", err)
	}

	if _, err := tx.ExecContext(ctx, "INSERT INTO templ_template_comms (id, name, description, is_send_report, reminder_days_id, time_send_report, header_logo_id, owner_entity_id, updated_at, created_at, updated_by, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		1, "test_name", "test desc", false, 1, "2024-01-25T11:09:25.1034333", 1, 1, "2024-01-25T11:09:25.1034333", "2024-01-25T11:09:25.1034333", 1, 1); err != nil {
		t.Fatalf("failed to insert row: %s", err)
	}
}

func TestGetCommunicationTemplateByIDHandler(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*4)
	defer cancel()
	handler, pgRepo, conn, _ := NewTestServer(ctx, t)

	testCases := map[string]tests.TestCase{
		"get communication template by id success": {
			Expected:   `{"id":1,"owner_entity_id":1,"header_logo_id":1,"reminder_days_id":1,"is_send_report":false,"name":"test_name","description":"test desc","time_send_report":"2024-01-25T11:09:25.103433Z","created_at":"2024-01-25T11:09:25.103433Z","updated_at":"2024-01-25T11:09:25.103433Z","created_by":1,"updated_by":1}`,
			StatusCode: http.StatusOK,
			Method:     "GET",
			Setup: func() func() {
				tx, err := tests.NewTx(conn)
				if err != nil {
					t.Fatalf("failed to open tx: %s", err)
				}

				pgRepo.DB = tx

				addCommunicationTemplateForTest(ctx, t, tx)

				return func() {
					pgRepo.DB = conn
					if err := tx.Rollback(); err != nil {
						t.Fatalf("failed to rollback tx: %s", err)
					}
				}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/1"
			},
		},
		"get communication template by id not found": {
			Expected:   `{"error":"communication template records not found"}`,
			StatusCode: http.StatusNotFound,
			Method:     "GET",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/15"
			},
		},
		"get communication template by id error id must be an integer": {
			Expected:   `{"error":"request validation failed","details":[{"field":"id","message":"status id must be an integer"}]}`,
			StatusCode: http.StatusUnprocessableEntity,
			Method:     "GET",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/dsfj"
			},
		},
		"get communication template by id internal error": {
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
				return "/api/v1/template/communication/1"
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
