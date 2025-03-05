package server

import (
	"github.com/atotto/clipboard"
	"github.com/lemonade-command/lemonade/lemon"
)

type Clipboard struct {
	token     string
	allowRead bool // Allow reading from clipboard, else return the last write
	lastWrite string
}

func (c *Clipboard) Copy(text string, _ *struct{}) error {
	<-connCh
	// Logger instance needs to be passed here somehow?

	text, err := lemon.DecryptMessage(c.token, text)
	if err != nil {
		return err
	}
	c.lastWrite = text
	return clipboard.WriteAll(lemon.ConvertLineEnding(text, LineEndingOpt))
}

func (c *Clipboard) Paste(_ struct{}, resp *string) error {
	<-connCh
	if c.allowRead {
		text, err := clipboard.ReadAll()
		if err != nil {
			return err
		}
		text, err = lemon.EncryptMessage(c.token, text)
		if err != nil {
			return err
		}
		*resp = text
	} else {
		*resp = c.lastWrite
	}
	return nil
}
