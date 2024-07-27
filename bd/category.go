package bd

import (
	"database/sql"
	"fmt"

	//"strconv"
	//"strings"

	"github.com/Naiggo/conectarc/models"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/Naiggo/conectarc/tools"
)

func InsertCategory(category models.Category) (int64, error) {
	fmt.Println("Starting InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentence := "INSERT INTO category (Categ_Name, Categ_Path) VALUS ('" + category.CategName + "','" + category.CategPath + "')"

	var result sql.Result
	result, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println("InsertCategory -> Successful excecution")
	return LastInsertedId, err
}
