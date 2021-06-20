# All Tasks

+ [x] Ipv6対応
+ [ ] セキュリティ向上
  + [参考1:NginxでHTTP2を有効にする](https://qiita.com/Aruneko/items/8c11f9e45a33457c3c1f)
  + [参考2:Nginx導入時やること](https://qiita.com/kidach1/items/985efebba639713c562e)
+ [x] HTTP/2対応
  + `http2_push index.css`とすると一応h2にはなるが，思ったのと違う
  + <解決> `localhost` にIPv6で接続していたらしい．なので，`listen [::]:443 ssl http2 default_server;` を追加するとh2で話すように． 
+ [ ] HTTP/3対応？
+ [ ] Webアプリケーションとの連携
  + [ ] Webアプリ用コンテナのDocker環境整備
  + [ ] NginxからWebアプリへのリバースプロキシ設定
+ [ ] 開発環境/本番環境の切り分け
  + [ ] Dockerfileを個別に
+ [ ] Chocolateyを非管理者でも実行できるようにする
+ [ ] `docker-compose.yml`の設定を最適化
  + [ ] volumesマウントの設定がイマイチよく分かってない
    + [ ] 証明書について，`Dockerfile` で `COPY` がうまくいかないので仕方なくvolumesを使っているが，Read-Onlyにしたい．
+ [x] Dockerfile中でGO Modulesがうまく機能しない
  + [ ] `RUN go mod init github.com/~~`で正常に生成されない
  + [ ] したがって`go build`もエラー
  + [x] <解決>　単純に"\"で連結するのが失敗してた
+ [ ] HTTP/3だとうまくいかない(`https://localhost/proxy`へのアクセス)
  + [ ] ERR_CONNECTION_CLOSED (Alt-svcが"h3-34="のとき)
  + [ ] 404 Not Found(Alt-svcが"v1="のとき)
  + [ ] h3のnginxコンテナの構成がalpineとは違う．別の設定ファイルを読み込んでるしLinuxでもない？