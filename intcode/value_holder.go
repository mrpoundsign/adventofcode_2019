package intcode

// import "log"

func NewValueHolder(input int64) *ValueHolder {
	return &ValueHolder{value: input}
}

type ValueHolder struct {
	value   int64
	outputs []int64
}

func (v ValueHolder) ReadValue() (int64, error) {
	return v.value, nil
}

func (v *ValueHolder) WriteValue(i int64) error {
	v.value = i
	v.outputs = append(v.outputs, i)
	return nil
}

func (v ValueHolder) Outputs() []int64 {
	return v.outputs
}

func (v ValueHolder) Fail() {}
func (v ValueHolder) Exit() {}
