package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/MovieMaker93/note-cli/cmd/ui/textinput"
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
type Note struct {
	Title   string
	Exit    bool
	Content string
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

		fmt.Printf("%s\n", logoStyle.Render(logo))
		currentDate := time.Now()

		options := Options{
			Title:   &textinput.Output{},
			Content: &textinput.Output{},
		}

		note := Note{
			Title:   flagTitle,
			Content: flagContent,
		}

		if note.Title == "" {
			tprogram := tea.NewProgram(
				textinput.InitialTextInputModel(options.Title, "What is the title?"),
			)
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Title contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			if options.Title.Output == "" {

			}
			note.ExitCLI(tprogram)
			note.Title = options.Title.Output
			err := cmd.Flag("title").Value.Set(note.Title)
			if err != nil {
				log.Fatal("failed to set the name flag value", err)
			}
		}

		if note.Content == "" {
			tprogram := tea.NewProgram(
				textinput.InitialTextInputModel(options.Content, "What is the content?"),
			)
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Content contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			note.ExitCLI(tprogram)
			note.Content = options.Content.Output
			err := cmd.Flag("content").Value.Set(note.Content)
			if err != nil {
				log.Fatal("failed to set the name flag value", err)
			}
		}

		formattetDate := currentDate.Format("02-01-2006")
		contentOfFile := fmt.Sprintf(
			"#refine \n%s\n # Links \n\n [[%s]]",
			note.Content,
			formattetDate,
		)

		file, err := os.OpenFile(note.Title+".md", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		defer file.Close()
		if err != nil {
			// Handle the error
			fmt.Println("Error opening file:", err)
			cobra.CheckErr("Error opening file")
			return
		}

		file.WriteString(contentOfFile)
		fmt.Println("File created successfly!")

		filePath := dir + "/" + note.Title + ".md"
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

// ExitCLI checks if the Project has been exited, and closes
// out of the CLI if it has
func (p *Note) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}
