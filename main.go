package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

func strlen(s string, c chan int) {
	c <- len(s)
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func errorHandler() error {
	return errors.New("something bad happened here")
}

type Sonion struct {
	Bar string
	Baz string
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>

	go test(100)
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
	}
}
