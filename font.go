package figure

import (
	"bufio"
	"bytes"
	"embed"
	"io"
	"path"
	"strings"
)

const defaultFont = "standard"

// Fonts 嵌入配置文件模板到程序内部
//
//go:embed fonts/*
var Fonts embed.FS

type font struct {
	name      string
	height    int
	baseLine  int
	hardBlank byte
	reverse   bool
	letters   [][]string
}

func newFont(name string) (font font) {
	font.setName(name)
	fontBytes, err := Fonts.ReadFile(path.Join("fonts", font.name+".flf"))
	if err != nil {
		panic(err)
	}
	fontBytesReader := bytes.NewReader(fontBytes)
	scanner := bufio.NewScanner(fontBytesReader)
	font.setAttributes(scanner)
	font.setLetters(scanner)
	return font
}

func newFontFromReader(reader io.Reader) (font font) {
	scanner := bufio.NewScanner(reader)
	font.setAttributes(scanner)
	font.setLetters(scanner)
	return font
}

func (font *font) setName(name string) {
	font.name = name
	if len(name) < 1 {
		font.name = defaultFont
	}
}

func (font *font) setAttributes(scanner *bufio.Scanner) {
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, signature) {
			font.height = getHeight(text)
			font.baseLine = getBaseLine(text)
			font.hardBlank = getHardBlank(text)
			font.reverse = getReverse(text)
			break
		}
	}
}

func (font *font) setLetters(scanner *bufio.Scanner) {
	font.letters = append(font.letters, make([]string, font.height)) //TODO: set spaces from flf
	for i := range font.letters[0] {
		font.letters[0][i] = "  "
	} //TODO: set spaces from flf
	letterIndex := 0
	for scanner.Scan() {
		text, cutLength, letterIndexInc := scanner.Text(), 1, 0
		if lastCharLine(text, font.height) {
			font.letters = append(font.letters, []string{})
			letterIndexInc = 1
			if font.height > 1 {
				cutLength = 2
			}
		}
		if letterIndex > 0 {
			appendText := ""
			if len(text) > 1 {
				appendText = text[:len(text)-cutLength]
			}
			font.letters[letterIndex] = append(font.letters[letterIndex], appendText)
		}
		letterIndex += letterIndexInc
	}
}

func (font *font) evenLetters() {
	var longest int
	for _, letter := range font.letters {
		if len(letter) > 0 && len(letter[0]) > longest {
			longest = len(letter[0])
		}
	}
	for _, letter := range font.letters {
		for i, row := range letter {
			letter[i] = row + strings.Repeat(" ", longest-len(row))
		}
	}
}
