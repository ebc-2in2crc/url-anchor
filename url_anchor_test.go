package main

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"

	"github.com/jarcoal/httpmock"
)

const testingURL = "https://www.testingURL.co.jp"

func TestRun(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder("GET", testingURL,
		func(request *http.Request) (response *http.Response, err error) {
			resp := httpmock.NewBytesResponse(200, []byte("<title>タイトル</title>"))
			return resp, nil
		})

	params := []struct {
		args     []string
		stdin    string
		exitCode int
		expect   string
	}{
		{args: []string{"", testingURL}, exitCode: 0, expect: `<a href="https://www.testingURL.co.jp">タイトル</a>`},
		{args: []string{"", "-m", testingURL}, exitCode: 0, expect: `[タイトル](https://www.testingURL.co.jp)`},
		{args: []string{"", "-r", testingURL}, exitCode: 0, expect: "`タイトル <https://www.testingURL.co.jp>`_"},
		{args: []string{"", "-r", "-"}, stdin: testingURL, exitCode: 0, expect: "`タイトル <https://www.testingURL.co.jp>`_"},
	}

	for _, p := range params {
		resetFlag()
		c, outStream, _ := newCLI(p.stdin)

		exitCode := c.run(p.args)

		if exitCode != p.exitCode {
			t.Errorf("run(%s): Output = %q; want %q", p.args, exitCode, p.exitCode)
		}

		expect := p.expect + "\n"
		if outStream.String() != expect {
			t.Errorf("run(%s): Output = %q; want %q", p.args, outStream.String(), expect)
		}
	}
}

func resetFlag() {
	markdownOpt = false
	reSTOpt = false
}

func newCLI(stdin string) (c cli, outStream, errStream *bytes.Buffer) {
	inStream := bytes.NewBufferString(stdin)
	outStream = bytes.NewBuffer(nil)
	errStream = bytes.NewBuffer(nil)
	c = cli{inStream: inStream, outStream: outStream, errStream: errStream}
	return
}

func TestFetchHTMLTitle(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()

	httpmock.RegisterResponder("GET", `https://hoge.com`,
		func(request *http.Request) (response *http.Response, err error) {
			resp := httpmock.NewBytesResponse(200, []byte("<title>タイトル</title>"))
			return resp, nil
		})

	params := []struct {
		url    string
		expect string
	}{
		{url: "https://hoge.com", expect: "タイトル"},
	}

	for _, p := range params {
		actual, err := fetchHTMLTitle(p.url)
		if err != nil {
			t.Errorf("got: %v\nwant: nil", err)
		}

		if actual != p.expect {
			t.Errorf("fetchHTMLTitle(%s): Output = %q; want %q", p.url, actual, p.expect)
		}
	}
}

func TestTransformReader(t *testing.T) {
	params := []struct {
		ct      string
		encoder *encoding.Encoder
	}{
		{ct: "text/html;", encoder: nil},
		{ct: "text/html; charset=Shift_JIS", encoder: japanese.ShiftJIS.NewEncoder()},
		{ct: "text/html; charset=sjis", encoder: japanese.ShiftJIS.NewEncoder()},
		{ct: "text/html; charset=EUC-JP", encoder: japanese.EUCJP.NewEncoder()},
	}

	for _, p := range params {
		expect := []byte("テキスト")
		input, err := encode(expect, p.encoder)
		if err != nil {
			t.Errorf("got: %v\nwant: nil", err)
		}

		r := transformReader(p.ct, -1, bytes.NewReader(input))
		actual := make([]byte, len(expect))
		_, err = r.Read(actual)
		if err != nil {
			t.Errorf("got: %v\nwant: nil", err)
		}

		if reflect.DeepEqual(actual, expect) == false {
			t.Errorf("transformReader(%s): Output = %q; want %q", p.ct, actual, input)
		}
	}
}

func encode(b []byte, encoder *encoding.Encoder) ([]byte, error) {
	if encoder == nil {
		return b, nil
	}
	return encoder.Bytes(b)
}

func TestFormatURL(t *testing.T) {
	params := []struct {
		m      bool
		r      bool
		expect string
	}{
		{m: true, expect: "[title](URL)"},
		{r: true, expect: "`title <URL>`_"},
	}

	url := "URL"
	title := "title"
	for _, p := range params {
		markdownOpt = p.m
		reSTOpt = p.r
		actual := formatURL(url, title)
		if actual != p.expect {
			t.Errorf("formatURL(%s, %s): Output = %q; want %q", url, title, actual, p.expect)
		}
	}
}
