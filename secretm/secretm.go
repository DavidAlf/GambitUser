package secretm

import (
	"fmt"

	//awsgo "github.com/DavidAlf/GambitUser/awsgo"
	model "github.com/DavidAlf/GambitUser/models"
	//aws "github.com/aws/aws-sdk-go-v2/aws"
	//sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (model.SecretRDSJson, error) {
	var datosSecret model.SecretRDSJson

	fmt.Println("[GetSecret]> Opteniendo Secreto [" + secretName + "]")
	/*
		svc := sm.NewFromConfig(awsgo.Cfg)

		llaveSecret, err := svc.GetSecretValue(awsgo.Ctx, &sm.GetSecretValueInput{
			SecretId: aws.String(secretName),
		})

		if err != nil {
			fmt.Println("[GetSecret]> [ERROR] Opteniendo datos de secret" + err.Error())

			return datosSecret, err
		}

		json.Unmarshal([]byte(*llaveSecret.SecretString), &datosSecret)

		fmt.Print("[GetSecret]> Lectura Secret OKA [" + secretName + "]")

	*/
	datosSecret = model.SecretRDSJson{
		UserName: "root",
		Password: "rot12345",
		Engine:   "",
		Host:     "gambit.cf80eicso0uv.us-east-1.rds.amazonaws.com",
		Puerto:   3306,
	}

	return datosSecret, nil

}
