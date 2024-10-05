package gchess

import (
	"math"
)

/*
 * Rook
 */
type rook struct {
	white     bool
	initMoved bool
}

func (r *rook) canMove(board *board, start *spot, end *spot) bool {
	if start == end {
		return false
	} // same location (pointer comparison)

	if end.piece != nil && start.piece.isWhite() == end.piece.isWhite() {
		return false
	} // same side

	if math.Abs(float64(start.x-end.x))*math.Abs(float64(start.y-end.y)) != 0.0 {
		return false
	} // invalid move

	i := start.x
	j := start.y
	bex := end.x // bound of end x to check starting from i
	bey := end.y // bound of end y to check starting from j
	if i < bex {
		bex--
	} else if i > bex {
		bex++
	}
	if j < bey {
		bey--
	} else if j > bey {
		bey++
	}

	for i != bex || j != bey {
		if i < bex {
			i++
		} else if i > bex {
			i--
		}
		if j < bey {
			j++
		} else if j > bey {
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

func (r rook) isWhite() bool {
	return r.white
}

func (r rook) toUnicode() string {
	if r.white {
		return "♜"
	} else {
		return "♖"
	}
}
