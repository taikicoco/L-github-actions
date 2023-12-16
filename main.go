package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// RandomString は指定された長さのランダムな文字列を生成します。
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Data はJSONファイルに書き込むデータ構造です。
type Data struct {
	RandomString string `json:"randomString"`
}

func main() {
	// ランダムな文字列を生成
	randomStr := RandomString(10) // 10文字のランダムな文字列

	// Dataオブジェクトを作成
	data := Data{RandomString: randomStr}

	// JSONに変換
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ファイルに書き込み
	err = os.WriteFile("random.json", file, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Random JSON file created successfully.")
}
