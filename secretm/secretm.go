package secretm

import (
	"encoding/json"
	"fmt"

	awsgo "github.com/DavidAlf/GambitUser/awsgo"
	model "github.com/DavidAlf/GambitUser/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (model.SecretRDSJson, error) {
	var datosSecret model.SecretRDSJson

	fmt.Println("> Opteniendo Secreto [" + secretName + "]")

	svc := sm.NewFromConfig(awsgo.Cfg)

	llaveSecret, err := svc.GetSecretValue(awsgo.Ctx, &sm.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println("> [ERROR] Opteniendo datos de secret" + err.Error())

		return datosSecret, err
	}

	json.Unmarshal([]byte(*llaveSecret.SecretString), &datosSecret)

	fmt.Print("> Lectura Secret OKA [" + secretName + "]")

	return datosSecret, nil

}
