[English](README.md) | [日本語](README_ja.md)

# url-anchor

[![Build Status](https://travis-ci.com/ebc-2in2crc/url-anchor.svg?branch=master)](https://travis-ci.com/ebc-2in2crc/url-anchor)
![CI](https://github.com/ebc-2in2crc/url-anchor/actions/workflows/pr.yml/badge.svg)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![GoDoc](https://godoc.org/github.com/ebc-2in2crc/url-anchor?status.svg)](https://godoc.org/github.com/ebc-2in2crc/url-anchor)
[![Go Report Card](https://goreportcard.com/badge/github.com/ebc-2in2crc/url-anchor)](https://goreportcard.com/report/github.com/ebc-2in2crc/url-anchor)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/url-anchor.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/url-anchor.svg?label=version)

url-anchor は URL を HTML `<a>` タグに変換するプログラムです。

## Description

url-anchor は URL を与えると URL とページのタイトルを HTML の `<a>` タグに変換します。

デフォルトの動作は与えた URL とページのタイトルを HTML の タグにしますがオプションを指定すると Markdown のリンクと reStructuredText のリンクにすることもできます。

## Usage

```
$ url-anchor https://google.com
<a href="https://google.com">Google</a>

$ url-anchor -m https://google.com
[Google](https://google.com)

$ url-anchor -r https://google.com
`Google <https://google.com>`_

$ url-anchor -help
# ...
```

Docker を使うこともできます。

```
$ docker container run --rm ebc2in2crc/url-anchor https://google.com/
<a href="https://google.com/">Google</a>
```

URL は標準入力から読むことができます。

```
$ echo https://google.com | url-anchor -
<a href="https://google.com">Google</a>

$ cat << EOF | url-anchor -
https://google.com
https://www.whitehouse.gov
EOF
<a href="https://google.com">Google</a>
<a href="https://www.whitehouse.gov">The White House</a>
```

### Use a clipboard (Mac のみ)

クリップボードを使う便利な使い方。

1. URL をクリップボードにコピーします。
2. 以下のコマンドを実行すると実行結果をクリップボードにコピーします。

```
$ url-anchor $(pbpaste) | pbcopy

# $ pbpaste
# <a href="https://www.yahoo.co.jp/">Yahoo! JAPAN</a>
```

## Installation

### Developer

Go 1.16 or later.

```
go install github.com/ebc-2in2crc/url-anchor@latest
```

Go 1.15.

```
$ go get -u github.com/ebc-2in2crc/url-anchor/...
```

### User

次の URL からダウンロードします。

- [https://github.com/ebc-2in2crc/url-anchor/releases](https://github.com/ebc-2in2crc/url-anchor/releases)

Homebrew を使うこともできます (Mac のみ)

```sh
$ brew tap ebc-2in2crc/tap
$ brew install url-anchor
```

Docker を使うこともできます。

```
$ docker image pull ebc2in2crc/url-anchor
```

## Contribution

1. Fork this repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Rebase your local changes against the master branch
5. Run test suite with the `make test` command and confirm that it passes
6. Run `make fmt` and `make lint`
7. Create new Pull Request

## Licence

[MIT](https://github.com/ebc-2in2crc/url-anchor/blob/master/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
