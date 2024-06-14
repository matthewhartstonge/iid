## Google Instance ID (FCM Batch APNs Token Import)

Provides a client to bulk import existing iOS APNs tokens to Firebase Cloud
Messaging using Google's [Instance ID batch import endpoint.](https://developers.google.com/instance-id/reference/server#create_registration_tokens_for_apns_tokens) 

Note, this is a deprecated Google API, so there is no guarantee this will work in the future.

If I get a report it's no longer working, I'll archive the project. ✌️

## Installation

```shell
go get github.com/matthewhartstonge/iid
```

## Usage

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/matthewhartstonge/iid"
)

func main() {
	jsonCreds, err := os.ReadFile("path-to-firebase-service-account-json.credentials")
	if err != nil {
		panic(err)
	}
	
	client := iid.New("com.example.app", jsonCreds)
	res, err := client.BatchImport(ctx.Background(), []string{"MY_APNS_TOKEN"})
	if err != nil {
		panic(err)
	}

	for _, item := range res {
		fmt.Println(item)
	}
}
```