package httpserver

import (
	// auth
	authHTTP "thuchanhgolang/internal/auth/delivery/http"
	authMongo "thuchanhgolang/internal/auth/repository/mongo"
	authUsecase "thuchanhgolang/internal/auth/usecase"

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

	// users
	userHTTP "thuchanhgolang/internal/user/delivery/http"
	userMongo "thuchanhgolang/internal/user/repository/mongo"
	userUsecase "thuchanhgolang/internal/user/usecase"

	// JWT
	"thuchanhgolang/pkg/jwt"

	// Middleware
	"thuchanhgolang/internal/middleware"
)

func (srv HTTPServer) mapHandlers() {
	// JWT Manager
	jwtManager := jwt.NewManager(srv.jwtSecretKey)

	// Middleware (encrypter có thể nil tạm thời)
	authMiddleware := middleware.New(srv.l, jwtManager, nil)

	// Repositories
	authRepo := authMongo.NewRepository(srv.l, srv.database)
	shopRepo := shopMongo.NewRepository(srv.l, srv.database)
	regionRepo := regionMongo.NewRepository(srv.l, srv.database)
	branchRepo := branchMongo.NewRepository(srv.l, srv.database)
	departmentRepo := departmentMongo.NewRepository(srv.l, srv.database)
	userRepo := userMongo.NewRepository(srv.l, srv.database)

	// Usecases
	authUC := authUsecase.NewUsecase(srv.l, authRepo, jwtManager, srv.accessDuration)
	shopUC := shopUsecase.NewUsecase(srv.l, shopRepo)
	regionUC := regionUsecase.NewUsecase(srv.l, regionRepo)
	branchUC := branchUsecase.NewUsecase(srv.l, branchRepo)
	departmentUC := departmentUsecase.NewUsecase(srv.l, departmentRepo)
	userUC := userUsecase.NewUsecase(srv.l, userRepo, branchRepo, departmentRepo, regionRepo)

	// Handlers
	authH := authHTTP.New(srv.l, authUC)
	shopH := shopHTTP.New(srv.l, shopUC)
	regionH := regionHTTP.New(srv.l, regionUC)
	branchH := branchHTTP.New(srv.l, branchUC)
	departmentH := departmentHTTP.New(srv.l, departmentUC)
	userH := userHTTP.New(srv.l, userUC)

	// Routes
	api := srv.gin.Group("/api/v1")

	// Public routes (không cần token)
	authHTTP.MapRoutes(api.Group("/auth"), authH)

	// Protected routes (cần token)
	protected := api.Group("")
	protected.Use(authMiddleware.Auth())

	shopHTTP.MapRoutes(protected.Group("/shops"), shopH)
	regionHTTP.MapRoutes(protected.Group("/regions"), regionH)
	branchHTTP.MapRoutes(protected.Group("/branches"), branchH)
	departmentHTTP.MapRoutes(protected.Group("/departments"), departmentH)
	userHTTP.MapRoutes(protected.Group("/users"), userH)
}
