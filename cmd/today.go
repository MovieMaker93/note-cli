/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/MovieMaker93/note-cli/cmd/utils"
	zet "github.com/MovieMaker93/note-cli/cmd/zettelkasten"
	"github.com/spf13/cobra"
)

const (
	TODAY_PATH string = "DAILY_PATH"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Open today note",
	Long:  `Open or create today note in nvim`,
	Run: func(cmd *cobra.Command, args []string) {

		todayPath := os.Getenv(TODAY_PATH)
		dir, err := utils.GoToVaultDirectory(todayPath)
		if err != nil {
			fmt.Println("Vault Path is not properly set")
			cobra.CheckErr("Vault Path is not properly set")
			os.Exit(1)
		}

		currentDate := time.Now()
		tomorrowDate := currentDate.AddDate(0, 0, 1).Format("02-01-2006")
		yesterdayDate := currentDate.AddDate(0, 0, -1).Format("02-01-2006")
		formattetDate := currentDate.Format("02-01-2006")

		todayNote := zet.TodayNote{
			CurrentDate: formattetDate,
			DayBefore:   yesterdayDate,
			DayAfter:    tomorrowDate,
		}

		zettelkasten := &zet.Zettelkasten{
			Exit:      false,
			TodayNote: &todayNote,
		}

		zettelkasten.CreateTodayNote()

		filePath := dir + "/" + formattetDate + ".md"

		// arguments := []string{"+normal ggzzi"}
		nvim := exec.Command("nvim", "+normal ggzzi", filePath)

		// Set the command's standard output and error to the current process' output
		nvim.Stdout = os.Stdout
		nvim.Stderr = os.Stderr
		error := nvim.Run()

		if error != nil {
			fmt.Println("Error executing command:", err)
			os.Exit(1) // Exit with a non-zero status indicating an error
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
