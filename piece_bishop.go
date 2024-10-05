package gchess

import (
	"math"
)

/*
 * Bishop
 */
type bishop struct {
	white bool
}

func (b *bishop) canMove(board *board, start *spot, end *spot) bool {
	if start == end {
		return false
	} // same location (pointer comparison)

	if end.piece != nil && start.piece.isWhite() == end.piece.isWhite() {
		return false
	} // same side

	if math.Abs(float64(start.x-end.x)) != math.Abs(float64(start.y-end.y)) {
		return false
	} // invalid move

	i := start.x
	j := start.y
	bex := end.x // bound of end x to check starting from i
	bey := end.y // bound of end y to check starting from j
	if i < bex {
		bex--
	} else {
		bex++
	}
	if j < bey {
		bey--
	} else {
		bey++
	}

	for i != bex && j != bey {
		if i < bex {
			i++
		} else {
			i--
		}
		if j < bey {
			j++
		} else {
			j--
		}

		if i < 8 && j < 8 {
			if board.boxes[i][j].piece != nil {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func (b bishop) isWhite() bool {
	return b.white
}

func (b bishop) toUnicode() string {
	if b.white {
		return "♝"
	} else {
		return "♗"
	}
}
