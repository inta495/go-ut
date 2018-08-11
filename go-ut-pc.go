// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "github.com/jacobsa/go-serial/serial"
import "time"

func main() {
	/* define option */
	options := serial.OpenOptions{
		PortName: "COM13",
		BaudRate: 9600,
		DataBits: 8,
    	StopBits: 1,
    	InterCharacterTimeout: 100,
    	MinimumReadSize: 4,
	}

	/* make connection */
	port, err := serial.Open(options)
	if err != nil {
		fmt.Println("serial.Open: %v", err)
		/*  open port */
		for{
			/* try open every 1 second */
			port, err = serial.Open(options)
			if err != nil {
				fmt.Println("serial.Open: %v", err)
			} else {
				break
			}
			time.Sleep(1000 * time.Millisecond)
		}
	} 	
	defer port.Close()

	buf := make([]byte, 8)
	buf[0] = 170
	buf[1] = 81
	buf[2] = 2
	buf[3] = 0
	buf[4] = 0
	buf[5] = 0
	buf[6] = 83
	buf[7] = 187

	fmt.Println(buf)
	port.Write(buf)

	//buf = make([]byte, 3)
	time.Sleep(200 * time.Millisecond)
	buf = make([]byte, 128)
	n, err := port.Read(buf)
	if err != nil {
		fmt.Println("Error reading from TELE: ", err)
	} else{
		buf = buf[:n]
		fmt.Println("Rx: ", buf)
	}

	fmt.Println("Hello, 世界")
	for{
		time.Sleep(10000 * time.Millisecond)
	}
}