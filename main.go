package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	_ "github.com/stacktitan/ldapauth"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

/*
		Books I'm using:
		- Security With Go
	    - Black Hat Go

		Also, helpful kommands
		- go fmt (reformatting)
		- golint (reports mistakes) (go get -u golang.org/x/lint/golint)
		- go vet (identifies issues with code)
*/
func test(num int) {
	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", num/i)
	}
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func errorHandler(er any) error {
	return errors.New("something bad happened here, check this out: " + er.(string))
}

func strlen(s string, c chan int) {
	c <- len(s)
}

var ip = "192.168.3.219"
var minPort = 1
var maxPort = 1024
var sshPort = 22
var RDPPort = 455

func testTCPConnection(ip string, port int, doneChan chan bool, file *os.File) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second)
	if err == nil {
		log.Printf("Port %d: Open\n", port)

		// Write bytes to the file
		byteSlicePortFile := []byte("Open Ports: " + strconv.Itoa(port) + "\n")
		bytesWrittenPortFile, err := file.Write(byteSlicePortFile)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Port Bytes written:", bytesWrittenPortFile)
	}
	doneChan <- true
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>

	/*go test(100)
	time.Sleep(2 * time.Second)

	c := make(chan int)
	go strlen("Hai", c)
	go strlen("Goodbai", c)
	x, y := <-c, <-c
	fmt.Println(x, y, x*y)
	time.Sleep(4 * time.Second)

	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	//FYI, later look into cross-compilation & a build cmdlet that works on Windows
	// Look at ch. on files in Black Hat Go
	fmt.Printf("Hello, Black Hat Gophers!\n")

	f := Sonion{"Kris D", "Deltarune"}
	b, _ := json.Marshal(f)
	fmt.Println(string(b))
	json.Unmarshal(b, &f)

	if err := errorHandler(); err != nil {
		fmt.Println(err)
	}*/

	//									   //
	// This is where the real stuff begins //
	//									   //

	//TIP PT 1. Create + Drop file (test)

	// Create file
	newFile, err := os.Create("drop_here/test.txt")

	if err != nil {
		err := errorHandler(err)
		if err != nil {
			return
		}
		log.Fatal(err)
	}
	log.Println(newFile)
	defer newFile.Close()

	// Write bytes to the file
	byteSlice := []byte("Add + Edit\nPooped out this file")
	bytesWritten, err := newFile.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bytes written:", bytesWritten)

	// Getting file info
	fileInfo, err := os.Stat("check_here/killswitch.txt")
	if err != nil {
		err := errorHandler(err)
		if err != nil {
			return
		}
		log.Fatal(err)
	}

	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("File size: ", fileInfo.Size())
	fmt.Println("File Last Modified: ", fileInfo.ModTime())

	// Checking if file exists
	fileInfo2, err := os.Stat("check_here/onswitch.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	log.Println("File exists, File Name:", fileInfo2.Name())
	//log.Println(fileInfo)

	// Writing to a file
	file1, err := os.OpenFile("check_here/edit.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	// Write bytes to the file
	byteSlice2 := []byte("Wrote bytes to here\nWorm wuz here")
	bytesWritten2, err := file1.Write(byteSlice2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bytes written:", bytesWritten2)

	//TIP PT 2a. Scan ports

	// Btw, the stuff down here I can't test because my laptop is being a chud
	doneChannel := make(chan bool)
	// Write open ports to file
	portFile, err := os.Create("drop_here/ports.txt")
	if err != nil {
		err := errorHandler(err)
		if err != nil {
			return
		}
		log.Fatal(err)
	}
	log.Println(portFile)
	defer portFile.Close()

	for port := minPort; port <= maxPort; port++ {
		go testTCPConnection(ip, port, doneChannel, portFile)
	}

	// Wait for routines to finish
	for port := minPort; port <= maxPort; port++ {
		<-doneChannel
	}

	//TIP PT 2b. Duplication + Send

	//TIP PT 3. SSH/RDP into ports

	//TIP PT 4a. Read files + Functionality/Determination

	//TIP PT 4b. Modify + Drop files
}
