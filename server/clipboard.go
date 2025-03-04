package server

import (
	"github.com/atotto/clipboard"
	"github.com/lemonade-command/lemonade/lemon"
)

type Clipboard struct {
	token string
}

func (c *Clipboard) Copy(text string, _ *struct{}) error {
	<-connCh
	// Logger instance needs to be passed here somehow?

	text, err := lemon.DecryptMessage(c.token, text)
	if err != nil {
		return err
	}
	return clipboard.WriteAll(lemon.ConvertLineEnding(text, LineEndingOpt))
}

func (c *Clipboard) Paste(_ struct{}, resp *string) error {
	<-connCh
	t, err := clipboard.ReadAll()
	t, err = lemon.EncryptMessage(c.token, t)
	if err != nil {
		return err
	}
	*resp = t
	return err
}
