package epub

import (
	"io"
	"testing"
)

func TestEpub(t *testing.T) {
	book, err := Open("testdata/meta.epub")
	if err != nil {
		t.Fatal(err)
	}
	defer book.Close()

	t.Run("Files", func(t *testing.T) {
		expected := 17
		files := book.Files()
		if len(files) != expected {
			t.Error("unexepected number of files returned:", len(files))
		}
	})

	t.Run("Each", func(t *testing.T) {
		count := 0
		expectedSections := 8
		book.Each(func(title string, xhtml io.ReadCloser) {
			count++
			assertNotNil(t, xhtml, "failed to read section contents")
		})
		if count != expectedSections {
			t.Error("unexpected number of sections:", count)
		}
	})
}
