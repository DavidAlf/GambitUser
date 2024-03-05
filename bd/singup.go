package bd

import (
	"fmt"

	model "github.com/DavidAlf/GambitUser/models"
	tool "github.com/DavidAlf/GambitUser/tools"

	/*_ "github.com/go-sql-driver/mysql"*/
	_ "github.com/microsoft/go-mssqldb"
)

func SingUp(sing model.UserSingUp) error {

	fmt.Println("[SingUp]> Comienza Registro en BD")

	db, err := DBConnect()

	if err != nil {
		return err
	}

	defer db.Close()

	stringSQL := "INSERT INTO users(User_email, User_UUID, User_DateAdd) VALUES ('" + sing.UserEmail + "', '" + sing.UserUUID + "', convert(datetime,'" + tool.FechaSQLServer() + "'));"

	fmt.Println(stringSQL)

	_, err = db.Exec(stringSQL)

	if err != nil {
		fmt.Println("[SingUp]> [ERROR] Error insertando la persona " + err.Error())

		return err
	}

	fmt.Println("[SingUP]> Ejecucion Exitosa")

	return nil
}
