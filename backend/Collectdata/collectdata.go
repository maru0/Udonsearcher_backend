package collectdata

import (
	"context"
	"fmt"
	"os"

	//"fmt"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo/"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func loadenv() (string, string) {
	// DB情報の抽出
	checks := godotenv.Load()
	if checks != nil {
		fmt.Println("Error loading")
	}

	hosts := os.Getenv("testdb")
	pass := os.Getenv("DB_PASSWORD")
	return hosts, pass
}

func Collectdata() {
	// DB接続
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer: 関数内で最後に処理を走らせる
	defer cancel()
	hosts, dbport := loadenv()
	dbURL := "mongodb://" + hosts + ":" + dbport
	c, _ := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	defer c.Disconnect(ctx)

	// 先にデータを整形
}
