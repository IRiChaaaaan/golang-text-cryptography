package main

import (
	"fmt"
	"io/ioutil"
	"test/cryptography"
)

func main() {
	// txtデータを暗号化
	// cryptography.EncodeFile("data/words01.txt", "data/data01.txt")

	// 暗号データを複合
	textFile, err := ioutil.ReadFile("data/data01.txt")
	if err != nil {
		panic(err)
	}
	str := cryptography.Decoder(textFile)
	fmt.Println(string(str))
}
