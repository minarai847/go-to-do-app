# Frontend - React Todo App

## プロジェクト構成

```
frontend/
├── public/           # 静的ファイル
├── src/
│   ├── components/  # 再利用可能なコンポーネント
│   ├── pages/       # ページコンポーネント
│   ├── services/    # API呼び出し
│   ├── hooks/       # カスタムフック
│   ├── utils/       # ユーティリティ関数
│   ├── types/       # TypeScript型定義
│   └── contexts/     # React Context
└── package.json
```

## セットアップ

```bash
npm install
```

## 開発サーバーの起動

```bash
npm start
```

## 環境変数

`.env`ファイルを作成して、以下を設定：

```
REACT_APP_API_URL=http://localhost:8080
```

## ビルド

```bash
npm run build
```
