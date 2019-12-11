package intcode

type ValueHolder struct {
	value int
}

func (v ValueHolder) ReadInput() (int, error) {
	return v.value, nil
}

func (v *ValueHolder) WriteOutput(i int) {
	v.value = i
}

func (v ValueHolder) Log(...interface{}) {

}
