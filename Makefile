.PHONY: gen run-server run-server2 run-client

# コード生成
gen:
	go tool gqlgen
	cd client && go tool gqlgenc

# サーバー実行
run-server:
	PORT=8080 go run ./cmd/server/main.go

run-server2:
	PORT=8081 go run ./cmd/server/main.go

# クライアント実行
run-client:
	go run ./cmd/client/main.go
