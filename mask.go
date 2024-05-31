package puzzlecaptcha

import (
	"image"
	"image/color"
	"math"
)

type PuzzleMaskType int

const (
	HoleType PuzzleMaskType = iota
	BaseType
)

var (
	// Alpha represents the distance from the intersection points of the circle
	// with the inner rectangle to the outer rectangle.
	Alpha = 1.8
)

type PuzzleMask struct {
	whole, hole image.Rectangle
	masktype    PuzzleMaskType

	// binary number 0b0001
	// These four bits respectively represent up, down, left, and right.
	Holes int
	// binary number 0b0001
	// corresponds to the number of holes. 1 represents convex, 0 represents concave.
	notches int

	r, gap           float64
	alphadiff        uint8
	offset           float64
	cc               [4]image.Point
	offsetDirections [4][2]float64
}

func NewPuzzleMask(whole, hole image.Rectangle, options ...Option) *PuzzleMask {
	m := &PuzzleMask{whole: whole, hole: hole}
	for _, o := range options {
		o(m)
	}

	m.r = float64(min(m.hole.Max.X-m.hole.Min.X, m.hole.Max.Y-m.hole.Min.Y)) / (2*Alpha + 4)
	m.gap = m.r * Alpha

	// initial circle center
	m.cc = [4]image.Point{
		image.Pt((m.hole.Min.X+m.hole.Max.X)/2, m.hole.Max.Y-int(m.gap)),
		image.Pt((m.hole.Min.X+m.hole.Max.X)/2, m.hole.Min.Y+int(m.gap)),
		image.Pt(m.hole.Min.X+int(m.gap), (m.hole.Min.Y+m.hole.Max.Y)/2),
		image.Pt(m.hole.Max.X-int(m.gap), (m.hole.Min.Y+m.hole.Max.Y)/2),
	}
	m.offset = math.Abs(m.gap - m.r)
	m.offsetDirections = [4][2]float64{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

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
	for i, holes, notches := 0, m.Holes, m.notches; holes > 0; holes, i, notches = holes>>1, i+1, notches>>1 {
		if holes%2 == 0 {
			continue
		}
		var (
			cc               color.Color
			xoffset, yoffset = m.offset, m.offset
		)

		if notches%2 == 0 {
			xoffset *= m.offsetDirections[i][0]
			yoffset *= m.offsetDirections[i][1]
			cc = color.Alpha{255 - m.alphadiff}
		} else {
			xoffset *= -m.offsetDirections[i][0]
			yoffset *= -m.offsetDirections[i][1]
			cc = color.Alpha{0 + m.alphadiff}
		}

		xx, yy, rr := float64(x)-(float64(m.cc[i].X)+xoffset), float64(y)-(float64(m.cc[i].Y)+yoffset), float64(m.r)
		if xx*xx+yy*yy < rr*rr {
			return cc
		}
	}

	if x < m.hole.Min.X+int(m.gap) || x > m.hole.Max.X-int(m.gap) || y < m.hole.Min.Y+int(m.gap) || y > m.hole.Max.Y-int(m.gap) {
		return color.Alpha{255 - m.alphadiff}
	}

	return color.Alpha{0 + m.alphadiff}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
