# gqlgenctx

gqlgenを使用したctxのエラーの伝播を検証するためのリポジトリです。

## シーケンス図

以下はこのプロジェクトの動作フローを示すシーケンス図です：

```mermaid
sequenceDiagram
    participant Client as クライアント
    participant GQLServer1 as GraphQLサーバー1<br>(ポート8080)
    participant GQLServer2 as GraphQLサーバー2<br>(ポート8081)
    
    Client->>+GQLServer1: hello クエリ
    Note over GQLServer1: Helloリゾルバーが処理を開始
    
    GQLServer1->>+GQLServer2: hello2 GraphQLクエリ
    Note over GQLServer2: Hello2リゾルバーが処理開始<br>5秒間待機
    Note over GQLServer2: コンテキストエラーチェック
    GQLServer2-->>-GQLServer1: "hello2"レスポンス
    
    Note over GQLServer1: レスポンスをログに出力
    GQLServer1-->>-Client: "world"
    
    Note right of Client: クライアント処理完了
```

## セットアップと実行

このプロジェクトではMakefileを使って各種操作を行います。

### 実行

```bash
# コード生成
make gen

# GraphQLサーバーを実行
make run-server
make run-server2

# クライアントを実行（サーバーが起動していることが前提）
make run-client
```


## サーバー情報

### GraphQLサーバー

GraphQLサーバーが起動したら、ブラウザで http://localhost:8080 と http://localhost:8081 にアクセスして、GraphQLプレイグラウンドを使用できます。

#### サンプルクエリ

```graphql
query {
  hello
}
```