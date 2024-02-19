package bd

import (
	"os"

	models "github.com/DavidAlf/GambitUser/models"
	secretm "github.com/DavidAlf/GambitUser/secretm"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
