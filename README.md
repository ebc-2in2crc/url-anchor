[English](README.md) | [日本語](README_ja.md)

# url-anchor

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

`url-anchor` converts URL to HTML `<a>` tag.

## Description

`url-anchor` converts URL and page titles to HTML `<a>` tag when you give a URL.

You can also convert a Markdown link to a reStructuredText link by specifying an option.

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

## Installation

### Developer

```
$ go get -u github.com/ebc-2in2crc/url-anchor/...
```

### User

Download from the following url.

- [https://github.com/ebc-2in2crc/url-anchor/releases](https://github.com/ebc-2in2crc/url-anchor/releases)

Or, you can use Homebrew (Only macOS).

```sh
$ brew tap ebc-2in2crc/tap
$ brew install url-anchor
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
