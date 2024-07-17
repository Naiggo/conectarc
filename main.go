package main

import (
	"context"
	"os"
	"strings"

	"github.com/Naiggo/conectarc/awsgo"
	"github.com/Naiggo/conectarc/bd"
	"github.com/Naiggo/conectarc/handlers"
	_ "github.com/Naiggo/conectarc/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjectarLamnda)
}

func EjectarLamnda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InitializeAWS()

	if !ParamsValidation() {
		panic("Error on params. You should send SecrentName and UrlPrefix")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Handlers(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil
}

func ParamsValidation() bool {
	_, getParam := os.LookupEnv("SecretName")
	if !getParam {
		return getParam
	}

	_, getParam = os.LookupEnv("UrlPrefix")
	if !getParam {
		return getParam
	}

	return getParam
}

func ParamsValidation1() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam
}
