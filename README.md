# mkcert-test

mkcert を使ってローカルに HTTPS 環境を建てる

タスクリストは[TASKs.md](./TASKs.md)で管理

### 開発環境

+ Windows 10 Home
  + version: 2004
  + OS build: 19041.1052
+ WSL2 distro
  + Ubuntu-20.04
+ Docker for Windows
  + version: 3.4.0 (65384)

### 依存ツール一覧

+ Docker
+ docker-compose
+ mkcert

## 依存ツール

### Docker

```bash
> docker -v
Docker version 20.10.7, build f0df350
```

### docker-compose

```bash
> docker-compose -v
docker-compose version 1.29.2, build 5becea4c
```

### mkcert

[Chocolatey](https://chocolatey.org/install) を使ってインストール

管理者権限で PowerShell を立ち上げて以下

```bash
> Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))

# chocolatey で mkcert をインストール
> choco install mkcert

# mkcert でローカルの CA を建てる
> mkcert --install
```

mkcert での証明書発行は以下

```bash
# 証明書の欲しいディレクトリに移動
> cd [path/to/certificates]

# mkcert で証明書発行(ホスト名は複数指定可能)
> mkcert localhost 127.0.0.1

# 以下のファイルが生成される (<ホスト名>.pem が証明書，<ホスト名>-key.pem が秘密鍵．'+1'はファイル名になっているホスト以外のホスト数)
> ls
localhost+1.pem localhost+1-key.pem
```

## HTTP/2対応

利用する Nginx のイメージが `latest` だとOpenSSLのバージョンの都合でHTTP/2に対応できないらしい ([参考](https://qiita.com/ktateish/items/49ccb8d1cf622f65c8fc))．古い記事なのでもう関係ないかもだが

そのためイメージには `alpine` の方を使う

<details><summary>解決した話</summary><div>

でも現状HTTP/2を話してくれない…
+ HTTP/2対応のNginxイメージ(`ehekatl/docker-nginx-http2`)でもだめだった
+ [このページ](https://tech.recruit-mp.co.jp/infrastructure/post-12795/) で使われていた `nginx:1.13.5-alpine` でもだめ
  
結局，IPv6 でアクセスしてたらしく設定を追加したら疎通

</div></details>

## HTTP から HTTPS へのリダイレクト

`default.conf` 内で HTTP:80 を同 URI の HTTPS にリダイレクトしている

`nginx` コンテナのポートも `default.conf` のポート設定も，HTTP:80は削除してもいいかも

## localhostの証明書だけChromeの警告を無効にする

**これは特に必要ないことが分かった**

mkcert を採用したため，https-portalとは事情が違うみたい(本来https-portalのときに必要とされてた設定)

ローカルのHTTPS環境で使う証明書は"オレオレ証明書"なので，警告が出てウザい


<details><summary>無効化の手順は残しておく</summary><div>

次のフラグを `Enable` にするとlocalhostに限り警告が表示されなくなる
```
chrome://flags/#allow-insecure-localhost
```

([参考:ChromeのSSL警告を、localhostの時だけ表示しないようにする](https://qiita.com/yanchi4425/items/76e502c41cbfb4f0542b))

</div></details>

## 参考記事

+ [Docker+nginxでHTTPSサーバを簡単に構築](https://qiita.com/yujimny/items/7615046be674895f9565)
  + NginxでHTTPSを扱う設定方法について
+ [Nginx版:mkcertを使ってローカル環境でもDockerでも楽々SSL](https://qiita.com/ProjectEuropa/items/bde085cbeff7b7e69295)
  + HTTPS環境をローカルに構築するために一連の流れが参考になった．
  + ブラウザで特に警告なく表示できることが採用の判断基準に合致
+ [Gin+Realize+docker-compose+Nginxで環境を整えてみる](https://qiita.com/ririkku/items/2ad76f6867c7b4078fc7)
  + とりあえず動かすGinのサンプルコードを利用させてもらった
+ [WebサーバーのHTTP/3対応をNginxのリバースプロキシでするためのDockerイメージが出来ました](https://qiita.com/nwtgck/items/ff633df298dfd9dc0887)
  + とりあえずHTTP/3を見据えた記事
+ [https://blog.homie.co.jp/entry/golang-air](https://blog.homie.co.jp/entry/golang-air)
  + ホットリロードのライブラリ`air`の使い方．`realize`はサポートが止まってるのか`go get`でエラーが出まくってコンテナがビルドできないので代替案を探した
+ [Let's EncryptでSSL対応したNginxのリバースプロキシを構築する](https://qiita.com/kenchan1193/items/850fd26a70a23fb141d3#docker%E3%82%B3%E3%83%B3%E3%83%86%E3%83%8A%E5%86%85%E3%81%AE%E3%83%AD%E3%83%BC%E3%82%AB%E3%83%AB%E3%83%9B%E3%82%B9%E3%83%88%E3%81%AB%E9%96%A2%E3%81%97%E3%81%A6)
  + 「ハマったこと」にある『「localhost」はコンテナ自身を指し同一ホストの他のコンテナにはたどり着けない』ということがヒントに．
+ [nginx のリバースプロキシを Docker-Compose で試してみる](https://neos21.net/blog/2020/06/24-01.html)
  + GitHubリポジトリの`default.conf`と`docker-compose.yml`を見て，`proxy_pass`の転送先ドメインをDockerのコンテナ名(としてのドメイン名)にしてよいことと，そのために`docker-compose.yml`側で設定は必要ないことを学んだ．