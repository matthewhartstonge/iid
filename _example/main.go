package main

import (
	"context"
	"fmt"
	"os"

	"github.com/matthewhartstonge/iid"
)

func main() {
	jsonCreds, err := os.ReadFile("path-to-firebase-service-account-json.credentials")
	if err != nil {
		panic(err)
	}

	client, err := iid.New("com.example.app", jsonCreds, iid.WithSandbox())
	if err != nil {
		panic(err)
	}

	res, err := client.BatchImport(context.Background(), []string{"MY_APNS_TOKEN"})
	if err != nil {
		panic(err)
	}

	for _, item := range res {
		fmt.Println(item)
	}
}
