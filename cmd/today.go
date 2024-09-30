/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"time"

	"github.com/MovieMaker93/note-cli/cmd/template/note"
	"github.com/MovieMaker93/note-cli/cmd/utils"
	"github.com/spf13/cobra"
)

type TodayNote struct {
	Exit        bool
	CurrentDate string
	DayBefore   string
	DayAfter    string
}

type TodayTemplater interface {
	Note() []byte
}

type Today struct {
	templater TodayTemplater
}

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Open today note",
	Long:  `Open today note in nvim`,
	Run: func(cmd *cobra.Command, args []string) {

		dir, err := utils.GoToVaultDirectory("Journal")
		if err != nil {
			fmt.Println("Vault Path is not properly set")
			cobra.CheckErr("Vault Path is not properly set")
			return
		}

		currentDate := time.Now()
		tomorrowDate := currentDate.AddDate(0, 0, 1).Format("02-01-2006")
		yesterdayDate := currentDate.AddDate(0, 0, -1).Format("02-01-2006")

		formattetDate := currentDate.Format("02-01-2006")
		todayInboxTemplate := Today{
			templater: note.TodayNoteTemplate{},
		}

		todayNote := TodayNote{
			CurrentDate: formattetDate,
			DayBefore:   yesterdayDate,
			DayAfter:    tomorrowDate,
		}

		filename := todayNote.CurrentDate + ".md"
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			file, err := os.Create(filename)
			defer file.Close()
			todayFileTemplate := template.Must(template.New(todayNote.CurrentDate + ".md").Parse(string(todayInboxTemplate.templater.Note())))
			err = todayFileTemplate.Execute(file, todayNote)
			if err != nil {
				// Handle the error
				fmt.Println("Error opening file:", err)
				cobra.CheckErr("Error opening file")
				return
			}
		}

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
