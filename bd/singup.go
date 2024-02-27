package bd

import (
	"fmt"

	model "github.com/DavidAlf/GambitUser/models"
	tool "github.com/DavidAlf/GambitUser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SingUp(sing model.UserSingUp) error {

	fmt.Println("[SingUp]> Comienza Registro en BD")

	err := DBConnect()

	if err != nil {
		return err
	}

	defer DB.Close()

	stringSQL := "INSERT INTO users(User_email, User_UUID, User_DateAdd) VALUES ('" + sing.UserEmail + "', '" + sing.UserUUID + "', '" + tool.FechaMySQL() + "');"

	fmt.Println(stringSQL)

	_, err = DB.Exec(stringSQL)

	defer fmt.Println("NO SE POR QUE FALLA " + err.Error())

	if err != nil {
		fmt.Println("[SingUp]> [ERROR] Error insertando la persona " + err.Error())

		return err
	}

	fmt.Println("[SingUP]> Ejecucion Exitosa" + err.Error())
	return err

}
