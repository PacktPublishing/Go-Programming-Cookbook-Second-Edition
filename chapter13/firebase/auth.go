package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

// Authenticate grabs oauth scopes using a generated
// service_account.json file from
// https://console.firebase.google.com/project/go-cookbook/settings/serviceaccounts/adminsdk
func Authenticate(ctx context.Context, collection string) (Client, error) {

	opt := option.WithCredentialsFile("/tmp/service_account.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing app")
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to intialize filestore")
	}
	return &firebaseClient{Client: client, collection: collection}, nil
}
