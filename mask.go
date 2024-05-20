package puzzlecaptcha

import (
	"image"
	"image/color"

	"golang.org/x/exp/constraints"
)

type PuzzleMaskType int

const (
	HoleType PuzzleMaskType = iota
	BaseType
)

type PuzzleMask struct {
	whole, hole image.Rectangle
	masktype    PuzzleMaskType

	Holes int // binary number 0b0001 四位分别代表上下左右

	r         float64
	xs, ys    []int
	alphadiff uint8
}

func NewPuzzleMask(whole, hole image.Rectangle, holes int, masktype PuzzleMaskType) *PuzzleMask {
	m := &PuzzleMask{whole: whole, hole: hole, Holes: holes, masktype: masktype}
	m.r = float64(min(m.hole.Max.X-m.hole.Min.X, m.hole.Max.Y-m.hole.Min.Y)) / 4
	m.xs = []int{m.hole.Max.X, m.hole.Min.X, (m.hole.Max.X + m.hole.Min.X) / 2, (m.hole.Max.X + m.hole.Min.X) / 2}
	m.ys = []int{(m.hole.Max.Y + m.hole.Min.Y) / 2, (m.hole.Max.Y + m.hole.Min.Y) / 2, m.hole.Min.Y, m.hole.Max.Y}

	switch m.masktype {
	case HoleType:
		m.alphadiff = 255
	case BaseType:
		fallthrough
	default:
		m.alphadiff = 0
	}
	return m
}

func (m *PuzzleMask) ColorModel() color.Model {
	return color.AlphaModel
}

func (m *PuzzleMask) Bounds() image.Rectangle {
	return image.Rectangle{Min: m.whole.Min, Max: m.whole.Max}
}

func (m *PuzzleMask) At(x, y int) color.Color {
	if x < m.hole.Min.X || x > m.hole.Max.X || y < m.hole.Min.Y || y > m.hole.Max.Y {
		return color.Alpha{255 - m.alphadiff}
	}
	for i, holes := 0, m.Holes; holes > 0; holes, i = holes>>1, i+1 {
		if holes%2 == 0 {
			continue
		}

		xx, yy, rr := float64(x-m.xs[i]), float64(y-m.ys[i]), float64(m.r)
		if xx*xx+yy*yy < rr*rr {
			return color.Alpha{255 - m.alphadiff}
		}
	}

	return color.Alpha{0 + m.alphadiff}
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
