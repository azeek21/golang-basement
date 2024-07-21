package drawer

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	drawer_parser "github.com/azeek21/blog/apps/drawer/parser"
)

type Drawer struct {
	Colors *drawer_parser.ColorMap
	Src    *drawer_parser.ImageSource
	Dst    *image.RGBA
	Scale  int
}

func NewDrawer() Drawer {
	return Drawer{
		Colors: drawer_parser.NewColorMap(),
		Src:    drawer_parser.NewImageSource(),
		Dst:    nil,
		Scale:  1,
	}
}

func (d *Drawer) DrawPngFromFile(in, out string) error {

	if err := d.ParseFile(in); err != nil {
		return err
	}

	// Actual drawing
	d.Dst = image.NewRGBA(image.Rect(0, 0, d.Src.W*d.Scale, d.Src.H*d.Scale))

	for y := 0; y < d.Src.H; y++ {
		for x := 0; x < d.Src.W; x++ {
			c := (*d.Colors)[string(d.Src.Bytes[y][x])]
			draw.Draw(d.Dst, image.Rect(x*d.Scale, y*d.Scale, x*d.Scale+d.Scale, y*d.Scale+d.Scale), &image.Uniform{*c}, image.Point{}, draw.Over)
		}
	}

	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()
	png.Encode(outFile, d.Dst)
	return nil
}
