package main

type DynamicMonkey struct {
	cachedResult    int
	id              string
	leftMonkey      Monkey
	hasCachedResult bool
	rightMonkey     Monkey
	operation       MonkeyOperation
}

func (dm *DynamicMonkey) ComputeUnknownValueOf(id string, otherSide int) int {
	if dm.leftMonkey.DependsOnMonkey(id) {
		rightValue := dm.rightMonkey.GetResult()
		return dm.leftMonkey.ComputeUnknownValueOf(id, ComputeInverseOperationWithLeftMissing(dm.operation, rightValue, otherSide))
	} else {
		leftValue := dm.leftMonkey.GetResult()
		return dm.rightMonkey.ComputeUnknownValueOf(id, ComputeInverseOperationWithRightMissing(dm.operation, leftValue, otherSide))
	}
}

func (dm *DynamicMonkey) DependsOnMonkey(id string) bool {
	return dm.id == id || dm.leftMonkey.DependsOnMonkey(id) || dm.rightMonkey.DependsOnMonkey(id)
}

func (dm *DynamicMonkey) GetResult() int {
	if !dm.hasCachedResult {
		dm.cachedResult = ComputeOperation(dm.operation, dm.leftMonkey.GetResult(), dm.rightMonkey.GetResult())
		dm.hasCachedResult = true
	}
	return dm.cachedResult
}

func NewDynamicMonkey(
	definition MonkeyDefinition,
	leftMonkey Monkey,
	rightMonkey Monkey,
) DynamicMonkey {
	return DynamicMonkey{
		id:          definition.id,
		leftMonkey:  leftMonkey,
		rightMonkey: rightMonkey,
		operation:   definition.operation,
	}
}
