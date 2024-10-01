package note

import (
	_ "embed"
)

//go:embed  files/refine.tmpl
var refineNoteTemplate []byte

type RefineNoteTemplate struct{}

func (i RefineNoteTemplate) Note() []byte {
	return refineNoteTemplate
}
