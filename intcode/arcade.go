package intcode

// import "log"

type gameObject int
type arMode int

const (
	goEmpty gameObject = iota
	goWall
	goBlock
	goHPaddle
	goBall
)

const (
	arModeX arMode = iota
	arModeY
	arModeObject
)

type Arcade struct {
	score     int64
	value     int64
	outputs   []int64
	objects   map[point]gameObject
	pos       point
	paddlePos point
	mode      int // false is color, true is rotate
}

func NewArcade(input int64) *Arcade {
	return &Arcade{value: input, pos: point{}, objects: map[point]gameObject{}}
}

func (v Arcade) ReadValue() (int64, error) {
	return v.value, nil
}

func (v *Arcade) WriteValue(i int64) error {
	switch v.mode {
	case 0:
		v.pos.X = int(i)
	case 1:
		v.pos.Y = int(i)
	case 2:
		if v.pos.X == -1 && v.pos.Y == 0 {
			v.score = i
			break
		}

		obj := gameObject(i)
		v.objects[v.pos] = obj

		switch obj {
		case goBall:
			if v.paddlePos.X < v.pos.X {
				v.value = 1
				break
			}
			if v.paddlePos.X > v.pos.X {
				v.value = -1
				break
			}
			v.value = 0
		case goHPaddle:
			v.paddlePos = v.pos
		}
	}
	v.mode = (v.mode + 1) % 3
	return nil
}

func (v Arcade) Outputs() []int64 {
	return v.outputs
}

func (v Arcade) ObjectCount() int {
	return len(v.objects)
}

func (v Arcade) BlockCount() int {
	num := 0
	for _, o := range v.objects {
		if o == goBlock {
			num++
		}
	}
	return num
}

func (v Arcade) Score() int64 {
	return v.score
}

func (v Arcade) Objects() map[point]gameObject {
	return v.objects
}

func (v Arcade) Fail() {}
func (v Arcade) Exit() {}
