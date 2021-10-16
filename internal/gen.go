package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	buf, err := os.ReadFile("gen.tmpl")
	if err != nil {
		log.Fatal("reading template file:", err)
	}

	intTmpl := template.Must(template.New("numeric").
		Parse(string(buf)))

	for i := GenTypeInvalid + 1; i < GenTypeEndTypes; i++ {
		b := &bytes.Buffer{}
		data := struct {
			GenType GenType
		}{
			GenType: i,
		}
		if err = intTmpl.Execute(b, data); err != nil {
			log.Fatal("executing template:", err)
		}

		if buf, err = format.Source(b.Bytes()); err != nil {
			log.Fatal("formatting output:", err)
		}

		if err = writeFile(bytes.NewReader(buf), i); err != nil {
			log.Fatal("writing go file:", err)
		}
	}
}

func writeFile(data io.Reader, genType GenType) error {
	lowerGenType := genType.StringLower()
	filename := fmt.Sprintf("../%sset/%s_set.go",lowerGenType, lowerGenType)

	if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
		return err
	}

	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, data); err != nil {
		return err
	}

	return nil
}
