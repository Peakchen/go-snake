package main

import (
	"errors"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/gonutz/ide/w32"

	"github.com/gonutz/d3d9"
)

type d3d9Graphics struct {
	window            uintptr
	font              *d3d9Font
	d3d               *d3d9.Direct3D
	device            *d3d9.Device
	presentParameters d3d9.PRESENT_PARAMETERS
	deviceIsLost      bool
	vertexData        []float32
	jobs              []renderJob
}

type renderJob struct {
	kind  renderJobKind
	count uint
}

type renderJobKind int

const (
	triangles renderJobKind = iota
	textTriangles
)

func newD3d9Graphics(window uintptr, ttfFontData []byte, fontHeightPix int) (*d3d9Graphics, error) {
	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	if err != nil {
		return nil, makeErr("d3d9.Create", err)
	}

	var createFlags uint32 = d3d9.CREATE_SOFTWARE_VERTEXPROCESSING
	caps, err := d3d.GetDeviceCaps(d3d9.ADAPTER_DEFAULT, d3d9.DEVTYPE_HAL)
	if err == nil && caps.DevCaps&d3d9.DEVCAPS_HWTRANSFORMANDLIGHT != 0 {
		createFlags = d3d9.CREATE_HARDWARE_VERTEXPROCESSING
	}

	// get the maximum resolution of all monitors and use that as the initial
	// back buffer size
	var backBufW, backBufH uint32 = 512, 512
	n := d3d.GetAdapterCount()
	for i := uint(0); i < n; i++ {
		mode, err := d3d.GetAdapterDisplayMode(i)
		if err == nil {
			if mode.Width > backBufW {
				backBufW = mode.Width
			}
			if mode.Height > backBufH {
				backBufH = mode.Height
			}
		}
	}

	pp := d3d9.PRESENT_PARAMETERS{
		Windowed:         1,
		HDeviceWindow:    d3d9.HWND(window),
		SwapEffect:       d3d9.SWAPEFFECT_DISCARD,
		BackBufferWidth:  backBufW,
		BackBufferHeight: backBufH,
		BackBufferFormat: d3d9.FMT_A8R8G8B8,
		BackBufferCount:  1,
	}
	device, actualPP, err := d3d.CreateDevice(
		d3d9.ADAPTER_DEFAULT,
		d3d9.DEVTYPE_HAL,
		d3d9.HWND(window),
		createFlags,
		pp,
	)
	if err != nil {
		return nil, makeErr("d3d9.CreateDevice", err)
	}
	pp = actualPP

	if err := setRenderState(device); err != nil {
		return nil, err
	}

	font, err := newD3d9Font(ttfFontData, fontHeightPix, device)
	if err != nil {
		return nil, makeErr("create font", err)
	}

	g := &d3d9Graphics{
		window:            window,
		font:              font,
		d3d:               d3d,
		device:            device,
		presentParameters: pp,
		deviceIsLost:      false,
	}
	return g, nil
}

type recordFirstError struct {
	err error
}

func (e *recordFirstError) add(err error) {
	if e.err == nil && err != nil {
		e.err = err
	}
}

func setRenderState(device *d3d9.Device) error {
	var e recordFirstError

	e.add(device.SetRenderState(d3d9.RS_CULLMODE, d3d9.CULL_CCW))
	e.add(device.SetRenderState(d3d9.RS_ALPHATESTENABLE, 0))
	e.add(device.SetRenderState(d3d9.RS_ALPHABLENDENABLE, 1))
	e.add(device.SetRenderState(d3d9.RS_SRCBLEND, d3d9.BLEND_SRCALPHA))
	e.add(device.SetRenderState(d3d9.RS_DESTBLEND, d3d9.BLEND_INVSRCALPHA))

	// The following texture stage state is used for font rendering. The font
	// texture is alpha-only. The diffuse color of a vertex is used when
	// rendering the font glyphs, but the diffuse color's alpha channel is
	// multiplied with the font texture's alpha.
	e.add(device.SetTextureStageState(0, d3d9.TSS_COLOROP, d3d9.TOP_SELECTARG1))
	e.add(device.SetTextureStageState(0, d3d9.TSS_COLORARG1, d3d9.TA_CURRENT))
	e.add(device.SetTextureStageState(0, d3d9.TSS_ALPHAOP, d3d9.TOP_MODULATE))
	e.add(device.SetTextureStageState(0, d3d9.TSS_ALPHAARG1, d3d9.TA_CURRENT))
	e.add(device.SetTextureStageState(0, d3d9.TSS_ALPHAARG2, d3d9.TA_TEXTURE))
	e.add(device.SetTextureStageState(1, d3d9.TSS_COLOROP, d3d9.TOP_MODULATE))
	e.add(device.SetTextureStageState(1, d3d9.TSS_COLORARG1, d3d9.TA_CURRENT))
	e.add(device.SetTextureStageState(1, d3d9.TSS_COLORARG2, d3d9.TA_TEXTURE))

	if e.err != nil {
		return makeErr("error setting render state", e.err)
	}
	return nil
}

