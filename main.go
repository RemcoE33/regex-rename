package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func main() {
	a := os.Args
	if len(a) == 1 {
		fmt.Println("please provide regex")
		return
	}
	rx := strings.Trim(a[1], " ")

	if rx == "" {
		fmt.Println("please provide regex")
		return
	}

	dir := ""

	if len(a) == 3 {
		dir = a[2]
	}

	if dir == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic("could not get working directorty")
		}
		dir = wd
		fmt.Println(" --> gettings files from current directorty")
	} else {
		if !path.IsAbs(dir) {
			wd, err := os.Getwd()
			if err != nil {
				panic("could not get working directorty")
			}
			dir = path.Join(wd, dir)
			fmt.Println(" --> getting files from:", dir)
		}
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		e := fmt.Sprintf("could not read files from: %s", dir)
		fmt.Println(err)
		panic(e)
	}

	fmt.Println(" --> using:", rx, " to rename the files")
	regex := regexp.MustCompile(rx)

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		name := f.Name()

		if !regex.MatchString(name) {
			return
		}

		corrected := regex.FindStringSubmatch(name)[1]

		fmt.Println(name, " -> ", corrected)

		oldPath := path.Join(dir, name)
		newPath := path.Join(dir, corrected)

		if err := os.Rename(oldPath, newPath); err != nil {
			fmt.Println("could not rename:", name, " to: ", corrected)
		}
	}
}
