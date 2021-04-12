package getmoe

import "fmt"

type BoardManager struct {
	Boards []Board
}

func LoadBoards(config map[string]BoardConfiguration) (BoardManager, error) {
	m := BoardManager{}

	for name, board := range config {
		b, err := NewBoard(name, board)
		if err != nil {
			return BoardManager{}, fmt.Errorf("getmoe: unable to create board '%s': %s", name, err)
		}
		m.Boards = append(m.Boards, b)
	}

	return m, nil
}
