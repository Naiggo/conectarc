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
