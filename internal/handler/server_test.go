package handler

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/engagerocketco/go-common/config"
	"github.com/engagerocketco/go-common/logger"
	"github.com/engagerocketco/go-common/ns"
	tu "github.com/engagerocketco/go-common/tests"
	"github.com/engagerocketco/templates-api-svc/internal/repository/postgres"
	"github.com/engagerocketco/templates-api-svc/internal/service/natsservice"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MockEntityService struct {
	Req *ns.GetEntityRequest
	Res *ns.GetEntityResponse
	Err error
}

type MockPermissionService struct {
	Req *ns.AccountInfoRequest
	Res *ns.AccountInfoResponse
	Err error
}

func (m MockEntityService) GetEntityByID(ctx context.Context, r *ns.GetEntityRequest) (*ns.GetEntityResponse, error) {
	return m.Res, m.Err
}

func (m MockEntityService) GetEntities(ctx context.Context, r *ns.GetEntitiesRequest) ([]ns.Entity, error) {
	return nil, nil
}

func (m MockEntityService) GetWorkspaceID(ctx context.Context, r *ns.GetEntityRequest) (int, error) {
	fmt.Println("not implemented")
	return 0, nil
}

func (m MockPermissionService) Validate(ctx context.Context, request *ns.ValidationPermissionRequestV1) (bool, error) {
	panic("not implemented")
}
func (m MockPermissionService) GetUserDivisions(ctx context.Context, req ns.UserDivisionsRequestV1) (*ns.UserDivisionsResponseV1, error) {
	panic("not implemented")
}
func (m MockPermissionService) GetAccountInfo(ctx context.Context, r *ns.AccountInfoRequest) (*ns.AccountInfo, error) {
	if r.Email == "test1@mail.com" {
		return &ns.AccountInfo{
			ID: 1,
		}, nil
	}

	return nil, nil
}

// NewTestServer will return configured handlers with postgres repository
// and current connection which is using in the repo.
func NewTestServer(ctx context.Context, t *testing.T) (http.Handler, *postgres.PostgresRepo, *tu.Conn, *MockEntityService) {

	log, err := logger.NewLogger("development")
	if err != nil {
		t.Fatalf("failed to create logger")
	}

	cfg, err := config.New()
	if err != nil {
		t.Fatalf("config failed: %s\n", err)
	}

	conn := tu.NewTestPgConn(ctx, t, cfg)
	db, err := sqlx.Connect("postgres", cfg.PostgresConfig.ConnectionString())
	if err != nil {
		t.Fatalf("create pg conn: %v", err)
	}

	postgresRepo, err := postgres.New(db, nil, log)
	if err != nil {
		t.Fatalf("init pg repo: %v", err)
	}

	//configure services
	mockEntityService := &MockEntityService{}
	mockPermissionService := &MockPermissionService{}
	natsService := natsservice.NewNatsService(mockEntityService, mockPermissionService, log)
	templateService := templateservice.New(postgresRepo, natsService, log)

	//create a new handler
	server := NewServer(cfg, templateService, log)

	repo, ok := postgresRepo.(*postgres.PostgresRepo)

	if !ok {
		t.Fatalf("must be postgres repo: %v", repo)
	}

	return server.Handler, repo, conn, mockEntityService
}
