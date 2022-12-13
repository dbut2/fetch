package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var day int
var year int
var template string
var ide string
var files string

func main() {
	flag.IntVar(&day, "day", today(), "day to download puzzle for")
	flag.IntVar(&year, "year", time.Now().Year(), "year to download puzzle for")
	flag.StringVar(&template, "template", "template", "template folder")
	flag.StringVar(&ide, "ide", "goland", "ide command to open files, must support opening files like \"$ {IDE} example\"")
	flag.StringVar(&files, "files", "puzzle.md", "comma seperated list of files to open automatically")
	flag.Parse()

	if year < 100 {
		year += 2000
	}

	// create necessary directories
	dir := fmt.Sprintf("%4d/%02d/", year, day)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err.Error())
		log.Fatalln(err.Error())
	}
	puzzle := dir + "puzzle.md"
	input := dir + "input.txt"

	// copy template
	exec.Command("cp", "-Rn", template+"/", dir).Run()

	// run aoc tool to download puzzle and input
	aoc := exec.Command("aoc", "download", "--overwrite", fmt.Sprintf("--day=%d", day), fmt.Sprintf("--year=%d", year), "--puzzle-file="+puzzle, "--input-file="+input)
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	aoc.Stdout = stdout
	aoc.Stderr = stderr
	err = aoc.Run()
	if err != nil {
		io.Copy(os.Stdout, stdout)
		io.Copy(os.Stderr, stderr)
		log.Fatalln(err.Error())
	}

	// open files in ide
	for _, file := range strings.Split(files, ",") {
		file = dir + file
		_, err = os.Stat(file)
		if err != nil {
			continue
		}
		exec.Command(ide, file).Run()
	}
}

// today date, or yesterday if ahead of ETC
func today() int {
	d := time.Now().Day()
	tz := time.FixedZone("ETC", -5*60*60)
	if time.Now().Day() > time.Now().In(tz).Day() {
		d--
	}
	return d
}
