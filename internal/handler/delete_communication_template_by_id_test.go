package handler

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/engagerocketco/go-common/tests"
)

func TestDeleteCommunicationTemplateHandler(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*4)
	defer cancel()
	handler, pgRepo, conn, _ := NewTestServer(ctx, t)

	testCases := map[string]tests.TestCase{
		"delete communication template by id success": {
			Expected:   `null`,
			StatusCode: http.StatusOK,
			Method:     "DELETE",
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
		"delete communication template by id not found": {
			Expected:   `{"error":"communication template records not found"}`,
			StatusCode: http.StatusNotFound,
			Method:     "DELETE",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/15"
			},
		},
		"delete communication template by id error id must be an integer": {
			Expected:   `{"error":"request validation failed","details":[{"field":"id","message":"status id must be an integer"}]}`,
			StatusCode: http.StatusUnprocessableEntity,
			Method:     "DELETE",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/dsfj"
			},
		},
		"delete communication template by id internal error": {
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
