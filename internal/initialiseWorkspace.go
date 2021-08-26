package internal

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
)

func InitialiseWorkspace() error {

	if _, err := os.Stat("./portal"); !os.IsNotExist(err) {
		return fmt.Errorf("Portal directory already exists, cannot initialise kobi project here.")
	}

	fmt.Println("Cloning Kong Portal Template...")

	_, err := git.PlainClone("./portal", false, &git.CloneOptions{
		URL:      "https://github.com/Kong/kong-portal-templates",
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Printf("Failed to connect to Kong Portal github repository: %s", err.Error())
		return err
	}

	fmt.Println("Kong Portal Template cloned.")

	fmt.Println("Creating BIAN workspace...")

	err = os.RemoveAll("./portal/.git")
	err = os.RemoveAll("./portal/.github")
	if err != nil {
		return err
	}

	err = os.Rename("./portal/workspaces/default", "./portal/workspaces/bian")
	if err != nil {
		return err
	}

	fmt.Println("Done.")

	return nil

}
