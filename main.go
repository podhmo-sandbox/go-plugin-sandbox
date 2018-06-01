package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"plugin"
)

// Greeter :
type Greeter interface {
	Greet()
}

var (
	modeOpt = flag.String("m", "hello", "mode(hello or bye)")
)

func main() {
	flag.Parse()
	mod := fmt.Sprintf("./%s/greeter.so", *modeOpt)
	if err := run(mod); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run(mod string) error {
	p, err := plugin.Open(mod)
	if err != nil {
		return err
	}

	lookup, err := p.Lookup("Greeter")
	if err != nil {
		return err
	}

	g, ok := lookup.(Greeter)
	if !ok {
		return errors.New("unexpected type from module symbol")
	}

	g.Greet()
	return nil
}
