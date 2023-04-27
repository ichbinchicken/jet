package helper

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

func AnyToString(any interface{}) string {
	anyString, ok := any.(string)
	if !ok {
		panic("failure on casting to string from any type")
	}
	return anyString
}

func Compress(raw []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	compressor, err := zlib.NewWriterLevel(buf, zlib.BestSpeed)
	if err != nil {
		return nil, err
	}

	_, _ = compressor.Write(raw) // error is propagated through Close
	err = compressor.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Decompress(input []byte) ([]byte, error) {
	decompressor, err := zlib.NewReader(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}
	defer decompressor.Close()

	var out bytes.Buffer
	if _, err = out.ReadFrom(decompressor); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func ReadStdin() string {
	reader := bufio.NewReader(os.Stdin)
	var result string
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		result += input
	}
	return result
}
