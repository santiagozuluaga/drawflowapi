package config

import (
	"fmt"
	"os"

	"github.com/go-chi/cors"
)

var (
	JWT_SECRET   = []byte("1a3e1095c042ape24f5b3f38ec5e0e87")
	PORT         = ":5000"
	DATABASE_URI = "localhost:9080"
)

func init() {
	fmt.Println("GET ENV VARS")
	if os.Getenv("JWT_SECRET") != "" {
		JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
	}
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	if os.Getenv("DATABASE_URI") != "" {
		DATABASE_URI = os.Getenv("DATABASE_URI")
	}
}

func CORS() cors.Options {
	handler := cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Access-Control-Allow-Origin",
			"Accept",
			"Authorization",
			"Content-Type",
		},
	}
	return handler
}
