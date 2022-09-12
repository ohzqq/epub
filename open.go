package epub

import (
	"archive/zip"
	"io"
)

//Open open a epub file
func Open(fn string) (*Book, error) {
	fd, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}

	bk, err := bookFromZipReader(&fd.Reader)
	if err != nil {
		_ = fd.Close()
		return nil, err
	}

	bk.closer = fd
	return bk, nil
}

func bookFromZipReader(zr *zip.Reader) (*Book, error) {
	bk := Book{fd: zr}

	mt, err := bk.readBytes("mimetype")
	if err != nil {
		return nil, err
	}

	if err == nil {
		bk.Mimetype = string(mt)
		err = bk.readXML("META-INF/container.xml", &bk.Container)
	}
	if err == nil {
		err = bk.readXML(bk.Container.Rootfile.Path, &bk.Opf)
	}
	if err == nil {
		err = bk.readXML("META-INF/encryption.xml", &bk.Encryption)
	}

	for _, mf := range bk.Opf.Manifest {
		if mf.ID == bk.Opf.Spine.Toc {
			err = bk.readXML(bk.filename(mf.Href), &bk.Ncx)
			break
		}
	}

	return &bk, nil
}

func Read(r io.ReaderAt, size int64) (*Book, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, err
	}

	return bookFromZipReader(zr)
}

//OpenDir opens an OPF file
func OpenDir(fn string) (*Book, error) {
	bk := Book{directory: fn}
	mt, err := bk.readBytes("mimetype")
	if err == nil {
		bk.Mimetype = string(mt)
		err = bk.readXML("META-INF/container.xml", &bk.Container)
	}
	if err == nil {
		err = bk.readXML(bk.Container.Rootfile.Path, &bk.Opf)
	}
	if err == nil {
		err = bk.readXML("META-INF/encryption.xml", &bk.Encryption)
	}

	for _, mf := range bk.Opf.Manifest {
		if mf.ID == bk.Opf.Spine.Toc {
			err = bk.readXML(bk.filename(mf.Href), &bk.Ncx)
			break
		}
	}

	return &bk, err
}
