package db

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebase() (*firebase.App, error) {
	conf := &firebase.Config{
		DatabaseURL: "https://moolair-blog-b027b-default-rtdb.firebaseio.com",
	}
	opt := option.WithCredentialsFile("firebase-key.json")
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
