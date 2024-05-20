package puzzlecaptcha

import (
	"image"
	"image/draw"
	"math/rand"
)

type PuzzleCaptcha interface {
	Generate() (image.Image, image.Image)
}

type puzzleCaptcha struct {
	origin   image.Image
	holerect image.Rectangle
	holes    int
}

func NewPuzzleCaptcha(origin image.Image, holesize, maxholes int) PuzzleCaptcha {
	holeX, holeY := rand.Intn(origin.Bounds().Dx()-holesize), rand.Intn(origin.Bounds().Dy()-holesize)
	holes := 0b0000
	for i, max := 0, 1; i < 4 && max < maxholes; i++ {
		v := rand.Intn(100) % 2
		if v == 1 {
			max++
		}
		holes = holes<<1 + v
	}
	if holes == 0 {
		holes = 0b0001
	}

	return &puzzleCaptcha{
		origin:   origin,
		holerect: image.Rect(holeX, holeY, holeX+holesize, holeY+holesize),
		holes:    holes,
	}
}

func (c *puzzleCaptcha) Generate() (image.Image, image.Image) {
	return c.draw(BaseType), c.draw(HoleType)
}

func (c *puzzleCaptcha) draw(masktype PuzzleMaskType) image.Image {
	mask := NewPuzzleMask(c.origin.Bounds(), c.holerect, c.holes, masktype)
	dst := image.NewRGBA(c.origin.Bounds())
	draw.DrawMask(dst, dst.Bounds(), c.origin, image.Pt(0, 0), mask, mask.Bounds().Min, draw.Over)
	return dst
}
