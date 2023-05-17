# rebabel-websocket
Re:babelのWebSocketサーバー。
## 構成
```
rebabel-websocket
 ├──api(Echo)
 |   ├──domain(Entities)
 |   |   └──example.go
 |   ├──infrastructure(Frameworks & Drivers)
 |   |   ├──router.go
 |   |   └──sqlhandler.go
 |   ├──interfaces(Interface)
 |   |   ├──controllers
 |   |   |   ├──example_controller.go
 |   |   |   ├──context.go
 |   |   |   ├──error.go
 |   |   |   └──token.go
 |   |   └──database
 |   |       ├──example_repository.go
 |   |       └──sqlhandler.go
 |   └──usecase(Usecases)
 |       ├──example_interactor.go
 |       └──example_repository.go
 ├──docker
 |   └──go
 |      └──Dockerfile
 ├──migration(SQL, マイグレーション)
 |   ├──000001_create_example.up.sql
 |   └──000001_create_example.down.sql
 ├──test
 ├──.air.toml
 ├──.env.example
 ├──.gitignore
 ├──docker-compose.yaml
 ├──go.mod
 ├──go.sum
 ├──main.go
 └──README.md
```
## 環境構築
1.コンテナを起動
```
docker compose up -d --build
```
2.コンテナに入る
```
docker container exec -it rebabel-websocket-app-1 bash
```
3.マイグレーションの作成(例：usersテーブル)
```
migrate create -ext sql -dir migration -seq create_users
```
4.マイグレーションの実行
```
migrate -database="mysql://root:root@tcp(host.docker.internal:3306)/rebabel-database?multiStatements=true" -path=migration up
```
5.サーバー(air)を起動
```
air
```
