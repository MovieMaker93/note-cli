package note

import (
	_ "embed"
)

//go:embed  files/today.tmpl
var todayNoteTemplate []byte

type TodayNoteTemplate struct{}

func (i TodayNoteTemplate) Note() []byte {
	return todayNoteTemplate
}
