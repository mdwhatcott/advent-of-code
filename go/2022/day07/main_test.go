package day06

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"advent/lib/util"
)

var sampleLines = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

var (
	sample = ParseTerminalSession(sampleLines)
	input  = ParseTerminalSession(util.InputLines())
)

func TestDay06Part1(t *testing.T) {
	should.So(t, sample.TotalSizeOfSmallDirectories(), should.Equal, 95437)
	should.So(t, input.TotalSizeOfSmallDirectories(), should.Equal, 1118405)
}
func TestDay06Part2(t *testing.T) {
	should.So(t, sample.SizeOfSingleDirectoryToDelete(), should.Equal, 24933642)
	should.So(t, input.SizeOfSingleDirectoryToDelete(), should.Equal, 12545514)
}

func ParseTerminalSession(lines []string) *Dir {
	root := NewDir(nil, "/")
	var at *Dir
	for _, line := range lines {
		fields := strings.Fields(line)
		if line == "$ cd /" {
			at = root
		} else if line == "$ ls" {
			continue
		} else if line == "$ cd .." {
			at = at.Parent
		} else if strings.HasPrefix(line, "$ cd ") {
			at = at.Dirs[fields[2]]
		} else if strings.HasPrefix(line, "dir ") {
			at.Dirs[fields[1]] = NewDir(at, fields[1])
		} else {
			at.Files[fields[1]] = util.ParseInt(fields[0])
		}
	}
	return root
}

type Dir struct {
	Parent *Dir
	Name   string
	Dirs   map[string]*Dir
	Files  map[string]int
}

func NewDir(parent *Dir, name string) *Dir {
	return &Dir{
		Parent: parent,
		Name:   name,
		Dirs:   make(map[string]*Dir),
		Files:  make(map[string]int),
	}
}

func (this *Dir) Size() (sum int) {
	for _, dir := range this.Dirs {
		sum += dir.Size()
	}
	for _, file := range this.Files {
		sum += file
	}
	return sum
}
func (this *Dir) TotalSizeOfSmallDirectories() (result int) {
	size := this.Size()
	if size <= 100000 {
		result += size
	}
	for _, dir := range this.Dirs {
		result += dir.TotalSizeOfSmallDirectories()
	}
	return result
}
func (this *Dir) SizeOfSingleDirectoryToDelete() int {
	const (
		capacity = 70_000_000
		needed   = 30_000_000
	)
	var (
		currentlyFree   = capacity - this.Size()
		stillMustDelete = needed - currentlyFree
	)
	return util.Min[int](this.listDirsAbove(stillMustDelete)...)
}
func (this *Dir) listDirsAbove(minSize int) (results []int) {
	for _, dir := range this.Dirs {
		results = append(results, dir.listDirsAbove(minSize)...)
	}
	size := this.Size()
	if size >= minSize {
		results = append(results, size)
	}
	return results
}
func (this *Dir) DebugRendering(w io.Writer, depth int) {
	prefix := strings.Repeat(" ", depth*2)
	_, _ = io.WriteString(w, prefix)
	_, _ = io.WriteString(w, "- ")
	_, _ = io.WriteString(w, this.Name)
	_, _ = io.WriteString(w, " (dir)\n")
	for _, dir := range this.Dirs {
		dir.DebugRendering(w, depth+1)
	}
	for name, file := range this.Files {
		_, _ = io.WriteString(w, prefix)
		_, _ = io.WriteString(w, "  ")
		_, _ = io.WriteString(w, "- ")
		_, _ = io.WriteString(w, name)
		_, _ = fmt.Fprintf(w, " (file, size=%d)\n", file)
	}
}
