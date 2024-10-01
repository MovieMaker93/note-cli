package note

import (
	_ "embed"
)

//go:embed files/consume.tmpl
var consumeNoteTemplate []byte

type ConsumeNoteTemplate struct{}

func (i ConsumeNoteTemplate) Note() []byte {
	return consumeNoteTemplate
}
