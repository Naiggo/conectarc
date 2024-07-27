package bd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Naiggo/conectarc/secretsm"

	"github.com/Naiggo/conectarc/models"

	_ "github.com/go-sql-driver/mysql"
)

var (
	SecretModel models.SecretRDSJson
	err         error
	Db          *sql.DB
	dbNameRds   = "conectar"
)

func ReadSecret() error {
	SecretModel, err = secretsm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println("DbConnect Error: ", err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("DbPing Error: ", err.Error())
		return err
	}

	fmt.Println("Successful Connection to DB")

	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = dbNameRds

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		dbUser, authToken, dbEndpoint, dbName)
	fmt.Println("Connection String created ", dbName)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("Starting UserIsAdmin validation")

	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}

	defer Db.Close()

	sentence := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"
	fmt.Println(sentence)

	rows, err := Db.Query(sentence)
	if err != nil {
		return false, err.Error()
	}

	var value string
	rows.Next()
	rows.Scan(&value)

	fmt.Println("UserIsAdmin > Successful Excecution - returned value " + value)
	if value == "1" {
		return true, ""
	}

	return false, "User is not Admin"
}
