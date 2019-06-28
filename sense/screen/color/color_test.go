package color

import (
	"image/color"
	"testing"
)

func TestNew(t *testing.T) {
	if got := New(255, 255, 255); got != 65535 {
		t.Logf("New got: %v, want: 65535", got)
		t.Fail()
	}
}

func TestPixel565_RGBA(t *testing.T) {
	t.Run("white", func(t *testing.T) {
		white := New(255, 255, 255)
		r, g, b, a := white.RGBA()
		var failed bool
		if r != 0xffff {
			t.Logf("r: %v", r)
			failed = true
		}
		if g != 0xffff {
			t.Logf("g: %v", g)
			failed = true
		}
		if b != 0xffff {
			t.Logf("b: %v", b)
			failed = true
		}
		if a != 0xffff {
			t.Logf("a: %v", a)
			failed = true
		}
		if failed {
			t.Fail()
		}
	})
	t.Run("red", func(t *testing.T) {
		red := New(255, 0, 0)
		r, g, b, a := red.RGBA()
		switch {
		case r != 0xffff:
			t.Logf("r: %v", r)
			t.Fail()
		case g != 0:
			t.Logf("g: %v", g)
			t.Fail()
		case b != 0:
			t.Logf("b: %v", b)
			t.Fail()
		case a != 0xffff:
			t.Logf("a: %v", a)
			t.Fail()
		}
	})
}

var rgb565PixelTests = []struct {
	rgb    color.RGBA
	rgb565 Pixel565
}{
	{rgb: color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, rgb565: 0xf800},
	{rgb: color.RGBA{R: 0x80, G: 0x00, B: 0x00, A: 0xff}, rgb565: 0x8000},
	{rgb: color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}, rgb565: 0x07e0},
	{rgb: color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0xff}, rgb565: 0x0400},
	{rgb: color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}, rgb565: 0x001f},
	{rgb: color.RGBA{R: 0x00, G: 0x00, B: 0x80, A: 0xff}, rgb565: 0x0010},
	{rgb: color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, rgb565: 0x0000},

	{rgb: color.RGBA{R: 0x05, G: 0x0a, B: 0x0b, A: 0xff}, rgb565: 0x0041},
	{rgb: color.RGBA{R: 0x0e, G: 0x21, B: 0x26, A: 0xff}, rgb565: 0x0904},
	{rgb: color.RGBA{R: 0x5a, G: 0xda, B: 0xff, A: 0xff}, rgb565: 0x5edf},
}

func TestRGB565Model(t *testing.T) {
	for _, test := range rgb565PixelTests {
		got := RGB565Model.Convert(test.rgb)
		want := test.rgb565
		if got != want {
			t.Errorf("unexpected RGB565 value for %+v: got: %016b, want: %016b", test.rgb, got, want)
		}
	}
}

func TestPixel565RGBA(t *testing.T) {
	for _, test := range rgb565PixelTests {
		got := color.RGBAModel.Convert(test.rgb565).(color.RGBA)
		got.R &= 0xf8
		got.G &= 0xfc
		got.B &= 0xf8
		want := test.rgb
		want.R &= 0xf8
		want.G &= 0xfc
		want.B &= 0xf8
		if got != want {
			t.Errorf("unexpected RGBA value for %016b: got: %+v, want: %+v", test.rgb565, got, want)
		}
	}
}
