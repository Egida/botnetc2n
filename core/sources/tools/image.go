package tools
// This is my watered down Pixterm.
// Still uses the same code, just some slight changes.
import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif" 

	_ "image/jpeg" 
	_ "image/png"  
	"strings"

	"github.com/disintegration/imaging"
	"github.com/lucasb-eyer/go-colorful"
	_ "golang.org/x/image/bmp"  
	_ "golang.org/x/image/tiff" 
	_ "golang.org/x/image/webp" 
)
const lowerHalfBlock = "\u2580"
const fullBlock = "\u2588"
const darkShadeBlock = "\u2593"
const mediumShadeBlock = "\u2592"
const lightShadeBlock = "\u2591"
const (
	ScaleModeResize = ScaleMode(iota)
	ScaleModeFill
	ScaleModeFit
)
const (
	NoDithering = DitheringMode(iota)
	DitheringWithBlocks
	DitheringWithChars
)
const (
	BlockSizeY = 8
	BlockSizeX = 4
)
var (
	ErrImageDownloadFailed = errors.New("image download failed")
	ErrHeightNonMoT = errors.New("height must be a Multiple of Two value")
	ErrInvalidBoundsMoT = errors.New("height or width must be >=2")
	ErrOutOfBounds = errors.New("out of bounds")
	errUnknownScaleMode = errors.New("unknown scale mode")
	errUnknownDitheringMode = errors.New("unknown dithering mode")
)
type ScaleMode uint8
type DitheringMode uint8
type ANSIpixel struct {
	Brightness uint8
	R, G, B    uint8
	upper      bool
	source     *ANSImage
}
type ANSImage struct {
	h, w      int
	maxprocs  int
	bgR       uint8
	bgG       uint8
	bgB       uint8
	dithering DitheringMode
	pixmap    [][]*ANSIpixel
}
func (ap *ANSIpixel) Render() string {
	return ap.RenderExt(false, false)
}
func (ap *ANSIpixel) RenderExt(renderGoCode, disableBgColor bool) string {
	backslash033 := "\033"
	if renderGoCode {
		backslash033 = "\\033"
	}
	if ap.source.dithering == NoDithering {
		var renderStr string
		if ap.upper {
			renderStr = fmt.Sprintf(
				"%s[38;2;%d;%d;%dm",
				backslash033,
				ap.R, ap.G, ap.B,
			)
		} else {
			renderStr = fmt.Sprintf(
				"%s[48;2;%d;%d;%dm%s",
				backslash033,
				ap.R, ap.G, ap.B,
				lowerHalfBlock,
			)
		}
		return renderStr
	}

	block := " "
	if ap.source.dithering == DitheringWithBlocks {
		switch bri := ap.Brightness; {
		case bri > 204:
			block = fullBlock
		case bri > 152:
			block = darkShadeBlock
		case bri > 100:
			block = mediumShadeBlock
		case bri > 48:
			block = lightShadeBlock
		}
	} else if ap.source.dithering == DitheringWithChars {
		switch bri := ap.Brightness; {
		case bri > 230:
			block = "#"
		case bri > 207:
			block = "&"
		case bri > 184:
			block = "$"
		case bri > 161:
			block = "X"
		case bri > 138:
			block = "x"
		case bri > 115:
			block = "="
		case bri > 92:
			block = "+"
		case bri > 69:
			block = ";"
		case bri > 46:
			block = ":"
		case bri > 23:
			block = "."
		}
	} else {
		panic(errUnknownDitheringMode)
	}

	bgColorStr := fmt.Sprintf(
		"%s[48;2;%d;%d;%dm",
		backslash033,
		ap.source.bgR, ap.source.bgG, ap.source.bgB,
	)
	if disableBgColor {
		bgColorStr = ""
	}
	return fmt.Sprintf(
		"%s%s[38;2;%d;%d;%dm%s",
		bgColorStr,
		backslash033,
		ap.R, ap.G, ap.B,
		block,
	)
}
func (ai *ANSImage) Height() int {
	return ai.h
}
func (ai *ANSImage) Width() int {
	return ai.w
}
func (ai *ANSImage) DitheringMode() DitheringMode {
	return ai.dithering
}
func (ai *ANSImage) SetMaxProcs(max int) {
	ai.maxprocs = max
}
func (ai *ANSImage) GetMaxProcs() int {
	return ai.maxprocs
}
func (ai *ANSImage) SetAt(y, x int, r, g, b, brightness uint8) error {
	if y >= 0 && y < ai.h && x >= 0 && x < ai.w {
		ai.pixmap[y][x].R = r
		ai.pixmap[y][x].G = g
		ai.pixmap[y][x].B = b
		ai.pixmap[y][x].Brightness = brightness
		ai.pixmap[y][x].upper = ((ai.dithering == NoDithering) && (y%2 == 0))
		return nil
	}
	return ErrOutOfBounds
}
func (ai *ANSImage) GetAt(y, x int) (*ANSIpixel, error) {
	if y >= 0 && y < ai.h && x >= 0 && x < ai.w {
		return &ANSIpixel{
				R:          ai.pixmap[y][x].R,
				G:          ai.pixmap[y][x].G,
				B:          ai.pixmap[y][x].B,
				Brightness: ai.pixmap[y][x].Brightness,
				upper:      ai.pixmap[y][x].upper,
				source:     ai.pixmap[y][x].source,
			},
			nil
	}
	return nil, ErrOutOfBounds
}
func (ai *ANSImage) Render() string {
	return ai.RenderExt(false, false)
}
func (ai *ANSImage) RenderExt(renderGoCode, disableBgColor bool) string {
	type renderData struct {
		row    int
		render string
	}

	backslashN := "\n"
	backslash033 := "\033"
	if renderGoCode {
		backslashN = "\\n"
		backslash033 = "\\033"
	}
	if ai.dithering == NoDithering {
		rows := make([]string, ai.h/2)
		for y := 0; y < ai.h; y += ai.maxprocs {
			ch := make(chan renderData, ai.maxprocs)
			for n, r := 0, y+1; (n <= ai.maxprocs) && (2*r+1 < ai.h); n, r = n+1, y+n+1 {
				go func(r, y int) {
					var str string
					for x := 0; x < ai.w; x++ {
						str += ai.pixmap[y][x].RenderExt(renderGoCode, disableBgColor)   
						str += ai.pixmap[y+1][x].RenderExt(renderGoCode, disableBgColor) 
					}
					str += fmt.Sprintf("%s[0m%s", backslash033, backslashN) 
					ch <- renderData{row: r, render: str}
				}(r, 2*r)
			}
			for n, r := 0, y+1; (n <= ai.maxprocs) && (2*r+1 < ai.h); n, r = n+1, y+n+1 {
				data := <-ch
				if renderGoCode {
					data.render = fmt.Sprintf(`fmt.Print("%s")%s`, data.render, "\n")
				}
				rows[data.row] = data.render
				
			}
		}
		return strings.Join(rows, "\r")
	}

	rows := make([]string, ai.h)
	for y := 0; y < ai.h; y += ai.maxprocs {
		ch := make(chan renderData, ai.maxprocs)
		for n, r := 0, y; (n <= ai.maxprocs) && (r+1 < ai.h); n, r = n+1, y+n+1 {
			go func(y int) {
				var str string
				for x := 0; x < ai.w; x++ {
					str += ai.pixmap[y][x].RenderExt(renderGoCode, disableBgColor)
				}
				str += fmt.Sprintf("%s[0m%s", backslash033, backslashN) 
				ch <- renderData{row: y, render: str}
			}(r)
		}
		for n, r := 0, y; (n <= ai.maxprocs) && (r+1 < ai.h); n, r = n+1, y+n+1 {
			data := <-ch
			if renderGoCode {
				data.render = fmt.Sprintf(`fmt.Print("%s")%s`, data.render, "\n")
			}
			rows[data.row] = data.render
		}
	}
	// If you wanted to be snazzy, you could do most of the main file shit in here.
	return strings.Join(rows, "\r")
}
func (ai *ANSImage) Draw() {
	ai.DrawExt(false, false)
}
func (ai *ANSImage) DrawExt(renderGoCode, disableBgColor bool) {
	fmt.Print(ai.RenderExt(renderGoCode, disableBgColor))
}
func New(h, w int, bg color.Color, dm DitheringMode) (*ANSImage, error) {
	if (dm == NoDithering) && (h%2 != 0) {
		return nil, ErrHeightNonMoT
	}

	if h < 2 || w < 2 {
		return nil, ErrInvalidBoundsMoT
	}

	r, g, b, _ := bg.RGBA()
	ansimage := &ANSImage{
		h: h, w: w,
		maxprocs:  1,
		bgR:       uint8(r),
		bgG:       uint8(g),
		bgB:       uint8(b),
		dithering: dm,
		pixmap:    nil,
	}

	ansimage.pixmap = func() [][]*ANSIpixel {
		v := make([][]*ANSIpixel, h)
		for y := 0; y < h; y++ {
			v[y] = make([]*ANSIpixel, w)
			for x := 0; x < w; x++ {
				v[y][x] = &ANSIpixel{
					R:          0,
					G:          0,
					B:          0,
					Brightness: 0,
					source:     ansimage,
					upper:      ((dm == NoDithering) && (y%2 == 0)),
				}
			}
		}
		return v
	}()

	return ansimage, nil
}
func NewScaledFromImage(image image.Image, y, x int, bg color.Color, sm ScaleMode, dm DitheringMode) (*ANSImage, error) {
	switch sm {
	case ScaleModeResize:
		image = imaging.Resize(image, x, y, imaging.Lanczos)
	case ScaleModeFill:
		image = imaging.Fill(image, x, y, imaging.Center, imaging.Lanczos)
	case ScaleModeFit:
		image = imaging.Fit(image, x, y, imaging.Lanczos)
	default:
		panic(errUnknownScaleMode)
	}

	return createANSImage(image, bg, dm)
}

