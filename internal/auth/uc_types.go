package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

// RegisterInput là input để đăng ký user mới từ HTTP layer
type RegisterInput struct {
	Username string
	Password string
	Email    string
	ShopID   primitive.ObjectID // Shop mà user thuộc về
}

// RegisterOutput là kết quả sau khi đăng ký thành công
type RegisterOutput struct {
	ID       primitive.ObjectID
	Username string
	Email    string
	ShopID   primitive.ObjectID
	Token    string // JWT token
}

// LoginInput là input để đăng nhập từ HTTP layer
type LoginInput struct {
	Username string
	Password string
}

// LoginOutput là kết quả sau khi đăng nhập thành công
type LoginOutput struct {
	ID       primitive.ObjectID
	Username string
	Email    string
	Token    string // JWT token
}
