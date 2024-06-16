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
```

## Firebase Credentials

You will require using Firebase service account credentials which can be 
downloaded from the Google Cloud Console as a JSON encoded file.

You will need to:

- Choose your Firebase Project from the Project dropdown selector (top left).
- Select and click on `IAM & Admin` > `Service accounts` from the side menu.
- Click on your `Firebase Admin SDK Service Agent` account within the presented table.
- Click on the 'keys' tab.
- Click 'Add Key' > 'Create New Key'.
- Choose key type `JSON`, and click 'create'.
- Download and use the JSON file as your service account credentials.
