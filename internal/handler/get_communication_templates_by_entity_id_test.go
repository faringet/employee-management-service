package handler

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/engagerocketco/go-common/tests"
)

func TestGetCommunicationTemplatesByEntityIDHandler(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*4)
	defer cancel()
	handler, pgRepo, conn, _ := NewTestServer(ctx, t)

	testCases := map[string]tests.TestCase{
		"get communication templates by entity id success": {
			Expected:   `[{"id":1,"owner_entity_id":1,"header_logo_id":1,"reminder_days_id":1,"is_send_report":false,"name":"test_name","description":"test desc","time_send_report":"2024-01-25T11:09:25.103433Z","created_at":"2024-01-25T11:09:25.103433Z","updated_at":"2024-01-25T11:09:25.103433Z","created_by":1,"updated_by":1}]`,
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
				return "/api/v1/template/communication/entity/1"
			},
		},
		"get communication templates by entity id not found": {
			Expected:   `{"error":"communication template records not found"}`,
			StatusCode: http.StatusNotFound,
			Method:     "GET",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/entity/15"
			},
		},
		"get communication templates by entity error id must be an integer": {
			Expected:   `{"error":"request validation failed","details":[{"field":"id","message":"status id must be an integer"}]}`,
			StatusCode: http.StatusUnprocessableEntity,
			Method:     "GET",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/entity/ss"
			},
		},
		"get communication templates by id internal error": {
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
			Body: func() io.Reader {
				return strings.NewReader(`{"id" : 15}`)
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/entity/15"
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
