package utils

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func GetSecret(secretName string, projectID string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretName),
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", err
	}
	payload := result.Payload.Data
	return string(payload), nil
}
