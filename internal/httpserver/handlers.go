package httpserver

import (
	shopHTTP "thuchanhgolang/internal/shop/delivery/http"
	shopMongo "thuchanhgolang/internal/shop/repository/mongo"
	shopUsecase "thuchanhgolang/internal/shop/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// // jwtManager := jwt.NewManager(srv.jwtSecretKey)
	// Repositories
	shopRepo := shopMongo.NewRepository(srv.l, srv.database)

	// Usecases
	shopUC := shopUsecase.NewUsecase(srv.l, shopRepo)

	// Handlers
	shopH := shopHTTP.New(srv.l, shopUC)

	// // Middlewares
	// // mw := middleware.New(srv.l, jwtManager, srv.encrypter)

	api := srv.gin.Group("/api/v1")

	shopHTTP.MapRoutes(api.Group("/shops"), shopH)
}
