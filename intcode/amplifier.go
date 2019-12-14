package intcode

import "errors"

type Amplifier struct {
	initialInput []int64
	inputCount   int
	value        int64
	send         chan int64
	recv         chan int64
	complete     bool
}

func (a Amplifier) Phase() int64 {
	return a.initialInput[0]
}

func (a *Amplifier) ReadValue() (int64, error) {
	if a.inputCount < len(a.initialInput) {
		defer func() { a.inputCount++ }()
		a.value = a.initialInput[a.inputCount]
	} else {
		value, ok := <-a.recv
		if !ok {
			return a.value, errors.New("upstream previous closed")
		}
		a.value = value

	}

	return a.value, nil
}

func (a *Amplifier) WriteValue(i int64) error {
	a.value = i
	a.send <- a.value
	return nil
}

func (a *Amplifier) Value() int64 {
	return a.value
}

func (a Amplifier) Complete() bool {
	return a.complete
}

func (a *Amplifier) Fail() {
	if !a.complete {
		a.Exit()
	}
}

func (a *Amplifier) Exit() {
	close(a.send)
	a.complete = true
}

func (a *Amplifier) SetPrevAmp(input Amplifier) {
	a.recv = input.send
}

func NewAmplifier(phaseInput []int64) *Amplifier {
	return &Amplifier{initialInput: phaseInput, send: make(chan int64, 1)}
}

func NewChainAmplifier(phase int64, input *Amplifier) *Amplifier {
	return &Amplifier{initialInput: []int64{phase}, send: make(chan int64, 1), recv: input.send}
}
