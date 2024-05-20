package puzzlecaptcha

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_Captcha(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 500, 500))
	c := NewPuzzleCaptcha(img, 50, 2)
	base, hole := c.Generate()

	f, err := os.Create("base.png")
	if err != nil {
		t.Errorf("create file failed: %v", err)
		return
	}
	if err := png.Encode(f, base); err != nil {
		t.Errorf("encode png image failed: %v", err)
		return
	}

	f, err = os.Create("hole.png")
	if err != nil {
		t.Errorf("create file failed: %v", err)
		return
	}
	if err := png.Encode(f, hole); err != nil {
		t.Errorf("encode png image failed: %v", err)
		return
	}
}
