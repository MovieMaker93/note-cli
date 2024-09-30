package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/MovieMaker93/note-cli/cmd/template/note"
	"github.com/MovieMaker93/note-cli/cmd/ui/textinput"
	"github.com/MovieMaker93/note-cli/cmd/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

type Options struct {
	Title   *textinput.Output
	Content *textinput.Output
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

// A Project contains the data for the project folder
// being created, and methods that help with that process
type NoteInbox struct {
	Title       string
	Exit        bool
	Content     string
	CurrentDate string
}

type NoteTemplater interface {
	Note() []byte
}

type Note struct {
	templater NoteTemplater
}

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

// createCmd represents the create command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new Inbox Note",
	Long:  `Crate new note with a specific title in Inbox Folder`,
	Run: func(cmd *cobra.Command, args []string) {

		flagTitle := cmd.Flag("title").Value.String()
		flagContent := cmd.Flag("content").Value.String()

		dir, err := utils.GoToVaultDirectory("Inbox")
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
		}

		noteInbox := NoteInbox{
			Title:   flagTitle,
			Content: flagContent,
		}

		if noteInbox.Title == "" {
			tprogram := tea.NewProgram(
				textinput.InitialTextInputModel(options.Title, "What is the title?"),
			)
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Title contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			if options.Title.Output == "" {

			}
			noteInbox.ExitCLI(tprogram)
			noteInbox.Title = options.Title.Output
			err := cmd.Flag("title").Value.Set(noteInbox.Title)
			if err != nil {
				log.Fatal("failed to set the name flag value", err)
			}
		}

		if noteInbox.Content == "" {
			tprogram := tea.NewProgram(
				textinput.InitialTextInputModel(options.Content, "What is the content?"),
			)
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Content contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			noteInbox.ExitCLI(tprogram)
			noteInbox.Content = options.Content.Output
			err := cmd.Flag("content").Value.Set(noteInbox.Content)
			if err != nil {
				log.Fatal("failed to set the name flag value", err)
			}
		}

		formattetDate := currentDate.Format("02-01-2006")
		noteInboxTemplate := Note{
			templater: note.InboxNoteTemplate{},
		}
		noteInbox.CurrentDate = formattetDate

		title := strings.TrimSpace(noteInbox.Title)
		file, err := os.OpenFile(title+".md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		defer file.Close()

		inboxFileTemplate := template.Must(template.New(title + ".md").Parse(string(noteInboxTemplate.templater.Note())))
		err = inboxFileTemplate.Execute(file, noteInbox)

		// utils.OpenAndWriteToFile(title, contentOfFile)
		if err != nil {
			// Handle the error
			fmt.Println("Error opening file:", err)
			cobra.CheckErr("Error opening file")
			return
		}

		filePath := dir + "/" + title + ".md"
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
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("title", "t", "", "Title of the note to create")
	newCmd.Flags().StringP("content", "c", "", "Content of the note")
}

// ExitCLI checks if the Note has been created, and closes
// out of the CLI if it has
func (p *NoteInbox) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
