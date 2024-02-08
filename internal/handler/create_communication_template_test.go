package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/engagerocketco/go-common/tests"
	"github.com/engagerocketco/templates-api-svc/internal/handler/endpoints"
)

func addDataForTest(ctx context.Context, t *testing.T, tx *tests.Tx) {
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
}

func addTestFieldToBodyCheck(fieldName string) func(reader io.Reader, t *testing.T) {
	return func(reader io.Reader, t *testing.T) {
		resp := &endpoints.CommunicationTemplateResponse{}
		err := json.NewDecoder(reader).Decode(resp)
		if err != nil {
			t.Fatalf("decode: createTemplateComms: %v", err)
		}

		if resp.ID != 1 {
			t.Fatalf("id must be: 1, but got %d", resp.ID)
		}

		if resp.OwnerEntityID != 1 {
			t.Fatalf("owner entity id must be: 1, but got %d", resp.OwnerEntityID)
		}

		if resp.HeaderLogoID != 1 {
			t.Fatalf("header logo id must be: 1, but got %d", resp.HeaderLogoID)
		}

		if resp.ReminderDaysID != 1 {
			t.Fatalf("reminder days id must be: 1, but got %d", resp.ReminderDaysID)
		}

		if !*resp.IsSendReport {
			t.Fatalf("is send report must be: true, but got %v", resp.IsSendReport)
		}

		if *resp.Name != fieldName {
			t.Fatalf("is send report: %s, but got %s", fieldName, *resp.Name)
		}

		if *resp.Description != "test description" {
			t.Fatalf("description must be: test description, but got %s", *resp.Description)
		}

		parsedTime := time.Date(2012, 4, 23, 18, 25, 43, 511, time.UTC)
		if resp.TimeSendReport.Equal(parsedTime) {
			t.Fatalf("time send report must be: %s, but got %s", parsedTime, *resp.TimeSendReport)
		}

		if *resp.CreatedBy != 1 {
			t.Fatalf("create by must be: 1, but got %d", *resp.CreatedBy)
		}

		if *resp.UpdatedBy != 1 {
			t.Fatalf("update by must be: 1, but got %d", *resp.UpdatedBy)
		}

		timeNow := time.Now().In(resp.UpdatedAt.Location()).Add(1 * time.Second)

		if timeNow.Before(*resp.CreatedAt) {
			t.Fatalf("created at must be in delta 1 second, but got %s", *resp.CreatedAt)
		}

		if timeNow.Before(*resp.UpdatedAt) {
			t.Fatalf("updated at must be in delta 1 second, but got %s", *resp.CreatedAt)
		}
	}
}

func TestCreateCommunicationTemplateHandler(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*4)
	defer cancel()
	handler, pgRepo, conn, _ := NewTestServer(ctx, t)

	testCases := map[string]tests.TestCase{
		"create communication template unable to decode": {
			Expected:   `{"error":"unable to decode the request body"}`,
			StatusCode: http.StatusBadRequest,
			Method:     "POST",
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
				return "/api/v1/template/communication"
			},
		},
		"create communication template success": {
			Expected:   addTestFieldToBodyCheck("test1"),
			StatusCode: http.StatusOK,
			Method:     "POST",
			Setup: func() func() {
				tx, err := tests.NewTx(conn)
				if err != nil {
					t.Fatalf("failed to open tx: %s", err)
				}

				pgRepo.DB = tx

				addDataForTest(ctx, t, tx)

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
				return "/api/v1/template/communication"
			},
		},
		"create communication template empty required fields": {
			Expected:   `{"error":"incorrect request format","details":[{"field":"OwnerEntityID","message":"Key: 'CreateCommunicationTemplateRequest.OwnerEntityID' Error:Field validation for 'OwnerEntityID' failed on the 'required' tag"},{"field":"HeaderLogoID","message":"Key: 'CreateCommunicationTemplateRequest.HeaderLogoID' Error:Field validation for 'HeaderLogoID' failed on the 'required' tag"},{"field":"ReminderDaysID","message":"Key: 'CreateCommunicationTemplateRequest.ReminderDaysID' Error:Field validation for 'ReminderDaysID' failed on the 'required' tag"},{"field":"IsSendReport","message":"Key: 'CreateCommunicationTemplateRequest.IsSendReport' Error:Field validation for 'IsSendReport' failed on the 'required' tag"},{"field":"Name","message":"Key: 'CreateCommunicationTemplateRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"},{"field":"Description","message":"Key: 'CreateCommunicationTemplateRequest.Description' Error:Field validation for 'Description' failed on the 'required' tag"},{"field":"TimeSendReport","message":"Key: 'CreateCommunicationTemplateRequest.TimeSendReport' Error:Field validation for 'TimeSendReport' failed on the 'required' tag"},{"field":"CreatedBy","message":"Key: 'CreateCommunicationTemplateRequest.CreatedBy' Error:Field validation for 'CreatedBy' failed on the 'required' tag"},{"field":"UpdatedBy","message":"Key: 'CreateCommunicationTemplateRequest.UpdatedBy' Error:Field validation for 'UpdatedBy' failed on the 'required' tag"}]}`,
			StatusCode: http.StatusUnprocessableEntity,
			Method:     "POST",
			Setup: func() func() {
				return func() {}
			},
			Body: func() io.Reader {
				return strings.NewReader(`{}`)
			},
			AddHeaders: tests.AddAuthToken,
			GetUrl: func(t *testing.T) string {
				return "/api/v1/template/communication"
			},
		},
		"create communication template internal server error": {
			Expected:   `{"error":"internal server error"}`,
			StatusCode: http.StatusInternalServerError,
			Method:     "POST",
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
				return "/api/v1/template/communication"
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
