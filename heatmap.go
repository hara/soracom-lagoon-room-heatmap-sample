package roomheatmap

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/draw"
	"image/png"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/math/fixed"
)

type Metric struct {
	Temperature float32
	Humidity    float32
}

type HeatMap struct {
	Room1 *Metric
}

//go:embed floormap.png
var floormap []byte

func (hm *HeatMap) Render() ([]byte, error) {
	bg, err := png.Decode(bytes.NewReader(floormap))
	if err != nil {
		return nil, err
	}

	dstimg := image.NewRGBA(bg.Bounds())
	draw.Draw(dstimg, dstimg.Bounds(), bg, image.Point{0, 0}, draw.Src)

	ft, err := truetype.Parse(gobold.TTF)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(ft, &truetype.Options{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	x, y := 320, 520
	dot := fixed.P(x, y)

	fd := font.Drawer{
		Dst:  dstimg,
		Src:  image.Black,
		Face: face,
		Dot:  dot,
	}

	text := fmt.Sprintf("%v°C, %v%%", hm.Room1.Temperature, hm.Room1.Humidity)
	fd.DrawString(text)

	outb := new(bytes.Buffer)
	if err := png.Encode(outb, dstimg); err != nil {
		return nil, err
	}

	return outb.Bytes(), nil
}
