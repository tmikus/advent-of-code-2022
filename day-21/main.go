package main

import (
	"bufio"
	"os"
)

func buildMonkeyDefinitionsMap(definitions []MonkeyDefinition) map[string]MonkeyDefinition {
	result := make(map[string]MonkeyDefinition)
	for _, definition := range definitions {
		result[definition.id] = definition
	}
	return result
}

func compileMonkey(
	definitionMap *map[string]MonkeyDefinition,
	staticMonkeys *[]StaticMonkey,
	dynamicMonkeys *[]DynamicMonkey,
	monkeysMap *map[string]Monkey,
	definition MonkeyDefinition,
) Monkey {
	if definition.isStatic {
		return compileStaticMonkey(
			staticMonkeys,
			monkeysMap,
			definition,
		)
	}
	return compileDynamicMonkey(
		definitionMap,
		staticMonkeys,
		dynamicMonkeys,
		monkeysMap,
		definition,
	)
}

func compileDynamicMonkey(
	definitionMap *map[string]MonkeyDefinition,
	staticMonkeys *[]StaticMonkey,
	dynamicMonkeys *[]DynamicMonkey,
	monkeysMap *map[string]Monkey,
	definition MonkeyDefinition,
) Monkey {
	leftMonkey := getOrCompileMonkey(
		definitionMap,
		staticMonkeys,
		dynamicMonkeys,
		monkeysMap,
		definition.leftMonkeyId,
	)
	rightMonkey := getOrCompileMonkey(
		definitionMap,
		staticMonkeys,
		dynamicMonkeys,
		monkeysMap,
		definition.rightMonkeyId,
	)
	dm := NewDynamicMonkey(leftMonkey, rightMonkey, definition.operation)
	(*dynamicMonkeys)[definition.index] = dm
	(*monkeysMap)[definition.id] = &(*dynamicMonkeys)[definition.index]
	return &(*dynamicMonkeys)[definition.index]
}

func compileStaticMonkey(
	staticMonkeys *[]StaticMonkey,
	monkeysMap *map[string]Monkey,
	definition MonkeyDefinition,
) Monkey {
	sm := NewStaticMonkey(definition.value)
	(*staticMonkeys)[definition.index] = sm
	(*monkeysMap)[definition.id] = &(*staticMonkeys)[definition.index]
	return &(*staticMonkeys)[definition.index]
}

func getOrCompileMonkey(
	definitionMap *map[string]MonkeyDefinition,
	staticMonkeys *[]StaticMonkey,
	dynamicMonkeys *[]DynamicMonkey,
	monkeysMap *map[string]Monkey,
	id string,
) Monkey {
	foundMonkey, monkeyExists := (*monkeysMap)[id]
	if monkeyExists {
		return foundMonkey
	}
	definition, ok := (*definitionMap)[id]
	if !ok {
		panic("Monkey not found!")
	}
	return compileMonkey(definitionMap, staticMonkeys, dynamicMonkeys, monkeysMap, definition)
}

func readMonkeyDefinitions() []MonkeyDefinition {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]MonkeyDefinition, 0)
	for scanner.Scan() {
		definition := ParseMonkeyDefinition(scanner.Text())
		definition.index = len(lines)
		lines = append(lines, definition)
	}
	return lines
}

func main() {
	definitions := readMonkeyDefinitions()
	definitionsMap := buildMonkeyDefinitionsMap(definitions)
	staticMonkeys := make([]StaticMonkey, len(definitions))
	dynamicMonkeys := make([]DynamicMonkey, len(definitions))
	monkeysMap := make(map[string]Monkey)
	rootMonkey := getOrCompileMonkey(&definitionsMap, &staticMonkeys, &dynamicMonkeys, &monkeysMap, "root")
	println("Part 1 result:", rootMonkey.GetResult())
}
