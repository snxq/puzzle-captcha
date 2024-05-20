package puzzlecaptcha

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_Mask(t *testing.T) {
	m := NewPuzzleMask(image.Rect(0, 0, 500, 500), image.Rect(100, 150, 150, 200), 0b1110, BaseType)
	f, err := os.Create("mask.png")
	if err != nil {
		t.Errorf("create file failed: %v", err)
		return
	}
	if err := png.Encode(f, m); err != nil {
		t.Errorf("encode png image failed: %v", err)
		return
	}
}
