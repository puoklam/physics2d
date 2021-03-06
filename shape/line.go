package shape

import "github.com/puoklam/physics2d/math/vector"

type Line struct {
	Start, End *vector.Vector2D
	Color      [3]int
}

func (l Line) Distance() float64 {
	d := vector.Sub(l.End, l.Start)
	return vector.Magnitude(d)
}

func (l Line) DistanceSquared() float64 {
	d := vector.Sub(l.End, l.Start)
	return vector.MagnitudeSquared(d)
}

func (l Line) Slope() float64 {
	dy := l.End.Y - l.Start.Y
	dx := l.End.X - l.Start.X
	return dy / dx
}

func NewLine(s, e *vector.Vector2D) *Line {
	return &Line{vector.Copy(s), vector.Copy(e), [3]int{0, 0, 0}}
}
