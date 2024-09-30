package note

import (
	_ "embed"
)

//go:embed  files/inbox.tmpl
var inboxNoteTemplate []byte

type InboxNoteTemplate struct{}

func (i InboxNoteTemplate) Note() []byte {
	return inboxNoteTemplate
}
