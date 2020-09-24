package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

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
		zipURL := *tag.ZipballURL
		fmt.Println("zip url:", zipURL)

		resp, err := http.Get(zipURL)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
		}
		defer resp.Body.Close()
		fmt.Println("status:", resp.Status)
		if resp.StatusCode != 200 {
			fmt.Printf("error: %s", err.Error())
		}

		out, err := os.Create("ghtest_repo.zip")
		if err != nil {
			fmt.Printf("error: %s", err.Error())
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
		}
	}
}
