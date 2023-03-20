package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	// png は先頭8バイトがシグネチャで固定バイト列
	file.Seek(8, 0)
	var offset int64 = 8
	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func textChunk(text string) io.Reader {
	byteText := []byte(text)
	crc := crc32.NewIEEE()
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteText)))
	writer := io.MultiWriter(&buffer, crc)
	io.WriteString(writer, "teXt")
	writer.Write(byteText)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func main() {
	file, err := os.Open("test.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("result.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")

	io.Copy(newFile, chunks[0])
	io.Copy(newFile, textChunk("Sincerely.sasakiy84"))
	for _, chunk := range chunks {
		dumpChunk(chunk)
		io.Copy(newFile, chunk)
	}
}
