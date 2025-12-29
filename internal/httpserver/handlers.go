package httpserver

import (
	regionHTTP "thuchanhgolang/internal/region/delivery/http"
	regionMongo "thuchanhgolang/internal/region/repository/mongo"
	regionUsecase "thuchanhgolang/internal/region/usecase"
	shopHTTP "thuchanhgolang/internal/shop/delivery/http"
	shopMongo "thuchanhgolang/internal/shop/repository/mongo"
	shopUsecase "thuchanhgolang/internal/shop/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// // jwtManager := jwt.NewManager(srv.jwtSecretKey)
	// Repositories
	shopRepo := shopMongo.NewRepository(srv.l, srv.database)
	regionRepo := regionMongo.NewRepository(srv.l, srv.database)

	// Usecases
	shopUC := shopUsecase.NewUsecase(srv.l, shopRepo)
	regionUC := regionUsecase.NewUsecase(srv.l, regionRepo)

	// Handlers
	shopH := shopHTTP.New(srv.l, shopUC)
	regionH := regionHTTP.New(srv.l, regionUC)

	// // Middlewares
	// // mw := middleware.New(srv.l, jwtManager, srv.encrypter)

	api := srv.gin.Group("/api/v1")

	shopHTTP.MapRoutes(api.Group("/shops"), shopH)
	regionHTTP.MapRoutes(api.Group("/regions"), regionH)
}
