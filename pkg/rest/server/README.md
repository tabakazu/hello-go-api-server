# REST API Server のコード自動生成に関するドキュメント

## コードの自動生成

作業は本リポジトリのルートディレクトリで行う

### ogen cli のインストール

```sh
GOBIN=$(PWD)/bin go install github.com/ogen-go/ogen/cmd/ogen@latest
```

### コードの生成

```sh
bin/ogen --target pkg/rest/server -package server --clean api/openapi/dist/swagger.gen.yaml
```
