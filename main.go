package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
)

func main() {
	client := github.NewClient(nil)
	ctx := context.Background()

	repoOpts := &github.RepositoryListOptions{}
	repos, _, err := client.Repositories.List(ctx, "shipyardapp", repoOpts)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	_ = repos

	tagOpts := &github.ListOptions{}
	tags, _, err := client.Repositories.ListTags(ctx, "shipyardapp", "ghtest", tagOpts)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	for _, tag := range tags {
		fmt.Printf("tag: %+v\n", tag)
	}
}
