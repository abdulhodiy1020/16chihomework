package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	textFilename := "text.txt"
	text1Filename := "text1.txt"

	textFile, err := os.Open(textFilename)
	if err != nil {
		fmt.Printf("Manba faylni ochishda xatolik: %v\n", err)
		return
	}
	defer textFile.Close()

	text1File, err := os.Create(text1Filename)
	if err != nil {
		fmt.Printf("Hedef faylni yaratishda xatolik: %v\n", err)
		return
	}
	defer text1File.Close()

	_, err = io.Copy(text1File, textFile)
	if err != nil {
		fmt.Printf("Fayl ko'chirishda xatolik: %v\n", err)
		return
	}

	fmt.Println("Fayl ko'chirildi.")
}
