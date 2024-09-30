package utils

import (
	"os"
	"os/exec"

	_ "github.com/joho/godotenv/autoload"

	"github.com/spf13/cobra"
)

const (
	VAULT       string = "VAULT"
	ERROR_VAULT string = "Vault Obsidian Path env need to be set up"
)

func GoToVaultDirectory(relativePath string) (string, error) {

	vaultPath := os.Getenv(VAULT)

	if vaultPath == "" {
		cobra.CheckErr(ERROR_VAULT)
	}

	// Go to Vault Directory
	dir := vaultPath + relativePath
	if err := os.Chdir(dir); err != nil {
		return "", err
	}
	return dir, nil
}

func ExecuteCommand(command string, args string, filePath string) error {
	cmd := exec.Command(command, args, filePath)

	// Set the command's standard output and error to the current process' output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func OpenAndWriteToFile(fileName string, contentOfFile string) error {
	file, err := os.OpenFile(fileName+".md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	if _, error := file.WriteString(contentOfFile); error != nil {
		return error
	}
	return nil
}
