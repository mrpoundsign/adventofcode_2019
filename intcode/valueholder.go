package intcode

type ValueHolder struct {
	value int
}

func (v ValueHolder) ReadValue() int {
	return v.value
}

func (v *ValueHolder) WriteValue(i int) error {
	v.value = i
	return nil
}

func (v ValueHolder) Fail() {}
func (v ValueHolder) End()  {}
