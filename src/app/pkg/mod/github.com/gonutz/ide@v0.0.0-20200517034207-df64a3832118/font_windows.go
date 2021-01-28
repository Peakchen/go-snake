package main

import (
	"unicode"
	"unicode/utf8"

	"github.com/gonutz/binpacker"
	"github.com/gonutz/d3d9"
	"github.com/gonutz/truetype"
)

type d3d9Font struct {
	// info and scale contain the truetype font definition
	info  *truetype.FontInfo
	scale float64
	// grayTexture is a textureSize-by-textureSize gray image containing all
	// glyphs in one image, packed by the packer. It is stored as a flat array.
	// texture and uvScale are the corresponding D3D9 texture and a factor to
	// get the UV coordinates from the glyph data.
	// mustUploadTexture is true if grayTexture and texture are out of sync,
	// meaning the D3D9 texture must be re-uploaded to graphics memory.
	// mustResizeTexture is true if texture must also be resized.
	grayTexture       []byte
	textureSize       int
	packer            *binpacker.Packer
	texture           *d3d9.Texture
	uvScale           float32
	mustUploadTexture bool
	mustResizeTexture bool
	// ascend, descend and lineGap are the vertical metrics of the font
	// - ascend : height above the baseline to which the font can extend
	// - descend: height below the baseline to which the font can extend
	// - lineGap: additional space between two lines
	// ascend and lineGap are >= 0 and descend is <= 0
	//
	// here is an illustration with the letters: A p
	//
	// line -> -----------------------   ---
	// start        xx                    |
	//             x  x                   |
	//             x  x                   |
	//            x    x                  | ascend
	//            xxxxxx     xxxxx        |
	//           x      x    x    x       |
	// base      x      x    x    x       |
	// line -> -x--------x---xxxx-----   ---
	//                       x            |
	//                       x            | descend
	//                       x            |
	//         -----------------------   ---
	// next                               | line gap
	// line -> -----------------------   ---
	// start
	ascend, descend, lineGap int
	// glyphs grows depending on demanded glyphs, it is indexed with the values
	// from runeToGlyphIndex
	glyphs           []glyph
	runeToGlyphIndex map[rune]int
}

type glyph struct {
	// u0, u1, v0, v1 are the texture coordinates as in this illustration
	//
	//   u0     u1
	//   |       |
	//   v       v
	//   +-------+ <- v0
	//   |       |
	//   |       |
	//   |       |
	//   +-------+ <- v1
	u0, u1, v0, v1 float32
	// advance is the distance to the right that the cursor needs to be advanced
	// after this glyph
	advance int
	// xOffset and yOffset are the offset to render this glyph, relative to the
	// current cursor position
	xOffset, yOffset int
}

func newD3d9Font(ttf []byte, heightPix int, device *d3d9.Device) (*d3d9Font, error) {
	info, err := truetype.InitFont(ttf, 0)
	if err != nil {
		return nil, makeErr("unable to decode TTF font data", err)
	}
	scale := info.ScaleForPixelHeight(float64(heightPix))
	ascend, descend, lineGap := info.GetFontVMetrics()

	const textureSize = 128
	texture, err := device.CreateTexture(
		textureSize,
		textureSize,
		1,
		d3d9.USAGE_SOFTWAREPROCESSING,
		d3d9.FMT_A8,
		d3d9.POOL_MANAGED,
		0,
	)

	font := &d3d9Font{
		info:             info,
		scale:            scale,
		packer:           binpacker.New(textureSize, textureSize),
		grayTexture:      make([]byte, textureSize*textureSize),
		textureSize:      textureSize,
		texture:          texture,
		uvScale:          1 / float32(textureSize),
		ascend:           round(float64(ascend) * scale),
		descend:          round(float64(descend) * scale),
		lineGap:          round(float64(lineGap) * scale),
		runeToGlyphIndex: make(map[rune]int),
	}

	err = font.addNilGlyph()
	if err != nil {
		texture.Release()
		return nil, makeErr("add nil glyph", err)
	}

	return font, nil
}

func (f *d3d9Font) close() {
	f.texture.Release()
}

func (f *d3d9Font) addNilGlyph() error {
	const glyphIndex = 0
	pixels, width, height := f.info.GetGlyphBitmapSubpixel(
		0, f.scale, 0, 0, glyphIndex, 0, 0,
	)

	rect, err := f.packer.Insert(width, height)
	if err != nil {
		return err
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			f.grayTexture[rect.X+x+(rect.Y+y)*f.textureSize] = pixels[x+y*width]
		}
	}
	f.mustUploadTexture = true

	advance, _ := f.info.GetGlyphHMetrics(glyphIndex)
	x0, y0, _, _ := f.info.GetGlyphBitmapBox(glyphIndex, f.scale, f.scale)

	f.glyphs = append(f.glyphs, glyph{
		u0:      float32(rect.X) * f.uvScale,
		u1:      float32(rect.X+rect.Width) * f.uvScale,
		v0:      float32(rect.Y) * f.uvScale,
		v1:      float32(rect.Y+rect.Height) * f.uvScale,
		advance: round(float64(advance) * f.scale),
		xOffset: x0,
		yOffset: y0,
	})

	return nil
}

func (f *d3d9Font) recreateResourcesAfterDeviceReset(device *d3d9.Device) error {
	f.texture.Release()
	texture, err := device.CreateTexture(
		uint(f.textureSize),
		uint(f.textureSize),
		1,
		d3d9.USAGE_SOFTWAREPROCESSING,
		d3d9.FMT_A8,
		d3d9.POOL_MANAGED,
		0,
	)
	if err != nil {
		return makeErr("recreate font texture", err)
	}
	f.texture = texture
	f.mustUploadTexture = true
	return nil
}

