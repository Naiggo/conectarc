package secretsm

import (
	"encoding/json"
	"fmt"

	"github.com/Naiggo/conectarc/awsgo"
	"github.com/Naiggo/conectarc/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println("Asking for Secret " + nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)

	fmt.Println("Read Secret OK" + nameSecret)

	return secretData, nil
}
