package main

import (
	"github.com/hspearman/library-inventory/internal"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// TODO: Fix this
	// We delay until redis starts up (ugh)
	delay, err := strconv.Atoi(os.Getenv("DELAY"))
	if err != nil {
		delay = 0
	}

	time.Sleep(time.Duration(delay) * time.Second)

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
