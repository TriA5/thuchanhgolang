package main

import (
	"log"
	"thuchanhgolang/config"
	"thuchanhgolang/internal/appconfig/mongo"
	"thuchanhgolang/internal/httpserver"
	pkgLog "thuchanhgolang/pkg/log"
	"time"
)

func main() {
	// Load config  (chạy main.go thì dòng này chạy thứ 1)
	//lần thứ 3 dòng này nhận được cấu hình từ config.go (URI và DBName đã được load từ file .env)
	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}

	// Connect to MongoDB
	//hàm này chạy tiếp thứ 4 để kết nối đến MongoDB với URI lấy từ cấu hình và nhảy vào connect
	client, err := mongo.Connect(cfg.Mongo.URI)

	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect(client)

	// Get Namedatabase
	db := client.Database(cfg.Mongo.DBName)

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	log.Println("Connected to MongoDB successfully!")

	// Sử dụng db để làm việc với collections
	// collection := db.Collection("your-collection")
	srv := httpserver.New(l, httpserver.Config{
		Port:           cfg.HTTPServer.Port,
		Database:       db,
		JWTSecretKey:   cfg.JWT.SecretKey,
		AccessDuration: time.Duration(cfg.JWT.AccessDuration) * time.Second,
	})
	srv.Run()
}
