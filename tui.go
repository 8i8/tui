package tui

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

func NewScreen() error {
	if term.IsTerminal(0) || term.IsTerminal(1) {
		fmt.Println("yes")
	}
	oldState, err := term.MakeRaw(0)
	if err != nil {
		return err
	}
	defer term.Restore(0, oldState)
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	t := term.NewTerminal(screen, "")
	t.SetPrompt(string(t.Escape.Red)+"->"+string(t.Escape.Reset))

	rePrefix := string(t.Escape.Cyan) + "Human says:" + string(t.Escape.Reset)

	for {
		line, err := t.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if line == "" {
			continue
		}
		fmt.Fprintln(t, rePrefix, line)
	}
}
