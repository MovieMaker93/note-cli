package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/MovieMaker93/note-cli/cmd/ui/textinput"
	"github.com/MovieMaker93/note-cli/cmd/utils"
	zet "github.com/MovieMaker93/note-cli/cmd/zettelkasten"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

type Options struct {
	Title   *textinput.Output
	Content *textinput.Output
	Type    *textinput.Output
}

const logo = `
	 _        _______ _________ _______    _______  _       _________
( (    /|(  ___  )\__   __/(  ____ \  (  ____ \( \      \__   __/
|  \  ( || (   ) |   ) (   | (    \/  | (    \/| (         ) (   
|   \ | || |   | |   | |   | (__      | |      | |         | |   
| (\ \) || |   | |   | |   |  __)     | |      | |         | |   
| | \   || |   | |   | |   | (        | |      | |         | |   
| )  \  || (___) |   | |   | (____/\  | (____/\| (____/\___) (___
|/    )_)(_______)   )_(   (_______/  (_______/(_______/\_______/

`

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

const (
	INBOX_PATH string = "NEW_NOTE_PATH"
)

// createCmd represents the create command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new Inbox Note",
	Long:  `Crate new note with a specific title, content, type in Inbox Folder`,
	Run: func(cmd *cobra.Command, args []string) {

		flagTitle := cmd.Flag("title").Value.String()
		flagContent := cmd.Flag("content").Value.String()
		flagType := cmd.Flag("type").Value.String()

		inboxPath := os.Getenv(INBOX_PATH)
		dir, err := utils.GoToVaultDirectory(inboxPath)
		if err != nil {
			fmt.Println("Vault Path is not properly set")
			cobra.CheckErr("Vault Path is not properly set")
			return
		}

		fmt.Printf("%s\n", logoStyle.Render(logo))
		currentDate := time.Now()

		options := Options{
			Title:   &textinput.Output{},
			Content: &textinput.Output{},
			Type:    &textinput.Output{},
		}

		zettelNote := zet.ZettelNote{
			Title:   flagTitle,
			Content: flagContent,
			Type:    flagType,
		}

		zettelkasten := &zet.Zettelkasten{
			Exit:       false,
			ZettelNote: &zettelNote,
		}

		if zettelNote.Title == "" {
			tprogram := tea.NewProgram(
				textinput.InitialTextInputModel(options.Title, "What is the title?", zettelkasten),
			)
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Title contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			zettelkasten.ExitCLI(tprogram)
			zettelNote.Title = options.Title.Output
			err := cmd.Flag("title").Value.Set(zettelNote.Title)
			if err != nil {
				log.Fatal("failed to set the name flag value", err)
			}
		}

		if zettelNote.Content == "" {
			tprogram := tea.NewProgram(
				textinput.InitialTextInputModel(options.Content, "What is the content?", zettelkasten),
			)
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Content contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			zettelkasten.ExitCLI(tprogram)
			zettelNote.Content = options.Content.Output
			err := cmd.Flag("content").Value.Set(zettelNote.Content)
			if err != nil {
				log.Fatal("failed to set the content flag value", err)
			}
		}

		if zettelNote.Type == "" {
			tprogram := tea.NewProgram(
				textinput.InitialTextInputModel(options.Type, "Allowed values: consume or refine", zettelkasten),
			)
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Type contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			zettelkasten.ExitCLI(tprogram)
			zettelNote.Type = options.Type.Output
			err := cmd.Flag("type").Value.Set(zettelNote.Type)
			if err != nil {
				log.Fatal("failed to set the type flag value", err)
			}

		}

		formattetDate := currentDate.Format("02-01-2006")
		zettelNote.CurrentDate = formattetDate
		title := strings.TrimSpace(zettelNote.Title)
		zettelkasten.ZettelNote = &zettelNote

		switch zettelNote.Type {
		case "refine":
			zettelkasten.CreateRefineNote()
		case "consume":
			zettelkasten.CreateConsumeNote()
		default:
			fmt.Println("No correct values selected for type. Allowed values: refine or consume")
			return
		}

		if err != nil {
			// Handle the error
			cobra.CheckErr("Error opening file")
			os.Exit(1)
		}

		filePath := dir + "/" + title + ".md"
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
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("title", "t", "", "Title of the note to create")
	newCmd.Flags().StringP("content", "c", "", "Content of the note")
	newCmd.Flags().String("type", "", "Kind of note to create. Allowed values: consume or refine")
}
