package iid

const (
	maxTokens      = 100
	uriBatchImport = "https://iid.googleapis.com/iid/v1:batchImport"
)

var firebaseScopes = []string{
	"https://www.googleapis.com/auth/firebase",
	"https://www.googleapis.com/auth/firebase.messaging",
}
