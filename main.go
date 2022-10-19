package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

func main() {
	littleEnd, bigEnd, numBytes := hexToEndian("ff00000000000000000000000000000000000000000000000000000000000000")
	endianToHex(bigEnd, littleEnd, numBytes)
}

func hexToEndian(hexNum string) (*big.Int, *big.Int, int) { // буде виводити hex to little endian і hex to big endian
	sliceByteNum, _ := hex.DecodeString(hexNum)
	numBytes := len(sliceByteNum)
	fmt.Println("Number of bytes:", numBytes)

	revSlice := []byte{}
	for i := range sliceByteNum {
		revSlice = append(revSlice, sliceByteNum[len(sliceByteNum)-1-i])
	}

	littleEndian := new(big.Int)
	littleEndian = littleEndian.SetBytes(revSlice)
	fmt.Println("little endian :", littleEndian)

	bigEndian := new(big.Int)
	bigEndian, _ = bigEndian.SetString(hexNum, 16)
	fmt.Println("big endian:", bigEndian)
	return littleEndian, bigEndian, numBytes

}

func endianToHex(bigEndian *big.Int, littleEndian *big.Int, numBytes int) {
	fmt.Printf("big endian to hex    %X\n", bigEndian)
	hexLittleEnd := fmt.Sprintf("%X\n", littleEndian)
	sliceByteNum, _ := hex.DecodeString(hexLittleEnd)
	revSlice := []byte{}
	for i := range sliceByteNum {
		revSlice = append(revSlice, sliceByteNum[len(sliceByteNum)-1-i])
	}
	hexNum := fmt.Sprintf("%X", revSlice)
	if numBytes != len(revSlice) {
		notEnoughBytes := (numBytes - len(revSlice)) * 2
		for i := 0; i < notEnoughBytes; i++ {
			hexNum += "0"
		}
	}
	fmt.Println("little endian to hex", hexNum)

}
