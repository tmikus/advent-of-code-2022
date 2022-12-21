package main

type StaticMonkey struct {
	value int
}

func (sm *StaticMonkey) GetResult() int {
	return sm.value
}

func NewStaticMonkey(value int) StaticMonkey {
	return StaticMonkey{value: value}
}
