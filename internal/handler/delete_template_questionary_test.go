package handler

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/engagerocketco/go-common/tests"
)

func TestDeleteTemplateQuestionaryByIDHandler(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	handler, postgresRepo, conn, _ := NewTestServer(ctx, t)

	var tqID int

	testCases := map[string]tests.TestCase{
		"delete template questionary by id": {
			Expected:   nil,
			StatusCode: http.StatusOK,
			Method:     "DELETE",
			Setup: func() func() {
				tx, err := tests.NewTx(conn)
				if err != nil {
					t.Fatalf("can not open tx: %s", err)
				}

				postgresRepo.DB = tx

				// tqID = prepare(t, tx) //TODO?

				return func() {
					if err := tx.Rollback(); err != nil {
						t.Fatalf("can not rollback changes: %s", err)
					}
				}
			},
			AddHeaders: tests.AddAuthToken,
			RequestCtx: tests.AddEmailClaim(ctx, "test1@mail.com"),
			GetUrl: func(t *testing.T) string {
				return fmt.Sprintf("/api/v1/template/questionary/%d", tqID)
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
