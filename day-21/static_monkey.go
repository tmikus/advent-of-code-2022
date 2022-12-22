package main

type StaticMonkey struct {
	id    string
	value int
}

func (sm *StaticMonkey) ComputeUnknownValueOf(id string, otherSide int) int {
	if id == sm.id {
		return otherSide
	}
	return sm.value
}

func (sm *StaticMonkey) DependsOnMonkey(id string) bool {
	return sm.id == id
}

func (sm *StaticMonkey) GetResult() int {
	return sm.value
}

func NewStaticMonkey(definition MonkeyDefinition) StaticMonkey {
	return StaticMonkey{
		id:    definition.id,
		value: definition.value,
	}
}
