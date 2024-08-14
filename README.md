# example-pulumi-gcp-go

## installation

このリポジトリは vscode の dev container で開発することを前提としています。開発するには以下のツールと環境変数を設定する必要があります。

### tools

- vscode
- docker

### env

- EXAMPLE_PULUMI_GCP_GO_PULUMI_ACCESS_TOKEN  
  pulumi のアクセストークンです。 [pulumi の公式サイト](https://app.pulumi.com/signin)から取得できます。
- EXAMPLE_PULUMI_GCP_GO_GCP_PROJECT_ID  
  開発に使う GCP のプロジェクト名です。プロジェクトにアクセスするユーザーに必要な権限を詳細に管理する場合は[こちらのページ](https://www.pulumi.com/docs/clouds/gcp/get-started/begin/)を参照してください。
- EXAMPLE_PULUMI_GCP_GO_STACK_NAME  
  環境を識別する値を設定してください。通常は `prod`, `stg`, `dev`, `local` のような値を設定します。

## commands

### GCP にログインする

pulumi で GCP アクセスするには事前に gcp cli にログインする必要があります。

```sh
gcloud auth application-default login --no-launch-browser
```

### デプロイする

```sh
# module1
./.cmd/deploy-module1.sh

#module2
./.cmd/deploy-module2.sh
```

### 削除する

```sh
# module1
./.cmd/destroy-module1.sh

#module2
./.cmd/destroy-module2.sh
```
