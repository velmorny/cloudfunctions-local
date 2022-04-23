# Cloud Functions (PubSub event) ローカル開発用サンプル

参考：[ローカルでの開発](https://cloud.google.com/functions/docs/local-development)

## サンプルの実行手順

[task](https://taskfile.dev/)を使用して実行している

### インストール

- `task install`

### Pub/Subエミュレータの起動

- `gcloud beta emulators pubsub start --project=localprj --host-port='localhost:8043'`

taskを使って、`task start-pubsub`で上記コマンドを実行させると停止時にプロセスが残る

その場合は`sudo lsof -i:8043`でプロセスを調べてkill

### Pub/Subのトピックとサブスクリプションの作成

- `create-pubsub-resouces`

### Cloud Functionsのビルド

- `task build`

### Cloud Functionsの起動

- `task run-pubsub-function`

### Cloud Functionsを実行

- `publish-test`
