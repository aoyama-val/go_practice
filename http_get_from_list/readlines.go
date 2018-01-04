package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(filePath string) []string {
	// ファイルを開く
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to open '%s': %v\n", filePath, err)
		os.Exit(1)
	}

	// 関数return時に閉じる
	defer f.Close()

	// Scannerで読み込む
	lines := make([]string, 0, 128) // ある程度行数が事前に見積もれるようであれば、makeで初期capacityを指定して予めメモリを確保しておくことが望ましい
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "Scan error in '%s': %v\n", filePath, err)
	}

	return lines
}
