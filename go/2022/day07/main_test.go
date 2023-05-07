package day06

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"

	"github.com/mdwhatcott/advent-of-code/go/lib/parse"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

var (
	actualSession = util.InputLines()
	sampleSession = []string{
		"$ cd /",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}
)

func TestDay06(t *testing.T) {
	var sampleFiles = CalculateDirectorySizes(ParseFiles(sampleSession))
	var actualFiles = CalculateDirectorySizes(ParseFiles(actualSession))

	should.So(t, TotalSizeOfSmallDirectories(sampleFiles), should.Equal, 95437)
	should.So(t, TotalSizeOfSmallDirectories(actualFiles), should.Equal, 1118405)

	should.So(t, SizeOfDirectoryPreventingUpdate(sampleFiles), should.Equal, 24933642)
	should.So(t, SizeOfDirectoryPreventingUpdate(actualFiles), should.Equal, 12545514)
}
func ParseFiles(session []string) map[string]int {
	files := make(map[string]int)
	at := "/"
	for _, line := range session {
		fields := strings.Fields(line)
		if line == "$ cd /" {
			at = "/"
		} else if line == "$ cd .." {
			at = filepath.Join(filepath.Dir(at))
		} else if strings.HasPrefix(line, "$ cd ") {
			at = filepath.Join(at, fields[2])
		} else if strings.HasPrefix(line, "dir ") {
			continue
		} else {
			files[filepath.Join(at, fields[1])] = parse.Int(fields[0])
		}
	}
	return files
}
func CalculateDirectorySizes(files map[string]int) (directories map[string]int) {
	directories = make(map[string]int)
	for path, size := range files {
		for path != "/" {
			path = filepath.Dir(path)
			directories[path] += size
		}
	}
	return directories
}
func TotalSizeOfSmallDirectories(sizes map[string]int) (result int) {
	for _, size := range sizes {
		if size <= 100_000 {
			result += size
		}
	}
	return result
}
func SizeOfDirectoryPreventingUpdate(sizes map[string]int) int {
	const TotalCapacity = 70_000_000
	const SizeOfUpdate = 30_000_000
	currentlyUsed := sizes["/"]
	currentlyFree := TotalCapacity - currentlyUsed
	mustDeleteAtLeast := SizeOfUpdate - currentlyFree
	candidate := 0xFFFFFFFF
	for _, size := range sizes {
		if size >= mustDeleteAtLeast && size < candidate {
			candidate = size
		}
	}
	return candidate
}
