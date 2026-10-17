package main

import (
	"context"
	"fmt"
)

type Request struct{ UserID string }
type Response struct{ Status int }

// Обработчик (функция высшего порядка, как в net/http)
type Handler func(ctx context.Context, req Request) Response

// Middleware (декоратор, создающий цепочку)
type Middleware func(Handler) Handler

func LoggingMiddleware(next Handler) Handler {
	return func(ctx context.Context, req Request) Response {
		fmt.Println("Logging: Request received")
		return next(ctx, req)
	}
}

func AuthMiddleware(next Handler) Handler {
	return func(ctx context.Context, req Request) Response {
		if req.UserID == "" {
			fmt.Println("Auth: Unauthorized")
			return Response{Status: 401}
		}
		fmt.Println("Auth: User authenticated")
		return next(ctx, req)
	}
}

func MainHandler(ctx context.Context, req Request) Response {
	fmt.Printf("Main: Processing request for user %s\n", req.UserID)
	return Response{Status: 200}
}

// Функция для сборки цепочки
func Chain(h Handler, middlewares ...Middleware) Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func main() {
	chain := Chain(MainHandler, LoggingMiddleware, AuthMiddleware)
	
	chain(context.Background(), Request{UserID: "user123"}) // Пройдет всю цепочку
	fmt.Println("---")
	chain(context.Background(), Request{}) // Остановится на Auth
}