package intcode

// import "log"

type point struct {
	X int
	Y int
}

func (p *point) setDir(i int) {
	x := p.X
	y := p.Y
	switch i {
	case 0:
		// left
		p.Y = x
		p.X = 0 - y
	case 1:
		// right
		p.Y = 0 - x
		p.X = y
	}
}

type Rover struct {
	value   int64
	outputs []int64
	paints  map[point]int
	pos     point
	dir     point
	mode    bool // false is color, true is rotate
	color   int
}

func NewRover(input int64) *Rover {
	return &Rover{value: input, pos: point{}, dir: point{Y: -1}, paints: map[point]int{}}
}

func (v Rover) ReadValue() (int64, error) {
	return v.value, nil
}

func (v *Rover) WriteValue(i int64) error {
	switch v.mode {
	case false:
		v.color = int(i)
	case true:
		v.dir.setDir(int(i))
		// paint then move
		point := point{X: v.pos.X, Y: v.pos.Y}
		v.paints[point] = v.color
		v.pos.X += v.dir.X
		v.pos.Y += v.dir.Y
		c, ok := v.paints[v.pos]
		if ok {
			v.value = int64(c)
		} else {
			v.value = 0
		}
	}
	v.mode = !v.mode
	// v.value = i
	v.outputs = append(v.outputs, i)
	return nil
}

func (v Rover) Outputs() []int64 {
	return v.outputs
}

func (v Rover) PaintCount() int {
	return len(v.paints)
}

func (v Rover) Paints() map[point]int {
	return v.paints
}

func (v Rover) Fail() {}
func (v Rover) Exit() {}
