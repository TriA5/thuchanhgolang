package httpserver

import (
	branchHTTP "thuchanhgolang/internal/branch/delivery/http"
	branchMongo "thuchanhgolang/internal/branch/repository/mongo"
	branchUsecase "thuchanhgolang/internal/branch/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// jwtManager := jwt.NewManager(srv.jwtSecretKey)
	// Repositories
	branchRepo := branchMongo.NewRepository(srv.l, srv.database)

	// Usecases
	branchUC := branchUsecase.New(srv.l, branchRepo)

	// Handlers
	branchH := branchHTTP.New(srv.l, branchUC)

	// Middlewares
	// mw := middleware.New(srv.l, jwtManager, srv.encrypter)

	api := srv.gin.Group("/api/v1")

	branchHTTP.MapRoutes(api.Group("/branches"), branchH)
}
