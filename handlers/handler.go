package handlers

import (
	"fmt"
	//"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path, method, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Going to process " + path + " -> " + method)

	//id := request.PathParameters["id"]
	//idn, _ := strconv.Atoi(id)

	return 400, "Invalid Method"
}
