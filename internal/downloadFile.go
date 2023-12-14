package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/go-github/v37/github"
)

func DownloadFile(service string, output string, bianVersion string, apiType string) error {

	fileExtension := ""
	repoPath := ""

	if bianVersion == BIAN_VERSION_12 {
		fileExtension = FILE_EXTENSION_YAML
		if apiType == SEMANTIC_API {
			repoPath = REPO_PATH_12_SEMANTIC
		} else {
			repoPath = REPO_PATH_12_ISO
		}
	} else {
		fileExtension = FILE_EXTENSION_JSON
		repoPath = REPO_PATH_9_1
	}

	if output == "" {
		if bianVersion == BIAN_VERSION_12 {
			output = service + fileExtension
		} else {
			output = service + fileExtension
		}
	}

	path := repoPath + service + fileExtension

	client := github.NewClient(nil)
	file, _, _, err := client.Repositories.GetContents(context.Background(), BIAN_ORG, BIAN_REPO, path, &github.RepositoryContentGetOptions{})

	fmt.Println("Downloading API spec...")
	if err != nil {
		fmt.Printf("Failed to connect to BIAN github repository: %s", err.Error())
		return err
	}

	fmt.Println("Decoding API spec...")
	dec, err := base64.StdEncoding.DecodeString(*file.Content)
	if err != nil {
		fmt.Printf("Failed to decode file: %s", err.Error())
		return err
	}

	fmt.Println("Creating output file...")
	f, err := os.Create(output)
	if err != nil {
		fmt.Printf("Failed to create file: %s", err.Error())
		return err
	}
	defer f.Close()

	fmt.Println("Writing output file...")
	if _, err := f.Write(dec); err != nil {
		fmt.Printf("Failed to write file: %s", err.Error())
		return err
	}
	if err := f.Sync(); err != nil {
		fmt.Printf("Failed to write file: %s", err.Error())
		return err
	}

	fmt.Println("Success!")

	return nil

}
