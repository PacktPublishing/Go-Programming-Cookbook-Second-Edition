package main

import (
	"context"
	"io"
	"os"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter7/oauthstore"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func main() {
	conf := oauthstore.Config{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT"),
			ClientSecret: os.Getenv("GITHUB_SECRET"),
			Scopes:       []string{"repo", "user"},
			Endpoint:     github.Endpoint,
		},
		Storage: &oauthstore.FileStorage{Path: "token.txt"},
	}
	ctx := context.Background()
	token, err := oauthstore.GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}

	cli := conf.Client(ctx, token)
	resp, err := cli.Get("https://api.github.com/user")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
