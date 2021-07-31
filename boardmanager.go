package getmoe

import (
	"fmt"
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

func (m *BoardManager) executeCommand(b *Board, cmd DownloadConfiguration) error {
	posts, err := b.RequestAll(cmd.Tags)
	if err != nil {
		return err
	}

	// if err := os.MkdirAll(cmd.DestinationConfiguration.Directory, os.ModePerm); err != nil {
	// 	return err
	// }
	for _, p := range posts {
		fmt.Println("Saving", p.FileURL, p.Hash)
		// os.MkdirAll(saveDir, os.ModePerm)
		// if err := p.Save(saveDir); err != nil {
		// 	return err
		// }

		// if !quiet {
		// 	fName, _ := helper.FileURLUnescape(p.FileURL)
		// 	fmt.Println("Getting", fName[:32], "...")
		// }
	}
	return nil
}
