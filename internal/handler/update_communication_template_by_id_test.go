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

func TestUpdateCommunicationTemplateHandler(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*4)
	defer cancel()
	handler, pgRepo, conn, _ := NewTestServer(ctx, t)

	testCases := map[string]tests.TestCase{
		"patch communication template by id success": {
			Expected:   addTestFieldToBodyCheck("test3"),
			StatusCode: http.StatusOK,
			Method:     "PATCH",
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
			Body: func() io.Reader {
				return strings.NewReader(`{
   					"owner_entity_id" : 1,
    				"header_logo_id" : 1,
    				"reminder_days_id" : 1,
    				"is_send_report" : true,
   					"name" : "test3",
    				"description" : "test description",
    				"time_send_report" : "2012-04-23T18:25:43.511Z",
    				"updated_by" : 1
				}`)
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/1"
			},
		},
		"patch communication template by id unable to decode": {
			Expected:   `{"error":"unable to decode the request body"}`,
			StatusCode: http.StatusBadRequest,
			Method:     "PATCH",
			Setup: func() func() {
				return func() {}
			},
			Body: func() io.Reader {
				return strings.NewReader(`{
   					"owner_entity_id" : 5,
    				"header_logo_id" : 1,
    				"reminder_days_id" : 1,
    				"is_send_report" : true,
   					 "name" : "test",
    				"description" : "test description",
    				"time_send_report" : "2012-04-23T18:25:43.511Z",
    				"created_by" : 1,
    				"updated_by" : 1,
					"updated_at" : "2024-01-25T11:09:25.1034333",
					"created_at" : "2024-01-25T11:09:25.1034333",
				}`)
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/1"
			},
		},
		"patch communication template by id empty fields": {
			Expected:   `{"error":"Incorrect request format","details":[{"field":"OwnerEntityID","message":"Key: 'UpdateCommunicationTemplateByIDRequest.OwnerEntityID' Error:Field validation for 'OwnerEntityID' failed on the 'required' tag"},{"field":"HeaderLogoID","message":"Key: 'UpdateCommunicationTemplateByIDRequest.HeaderLogoID' Error:Field validation for 'HeaderLogoID' failed on the 'required' tag"},{"field":"ReminderDaysID","message":"Key: 'UpdateCommunicationTemplateByIDRequest.ReminderDaysID' Error:Field validation for 'ReminderDaysID' failed on the 'required' tag"},{"field":"IsSendReport","message":"Key: 'UpdateCommunicationTemplateByIDRequest.IsSendReport' Error:Field validation for 'IsSendReport' failed on the 'required' tag"},{"field":"Name","message":"Key: 'UpdateCommunicationTemplateByIDRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"},{"field":"Description","message":"Key: 'UpdateCommunicationTemplateByIDRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag"},{"field":"TimeSendReport","message":"Key: 'UpdateCommunicationTemplateByIDRequest.TimeSendReport' Error:Field validation for 'TimeSendReport' failed on the 'required' tag"},{"field":"UpdatedBy","message":"Key: 'UpdateCommunicationTemplateByIDRequest.UpdatedBy' Error:Field validation for 'UpdatedBy' failed on the 'required' tag"}]}`,
			StatusCode: http.StatusUnprocessableEntity,
			Method:     "PATCH",
			Setup: func() func() {
				return func() {}
			},
			Body: func() io.Reader {
				return strings.NewReader(`{}`)
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/1"
			},
		},
		"patch communication template by id not found": {
			Expected:   `{"error":"communication template records not found"}`,
			StatusCode: http.StatusNotFound,
			Method:     "PATCH",
			Setup: func() func() {
				return func() {}
			},
			Body: func() io.Reader {
				return strings.NewReader(`{
   					"owner_entity_id" : 1,
    				"header_logo_id" : 1,
    				"reminder_days_id" : 1,
    				"is_send_report" : true,
   					"name" : "test1",
    				"description" : "test description",
    				"time_send_report" : "2012-04-23T18:25:43.511Z",
    				"created_by" : 1,
    				"updated_by" : 1
				}`)
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/15"
			},
		},
		"patch communication template by id error id must be an integer": {
			Expected:   `{"error":"Request validation failed","details":[{"field":"id","message":"customer id must be an integer"}]}`,
			StatusCode: http.StatusUnprocessableEntity,
			Method:     "PATCH",
			Setup: func() func() {
				return func() {}
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication/dsfj"
			},
		},
		"patch communication template by id internal error": {
			Expected:   `{"error":"internal server error"}`,
			StatusCode: http.StatusInternalServerError,
			Method:     "PATCH",
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
				return strings.NewReader(`{
   					"owner_entity_id" : 1,
    				"header_logo_id" : 1,
    				"reminder_days_id" : 1,
    				"is_send_report" : true,
   					"name" : "test1",
    				"description" : "test description",
    				"time_send_report" : "2012-04-23T18:25:43.511Z",
    				"created_by" : 1,
    				"updated_by" : 1
				}`)
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
