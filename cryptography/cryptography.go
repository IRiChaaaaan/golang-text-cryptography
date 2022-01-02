package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// ファイルを暗号化
func EncodeFile(inFileName string, outFileName string) {
	// ファイル読み込み
	textFile, err := ioutil.ReadFile(inFileName)
	if err != nil {
		panic(err)
	}
	// 暗号化
	cryData := Encoder(textFile)
	// ファイルに保存
	err = ioutil.WriteFile(outFileName, cryData, 0644)
	if err != nil {
		panic(err)
	}
}

// ファイルを複合
func DecodeFile(inFileName string, outFileName string) {
	// ファイル読み込み
	textFile, err := ioutil.ReadFile(inFileName)
	if err != nil {
		panic(err)
	}
	// 複合
	cryData := Decoder(textFile)
	// ファイルに保存
	err = ioutil.WriteFile(outFileName, cryData, 0644)
	if err != nil {
		panic(err)
	}
}

// byteデータ列を暗号化
func Encoder(inData []byte) []byte {
	// 鍵の設定
	var key string = "s8fkelp29fk4ivjw"
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	// ガロアカウンタモードのブロック暗号
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	// 一度きりの初期化ベクトル作成
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}
	// 暗号化
	cipherText := gcm.Seal(nil, nonce, inData, nil)
	// 初期化ベクトルを先頭に付与
	cipherText = append(nonce, cipherText...)
	// Base64 Encode
	sEnc := base64.StdEncoding.EncodeToString(cipherText)
	return []byte(sEnc)
}

// 暗号byteデータ列を複合
func Decoder(cryData []byte) []byte {
	// 暗号化されたBase64テキストを変換
	cipherText, _ := base64.StdEncoding.DecodeString(string(cryData))
	// 鍵の設定
	var key string = "s8fkelp29fk4ivjw"
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	// ガロアカウンタモードのブロック暗号
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	//先頭に追加された初期化ベクトルを取り出す
	nonce := cipherText[:gcm.NonceSize()]
	plainByte, err := gcm.Open(nil, nonce, cipherText[gcm.NonceSize():], nil)
	if err != nil {
		fmt.Println(err)
	}
	return []byte(plainByte)
}
