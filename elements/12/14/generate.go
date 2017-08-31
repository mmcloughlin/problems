package main

import (
	"encoding/binary"
	"os"
)

func main() {
	x := uint64(1)
	for i := uint64(0); i>>16 == 0; i++ {
		x = 134775813*x + 1
		err := binary.Write(os.Stdout, binary.BigEndian, uint32(x))
		if err != nil {
			panic(err)
		}
	}
}
