package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Inbox Note",
	Long:  `Crate new note with a specific title in Inbox Folder`,
	Run: func(cmd *cobra.Command, args []string) {
		flagTitle := cmd.Flag("title").Value.String()

		if flagTitle == "" {
			cobra.CheckErr("Title must not be empty. Fill in the Title!")
		}

		vaultPath := os.Getenv("VAULT")

		if vaultPath == "" {
			cobra.CheckErr("Vaulth Obsidian Path need to be set up")
		}

		// Go to Vault Directory
		dir := vaultPath + "Inbox"
		err := os.Chdir(dir)
		if err != nil {
			fmt.Println("Vault Path is not properly set")
			cobra.CheckErr("Vault Path is not properly set")
			return
		}

		currentDate := time.Now()

		formattetDate := currentDate.Format("02-01-2006")
		contentOfFile := fmt.Sprintf("#refine \n\n # Links \n\n [[%s]]", formattetDate)

		file, err := os.OpenFile(flagTitle+".md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		defer file.Close()
		if err != nil {
			// Handle the error
			fmt.Println("Error opening file:", err)
			cobra.CheckErr("Error opening file")
			return
		}

		file.WriteString(contentOfFile)
		fmt.Println("File created successfly!")

		filePath := dir + "/" + flagTitle + ".md"
		// path := strings.ReplaceAll(filePath, " ", "\\ ")
		nvim := exec.Command("nvim", "+normal ggzzi", filePath)

		// Set the command's standard output and error to the current process' output
		nvim.Stdout = os.Stdout
		nvim.Stderr = os.Stderr
		err = nvim.Run()

		if err != nil {
			fmt.Println("Error executing command:", err)
			os.Exit(1) // Exit with a non-zero status indicating an error
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("title", "t", "", "Title of the note to create")
}
