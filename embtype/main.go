package embtype

import "fmt"

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type namedPoint struct {
	point
	Name string `json:"name"`
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
