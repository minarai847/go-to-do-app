# 📝 Todo App - React × Go(Echo) × PostgreSQL

本アプリケーションは、React（Vercel）× Go(Echo)（Render）× PostgreSQL を使用した  
フルスタック Todo 管理アプリです。

学習目的にとどまらず、**本番運用を意識した構成（CORS / 認証 / CSRF / デプロイ）** を採用しており、  
Web アプリの基礎からデプロイまで一通り経験できるプロジェクトです。

---

## 🚀 本番環境 URL

### Frontend（Vercel）
🔗 https://go-to-do-app.vercel.app/

### Backend API（Render）
🔗 https://（あなたの Render API URL）  
※ `/health` を作成している場合は: https://◯◯◯.onrender.com/health

---

## 🏗 使用技術（Tech Stack）

### Frontend
- React
- Axios
- Vercel（デプロイ・ホスティング）
- Cookie ベース認証（withCredentials）

### Backend
- Go (Echo)
- Clean Architecture（Controller / Usecase / Repository）
- GORM
- JSON API
- JWT or Cookie 認証
- CSRF トークン
- CORS 設定

### Infrastructure
- Vercel (CI/CD)
- Render（Go API, PostgreSQL）
- PostgreSQL 15

---

## ✨ 主な機能

### 👤 認証機能
- 新規ユーザー登録（Signup）
- ログイン / ログアウト
- JWT または Cookie によるセッション管理
- CSRF Token 発行（安全性向上）
- 認証済みユーザーのみがタスクを操作可能

### 📝 タスク管理（CRUD）
- タスク一覧取得
- タスクの作成
- タスク名の更新
- タスクの完了ステータス更新
- タスク削除

### 💡 その他
- リロードしてもログイン状態が保持される
- エラーメッセージをわかりやすくフロントに表示

---

## 📐 アプリ構成図

React (Vercel)
↓ HTTPS / JSON API
Go API (Render)
↓
PostgreSQL (Render)

yaml
コードをコピーする

---

## 📁 ディレクトリ構成（Backend）

go-to-do-app/
├── controller/ # ルーティング・リクエスト処理
├── usecase/ # ビジネスロジック
├── repository/ # DB アクセス処理
├── model/ # GORM モデル
├── validator/ # 入力バリデーション
├── router/ # Echo ルーティング設定
├── db/ # DB 接続
└── main.go # エントリーポイント

yaml
コードをコピーする

---

## 🔑 工夫したポイント（実務で評価される点）

### 1. **開発環境と本番環境で API URL を切り替え**
React 側で環境変数を利用し、  
ローカル開発（localhost）と本番（Render）を安全に切り替え。

### 2. **CORS・CSRF・Cookie の扱いを正しく実装**
フロントとバックエンドが別ドメインのため、  
本番環境では CORS・CSRF を正しく理解しないと動作しない。

このプロジェクトでは以下を対応：

- `withCredentials: true`
- Echo の CORS 設定に Vercel の URL を許可
- CSRF トークンの付与・検証

### 3. **Render の Postgres に接続し、AutoMigrate でテーブル生成**
ローカルと本番の DB の違いによるエラー（SQLSTATE 42P01）も解決。

### 4. **本番環境でデプロイ成功するまでの運用経験**
- Render の PORT バインディング
- DATABASE_URL の扱い
- デプロイ時のエラー調査（Network Error / CORS / DB）
- Vercel × Render の CI/CD

実務で必ず必要になるスキルを一通り経験。

---

## 🧪 動作手順（開発環境）

### Backend
go run main.go

shell
コードをコピーする

### Frontend
npm install
npm start

yaml
コードをコピーする

---

## 🎯 今後の改善予定

- タグ機能・優先度機能の追加
- UI/UX の改善（レスポンシブ対応）
- ユーザープロフィール機能
- Docker 化
- テストコード（Go / React）の追加

---

## 📄 ライセンス
MIT License

---
