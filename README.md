# iot-ak4

AKASHIをAWS Iot Buttonで勤怠つけるためのツール ついでにAthenaに情報を記録する

## Install

```shell
$ make main.zip
```

Lambda関数を作成してmain.zipをアップロードし、ハンドラを `main` に設定する。ほか、AthenaへのアクセスのためにIAMを設定する必要がある。

環境変数として APIトークンである`AK_TOKEN` と `AK_COOP_ID` が必要。

Athena書き込みに使う環境変数として以下が必要。

- `ATHENA_DB_NAME`
- `ATHENA_OUTPUT_LOCATION`
- `ATHENA_CATALOG`
- `ATHENA_WORKGROUP`
- `ATHENA_TABLE`

Athenaの書き込み先となるテーブルのスキーマは次のカラムを持っている必要がある。

- `created_at TIMESTAMP`
- `operation STRING`
- `year SMALLINT`
- `month TINYINT`
