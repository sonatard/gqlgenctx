package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/sonatard/gqlgenctx/client/generated"
)

func main() {
	// クライアントを初期化
	client := generated.NewClient(
		http.DefaultClient,
		"http://localhost:8080/query", // GraphQLサーバーのエンドポイント
		&clientv2.Options{},
	)

	// コンテキストを作成
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// GraphQLクエリを実行
	result, err := client.Hello(ctx)
	if err != nil {
		log.Fatalf("クエリの実行に失敗しました: %v", err)
	}

	// 結果を表示
	fmt.Printf("GraphQLレスポンス: %s\n", result.Hello)
}
