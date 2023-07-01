# WellPro_Testing

## modpoll
- ./modpoll -t3:int -b 9600 -p none -m rtu -a 1 -r 10 -c 1 /dev/ttyUSB0

## getting started
- go run . port baudrate slave_id
- example : go run . /dev/ttyUSB0 9600 1

## reference
- https://github.com/microrobotics/wellpro