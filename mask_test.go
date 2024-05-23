package puzzlecaptcha

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_Mask(t *testing.T) {
	m := NewPuzzleMask(
		image.Rect(0, 0, 500, 500),
		image.Rect(100, 150, 150, 200),
		[]Option{
			WithMaskType(0b1110), WithMaskType(0b1010), WithMaskType(BaseType),
		}...) // 0b1110, 0b1010, BaseType

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
