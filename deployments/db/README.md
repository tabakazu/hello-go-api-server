# データベース管理に関するドキュメント

## マイグレーション

作業は本リポジトリのルートディレクトリで行う

### migrate cli のインストール

```bash
GOBIN=$(PWD)/bin go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### データベースの作成

```bash
psql -h db -U $PGUSER -c "CREATE DATABASE myapp_dev;"
```

## データベースの削除

```bash
psql -h db -U $PGUSER -c "DROP DATABASE myapp_dev;"
```

### マイグレーションバージョンの確認

```bash
bin/migrate -path deployments/db/migrations/ -database $POSTGRES_DB_URL version
```

### マイグレーションファイルの作成

```bash
bin/migrate create -ext sql -dir deployments/db/migrations/ -seq create_table_name
```

### マイグレーションの適用

```bash
bin/migrate -path deployments/db/migrations/ -database $POSTGRES_DB_URL up

# stepを指定してマイグレーションを適用する場合
bin/migrate -path deployments/db/migrations/ -database $POSTGRES_DB_URL up 1
```

### マイグレーションのロールバック

```bash
bin/migrate -path deployments/db/migrations/ -database $POSTGRES_DB_URL down

# stepを指定してマイグレーションをロールバックする場合
bin/migrate -path deployments/db/migrations/ -database $POSTGRES_DB_URL down 1
```