func (g *d3d9Graphics) close() {
	g.font.close()
	g.device.Release()
	g.d3d.Release()
}

func (g *d3d9Graphics) rect(x, y, w, h int, argb uint32) {
	fx, fy := float32(x), float32(y)
	fx2, fy2 := float32(x+w), float32(y+h)

	var col float32 = *(*float32)(unsafe.Pointer(&argb))

	// add two triangles for the rectangle
	g.vertexData = append(g.vertexData,
		fx, fy, 0, 1, col, 0, 0,
		fx2, fy, 0, 1, col, 0, 0,
		fx, fy2, 0, 1, col, 0, 0,

		fx, fy2, 0, 1, col, 0, 0,
		fx2, fy, 0, 1, col, 0, 0,
		fx2, fy2, 0, 1, col, 0, 0,
	)
	g.addJob(triangles, 2)
}

func (g *d3d9Graphics) addJob(kind renderJobKind, count uint) {
	n := len(g.jobs)
	if n == 0 || g.jobs[n-1].kind != kind {
		g.jobs = append(g.jobs, renderJob{kind: kind, count: count})
	} else {
		// combine this job with the last job in the queue, they have the same
		// kind
		g.jobs[n-1].count += count
	}
}

func (g *d3d9Graphics) text(text []byte, textX, textY int, clip rectangle, argb8 uint32) {
	if len(text) == 0 {
		return
	}

	x, y := textX, textY
	right, bottom := clip.x+clip.w, clip.y+clip.h
	var col float32 = *(*float32)(unsafe.Pointer(&argb8))
	var last rune
	var glyphCount uint

	i := 0

	// first skip all lines that are not visible
	lineHeight := g.font.lineHeight()
	maxInvisibleY := clip.y - lineHeight
	for i < len(text) && y < maxInvisibleY {
		for i < len(text) {
			character, size := utf8.DecodeRune(text[i:])
			i += size
			if character == '\n' {
				y += lineHeight
				break
			}
		}
	}

	for i < len(text) {
		character, size := utf8.DecodeRune(text[i:])
		i += size

		if character == '\n' {
			x = textX
			y += lineHeight
			last = 0
			continue
		}

		if character == ' ' || character == '\t' {
			glyph := g.font.getGlyph(character)
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

		glyph := g.font.getGlyph(character)
		w := round(float64(glyph.u1-glyph.u0) * float64(g.font.textureSize))
		h := round(float64(glyph.v1-glyph.v0) * float64(g.font.textureSize))

		if last != 0 {
			x += g.font.xSpaceBetween(last, character)
		}
		x += glyph.xOffset
		y := y + g.font.ascend + glyph.yOffset

		if x+w >= clip.x && x < right && y+h >= clip.y && y < bottom {
			// clip partially visible glyphs
			x, y := x, y
			u0, u1, v0, v1 := glyph.u0, glyph.u1, glyph.v0, glyph.v1
			if x+w > right {
				xFraction := float32(right-x) / float32(w)
				w = right - x
				u1 = u0 + (u1-u0)*xFraction
			}
			if y+h > bottom {
				yFraction := float32(bottom-y) / float32(h)
				h = bottom - y
				v1 = v0 + (v1-v0)*yFraction
			}
			if x < clip.x {
				xFraction := float32(x+w-clip.x) / float32(w)
				w = x + w - clip.x
				x = clip.x
				u0 = u1 - (u1-u0)*xFraction
			}
			if y < clip.y {
				yFraction := float32(y+h-clip.y) / float32(h)
				h = y + h - clip.y
				y = clip.y
				v0 = v1 - (v1-v0)*yFraction
			}

			// correct x,y by 0.5 so texels align with pixels, see
			// https://msdn.microsoft.com/en-us/library/windows/desktop/bb219690(v=vs.85).aspx
			x0 := float32(x) - 0.5
			y0 := float32(y) - 0.5
			x1 := x0 + float32(w)
			y1 := y0 + float32(h)
			g.vertexData = append(
				g.vertexData,
				x0, y0, 0, 1, col, u0, v0,
				x1, y1, 0, 1, col, u1, v1,
				x0, y1, 0, 1, col, u0, v1,

				x0, y0, 0, 1, col, u0, v0,
				x1, y0, 0, 1, col, u1, v0,
				x1, y1, 0, 1, col, u1, v1,
			)
			glyphCount++
		}

		if y-g.font.ascend > bottom {
			// if we are already below the given screen rectangle we can stop
			break
		}

		if x > right {
			// we are right of the given screen rectangle so skip the rest of
			// the line
			for i < len(text) && character != '\n' {
				character, size = utf8.DecodeRune(text[i:])
				i += size
			}
			if i >= len(text) {
				break
			}
			i -= size // give the line break back for the next processing step
		}

		x += glyph.advance - glyph.xOffset
		last = character
	}

	if glyphCount > 0 {
		g.addJob(textTriangles, glyphCount*2)
	}

	return
}

func (g *d3d9Graphics) present() error {
	const (
		vertexFmt       = d3d9.FVF_XYZRHW | d3d9.FVF_DIFFUSE | d3d9.FVF_TEX1
		floatsPerVertex = 7
	)

	if g.deviceIsLost {
		pp, err := g.device.Reset(g.presentParameters)
		if err != nil {
			// the device is not yet ready for rendering again, this error is
			// not fatal, it may take some frames until the device is ready
			// again
			return nil
		} else {
			g.presentParameters = pp
			g.deviceIsLost = false
			if err := setRenderState(g.device); err != nil {
				return err
			}
			if err := g.font.recreateResourcesAfterDeviceReset(g.device); err != nil {
				return err
			}
		}
	}

	r, ok := w32.GetClientRect(g.window)
	if !ok {
		return errors.New("unable to query window size")
	}
	windowW := uint32(r.Right - r.Left)
	windowH := uint32(r.Bottom - r.Top)

	if windowW > g.presentParameters.BackBufferWidth ||
		windowH > g.presentParameters.BackBufferHeight {
		// if the window is larger than the back buffer, increase the back
		// buffer size; note that it is never shrunk
		if windowW > g.presentParameters.BackBufferWidth {
			g.presentParameters.BackBufferWidth = windowW
		}
		if windowH > g.presentParameters.BackBufferHeight {
			g.presentParameters.BackBufferHeight = windowH
		}
		pp, err := g.device.Reset(g.presentParameters)
		if err != nil {
			return err
		} else {
			g.presentParameters = pp
			if err := setRenderState(g.device); err != nil {
				return err
			}
			if err := g.font.recreateResourcesAfterDeviceReset(g.device); err != nil {
				return err
			}
		}
	}

	if err := g.device.SetViewport(
		d3d9.VIEWPORT{0, 0, uint32(windowW), uint32(windowH), 0, 1},
	); err != nil {
		return err
	}

	if err := g.device.SetFVF(vertexFmt); err != nil {
		return err
	}

	if err := g.font.update(g.device); err != nil {
		return err
	}

	// render the scene
	if err := g.device.BeginScene(); err != nil {
		return err
	}

	jobs := g.jobs
	vertexData := g.vertexData
	g.vertexData = g.vertexData[0:0]
	g.jobs = g.jobs[0:0]

	for _, job := range jobs {
		if job.kind == triangles {
			if err := g.device.SetTexture(0, nil); err != nil {
				return makeErr("error resetting texture", err)
			}
			if err := g.device.DrawPrimitiveUP(
				d3d9.PT_TRIANGLELIST,
				job.count,
				uintptr(unsafe.Pointer(&vertexData[0])),
				floatsPerVertex*4,
			); err != nil {
				return makeErr("error drawing rectangles", err)
			}
		} else if job.kind == textTriangles {
			if err := g.device.SetTexture(0, g.font.texture); err != nil {
				return makeErr("error setting font texture", err)
			}
			if err := g.device.DrawPrimitiveUP(
				d3d9.PT_TRIANGLELIST,
				job.count,
				uintptr(unsafe.Pointer(&vertexData[0])),
				floatsPerVertex*4,
			); err != nil {
				return makeErr("error drawing text", err)
			}
		} else {
			// this is only expected to happen during development, if we add a
			// new job kind but forget to handle it here
			panic("unexpected job kind")
		}
		// for now it is all triangles, so 3 vertices per job item
		vertexData = vertexData[job.count*3*floatsPerVertex:]
	}

	if err := g.device.EndScene(); err != nil {
		return err
	}

	presentErr := g.device.Present(
		&d3d9.RECT{0, 0, int32(windowW), int32(windowH)},
		nil, 0, nil,
	)
	if presentErr != nil {
		if presentErr.Code() == d3d9.ERR_DEVICELOST {
			g.deviceIsLost = true
			return nil
		} else {
			return presentErr
		}
	}

	return nil
}
