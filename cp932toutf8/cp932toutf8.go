// CP932.TXTを使ってCP932のファイルをUTF-8に変換する

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var table map[int]int

// Unicode code point to UTF-8
func code_to_utf8(code int) ([4]byte, int) {
	var buffer [4]byte

	if code <= 0x7F {
		buffer[0] = byte(code)
		return buffer, 1
	}
	if code <= 0x7FF {
		buffer[0] = byte(0xC0 | (code >> 6))   /* 110xxxxx */
		buffer[1] = byte(0x80 | (code & 0x3F)) /* 10xxxxxx */
		return buffer, 2
	}
	if code <= 0xFFFF {
		buffer[0] = byte(0xE0 | (code >> 12))         /* 1110xxxx */
		buffer[1] = byte(0x80 | ((code >> 6) & 0x3F)) /* 10xxxxxx */
		buffer[2] = byte(0x80 | (code & 0x3F))        /* 10xxxxxx */
		return buffer, 3
	}
	if code <= 0x10FFFF {
		buffer[0] = byte(0xF0 | (code >> 18))          /* 11110xxx */
		buffer[1] = byte(0x80 | ((code >> 12) & 0x3F)) /* 10xxxxxx */
		buffer[2] = byte(0x80 | ((code >> 6) & 0x3F))  /* 10xxxxxx */
		buffer[3] = byte(0x80 | (code & 0x3F))         /* 10xxxxxx */
		return buffer, 4
	}
	return buffer, 0
}

// Load CP932.TXT into table
func loadTable(filename string) map[int]int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	ret := make(map[int]int)
	for sc.Scan() {
		line := sc.Text()
		if line[0] == '#' {
			continue
		}
		fields := strings.Split(line, "\t")
		if len(fields) != 3 {
			fmt.Printf("len != 3, %d\n", len(fields))
			os.Exit(1)
		}
		var cp932 int
		var code int
		fmt.Sscanf(fields[0], "0x%x", &cp932)
		fmt.Sscanf(fields[1], "0x%x", &code)
		ret[cp932] = code
	}
	return ret
}

func cp932ToUtf8(in []byte) []byte {
	w := bytes.NewBuffer(make([]byte, 0, len(in)))
	i := 0
	for i < len(in) {
		b := in[i]
		var unicode int
		if b&0x80 > 0 {
			unicode = table[int(b)<<8+int(in[i+1])]
			i += 2
		} else {
			unicode = table[int(b)]
			i += 1
		}
		bytes, _ := code_to_utf8(unicode)
		_, err := w.Write(bytes[:])
		if err != nil {
			panic("Write")
		}
	}
	return w.Bytes()
}

func convertFile(in_filename string, out_filename string) {
	cp932bytes, _ := ioutil.ReadFile(in_filename)
	out, err := os.Create(out_filename)
	if err != nil {
		panic("Create")
	}
	defer out.Close()
	utf8bytes := cp932ToUtf8(cp932bytes)
	out.Write(utf8bytes)
}

func init() {
	table = loadTable("CP932.TXT")
}

func main() {
	convertFile("a.txt", "out.txt")
}
