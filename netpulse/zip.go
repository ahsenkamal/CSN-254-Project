package netpulse

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func zip(in []byte) []byte {
	out := new(bytes.Buffer)
	writer := gzip.NewWriter(out)
	writer.Write(in)
	writer.Close()
	return out.Bytes()
}

func unzip(in []byte) ([]byte, error) {
	var out []byte
	reader, err := gzip.NewReader(bytes.NewBuffer(in))
	if err != nil {
		return out, newError(errUnzip, err.Error())
	}
	reader.Close()
	out, err = ioutil.ReadAll(reader)
	if err != nil {
		return out, newError(errUnzipRead, err.Error())
	}
	return out, nil
}
