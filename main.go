package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {
	storageAccountName := os.Getenv("PLUGIN_STORAGE_ACCOUNT")
	storageAccountKey := os.Getenv("PLUGIN_STORAGE_ACCOUNT_KEY")
	container := os.Getenv("PLUGIN_CONTAINER")
	source := os.Getenv("PLUGIN_SOURCE")
	target := os.Getenv("PLUGIN_TARGET")

	if target == "" {
		target = path.Base(source)
	}

	args := []string{
		"--upload",
		"--storageaccountkey",
		storageAccountKey,
		"--remoteresource",
		target,
		storageAccountName,
		container,
		source,
	}
	cmd := exec.Command("blobxfer", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}
