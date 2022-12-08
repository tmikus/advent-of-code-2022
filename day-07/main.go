package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const COMMAND_PREFIX = "$ "

type Directory struct {
	directories map[string]Directory
	files       map[string]File
	name        string
	parent      *Directory
}

func newDirectory(name string) Directory {
	return Directory{
		directories: make(map[string]Directory),
		files:       make(map[string]File),
		name:        name,
		parent:      nil,
	}
}

func (directory *Directory) addDirectory(name string) *Directory {
	childDirectory := newDirectory(name)
	childDirectory.parent = directory
	directory.directories[name] = childDirectory
	return &childDirectory
}

func (directory *Directory) addFile(name string, size int) *File {
	childFile := newFile(name, size)
	directory.files[name] = childFile
	return &childFile
}

func (directory *Directory) computeSize() int {
	return directory.computeDirectoriesSize() + directory.computeFilesSize()
}

func (directory *Directory) computeDirectoriesSize() int {
	size := 0
	for _, dir := range directory.directories {
		size += dir.computeSize()
	}
	return size
}

func (directory *Directory) computeFilesSize() int {
	size := 0
	for _, file := range directory.files {
		size += file.size
	}
	return size
}

type File struct {
	name string
	size int
}

func newFile(name string, size int) File {
	return File{
		name: name,
		size: size,
	}
}

type State struct {
	allDirectories      []*Directory
	commands            []string
	currentCommandIndex int
	currentDirectory    *Directory
	rootDirectory       Directory
}

func parseCommand(state *State) {
	commandLine := state.commands[state.currentCommandIndex]
	if !strings.HasPrefix(commandLine, COMMAND_PREFIX) {
		panic(fmt.Sprint("Invalid command:", commandLine))
	}
	state.currentCommandIndex++
	commandParts := strings.Split(commandLine, " ")
	command := commandParts[1]
	args := commandParts[2:]
	switch command {
	case "cd":
		runCommandCd(args, state)
	case "ls":
		runCommandLs(state)
	}
}

func parseStructure(commands []string) []*Directory {
	state := State{
		allDirectories:      make([]*Directory, 0),
		rootDirectory:       newDirectory("/"),
		commands:            commands,
		currentCommandIndex: 0,
		currentDirectory:    nil,
	}
	for state.currentCommandIndex < len(state.commands) {
		parseCommand(&state)
	}
	return state.allDirectories
}

func runCommandCd(args []string, state *State) {
	dirName := args[0]
	if dirName == "/" {
		state.currentDirectory = &state.rootDirectory
		return
	}
	if dirName == ".." {
		state.currentDirectory = state.currentDirectory.parent
		return
	}
	directory := state.currentDirectory.directories[dirName]
	state.currentDirectory = &directory
}

func runCommandLs(state *State) {
	for state.currentCommandIndex < len(state.commands) {
		line := state.commands[state.currentCommandIndex]
		if strings.HasPrefix(line, COMMAND_PREFIX) {
			break
		}
		state.currentCommandIndex++
		lineParts := strings.Split(line, " ")
		typeOrSize := lineParts[0]
		name := lineParts[1]
		if typeOrSize == "dir" {
			directory := state.currentDirectory.addDirectory(name)
			state.allDirectories = append(state.allDirectories, directory)
		} else {
			size, _ := strconv.ParseInt(typeOrSize, 10, 32)
			state.currentDirectory.addFile(name, int(size))
		}
	}
}

func addDirectorySizes(allDirectories []*Directory) int {
	result := 0
	for _, dir := range allDirectories {
		size := dir.computeSize()
		if size <= 100000 {
			result += size
		}
	}
	return result
}

func readLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	commands := readLines()
	allDirectories := parseStructure(commands)
	sizes := addDirectorySizes(allDirectories)
	println("Day 1 result:", sizes)
}
