package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

var adapter *httpadapter.HandlerAdapter

func init() {
	// Standart Go Route'larını burada tanımlıyoruz
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello! This is a Serverless Go API running on AWS Lambda 🚀"))
	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Serverless project is healthy!"))
	})

	// Standart HTTP Mux'ı Lambda'nın anlayacağı dile çeviren adaptör
	adapter = httpadapter.New(mux)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Gelen isteği al, adaptörden geçir ve sonucu dön
	return adapter.ProxyWithContext(ctx, req)
}

func main() {
	// Lambda runtime bu satırı bekler
	lambda.Start(Handler)
}