func (f *d3d9Font) update(device *d3d9.Device) error {
	if f.mustResizeTexture {
		if err := f.recreateResourcesAfterDeviceReset(device); err != nil {
			return makeErr("resize font texture", err)
		}
	}

	if f.mustUploadTexture {
		mem, err := f.texture.LockRect(0, nil, d3d9.LOCK_DISCARD)
		if err != nil {
			return err
		}
		mem.SetAllBytes(f.grayTexture[:], f.textureSize)
		if err := f.texture.UnlockRect(0); err != nil {
			return err
		}
		f.mustUploadTexture = false
	}
	return nil
}

func (f *d3d9Font) resize(newSize int) {
	oldSize := f.textureSize

	if newSize <= oldSize {
		return
	}

	// copy over the old gray texture line by line
	oldGray := f.grayTexture
	newGray := make([]byte, newSize*newSize)
	for y := 0; y < oldSize; y++ {
		copy(
			newGray[y*newSize:],
			oldGray[y*oldSize:(y+1)*oldSize],
		)
	}

	// keep the texture coordinates right
	newUVScale := 1.0 / float32(newSize)
	scale := newUVScale / f.uvScale
	for i := range f.glyphs {
		f.glyphs[i].u0 *= scale
		f.glyphs[i].u1 *= scale
		f.glyphs[i].v0 *= scale
		f.glyphs[i].v1 *= scale
	}

	// update packer so that it knows the new size
	err := f.packer.Enlarge(newSize, newSize)
	if err != nil {
		// this is expected to never happen
		panic(err)
	}

	f.textureSize = newSize
	f.grayTexture = newGray
	f.uvScale = newUVScale
	f.mustResizeTexture = true
}

func (f *d3d9Font) xSpaceBetween(a, b rune) int {
	return round(
		f.scale * float64(f.info.GetCodepointKernAdvance(int(a), int(b))),
	)
}

func (f *d3d9Font) lineHeight() int {
	return f.ascend - f.descend + f.lineGap
}

func (f *d3d9Font) textHeight() int {
	return f.ascend - f.descend
}

// singleLineExtend is for text input that is known to be only one line, the
// calculation can then ignore all vertical offsets
func (f *d3d9Font) singleLineExtent(text []byte) (width, height int) {
	if len(text) == 0 {
		return
	}

	x := 0
	var last rune

	i := 0
	for i < len(text) {
		character, size := utf8.DecodeRune(text[i:])
		i += size

		if character == ' ' || character == '\t' {
			glyph := f.getGlyph(character)
			if character == '\t' {
				x += glyph.advance * 4
			} else {
				x += glyph.advance
			}
			last = 0
			continue
		}

		if unicode.IsControl(character) {
			last = 0
			continue
		}

		glyph := f.getGlyph(character)

		if last != 0 {
			x += f.xSpaceBetween(last, character)
		}

		x += glyph.advance
		last = character
	}

	width = x
	height = f.ascend - f.descend

	return
}

func (f *d3d9Font) extent(text []byte) (width, height int) {
	if len(text) == 0 {
		return
	}

	x := 0
	var last rune
	lineCount := 1

	i := 0
	for i < len(text) {
		character, size := utf8.DecodeRune(text[i:])
		i += size

		if character == '\n' {
			if x > width {
				width = x
			}
			lineCount++
			last = 0
			x = 0
			continue
		}

		if character == ' ' || character == '\t' {
			glyph := f.getGlyph(character)
			if character == '\t' {
				x += glyph.advance * 4
			} else {
				x += glyph.advance
			}
			last = 0
			continue
		}

		if unicode.IsControl(character) {
			last = 0
			continue
		}

		glyph := f.getGlyph(character)

		if last != 0 {
			x += f.xSpaceBetween(last, character)
		}

		x += glyph.advance
		last = character
	}
	// handle the last line
	if x > width {
		width = x
	}

	height = lineCount*f.lineHeight() - f.lineGap

	return
}

func (f *d3d9Font) getGlyph(r rune) *glyph {
	if i, ok := f.runeToGlyphIndex[r]; ok {
		return &f.glyphs[i]
	}

	// add the glyph

	glyphIndex := f.info.FindGlyphIndex(int(r))
	if glyphIndex == 0 {
		f.runeToGlyphIndex[r] = 0
		return &f.glyphs[0]
	}

	pixels, width, height := f.info.GetGlyphBitmapSubpixel(
		0, f.scale, 0, 0, glyphIndex, 0, 0,
	)

	rect, err := f.packer.Insert(width, height)
	for err == binpacker.ErrNoMoreSpace {
		f.resize(f.textureSize * 2)
		rect, err = f.packer.Insert(width, height)
	}
	if err != nil {
		// this is never expected to happen
		panic("cannot resize bin packer")
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			f.grayTexture[rect.X+x+(rect.Y+y)*f.textureSize] = pixels[x+y*width]
		}
	}
	f.mustUploadTexture = true

	advance, _ := f.info.GetGlyphHMetrics(glyphIndex)
	x0, y0, _, _ := f.info.GetGlyphBitmapBox(glyphIndex, f.scale, f.scale)

	f.glyphs = append(f.glyphs, glyph{
		u0:      float32(rect.X) * f.uvScale,
		u1:      float32(rect.X+rect.Width) * f.uvScale,
		v0:      float32(rect.Y) * f.uvScale,
		v1:      float32(rect.Y+rect.Height) * f.uvScale,
		advance: round(float64(advance) * f.scale),
		xOffset: x0,
		yOffset: y0,
	})
	index := len(f.glyphs) - 1
	f.runeToGlyphIndex[r] = index

	return &f.glyphs[index]
}
