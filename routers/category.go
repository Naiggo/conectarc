package routers

import (
	"encoding/json"
	"strconv"

	"github.com/Naiggo/conectarc/bd"
	"github.com/Naiggo/conectarc/models"
	//"github.com/aws/aws-lambda-go/events"
)

func InserCategory(body, user string) (int, string) {
	var category models.Category

	err := json.Unmarshal([]byte(body), &category)
	if err != nil {
		return 400, "Error on recieved data " + err.Error()
	}

	if len(category.CategName) == 0 {
		return 400, "You must specify the Name of the category"
	}

	if len(category.CategPath) == 0 {
		return 400, "You must specify the Path of the category"
	}

	isAdmin, msg := bd.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}

	result, err := bd.InsertCategory(category)
	if err != nil {
		return 400, "Error while trying to insert category " + category.CategName + " > " + err.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}
