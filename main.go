package main

import (
	"context"
	"errors"
	"fmt"

	"os"

	awsgo "github.com/DavidAlf/GambitUser/awsgo"
	"github.com/DavidAlf/GambitUser/bd"
	model "github.com/DavidAlf/GambitUser/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	if !ValidaParametros() {
		fmt.Println("[EjecutoLambda]> [ERROR] Error en los parametros, no encuentra el 'SecretName'")

		err := errors.New("[EjecutoLambda]> [ERROR] Error en los parametros, no encuentra el 'SecretName'")

		return event, err
	}

	var datos model.UserSingUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = [" + datos.UserEmail + "]")

		case "sub":
			datos.UserUUID = att
			fmt.Println("UserId = [" + datos.UserUUID + "]")
		}
	}

	err := bd.ReadSecret()

	if err != nil {
		fmt.Println("[EjecutoLambda]> [ERROR] Error al leer el secret" + err.Error())

		return event, err
	}

	err = bd.SingUp(datos)

	return event, err

}

func ValidaParametros() bool {
	var traeParametro bool

	fmt.Println("[ValidaParametros] OKA")
	_, traeParametro = os.LookupEnv("SecretName")

	return traeParametro
}
