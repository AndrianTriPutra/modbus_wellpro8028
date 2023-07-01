package wp2028

import (
	"errors"
	"time"

	"github.com/simonvetter/modbus"
)

func ReadDigital(port string, baudrate uint, id uint8) (uint16, error) {
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
		return 0, newErr
	}

	err = client.Open()
	if err != nil {
		newErr := errors.New("ERROR E2:" + err.Error())
		return 0, newErr
	}
	defer client.Close()
	client.SetUnitId(id)

	data, err := client.ReadRegister(10, modbus.INPUT_REGISTER)
	data, err = client.ReadRegister(10, modbus.INPUT_REGISTER)
	if err != nil {
		newErr := errors.New("ERROR TO:" + err.Error())
		return 0, newErr
	}

	return data, nil
}
