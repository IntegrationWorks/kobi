package internal

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/go-github/v37/github"
)

func DownloadFile(service string, output string) error {

	if output == "" {
		output = service + ".json"
	}

	path := REPO_PATH + service + FILE_EXTENSION

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
