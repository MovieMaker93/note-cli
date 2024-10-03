package zettelkasten

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/MovieMaker93/note-cli/cmd/template/note"
	"github.com/MovieMaker93/note-cli/cmd/utils"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var (
	dailyPath string = os.Getenv("DAILY_PATH")
)

type Zettelkasten struct {
	Exit       bool
	ZettelNote *ZettelNote
	TodayNote  *TodayNote
}

type ZettelNote struct {
	Title       string
	Content     string
	CurrentDate string
	Type        string
}

type TodayNote struct {
	CurrentDate string
	DayBefore   string
	DayAfter    string
}

type ZettelNoteTemplater interface {
	Note() []byte
}

type Note struct {
	templater ZettelNoteTemplater
}

type TodayTemplater interface {
	Note() []byte
}

type Today struct {
	templater TodayTemplater
}

// ExitCLI checks if the Note has been created, and closes
// out of the CLI if it has
func (p *Zettelkasten) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}

func (p *Zettelkasten) CreateConsumeNote() error {
	consumeNoteTemplate := Note{
		templater: note.ConsumeNoteTemplate{},
	}

	// Create new note
	fileConsume, _ := os.Create(p.ZettelNote.Title + ".md")
	// Open today note
	fmt.Println(dailyPath)
	utils.GoToVaultDirectory(dailyPath)
	fileToday, err := os.OpenFile(p.ZettelNote.CurrentDate+".md", os.O_RDWR|os.O_APPEND, 0644)

	if err != nil {
		cobra.CheckErr("No today note found. Create it before!")
	}

	defer fileConsume.Close()
	defer fileToday.Close()

	fileToday.WriteString("[[" + p.ZettelNote.Title + "]]")

	consumeTempl := template.Must(
		template.New(p.ZettelNote.Title + ".md").
			Parse(string(consumeNoteTemplate.templater.Note())),
	)
	return consumeTempl.Execute(fileConsume, p.ZettelNote)
}

func (p *Zettelkasten) CreateRefineNote() error {
	refineNoteTemplate := Note{
		templater: note.RefineNoteTemplate{},
	}

	file, _ := os.Create(p.ZettelNote.Title + ".md")
	defer file.Close()

	refineTempl := template.Must(
		template.New(p.ZettelNote.Title + ".md").Parse(string(refineNoteTemplate.templater.Note())),
	)
	return refineTempl.Execute(file, p.ZettelNote)
}

func (p *Zettelkasten) CreateTodayNote() {

	todayInboxTemplate := Today{
		templater: note.TodayNoteTemplate{},
	}

	filename := p.TodayNote.CurrentDate + ".md"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		defer file.Close()
		todayFileTemplate := template.Must(
			template.New(p.TodayNote.CurrentDate + ".md").
				Parse(string(todayInboxTemplate.templater.Note())),
		)
		err = todayFileTemplate.Execute(file, p.TodayNote)
		if err != nil {
			// Handle the error
			fmt.Println("Error opening file:", err)
			cobra.CheckErr("Error opening file")
			return
		}
	}

}
