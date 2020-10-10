package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	var files = []struct{
		Name, Body string
	}{
		{"main1.go",  "file1"},
		{"main2.go",  "file2"},
	}
	for _, file := range files{
		hdr := &tar.Header{
			Name:file.Name,
			Mode:0660,
			Size:int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)

	for{
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}

}