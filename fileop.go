package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func fileOp() {

	//Open file

	file, err := os.Open("my_file.txt")
	if err != nil {
		if os.IsNotExist(err) {
			file, _ = os.Create("my_file.txt")
		} else {
			log.Fatal(err)
		}
	}

	defer file.Close()

	//Write to file

	// file, err = os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)

	// bufferedWriter := bufio.NewWriter(file)
	// bs := []byte{97, 98, 99}

	// // writing the byte slice to the buffer in memory
	// bytesWritten, err := bufferedWriter.Write(bs)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)

	// bufferedWriter.Flush()

	//Read from file
	// file, _ = os.Open("my_file.txt")
	// byteSlice := make([]byte, 3)

	// io.ReadFull(file, byteSlice)

	// fmt.Printf("io : %s\n", byteSlice)

	// data, _ := ioutil.ReadFile("my_file.txt")
	// fmt.Printf("ioutil : %s\n", data)

	//Scan file line by line

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Println(scanner.Text())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
