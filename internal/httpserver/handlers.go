package httpserver

import (
	// branches
	branchHTTP "thuchanhgolang/internal/branch/delivery/http"
	branchMongo "thuchanhgolang/internal/branch/repository/mongo"
	branchUsecase "thuchanhgolang/internal/branch/usecase"

	// departments
	departmentHTTP "thuchanhgolang/internal/department/delivery/http"
	departmentMongo "thuchanhgolang/internal/department/repository/mongo"
	departmentUsecase "thuchanhgolang/internal/department/usecase"

	// regions
	regionHTTP "thuchanhgolang/internal/region/delivery/http"
	regionMongo "thuchanhgolang/internal/region/repository/mongo"
	regionUsecase "thuchanhgolang/internal/region/usecase"

	// shops
	shopHTTP "thuchanhgolang/internal/shop/delivery/http"
	shopMongo "thuchanhgolang/internal/shop/repository/mongo"
	shopUsecase "thuchanhgolang/internal/shop/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// // jwtManager := jwt.NewManager(srv.jwtSecretKey)
	// Repositories
	shopRepo := shopMongo.NewRepository(srv.l, srv.database)
	regionRepo := regionMongo.NewRepository(srv.l, srv.database)
	branchRepo := branchMongo.NewRepository(srv.l, srv.database)
	departmentRepo := departmentMongo.NewRepository(srv.l, srv.database)
	// Usecases
	shopUC := shopUsecase.NewUsecase(srv.l, shopRepo)
	regionUC := regionUsecase.NewUsecase(srv.l, regionRepo)
	branchUC := branchUsecase.NewUsecase(srv.l, branchRepo)
	departmentUC := departmentUsecase.NewUsecase(srv.l, departmentRepo)
	// Handlers
	shopH := shopHTTP.New(srv.l, shopUC)
	regionH := regionHTTP.New(srv.l, regionUC)
	branchH := branchHTTP.New(srv.l, branchUC)
	departmentH := departmentHTTP.New(srv.l, departmentUC)
	// // Middlewares
	// // mw := middleware.New(srv.l, jwtManager, srv.encrypter)

	api := srv.gin.Group("/api/v1")

	shopHTTP.MapRoutes(api.Group("/shops"), shopH)
	regionHTTP.MapRoutes(api.Group("/regions"), regionH)
	branchHTTP.MapRoutes(api.Group("/branches"), branchH)
	departmentHTTP.MapRoutes(api.Group("/departments"), departmentH)
}