func createANSImage(img image.Image, bg color.Color, dm DitheringMode) (*ANSImage, error) {
	var rgbaOut *image.RGBA
	bounds := img.Bounds()

	if _, _, _, a := bg.RGBA(); a >= 0xffff {
		rgbaOut = image.NewRGBA(bounds)
		draw.Draw(rgbaOut, bounds, image.NewUniform(bg), image.ZP, draw.Src)
		draw.Draw(rgbaOut, bounds, img, image.ZP, draw.Over)
	} else {
		if v, ok := img.(*image.RGBA); ok {
			rgbaOut = v
		} else {
			rgbaOut = image.NewRGBA(bounds)
			draw.Draw(rgbaOut, bounds, img, image.ZP, draw.Src)
		}
	}

	yMin, xMin := bounds.Min.Y, bounds.Min.X
	yMax, xMax := bounds.Max.Y, bounds.Max.X

	if dm == NoDithering {
		yMax = yMax - yMax%2 
	} else {
		yMax = yMax / BlockSizeY 
		xMax = xMax / BlockSizeX 
	}

	ansimage, err := New(yMax, xMax, bg, dm)
	if err != nil {
		return nil, err
	}

	if dm == NoDithering {
		for y := yMin; y < yMax; y++ {
			for x := xMin; x < xMax; x++ {
				v := rgbaOut.RGBAAt(x, y)
				if err := ansimage.SetAt(y, x, v.R, v.G, v.B, 0); err != nil {
					return nil, err
				}
			}
		}
	} else {
		pixelCount := BlockSizeY * BlockSizeX

		for y := yMin; y < yMax; y++ {
			for x := xMin; x < xMax; x++ {

				var sumR, sumG, sumB, sumBri float64
				for dy := 0; dy < BlockSizeY; dy++ {
					py := BlockSizeY*y + dy

					for dx := 0; dx < BlockSizeX; dx++ {
						px := BlockSizeX*x + dx

						pixel := rgbaOut.At(px, py)
						color, _ := colorful.MakeColor(pixel)
						_, _, v := color.Hsv()
						sumR += color.R
						sumG += color.G
						sumB += color.B
						sumBri += v
					}
				}

				r := uint8(sumR/float64(pixelCount)*255.0 + 0.5)
				g := uint8(sumG/float64(pixelCount)*255.0 + 0.5)
				b := uint8(sumB/float64(pixelCount)*255.0 + 0.5)
				brightness := uint8(sumBri/float64(pixelCount)*255.0 + 0.5)

				if err := ansimage.SetAt(y, x, r, g, b, brightness); err != nil {
					return nil, err
				}
			}
		}
	}

	return ansimage, nil
}