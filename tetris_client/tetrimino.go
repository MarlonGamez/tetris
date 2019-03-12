package tetrimino

type Coord struct {
	X int
	Y int
}

func (c &Coord) Add(other Coord) {
	newX := c.X + other.X
	newY := c.Y + other.Y

	return Coord{
		X: newX,
		Y: newY
	}
}

type TPiece int
const {
	IVal TPiece = 0
	QVal TPiece = 1
	JVal TPiece = 2
	LVal TPiece = 3
	SVal TPiece = 4
	ZVal TPiece = 5
	TVal TPiece = 6
}

type Tetrimino struct {
	Piece tType
	Pos Coord
	Cubes []Coord
}

const (
	I Tetrimino = New(IVal, Coord{1, 0}, Coord{0, 0}, Coord{1, 0}, Coord{2, 0})
	Q Tetrimino = New(SQVal, Coord{0, 0}, Coord{1, 0}, Coord{0, 1}, Coord{1, 1})
	J Tetrimino = New(JVal, Coord{0, 0}, Coord{-1, 0}, Coord{-1, -1}, Coord{1, 0})
	L Tetrimino = New(LVal, Coord{0, 0}, Coord{-1, 0}, Coord{1, 1},   Coord{1, 0})
	S Tetrimino = New(SVal, Coord{0, 0}, Coord{-1, 0}, Coord{0, 1},   Coord{1, 1})
	Z Tetrimino = New(ZVal, Coord{0, 0}, Coord{1, 0},  Coord{0, 1},   Coord{-1, -1})
	T Tetrimino = New(TVal, Coord{0, 0}, Coord{-1, 0}, Coord{1, 0},   Coord{0, 1})
)

func New(pType TPiece, offset1, offset2, offset3, offset4 Coord) Tetrimino {
	p = Coord{5, 1}
	cubes := make([]Coord, 4)
	cubes[0] = p.Add(offset1)
	cubes[1] = p.Add(offset2)
	cubes[2] = p.Add(offset3)
	cubes[3] = p.Add(offset4)

	return Tetrimino{pType, p, cubes}
}

// Check if two tetriminos are of the same type
func (t Tetrimino) Equals(other Tetrimino) bool {
	return t.Piece == other.Piece
}

// Rotate a tetrimino's cubes around its position
func (t Tetrimino) Rotate(dir int) {
	if t.Equals(I) {
		return
	}

	for i := 0; i < 4; i++ {
		if dir == 1 {
			t.Cubes[i] = clockwise(t.Pos, t.Cubes[i])
		} else {
			t.Cubes[i] = counterclockwise(t.Pos, t.Cubes[i])
		}
	}
}

// Clockwise rotation
func rotateC(center, point Coord) {
	newX = center.X + (point.Y - center.Y)
	newY = center.Y - (point.X - center.X)

	return Coord{newX, newY}
}

// Counter clockwise rotation
func rotateCC(center, point Coord) {
	newX = center.X - (point.Y - center.Y)
	newY = center.Y + (point.X - center.X)

	return Coord{newX, newY}
}
