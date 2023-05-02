package noter

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	notelyDirectory   = ".notely/"
	markdownExtension = ".md"
)

// Noter is an interface that helps do everything notely needs for note-taking
type Noter interface {
	// WriteNote writes a new note given the note name and contents.
	// Returns a non-nil error if it can't successfully write the note
	WriteNote(name string, contents string) error

	// LoadNote loads a note given the note name and returns its contents.
	// Returns a non-nil error if it can't successfully load the note.
	LoadNote(name string) (string, error)

	// ListNotes lists all the notes created with notely.
	// Returns a non-nil error if any are encountered.
	ListNotes() ([]string, error)

	// CheckNoteExists checks if a note already exists.
	// Returns true if the note exists and false if it doesn't.
	// Returns an error if one is encountered
	CheckNoteExists(name string) (bool, error)

	// DeleteNote attempts to delete a note. Returns an error if any encountered
	DeleteNote(name string) error
}

func NewNoter() Noter {
	return &noter{}
}

// noter implements Noter
type noter struct{}

func (n *noter) getNotelyPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting user home directory: %w", err)
	}

	return path.Join(home, notelyDirectory), nil
}

func (n *noter) buildNotePath(notelyPath, name string) string {
	return fmt.Sprintf("%s.md", path.Join(notelyPath, name))
}

func (n *noter) WriteNote(name string, contents string) error {
	// Check if the notely directory exists first
	notelyPath, err := n.getNotelyPath()
	if err != nil {
		return err
	}

	_, err = os.Stat(notelyPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		// create the notely directory if it doesn't exist
		os.MkdirAll(notelyPath, os.ModePerm)
	} else if err != nil {
		return fmt.Errorf("checking notely directory exists: %w", err)
	}

	// Write the note
	return os.WriteFile(n.buildNotePath(notelyPath, name), []byte(contents), os.ModePerm)
}

func (n *noter) LoadNote(name string) (string, error) {
	notelyPath, err := n.getNotelyPath()
	if err != nil {
		return "", err
	}

	notePath := n.buildNotePath(notelyPath, name)
	file, err := os.Open(notePath)
	if err != nil {
		return "", fmt.Errorf("opening file %q: %w", notePath, err)
	}

	contents, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("reading note %q: %w", notePath, err)
	}

	return string(contents), nil
}

func (n *noter) ListNotes() ([]string, error) {
	notes := []string{}

	notelyPath, err := n.getNotelyPath()
	if err != nil {
		return notes, err
	}

	err = fs.WalkDir(os.DirFS(notelyPath), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == markdownExtension {
			notes = append(notes, strings.Replace(filepath.Base(path), filepath.Ext(path), "", -1))
		}

		return nil
	})
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (n *noter) CheckNoteExists(name string) (bool, error) {
	notelyPath, err := n.getNotelyPath()
	if err != nil {
		return false, err
	}

	_, err = os.Stat(n.buildNotePath(notelyPath, name))
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (n *noter) DeleteNote(name string) error {
	notelyPath, err := n.getNotelyPath()
	if err != nil {
		return err
	}

	return os.Remove(n.buildNotePath(notelyPath, name))
}
