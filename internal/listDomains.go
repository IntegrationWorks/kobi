package internal

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/v37/github"
)

func ListServices() {

	client := github.NewClient(nil)

	tree, _, err := client.Git.GetTree(
		context.Background(),
		BIAN_ORG,
		BIAN_REPO,
		BIAN_BRANCH,
		true,
	)

	if err != nil {
		log.Printf("Failed to connect to BIAN github repository: %s", err.Error())
		return
	}

	for _, t := range tree.Entries {
		if strings.HasPrefix(t.GetPath(), REPO_PATH) {
			domainName := t.GetPath()
			domainName = strings.Replace(domainName, REPO_PATH, "", 1)
			domainName = strings.Replace(domainName, FILE_EXTENSION, "", 1)
			if domainName != "Readme.md" {
				fmt.Println(domainName)
			}

		}
	}
}
