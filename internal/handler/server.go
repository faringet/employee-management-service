package handler

import (
	"context"
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/engagerocketco/go-common/config"
	"github.com/engagerocketco/go-common/network"
	"github.com/engagerocketco/templates-api-svc/docs"
	"github.com/engagerocketco/templates-api-svc/internal/handler/transport"
	"github.com/engagerocketco/templates-api-svc/internal/service/templateservice"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	*http.Server
}

func NewServer(
	cfg *config.Config,
	templateService templateservice.Service,
	logger *zap.Logger,
	middlewares ...func(next http.Handler) http.Handler,
) *Server {
	srv := Server{
		&http.Server{Addr: fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort)},
	}

	docs.SwaggerInfo.Host = cfg.SwaggerConfig.Host

	transport.NewTransportLogger(logger)

	r := mux.NewRouter().PathPrefix("/api/v1/template").Subrouter()

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	r.Handle(
		"/questionary",
		network.ApplyMiddlewares(transport.MakeGetTemplTemplateQuestionaryHandler(templateService, logger), middlewares...)).
		Methods(http.MethodGet)
	r.Handle("/questionary/{id}",
		network.ApplyMiddlewares(transport.MakeGetTemplTemplateQuestionaryByIDHandler(templateService, logger), middlewares...)).
		Methods(http.MethodGet)
	r.Handle("/questionary",
		network.ApplyMiddlewares(transport.MakeCreateTemplTemplateQuestionaryHandler(templateService, logger), middlewares...)).
		Methods(http.MethodPost)
	r.Handle("/questionary/{id}",
		network.ApplyMiddlewares(transport.MakeUpdateTemplTemplateQuestionaryByIDHandler(templateService, logger), middlewares...)).
		Methods(http.MethodPut)
	r.Handle("/questionary/{id}",
		network.ApplyMiddlewares(transport.MakeDeleteTemplTemplateQuestionaryByIDHandler(templateService, logger), middlewares...)).
		Methods(http.MethodDelete)
	r.Handle(
		"/ping",
		transport.MakePingHandler(logger),
	).Methods(http.MethodGet)

	r.Handle(
		"/communication",
		network.ApplyMiddlewares(transport.MakeCreateCommunicationTemplateHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPost)

	r.Handle(
		"/communication/{id}",
		network.ApplyMiddlewares(transport.MakeGetCommunicationTemplateByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	r.Handle(
		"/communication/{id}/entity",
		network.ApplyMiddlewares(transport.MakeGetCommunicationTemplateByIDWithEntityHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	r.Handle(
		"/communication/entity/{id}",
		network.ApplyMiddlewares(transport.MakeGetCommunicationTemplatesByEntityIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	r.Handle(
		"/communication/{id}",
		network.ApplyMiddlewares(transport.MakeCommunicationUpdateTemplateByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPatch)

	r.Handle(
		"/communication/{id}",
		network.ApplyMiddlewares(transport.MakeDeleteCommunicationTemplateByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodDelete)

	r.Handle(
		"/survey_tags/{id}",
		network.ApplyMiddlewares(transport.MakeGetSurveyTagsByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	// UNCOMMNECT IF YOU DO NOT WANT TO USE PAGINATION OR RENAME ROUTE PATH
	// r.Handle(
	// 	"/survey_tags",
	// 	network.ApplyMiddlewares(transport.MakeGetSurveyTagsHandler(templateService, logger), middlewares...),
	// ).Methods(http.MethodGet)


	r.Handle(
		"/survey_tags",
		network.ApplyMiddlewares(transport.MakeGetSurveyTagsPaginationHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	r.Handle(
		"/survey_tags",
		network.ApplyMiddlewares(transport.MakeCreateSurveyTagsHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPost)

	r.Handle(
		"/survey_tags/range",
		network.ApplyMiddlewares(transport.MakeCreateSurveyTagsRangeHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPost)

	r.Handle(
		"/survey_tags/{id}",
		network.ApplyMiddlewares(transport.MakeUpdateSurveyTagsByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPatch)

	r.Handle(
		"/survey_tags/{id}",
		network.ApplyMiddlewares(transport.MakeDeleteSurveyTagsByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodDelete)

	r.Handle(
		"/templ_recomended_frequancy/{id}",
		network.ApplyMiddlewares(transport.MakeGetTemplRecomendedFrequancyByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	r.Handle(
		"/templ_recomended_frequancy",
		network.ApplyMiddlewares(transport.MakeGetTemplRecomendedFrequancyHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	r.Handle(
		"/templ_recomended_frequancy",
		network.ApplyMiddlewares(transport.MakeCreateTemplRecomendedFrequancyHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPost)

	r.Handle(
		"/templ_recomended_frequancy/range",
		network.ApplyMiddlewares(transport.MakeCreateTemplRecomendedFrequancyRangeHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPost)

	r.Handle(
		"/templ_recomended_frequancy/{id}",
		network.ApplyMiddlewares(transport.MakeUpdateTemplRecomendedFrequancyByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPatch)

	r.Handle(
		"/templ_recomended_frequancy/{id}",
		network.ApplyMiddlewares(transport.MakeDeleteTemplRecomendedFrequancyByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodDelete)

	r.Handle(
		"/templ_questionary_tags/{id}",
		network.ApplyMiddlewares(transport.MakeGetTemplQuestionaryTagsByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	r.Handle(
		"/templ_questionary_tags",
		network.ApplyMiddlewares(transport.MakeGetTemplQuestionaryTagsHandler(templateService, logger), middlewares...),
	).Methods(http.MethodGet)

	r.Handle(
		"/templ_questionary_tags",
		network.ApplyMiddlewares(transport.MakeCreateTemplQuestionaryTagsHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPost)

	r.Handle(
		"/templ_questionary_tags/range",
		network.ApplyMiddlewares(transport.MakeCreateTemplQuestionaryTagsRangeHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPost)

	r.Handle(
		"/templ_questionary_tags/{id}",
		network.ApplyMiddlewares(transport.MakeUpdateTemplQuestionaryTagsByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodPatch)

	r.Handle(
		"/templ_questionary_tags/{id}",
		network.ApplyMiddlewares(transport.MakeDeleteTemplQuestionaryTagsByIDHandler(templateService, logger), middlewares...),
	).Methods(http.MethodDelete)

	srv.Handler = r
	return &srv
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			errCh <- fmt.Errorf("server: failed to listen and serve: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		return s.Shutdown(context.Background())
	case err := <-errCh:
		return err
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
