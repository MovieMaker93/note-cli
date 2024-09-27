/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Open today note",
	Long:  `Open today note in nvim`,
	Run: func(cmd *cobra.Command, args []string) {

		vaultPath := os.Getenv("VAULT")

		if vaultPath == "" {
			cobra.CheckErr("Vaulth Obsidian Path need to be set up")
		}

		// Go to Vault Directory
		dir := vaultPath + "Journal"
		err := os.Chdir(dir)
		if err != nil {
			fmt.Println("Vault Path is not properly set")
			cobra.CheckErr("Vault Path is not properly set")
			return
		}

		currentDate := time.Now()

		formattetDate := currentDate.Format("02-01-2006")

		file, err := os.OpenFile(formattetDate+".md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		defer file.Close()
		if err != nil {
			// Handle the error
			fmt.Println("Error opening file:", err)
			cobra.CheckErr("Error opening file")
			return
		}

		filePath := dir + "/" + formattetDate + ".md"
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
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
