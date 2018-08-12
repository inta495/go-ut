// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "github.com/jacobsa/go-serial/serial"
import "time"
//import "bufio"
import "os"
import "os/exec"
//import "strings"
import "flag"
import "github.com/fatih/color"
import "github.com/labstack/echo"

func xor(b []byte) []byte {
	
	b = append(b,0)
	LEN := len(b)

	for i := 0; i < LEN-1; i++ {
		b[LEN-1] = b[i]^b[LEN-1]
 	}
	return b
}

func main() {
	/* XOR testing */
	b1 := make([]byte,5)
	b1[0] = 67
	b1[1] = 11
	b1[2] = 0
	b1[3] = 11
	b1[4] = 12
	fmt.Println(xor(b1))

	color.Yellow("Prints text in cyan.")
	/* define */
	portname   := flag.String("port", "", "COMx on windows or ttyUSBx on linux")
	baudrate   := flag.Int("baudrate", 9600, "baudrate")
	filenumber := flag.Int("file", -1, "file number")
	loop       := flag.Int("loop", -1, "update data every x second")
	
	/* clear terminal */
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
    cmd.Run()
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
	e := echo.New()
	e.Static("/", "page")
    // Start as a web server
    e.Start(":8000")

	/* for fun */
	fmt.Println("Hello, 世界")
	for{
		time.Sleep(10000 * time.Millisecond)
	}
}