package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
	"wellpro/wp2028"
)

func main() {
	flag.Usage = func() {
		log.Println("usage: go run . port baudrate slave_id")
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) != 3 {
		flag.Usage()
		log.Fatal("flag must be 3")
	}

	port := flag.Args()[0]
	baud := flag.Args()[1]
	id := flag.Args()[2]
	intbaud, _ := strconv.Atoi(baud)
	intid, _ := strconv.Atoi(id)
	baudrate := uint(intbaud)
	slave_id := uint8(intid)

	log.Printf("port     : %s", port)
	log.Printf("baudrate : %v", baudrate)
	log.Printf("slave_id : %v", slave_id)
	fmt.Println()

	for {
		//fmt.Println()
		data, err := wp2028.ReadDigital(port, baudrate, slave_id)
		if err != nil {
			log.Fatalf("FAILED ReadDigital:%s", err.Error())
		}
		//log.Printf("data   : %v", data)

		bit64L := asBits(uint64(data))
		var binary [8]uint8
		for i, bit := range bit64L {
			if i >= 8 {
				binary[i-8] = uint8(bit)
			}
		}
		//log.Printf("binary : %v", binary)
		//time.Sleep(1 * time.Second)

		err = wp2028.WriteDigital(port, baudrate, slave_id, binary)
		if err != nil {
			log.Fatalf("FAILED WriteDigital:%s", err.Error())
		}

		// err = wp2028.WriteShifting(port, baudrate, slave_id)
		// if err != nil {
		// 	log.Fatalf("FAILED WriteShifting:%s", err.Error())
		// }

		time.Sleep(1 * time.Second)
	}

}

func asBits(val uint64) []uint64 {
	var bits = []uint64{}
	for i := 0; i < 16; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}
