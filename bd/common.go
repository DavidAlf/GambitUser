package bd

import (
	"database/sql"
	"fmt"
	"os"

	models "github.com/DavidAlf/GambitUser/models"
	secretm "github.com/DavidAlf/GambitUser/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var DB *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DBConnect() error {
	DB, err = sql.Open("mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println("[DBConnect]>[ERROR] Error con el string de conexion a la bds " + err.Error())
		return err
	}

	err = DB.Ping()

	if err != nil {
		fmt.Println("[DBConnect]>[ERROR] Error con el ping de conexion a la bds " + err.Error())
		return err
	}

	fmt.Println("[DBConnect]>Se conecto a la BDs OKA")

	return err
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndPoint, dbName string

	dbUser = claves.UserName
	authToken = claves.Password
	dbEndPoint = claves.Host
	dbName = "gambit"

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndPoint, dbName)

	fmt.Println(dns)

	return dns
}
