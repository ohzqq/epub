package epub

import (
	"testing"
)

const (
	testEpub = "testdata/meta.epub"
)

func TestOpen(t *testing.T) {
	book, err := Open(testEpub)
	if err != nil {
		t.Fatal("error opening epub:", err)
	}
	defer book.Close()

	if book.Mimetype != EpubMimeType {
		t.Error("unexpected mime type:", book.Mimetype)
	}

	assertNotNil(t, book.Container, "failed to parse book container details")
	assertNotNil(t, book.Opf, "failed to parse book OPF details")
	assertNotNil(t, book.Ncx, "failed to parse book NCX details")
}

func assertNotNil(t *testing.T, v interface{}, msg string) {
	if v == nil {
		t.Error(msg)
	}
}
