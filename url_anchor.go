package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/mattn/go-encoding"
	"golang.org/x/net/html/charset"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/pkg/errors"
)

var re *regexp.Regexp

var (
	markdownOpt bool
	reSTOpt     bool
	versionOpt  bool
)

func init() {
	re = regexp.MustCompile("<title>.*?</title>")

	flag.BoolVar(&markdownOpt, "m", false, "URL to Markdown")
	flag.BoolVar(&reSTOpt, "r", false, "URL to reStructuredText")
	flag.BoolVar(&versionOpt, "version", false, "Show version")

	flag.Usage = func() {
		fmt.Printf(`NAME:
  url-anchor - URL to HTML <a> tag

USAGE:
  url-anchor [options] url

DESCRIPTION:
  url-anchor converts URL to HTML <a> tag.
  url-anchor also converts URL to Markdown or reStructuredText.

  e.g.
    $ url-anchor https://google.com
    <a href="https://google.com">Google</a>

OPTIONS:
`)
		flag.PrintDefaults()
	}
}

type cli struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

func (c *cli) run(args []string) int {
	os.Args = args
	flag.Parse()

	if versionOpt {
		fmt.Fprintf(c.outStream, "url-anchor version %s\n", version)
		return 0
	}

	url := flag.Arg(0)
	if len(url) == 0 {
		flag.Usage()
		return 1
	}

	if url != "-" {
		printFormatURL(url, c)
		return 0
	}

	scanner := bufio.NewScanner(c.inStream)
	for scanner.Scan() {
		url := scanner.Text()
		printFormatURL(url, c)
	}
	return 0
}

func printFormatURL(url string, c *cli) {
	title, err := fetchHTMLTitle(url)
	if err != nil {
		e := errors.Wrapf(err, "Failed to get HTML title")
		fmt.Fprintf(c.errStream, "%s\n", e.Error())
		return
	}

	s := formatURL(url, title)
	fmt.Fprintf(c.outStream, "%s\n", s)
}

func fetchHTMLTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to get %s", url)
	}
	defer resp.Body.Close()

	r := transformReader(resp.Header.Get("Content-Type"), resp.Body)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to read page")
	}

	return getTitleText(string(b)), nil
}

func transformReader(contentType string, r io.Reader) io.Reader {
	switch getCharsetName(contentType) {
	case "shift_jis", "sjis":
		return transform.NewReader(r, japanese.ShiftJIS.NewDecoder())
	case "euc-jp":
		return transform.NewReader(r, japanese.EUCJP.NewDecoder())
	default:
		br := bufio.NewReader(r)
		data, err := br.Peek(1024)
		if err != nil && err != io.EOF && err != bufio.ErrBufferFull {
			return r
		}

		enc, name, _ := charset.DetermineEncoding(data, contentType)
		if enc != nil {
			return enc.NewDecoder().Reader(br)
		}
		if name != "" {
			if enc := encoding.GetEncoding(name); enc != nil {
				return enc.NewDecoder().Reader(br)
			}
		}

		return r
	}
}

func getCharsetName(contentType string) string {
	if contentType == "" {
		return ""
	}

	ct := strings.ToLower(contentType)
	startIndex := strings.Index(contentType, "charset")
	if startIndex < 0 {
		return ""
	}

	charset := ct[startIndex+len("charset"):]
	return regexp.MustCompile("[= ]").ReplaceAllString(charset, "")
}

func getTitleText(text string) string {
	s := re.FindString(text)
	if len(s) == 0 {
		return ""
	}

	title := strings.TrimPrefix(
		strings.TrimSuffix(s, "</title>"),
		"<title>")

	return title
}

func formatURL(URL string, title string) string {
	switch {
	case markdownOpt:
		return fmt.Sprintf(`[%s](%s)`, title, URL)
	case reSTOpt:
		return fmt.Sprintf("`%s <%s>`_", title, URL)
	default:
		return fmt.Sprintf(`<a href="%s">%s</a>`, URL, title)
	}
}
