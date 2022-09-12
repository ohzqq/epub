package epub

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const EpubMimeType = "application/epub+zip"

// Book provides the epub book contents
type Book struct {
	Ncx        Ncx        `json:"ncx"`
	Opf        Opf        `json:"opf"`
	Container  Container  `json:"-"`
	Mimetype   string     `json:"-"`
	Encryption Encryption `json:"-"`

	fd        *zip.Reader
	closer    io.Closer
	directory string
}

//Open a resource file.
func (p *Book) Open(n string) (io.ReadCloser, error) {
	return p.open(p.filename(n))
}

//Files returns the list of resource files
func (p *Book) Files() []string {
	var fns []string
	for _, f := range p.fd.File {
		fns = append(fns, f.Name)
	}
	return fns
}

//Each provides an iterator over each book section passing its
//title and contents
func (p *Book) Each(fn func(string, io.ReadCloser)) error {
	for _, point := range p.Ncx.Points {
		xhtml, err := p.Open(point.Content.Src)
		if err != nil {
			return err
		}
		fn(point.Text, xhtml)
		xhtml.Close()
	}
	return nil
}

//Close file reader
func (p *Book) Close() {
	if p.closer != nil {
		p.closer.Close()
	}
}

//-----------------------------------------------------------------------------
func (p *Book) filename(n string) string {
	return path.Join(path.Dir(p.Container.Rootfile.Path), n)
}

func (p *Book) readXML(n string, v interface{}) error {
	fd, err := p.open(n)
	if err != nil {
		return nil
	}
	defer fd.Close()
	dec := xml.NewDecoder(fd)
	return dec.Decode(v)
}

func (p *Book) readBytes(n string) ([]byte, error) {
	fd, err := p.open(n)
	if err != nil {
		return nil, nil
	}
	defer fd.Close()

	return ioutil.ReadAll(fd)
}

func (p *Book) open(n string) (io.ReadCloser, error) {
	if p.directory != "" {
		filename := path.Join(path.Dir(p.directory+string(filepath.Separator)), n)
		return os.Open(filename)
	}

	for _, f := range p.fd.File {
		if f.Name == n {
			return f.Open()
		}
	}
	return nil, fmt.Errorf("file %s not exist", n)
}

// ZipReader returns the internal file descriptor
func (p *Book) ZipReader() *zip.Reader {
	return p.fd
}

// GetSMIL parses and returns the SMIL structure
func (p *Book) GetSMIL(resource string) (*SMIL, error) {
	var smil SMIL
	err := p.readXML(resource, &smil)

	return &smil, err
}
