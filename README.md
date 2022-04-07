# GO-ITDDD-05-REPOSITORY

zenn の記事「[Go でリポジトリを実装（「入門ドメイン駆動設計」Chapter5）](https://zenn.dev/msksgm/articles/20220408-go-itddd-05-repository)」のサンプルコードです。

# 実行環境

- Go
  - 1.18
- docker compose

# 実行方法

## コンテナを起動・マイグレーション

コンテナの起動

```bash:コンテナの起動
> make up
docker compose up -d
# 完了までまつ
```

```bash:マイグレーション
> make run-migration
docker compose exec app bash db-migration.sh
1/u user (9.199ms)
```

## 実行

test-user 登録 1 回目

```bash:test-user 登録 1回目
> make run
docker compose exec app go run main.go
2022/04/07 22:19:24 successfully connected to database
2022/04/07 22:19:24 test-user is successfully added in users table
```

test-user 登録 2 回目

```bash:test-user 登録 2回目
> make run
docker compose exec app go run main.go
2022/04/07 22:19:34 successfully connected to database
2022/04/07 22:19:34 main.CreateUser err: the user &{{userid} {username}} is already exists
```

# テスト

```bash
> make test
docker compose exec app go test ./...
?       github.com/msksgm/go-itddd-05-repository        [no test files]
ok      github.com/msksgm/go-itddd-05-repository/domain/model/user      0.003s

```
