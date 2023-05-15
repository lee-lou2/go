package configs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lee-lou2/go/pkg/mongo"
	"io/ioutil"
	"log"
	"os"
)

// LoadEnvironments 환경 변수 설정
func LoadEnvironments() {
	env := os.Getenv("GO_ENV")
	// .env 파일 불러오기
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	client, collection, err := mongo.GetCollection("configs", "env")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())
	data, _ := collection.Find("GO_" + env)
	// 환경 변수로 지정
	for key, value := range data.(map[string]interface{}) {
		strValue := fmt.Sprintf("%v", value)
		os.Setenv(key, strValue)
	}
	// 푸시 키 저장
	key, _ := collection.Find("GO_" + env + "_PUSH_KEY")
	file, _ := json.MarshalIndent(key.(map[string]interface{}), "", " ")
	if _, err := os.Stat("configs/keys"); os.IsNotExist(err) {
		_ = os.Mkdir("configs/keys", 0755)
	}
	_ = ioutil.WriteFile("configs/keys/serviceAccountKey.json", file, 0644)
}
