package getmoe

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

type BoardManager struct {
	Boards     map[string]*Board
	boardNames []string
}

func LoadBoards(config map[string]BoardConfiguration) (BoardManager, error) {
	m := BoardManager{
		Boards: make(map[string]*Board),
	}

	for name, board := range config {
		b, err := NewBoard(name, board)
		if err != nil {
			return BoardManager{}, fmt.Errorf("getmoe: unable to create board '%s': %w", name, err)
		}
		m.Boards[name] = b
		m.boardNames = append(m.boardNames, name)
	}

	return m, nil
}

func (m *BoardManager) ExecuteCommands(cmds []DownloadConfiguration) error {
	for _, cmd := range cmds {
		// Execute command on all boards if there are no boards specified.
		if len(cmd.Boards) == 0 {
			cmd.Boards = m.boardNames
		}

		for _, name := range cmd.Boards {
			if board, ok := m.Boards[name]; ok {
				if err := m.executeCommand(board, cmd); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type templateData struct {
	BoardName  string
	PostID     int
	PostAuthor string
	FilePath   string
	FileExt    string
	FileHash   string
}

func (m *BoardManager) executeCommand(b *Board, cmd DownloadConfiguration) error {
	posts, err := b.RequestAll(cmd.Tags)
	if err != nil {
		return err
	}

	tmpl, err := template.New("savepath").Parse(cmd.SavePath)
	if err != nil {
		return err
	}

	for _, p := range posts {
		var fname bytes.Buffer
		ext := filepath.Ext(p.FileURL)
		bname := strings.TrimSuffix(filepath.Base(p.FileURL), ext)
		data := templateData{
			BoardName:  b.name,
			PostID:     p.ID,
			PostAuthor: p.Author,
			FileHash:   p.Hash,
			FilePath:   bname,
			FileExt:    ext,
		}
		if err := tmpl.Execute(&fname, data); err != nil {
			return err
		}
		fmt.Println("Saving to", fname.String())
	}
	return nil
}
