package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"

	bufra "github.com/avvmoto/buf-readerat"
	"github.com/snabb/httpreaderat"
)

func main() {

	// client trace to log whether the request's underlying tcp connection was re-used
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Printf("reused %s conn? %v\n", info.Conn.RemoteAddr(), info.Reused)
		},
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	// ---

	url := "https://github.com/ogri-la/ziptest/releases/download/0.0.1/the-undermine-journal--5-6-20220531.zip"

	client := http.Client{}
	req, err := http.NewRequestWithContext(traceCtx, http.MethodGet, url, nil)
	if err != nil {
		panic("error creating request: " + err.Error())
	}

	// ---

	http_readerat, err := httpreaderat.New(&client, req, nil)
	if err != nil {
		panic("error creating HTTPReaderAt: " + err.Error())
	}

	buffered_http_readerat := bufra.NewBufReaderAt(http_readerat, 1024*1024)
	zip_reader, err := zip.NewReader(buffered_http_readerat, http_readerat.Size())
	if err != nil {
		panic("error creating zip.Reader: " + err.Error())
	}

	for _, f := range zip_reader.File {
		fmt.Println(f.Name)
		if f.Name == "TheUndermineJournal/TheUndermineJournal.toc" {
			fmt.Println("---")
			fh, err := f.Open()
			if err != nil {
				panic("error opening zipfile entry: " + err.Error())
			}

			bl, err := io.ReadAll(fh)
			if err != nil {
				panic("error reading bytes from zipfile entry: " + err.Error())
			}
			fmt.Println(string(bl))
			fmt.Println("---")
		}
	}
}
