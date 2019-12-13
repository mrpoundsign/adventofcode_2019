package intcode

type Amplifier struct {
	initialInput []int
	inputCount   int
	value        int
	send         chan int
	recv         chan int
	complete     bool
}

func (a Amplifier) Phase() int {
	return a.initialInput[0]
}

func (a *Amplifier) ReadValue() int {
	if a.inputCount < len(a.initialInput) {
		defer func() { a.inputCount++ }()
		a.value = a.initialInput[a.inputCount]
	} else {
		value, ok := <-a.recv
		if !ok {
			return a.value
		}
		a.value = value

	}

	return a.value
}

func (a *Amplifier) WriteValue(i int) error {
	a.value = i
	a.send <- a.value
	return nil
}

func (a *Amplifier) Value() int {
	return a.value
}

func (a Amplifier) Complete() bool {
	return a.complete
}

func (a *Amplifier) Fail() {
	a.Exit()
}

func (a *Amplifier) Exit() {
	close(a.send)
	a.complete = true
}

func (a *Amplifier) SetPrevAmp(input Amplifier) {
	a.recv = input.send
}

func NewAmplifier(phaseInput []int) *Amplifier {
	return &Amplifier{initialInput: phaseInput, send: make(chan int)}
}

func NewChainAmplifier(phase int, input *Amplifier) *Amplifier {
	return &Amplifier{initialInput: []int{phase}, send: make(chan int), recv: input.send}
}
