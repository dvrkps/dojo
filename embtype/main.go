package embtype

import "fmt"

type point struct {
	X int
	Y int
}

type namedPoint struct {
	point
	Name string
}

func newNamedPoint(name string, x, y int) namedPoint {
	return namedPoint{
		point: point{
			X: x,
			Y: y,
		},
		Name: name,
	}
}

func (p namedPoint) String() string {
	return fmt.Sprintf("%s x:%d y:%d", p.Name, p.X, p.Y)
}
