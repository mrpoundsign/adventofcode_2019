package intcode

type Amplifier struct {
	initialInput []int
	inputCount   int
	value        int
}

func (a Amplifier) Phase() int {
	return a.initialInput[0]
}

func (a *Amplifier) ReadValue() int {
	if a.inputCount < len(a.initialInput) {
		defer func() { a.inputCount++ }()
		a.value = a.initialInput[a.inputCount]
	}
	return a.value
}

func (a *Amplifier) WriteValue(i int) error {
	a.value = i
	return nil
}

func (a *Amplifier) Value() int {
	return a.value
}

func (a *Amplifier) Fail() {
	a.End()
}

func (a *Amplifier) End() {}

func NewAmplifier(phase int, input int) Amplifier {
	return Amplifier{initialInput: []int{phase, input}}
}
