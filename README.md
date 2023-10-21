# go-clean-architecture

このリポジトリは、Go言語を使用してClean Architectureの原則に基づいて設計されたサンプルプログラムを提供しています。

![Clean Architectureの図](https://github.com/mi-01-24fu/go-clean-architecture/assets/133470645/417fb5e2-d78f-4907-b3a2-3f10aa170d02)

## 概要

ユーザーリソースを提供するアプリケーションです。以下のエンドポイントを実装しています：

- `POST /users`: ユーザーを登録する
- `GET /users`: ユーザーのリストを取得する
- `GET /user/:id`: 特定のユーザーの詳細を取得する

各ユーザーは次の属性を持っています：
- `id`
- `firstName`
- `lastName`

## ディレクトリ構造
```
C:.
│  go.mod
│  go.sum
│  README.md
│
└─src
    ├─app
    │      main.go
    │
    ├─domain
    │      user.go
    │
    ├─infrastructure
    │      router.go
    │      sqlhandler.go
    │
    ├─interfaces
    │  ├─controllers
    │  │      context.go
    │  │      user_controller.go
    │  │
    │  └─database
    │          sql_handler.go
    │          user_repository.go
    │
    └─usecase
            user_interactor.go
            user_repository.go
```

## アーキテクチャのレイヤー

| ディレクトリ名    | レイヤー               |
|-----------------|--------------------|
| domain          | エンティティ         |
| infrastructure  | フレームワーク & ドライバ |
| interfaces      | インターフェース       |
| usecase         | ユースケース           |

## 初期構築
### DB作成
```
CREATE DATABASE sample;
```

### TABLE作成
```
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL
);
```

## 実行
### サーバー起動
```
go run src/app/main.go
```

### 動作確認
#### User登録
```
$ curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"FirstName": "FirstName", "LastName": "LastName"}' localhost:8080/users
```

#### User一覧
```
$ curl -i -H 'Content-Type:application/json' localhost:8080/users
```

#### User情報
```
$ curl -i -H 'Content-Type:application/json' localhost:8080/users/3
```

## 参考
- [クリーンアーキテクチャ完全に理解した](https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60) 
- [GitHub - bxcodec/go-clean-arch: Go (Golang) Clean Architecture based on Reading Uncle Bob's Clean Architecture](https://github.com/bxcodec/go-clean-arch) 
- [実践クリーンアーキテクチャ with Java](https://nrslib.com/clean-architecture-with-java/#outline__6_3)
- [Clean ArchitectureでAPI Serverを構築してみる - Qiita](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e) 
- [クリーンアーキテクチャ(The Clean Architecture翻訳)](https://blog.tai2.net/the_clean_architecture.html) 
