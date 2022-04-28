package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

// Structure for interacting with an Azure Blob Storage account
type AzureBlobStorage struct {
	AzureKey                        string `json:"AzureKey"`
	AzureBlobAccountName            string `json:"AzureBlobAccountName"`
	AzurePrimaryBlobServiceEndpoint string `json:"AzurePrimaryBlobServiceEndpoint"`
	AzureBlobContainer              string `json:"AzureBlobContainer"`
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

// Reads the configuration file for Azure Blob Storage and returns a structure to utilize it
func LoadBlobStorageConfiguration() (AzureBlobStorage, error) {
	var config AzureBlobStorage

	// try to load under main.go position
	configFile, err := os.Open("blob-configuration.json")

	if err != nil {
		// try to load under a child position
		configFile, err = os.Open("../blob-configuration.json")
	}

	if err != nil {
		return AzureBlobStorage{}, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}
