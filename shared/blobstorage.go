package shared

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

// Structure for interacting with an Azure Blob Storage account
type AzureBlobStorage struct {
	AzureKey                        string
	AzureBlobAccountName            string
	AzurePrimaryBlobServiceEndpoint string
	AzureBlobContainer              string
}

// Uploads bytes to blob storage with a given name and content type
func (a *AzureBlobStorage) UploadBytesToBlob(b []byte, n string, ct string) (string, error) {
	u, _ := url.Parse(fmt.Sprint(a.AzurePrimaryBlobServiceEndpoint, a.AzureBlobContainer, "/", n))
	credential, errC := azblob.NewSharedKeyCredential(a.AzureBlobAccountName, a.AzureKey)

	if errC != nil {
		return "", errC
	}

	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))
	ctx := context.Background()

	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: ct,
		},
	}

	_, errU := azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
	return blockBlobUrl.String(), errU
}

// Reads the environment variables for Azure Blob Storage and returns a structure to utilize it
func LoadBlobStorageConfiguration() (AzureBlobStorage, error) {
	// read environment variables
	azureKey := os.Getenv("AzureKey")
	azureBlobAccountName := os.Getenv("AzureBlobAccountName")
	azurePrimaryBlobServiceEndpoint := os.Getenv("AzurePrimaryBlobServiceEndpoint")
	azureBlobContainer := os.Getenv("AzureBlobContainer")

	// return configuration
	return AzureBlobStorage{
		AzureKey:                        azureKey,
		AzureBlobAccountName:            azureBlobAccountName,
		AzurePrimaryBlobServiceEndpoint: azurePrimaryBlobServiceEndpoint,
		AzureBlobContainer:              azureBlobContainer,
	}, nil
}
