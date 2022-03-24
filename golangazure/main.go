package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getKeyVaultClient() (client keyvault.BaseClient) {
	keyvaultClient := keyvault.New()
	clientCredentialConfig := auth.NewClientCredentialsConfig("39d4753c-c85e-4f38-bffa-2509f1dc30a8", "Rlf7Q~AK5kzVPmGIJBQp4DAC_YEYbn1ETqjZn", "e94d57ad-a0b7-44cf-883f-d7987a0bd244")

	// From SDK NewClientCredentialsConfig generates a object to azure control plane
	// (By default Resource is set to management.azure.net)
	// There below line was added to access the azure data plane
	// Which is required to access secrets in keyvault

	clientCredentialConfig.Resource = "https://vault.azure.net"
	authorizer, err := clientCredentialConfig.Authorizer()

	if err != nil {
		fmt.Printf("Error occured while creating azure KV authroizer %v ", err)

	}
	keyvaultClient.Authorizer = authorizer

	return keyvaultClient
}

func main() {

	keyvaultClient := getKeyVaultClient()
	vaultUri := fmt.Sprintf("https://%s.vault.azure.net", "testbenzkey")

	// GetSecretFromKeyvault
	GetSecret(vaultUri, keyvaultClient, "testnaja", "b975ed71815a4c03948d81b1a5e33dd6")
	// set secretVersion empty string ("") to receive the latest

	SetSecret(vaultUri, keyvaultClient, "testna", "wowsa")
	DeleteSecret(vaultUri, keyvaultClient, "testnaja3")

}

func GetSecret(vaultUri string, keyvaultClient keyvault.BaseClient, secretName string, version string) {

	res, err := keyvaultClient.GetSecret(context.Background(), vaultUri, secretName, version)

	if err != nil {
		fmt.Printf("Error occured Get Secret %s , %v", secretName, err)
	}

	fmt.Printf("Secret : %s , Value : %s", secretName, *res.Value)
}

func SetSecret(vaultUri string, keyvaultClient keyvault.BaseClient, secretName string, value string) {

	res, err := keyvaultClient.SetSecret(context.Background(), vaultUri, secretName, keyvault.SecretSetParameters{Value: &value})

	if err != nil {
		fmt.Printf("Error occured Set Secret %s , %v", secretName, err)
	}

	fmt.Printf("Added Secret : %s , Id : %s", secretName, *res.ID)
}

func DeleteSecret(vaultUri string, keyvaultClient keyvault.BaseClient, secretName string) {

	res, err := keyvaultClient.DeleteSecret(context.Background(), vaultUri, secretName)

	if err != nil {
		fmt.Printf("Error occured Delete Secret %s , %v", secretName, err)
	}

	fmt.Printf("Deleted Secret : %s , Id : %s", secretName, *res.ID)
}