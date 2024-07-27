package handlers

import (
	"fmt"
	"strconv"

	"github.com/Naiggo/conectarc/auth"
	"github.com/Naiggo/conectarc/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path, method, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Going to process " + path + " -> " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isValidRequest, statusCode, user := validAuthorization(path, method, headers)
	if !isValidRequest {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcessUsers(body, path, method, user, id, request)
	case "prod":
		return ProcessProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcessStock(body, path, method, user, idn, request)
	case "addr":
		return ProcessAddress(body, path, method, user, idn, request)
	case "cate":
		return ProcessCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcessOrders(body, path, method, user, idn, request)
	}

	return 400, "Invalid Method"
}

func validAuthorization(path, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token required"
	}

	tokenValid, err, msg := auth.TokenValidation(token)
	if !tokenValid {
		if err != nil {
			fmt.Println("ValidAutorization Token on error - " + err.Error())
			return false, 401, err.Error()
		}
		fmt.Println("ValidAutorization Token on error - " + msg)
		return false, 401, msg
	}

	fmt.Println("ValidAutorization Token OK")
	return true, 200, msg
}

func ProcessUsers(body, path, method, user, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}

func ProcessProducts(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}

func ProcessCategory(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InserCategory(body, user)
	}
	return 400, "Invalid Method"
}

func ProcessStock(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}

func ProcessAddress(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}

func ProcessOrders(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Invalid Method"
}
