package puzzlecaptcha

import (
	"image"
	"image/draw"
	"math/rand"
)

type PuzzleCaptcha interface {
	Generate() (image.Image, image.Image)
	HoleRect() image.Rectangle
}

type puzzleCaptcha struct {
	origin         image.Image
	holerect       image.Rectangle
	holes, notches int
}

func NewPuzzleCaptcha(origin image.Image, holesize, maxholes int) PuzzleCaptcha {
	holeX, holeY := rand.Intn(origin.Bounds().Dx()-holesize), rand.Intn(origin.Bounds().Dy()-holesize)
	holes := randomInt(4, maxholes)
	notches := randomInt(4, 4)

	return &puzzleCaptcha{
		origin:   origin,
		holerect: image.Rect(holeX, holeY, holeX+holesize, holeY+holesize),
		holes:    holes,
		notches:  notches,
	}
}

func randomInt(number, max int) int {
	result := 0b0000
	for i, m := 0, 0; i < number; i++ {
		if m >= max {
			result = result << 1
			continue
		}
		v := rand.Intn(100) % 2
		if v == 1 {
			m++
		}
		result = result<<1 + v
	}
	if result == 0 {
		result = 0b0001
	}
	return result
}

func (c *puzzleCaptcha) Generate() (image.Image, image.Image) {
	return c.draw(BaseType), c.draw(HoleType).(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(c.holerect)
}

func (c *puzzleCaptcha) HoleRect() image.Rectangle {
	return c.holerect
}

func (c *puzzleCaptcha) draw(masktype PuzzleMaskType) image.Image {
	options := []Option{
		WithHoles(c.holes),
		WithNotches(c.notches),
		WithMaskType(masktype),
	}
	mask := NewPuzzleMask(c.origin.Bounds(), c.holerect, options...)

	rect := c.origin.Bounds()
	if masktype == HoleType {
		rect = mask.Bounds()
	}
	dst := image.NewRGBA(rect)
	draw.DrawMask(dst, dst.Bounds(), c.origin, image.Pt(0, 0), mask, mask.Bounds().Min, draw.Over)
	return dst
}
