package frame

import (
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/mouse"
	"image"
	//		"golang.org/x/mobile/event/paint"
	"time"
)

func Region(a, b int64) int64 {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

func (f *Frame) Select(mp image.Point, ed screen.EventDeque, paintfn func()) {
	f.modified = false
	f.Drawsel(f.PtOfChar(f.P0), f.P0, f.P1, false)
	p1 := f.IndexOf(mp)
	p0 := p1
	pt0 := f.PtOfChar(p0)
	pt1 := f.PtOfChar(p1)
	f.Drawsel(pt0, p0, p1, true)

	clock60hz := time.NewTicker(time.Second / 60).C
	paintfn()

	reg := int64(0)
	for {
		sc := false
		if f.Scroll != nil {
			if mp.Y < f.r.Min.Y {
				f.Scroll(-(f.r.Min.Y - mp.Y) / (f.Dy() - 1))
				p0 = f.P1
				p1 = f.P0
				sc = true
			} else if mp.Y > f.r.Max.Y {
				f.Scroll((mp.Y - f.r.Max.Y) / (f.Dy() + 1))
				p0 = f.P0
				p1 = f.P1
				sc = true
			}
			if sc {
				if reg != Region(p1, p0) {
					p0, p1 = p1, p0
				}
				pt0 = f.PtOfChar(p0)
				pt1 = f.PtOfChar(p1)
				reg = Region(p1, p0)
			}
		}
		q := f.IndexOf(mp)
		if p1 != q {
			if reg != Region(q, p0) {
				if reg > 0 {
					f.Drawsel(pt0, p0, p1, false)
				} else if reg < 0 {
					f.Drawsel(pt1, p1, p0, false)
				}
				p1 = p0
				pt1 = pt0
				reg = Region(q, p0)
				if reg == 0 {
					f.Drawsel(pt0, p0, p1, true)
				}
			}
			qt := f.PtOfChar(q)
			if reg > 0 {
				if q > p1 {
					f.Drawsel(pt1, p1, q, true)
				} else if q < p1 {
					f.Drawsel(qt, q, p1, false)
				}
			} else if reg < 0 {
				if q > p1 {
					f.Drawsel(pt1, p1, q, false)
				} else {
					f.Drawsel(qt, q, p1, true)
				}
			}
			p1 = q
			pt1 = qt
		}
		f.modified = false
		if p0 < p1 {
			f.P0 = p0
			f.P1 = p1
		} else {
			f.P0 = p1
			f.P1 = p0
		}

		if sc {
			ed.Send(ScrollEvent{})
			f.Scroll(0)
		}

		switch e := ed.NextEvent().(type) {
		case ScrollEvent:
		case mouse.Event:
			if e.Button == 1 && e.Direction == 2 || e.Button == 2 || e.Button == 3 {
				ed.SendFirst(e)
				return
			}
			mp = image.Pt(int(e.X), int(e.Y))
		case interface{}:
			ed.SendFirst(e)
			return
		}
		select {
		case <-clock60hz:
			paintfn()
		default:
		}
	}
}

type ScrollEvent struct {
}

func (f *Frame) SelectPaint(p0, p1 image.Point, col image.Image) {
	if f.b == nil {
		panic("selectpaint: b == 0")
	}
	if f.r.Max.Y == p0.Y {
		return
	}
	h := f.Font.height
	q0, q1 := p0, p1
	q0.Y += h
	q1.Y += h
	n := (p1.Y - p0.Y) / h

	if n == 0 { // one line
		f.draw(f.b, image.Rectangle{p0, q1}, col, image.ZP)
	} else {
		if p0.X >= f.r.Max.X {
			p0.X = f.r.Max.X - 1
		}
		f.draw(f.b, image.Rect(p0.X, p0.Y, f.r.Max.X, q0.Y), col, image.ZP)
		if n > 1 {
			f.draw(f.b, image.Rect(f.r.Min.X, q0.Y, f.r.Max.X, p1.Y), col, image.ZP)
		}
		f.draw(f.b, image.Rect(f.r.Min.X, p1.Y, q1.X, q1.Y), col, image.ZP)
	}
}
