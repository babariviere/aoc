package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type File struct {
	isDirectory bool
	name        string
	size        int
	children    []*File
}

func (f *File) Size() int {
	total := f.size
	for _, child := range f.children {
		total += child.Size()
	}
	return total
}

func (f File) Dirs() []*File {
	var dirs []*File
	for _, child := range f.children {
		if child.isDirectory {
			dirs = append(dirs, child)
			dirs = append(dirs, child.Dirs()...)
		}
	}
	return dirs
}

func (f File) String() string {
	if !f.isDirectory {
		return fmt.Sprint(f.name, " ", f.size)
	}
	var sb strings.Builder
	sb.WriteString(f.name)
	sb.WriteByte('\n')
	for _, child := range f.children {
		s := child.String()
		for _, sub := range strings.Split(s, "\n") {
			if sub == "" {
				continue
			}
			sb.WriteString("|- ")
			sb.WriteString(sub)
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

type Path struct {
	name   string
	parent *Path
}

func (p Path) pathToList() []string {
	if p.parent == nil {
		return []string{p.name}
	}
	return append(p.parent.pathToList(), p.name)
}

func execCd(cwd *Path, path string) *Path {
	if path == ".." {
		return cwd.parent
	}
	return &Path{
		name:   path,
		parent: cwd,
	}
}

func ensureDir(path *Path, cwd *File) *File {
	fmt.Println("Create dir", path.name)
	if path.name == "/" {
		return cwd
	}
	for _, file := range cwd.children {
		if file.name == path.name {
			return file
		}
	}
	f := &File{
		name:        path.name,
		isDirectory: true,
	}
	cwd.children = append(cwd.children, f)
	return f
}

func gotoPath(path *Path, root *File) *File {
	p := path.pathToList()
	// skip root, which is the first item
	cwd := root
	for _, dir := range p[1:] {
		for _, child := range cwd.children {
			if child.name == dir {
				cwd = child
				break
			}
		}
	}
	return cwd
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	root := File{
		isDirectory: true,
		name:        "/",
	}
	path := &Path{
		name: "/",
	}
	cwd := &root
	for scanner.Scan() {
	begin:
		command := strings.Split(scanner.Text(), " ")
		fmt.Println(command)
		switch command[1] {
		case "cd":
			path = execCd(path, command[2])
			if command[2] != ".." {
				cwd = ensureDir(path, cwd)
			} else {
				cwd = gotoPath(path, &root)
			}

		case "ls":
			for scanner.Scan() {
				if scanner.Text()[0] == '$' {
					goto begin
				}
				line := strings.Split(scanner.Text(), " ")
				rawSize := line[0]
				name := line[1]
				if rawSize == "dir" {
					ensureDir(&Path{name: name, parent: path}, cwd)
					continue
				}
				size, _ := strconv.Atoi(rawSize)
				fmt.Println("Create file", name)
				cwd.children = append(cwd.children, &File{
					name: name,
					size: size,
				})
				fmt.Println("CWD", cwd.name)
				fmt.Println(cwd)
			}
		}
	}

	// Collect all dirs
	dirs := root.Dirs()
	totalSize := 70000000
	unusedSpace := totalSize - root.Size()
	toFree := 30000000 - unusedSpace
	var total, smallest int
	smallest = math.MaxInt
	for _, dir := range dirs {
		size := dir.Size()
		if size <= 100000 {
			total += size
		}
		if size >= toFree && size < smallest {
			smallest = size
		}
	}

	fmt.Println("Total", total)
	fmt.Println("Smallest", smallest)
}
