package main

type DynamicMonkey struct {
	cachedResult    int
	leftMonkey      Monkey
	hasCachedResult bool
	rightMonkey     Monkey
	operation       MonkeyOperation
}

func (dm *DynamicMonkey) GetResult() int {
	if !dm.hasCachedResult {
		dm.cachedResult = dm.operation(dm.leftMonkey.GetResult(), dm.rightMonkey.GetResult())
		dm.hasCachedResult = true
	}
	return dm.cachedResult
}

func NewDynamicMonkey(
	leftMonkey Monkey,
	rightMonkey Monkey,
	operation MonkeyOperation,
) DynamicMonkey {
	return DynamicMonkey{
		leftMonkey:  leftMonkey,
		rightMonkey: rightMonkey,
		operation:   operation,
	}
}
