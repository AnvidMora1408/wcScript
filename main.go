package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Verificar si se pasa un argumento
	if len(os.Args) < 2 {
		fmt.Println("Error: Se necesita especificar un archivo como argumento")
		fmt.Println("Uso: go run . <nombre_del_archivo>")
		os.Exit(1)
	}
	
	filename := os.Args[1]
	
	// Verificar si el archivo existe
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("El archivo %s no existe", filename)
	}
	
	filewords := Words(filename)
	fileLines := Lines(filename)
	fileBytes := Characters(filename)

	fmt.Printf("%v %v %v %s\n", fileLines, filewords, fileBytes, filename)
}

func Words(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	nwords := len(words)

	return nwords
}

func Lines(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var counter int = 0

	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return counter
}

func Characters(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanBytes)

	var bytes []string

	for scanner.Scan() {
		Byte := string(scanner.Bytes())
		bytes = append(bytes, Byte)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	nbytes := len(bytes)

	return nbytes
}