package main

import (
	"github.com/hspearman/library-inventory/internal"
	"log"
)

func main() {
	storageClient, err := internal.NewStorageClient()
	if err != nil {
		log.Fatalf("Failed to dial redis: %s", err.Error())
	}

	defer storageClient.Close()

	internal.SetDefaultInventory(storageClient)

	api := internal.NewAPI(storageClient)
	api.Configure()

	api.Echo.Logger.Fatal(api.Echo.Start(":1323"))
}
