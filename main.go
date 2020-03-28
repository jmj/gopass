package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/marcusolsson/tui-go"
)

type Password struct {
	Password string
	Name     string
	URL      string
}

var (
	passwords map[string]Password

	passgen bool
	print   string
)

func init() {
	passwords = make(map[string]Password)
}

func cleanup(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)

	for _, v := range passwords {
		encoder.Encode(v)
	}
}

func main() {
	flag.BoolVar(&passgen, "passgen", false, "Just generate a password and exit")
	flag.StringVar(&print, "print", "", "Print the password for the given name")
	flag.Parse()

	if passgen {
		c := &GenConfig{}

		pw, _ := GenPassword(c)
		fmt.Println(pw)
		os.Exit(0)
	}

	encdata, err := ioutil.ReadFile("passdb.dat")
	if err != nil {
		panic(err)
	}

	// Decrypt
	cleardata := encdata

	buf := bytes.NewBuffer(cleardata)
	decoder := gob.NewDecoder(buf)

	for {
		var holding Password
		e := decoder.Decode(&holding)
		if e != nil {
			if errors.Is(e, io.EOF) {
				break
			}
			panic(err)
		}

		passwords[holding.Name] = holding
	}

	defer func() {
		cleanup("passdb.dat")
	}()

	if print != "" {
		pw := GetPassword(print)
		fmt.Println(pw)
		os.Exit(0)
	}

	app, err := tui.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	mainWin := ui(app)

	FocusChain.SetActiveSet("main")
	app.SetFocusChain(&FocusChain)

	app.SetWidget(mainWin)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
