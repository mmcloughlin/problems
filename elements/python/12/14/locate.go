package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// FindUpper16 reads IPs from reader and finds a value of the upper 16 bits
// that is has fewer than 2^16 IPs under it.
func FindUpper16(r io.Reader) uint32 {
	counts := make([]int, 1<<16)
	var x uint32
	for {
		err := binary.Read(r, binary.BigEndian, &x)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		counts[x>>16]++
	}
	for x = 0; x < uint32(1)<<16; x++ {
		if counts[x] < 1<<16 {
			return x
		}
	}
	panic("seen all ips")
}

// FindMissingWithHi reads the file again and looks for a missing IP with
// given hi bits.
func FindMissingWithHi(r io.Reader, hi uint32) uint32 {
	seen := make([]bool, 1<<16)
	var x uint32
	for {
		err := binary.Read(r, binary.BigEndian, &x)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if x>>16 == hi {
			seen[x&0xff] = true
		}
	}

	for x = 0; x < uint32(1)<<16; x++ {
		if !seen[x] {
			return (hi << 16) | x
		}
	}

	panic("none")
}

var (
	filename = flag.String("ips", "", "IPs file")
)

func main() {
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hi := FindUpper16(f)
	f.Seek(0, io.SeekStart)

	ip := FindMissingWithHi(f, hi)

	fmt.Println(ip)
}
