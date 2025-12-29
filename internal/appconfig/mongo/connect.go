package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"thuchanhgolang/pkg/mongo"
)

const (
	ctxTimeout = 10 * time.Second
)

func Connect(uri string) (mongo.Client, error) {
	// 5.1 Tạo context với 10 giây timeout nếu các operation vượt quá thời gian này sẽ bị cancel
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	// 5.2 Tạo MongoDB client
	//Nhảy vào newClient
	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to create mongo client: %w", err)
	}

	// 5.3: Thực hiện kết nối
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	// 5.4: Ping để test
	err = client.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping to DB: %w", err)
	}

	return client, nil
}

func Disconnect(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
