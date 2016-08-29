package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	// read tag
	header := make([]byte, 10)
	file, _ := os.Open("sample.mp3")
	defer file.Close()
	file.Read(header)
	//fmt.Printf("bytes: %q\n", header)

	// id3v2 version
	fmt.Printf("ID3 v2.%d.%d\n", header[3], header[4])

	// flags
	fmt.Printf("flags = %d\n", header[5])

	// size
	size := synchSafeInt(header[6:10])
	fmt.Printf("size = %d\n", size)

	// extended header -- TODO
	//xheader := make([]byte, 0)
	//count, _ = file.Read(xheader)
	//fmt.Printf("read %d bytes: %q\n", count, xheader[:100])

	// extended header -- TODO
	xheader := make([]byte, 0)
	file.Read(xheader)
	//fmt.Printf("read %d bytes: %q\n", count, xheader[:100])

	// frames
	// frame id
	// frame size
	// flags

	var (
		fname     []byte
		fsize     []byte
		frameSize uint32
		fflag     []byte
		frame     []byte
		str       string
	)

	byteLoc := 10
	//for i := 0; i < 12; i++ {
	for {
		fname = make([]byte, 4)
		file.Read(fname)
		if binary.BigEndian.Uint32(fname) == 0 {
			rem := size - byteLoc - 4
			buffer := make([]byte, rem)
			file.Read(buffer)
			fmt.Println(buffer)

			break
		} else {
			fmt.Println(fname)
		}

		fsize = make([]byte, 4)
		file.Read(fsize)

		fflag = make([]byte, 2)
		file.Read(fflag)

		frameSize = binary.BigEndian.Uint32(fsize)
		frame = make([]byte, frameSize)
		file.Read(frame)

		// apic
		//  00 02 8D 47 00 00 00
		fmt.Printf("\n@%7d [%s] size = %d; flags = %d\n", byteLoc, fname, frameSize, fflag)
		byteLoc += 10 + int(frameSize)

		str = string(fname)
		if str == "APIC" {
			//func (f *File) Write(b []byte) (n int, err error)
			fmt.Printf(">>>%s, %d, \n", str, str == "APIC")
			filex, _ := os.Create("./output.jpg")
			defer filex.Close()
			// frame[0] = Text encoding
			// frame[1:10] = mime
			// frame[11] = x00
			// frame[12] = picture type        | 03
			// frame[13] = description and x00 | 00
			// frame[14...] = FF D8...
			fmt.Printf("%s, %#v\n", frame[1:11], frame[14:24])
			n, _ := filex.Write(frame[14:])
			fmt.Printf("wrote %d bytes", n)
		}
	}

	music := make([]byte, 100)
	file.Read(music)
	fmt.Println(music)

	//whatsup := make([]byte, 20)
	//file.Read(whatsup)
	//fmt.Printf("\n@%7d %#v\n", byteLoc, whatsup)
}

// big endian synch safe int
// input [0 10 44 110]
// 10 * 128^2 + 44 * 128 + 110
// output 169582
func synchSafeInt(buf []byte) (r int) {
	r = 0
	for _, v := range buf {
		r <<= 7
		r += (int(v) % 128)
	}

	return r
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
