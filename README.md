# iot-ak4

AKASHIをAWS Iot Buttonで勤怠つけるためのツール ついでにAthenaに情報を記録する

## Install

```shell
$ make main.zip
```

Lambda関数を作成してmain.zipをアップロードし、ハンドラを `main` に設定する。ほか、AthenaへのアクセスのためにIAMを設定する必要がある。

環境変数として APIトークンである`AK_TOKEN` と `AK_COOP_ID` が必要。
