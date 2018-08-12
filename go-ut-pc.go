// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "github.com/jacobsa/go-serial/serial"
import "time"
//import "bufio"
import "os"
//import "strings"
import "flag"
import "github.com/fatih/color"

func main() {
	color.Yellow("Prints text in cyan.")
	/* define */
	portname   := flag.String("port", "", "COMx on windows or ttyUSBx on linux")
	baudrate   := flag.Int("baudrate", 9600, "baudrate")
	filenumber := flag.Int("file", -1, "file number")
	loop       := flag.Int("loop", -1, "update data every x second")

	/* parse */
	flag.Parse()
	if len(*portname) == 0 {
		fmt.Println("Invalid portname")
		os.Exit(0)
	}

	/* print status */
	color.Yellow("portname  : %s", *portname)
	color.Yellow("baudrate  : %d", *baudrate)
	color.Yellow("filenumber: %d", *filenumber)
	color.Yellow("loop      : %d", *loop)

	/* define option */
	options := serial.OpenOptions{
		PortName: *portname,
		BaudRate: uint(*baudrate),
		DataBits: 8,
		StopBits: 1,
		InterCharacterTimeout: 100,
		MinimumReadSize: 1,
	}

	/* make connection */
	port, err := serial.Open(options)
	if err != nil {
		fmt.Println("serial.Open: %v", err)
		/*  open port */
		for i:=0 ; i<30 ; i++{
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

	/* Ping */
	buf := make([]byte, 2)
	buf = []byte{116, 10}
	fmt.Println(buf)
	port.Write(buf)
	/* wait for response */
	time.Sleep(200 * time.Millisecond)
	buf = make([]byte, 8)
	n, err := port.Read(buf)
	if err != nil {
		fmt.Println("Error reading from TELE: ", err)
	} else{
		buf = buf[:n]
		fmt.Println("Rx: ", buf)
	}

	/* check response */
	if buf[0] == 0 {
		
	}

	/* who are you */
	buf = make([]byte, 8)
	buf = []byte{116, 10}
	fmt.Println(buf)
	port.Write(buf)
	/* wait for response */
	time.Sleep(200 * time.Millisecond)
	buf = make([]byte, 16)
	n, err = port.Read(buf)
	if err != nil {
		fmt.Println("Error reading from TELE: ", err)
	} else{
		buf = buf[:n]
		fmt.Println("Rx: ", buf)
	}

	/* check response */
	if buf[0] == 0 {
		
	}

	/* file size */
	buf = make([]byte, 8)
	buf = []byte{116, 10}
	fmt.Println(buf)
	port.Write(buf)
	/* wait for response */
	time.Sleep(200 * time.Millisecond)
	buf = make([]byte, 16)
	n, err = port.Read(buf)
	if err != nil {
		fmt.Println("Error reading from TELE: ", err)
	} else{
		buf = buf[:n]
		fmt.Println("Rx: ", buf)
	}

	/* check response */
	if buf[0] == 0 {
		
	}

	if *loop >= 1 {
		go func() {
			for{						
				/* call file */
				//TODO
				time.Sleep(10000 * time.Millisecond)
			}
		}()

		go func () {
			for{
				/* receive data */
				//TODO

				/* channal back to global variable */
				//TODO
			}
		}()
	}
	
	/* serve web */
	//TODO


	/* for fun */
	fmt.Println("Hello, 世界")
	for{
		time.Sleep(10000 * time.Millisecond)
	}
}