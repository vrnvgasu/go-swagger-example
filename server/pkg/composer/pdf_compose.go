package composer

import (
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/go-openapi/runtime"
	"github.com/jung-kurt/gofpdf"
)

func ComposeFromFiles(files []io.ReadCloser) (io.ReadCloser, error) {
	var opt gofpdf.ImageOptions
	pdf := gofpdf.New("P", "mm", "A4", "")

	for i, f := range files {
		if f == nil {
			continue
		}
		m, format, err := image.DecodeConfig(f)
		if err != nil {
			continue
		}
		fR, ok := f.(*runtime.File)
		if !ok {
			continue
		}
		_, err = fR.Data.Seek(0, 0)
		if err != nil {
			continue
		}
		weigh, height := calculateRotation(m.Width, m.Height)
		if weigh > 0 {
			pdf.AddPage()
		}
		_ = pdf.RegisterImageReader("file"+strconv.Itoa(i), format, f)
		pdf.ImageOptions("file"+strconv.Itoa(i), 0, 0, weigh, height, true, opt, 0, "")
		f.Close()
	}

	fb := []byte{}
	rw := bytes.NewBuffer(fb)

	err := pdf.Output(rw)
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(rw), nil
}

func calculateRotation(w, h int) (width, height float64) {
	x := float32(w)
	y := float32(h)
	if x >= y {
		return 210, 0
	}
	f := x / 100
	percent := y / f
	diff := percent - 100
	if diff > 41 {
		return 0, 280
	}

	return 210, 0
}
