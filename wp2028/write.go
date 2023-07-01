package wp2028

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/simonvetter/modbus"
)

func WriteDigital(port string, baudrate uint, id uint8, value [8]uint8) error {
	url := "rtu://" + port
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      url,
		Speed:    baudrate,
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 1,
		Timeout:  100 * time.Millisecond,
	})

	if err != nil {
		newErr := errors.New("ERROR E1:" + err.Error())
		return newErr
	}

	err = client.Open()
	if err != nil {
		newErr := errors.New("ERROR E2:" + err.Error())
		return newErr
	}
	defer client.Close()
	client.SetUnitId(id)

	var rotate [8]uint8
	for i := 0; i < 8; i++ {
		rotate[i] = value[7-i]
	}
	//log.Printf("rotate : %v", rotate)

	for i := 0; i < 8; i++ {
		if rotate[i] == 1 {
			err = client.WriteCoil(uint16(i), true)
			if err != nil {
				newErr := errors.New("ERROR E3A:" + err.Error())
				return newErr
			}
		} else {
			err = client.WriteCoil(uint16(i), false)
			if err != nil {
				newErr := errors.New("ERROR E3B:" + err.Error())
				return newErr
			}
		}
	}
	return nil
}

func WriteShifting(port string, baudrate uint, id uint8) error {
	url := "rtu://" + port
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      url,
		Speed:    baudrate,
		DataBits: 8,
		Parity:   modbus.PARITY_NONE,
		StopBits: 1,
		Timeout:  100 * time.Millisecond,
	})

	if err != nil {
		newErr := errors.New("ERROR E1:" + err.Error())
		return newErr
	}

	err = client.Open()
	if err != nil {
		newErr := errors.New("ERROR E2:" + err.Error())
		return newErr
	}
	defer client.Close()
	client.SetUnitId(id)

	for i := 0; i <= 7; i++ {
		log.Printf("TURN ON pin %v", i)
		err = client.WriteCoil(uint16(i), true)
		if err != nil {
			log.Printf("ERROR E3A:%s", err.Error())
			// newErr := errors.New("ERROR E3:" + err.Error())
			// return newErr
		}
		time.Sleep(1 * time.Second)

		log.Printf("TURN OFF pin %v", i)
		err = client.WriteCoil(uint16(i), false)
		if err != nil {
			log.Printf("ERROR E3B:%s", err.Error())
			// newErr := errors.New("ERROR E3:" + err.Error())
			// return newErr
		}

		fmt.Println()
	}

	return nil
}
