package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/hypebeast/go-osc/osc"
)

var flagHost, flagAddress string
var flagPort int

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var flagInts, flagStrings, flagFloats arrayFlags

func init() {
	flag.StringVar(&flagHost, "host", "localhost", "osc host")
	flag.IntVar(&flagPort, "port", 57120, "port to use")
	flag.StringVar(&flagAddress, "addr", "/osc/address", "osc address")
	flag.Var(&flagInts, "i", "integer to send")
	flag.Var(&flagFloats, "f", "float to send")
	flag.Var(&flagStrings, "s", "integer to send")
}

func main() {
	flag.Parse()
	err := run()
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}

func run() (err error) {
	client := osc.NewClient(flagHost, flagPort)
	msg := osc.NewMessage(flagAddress)
	for _, u := range flagInts {
		v, err := strconv.Atoi(u)
		if err == nil {
			msg.Append(int32(v))
		} else {
			return err
		}
	}
	for _, u := range flagFloats {
		v, err := strconv.ParseFloat(u, 64)
		if err == nil {
			msg.Append(v)
		} else {
			return err
		}
	}
	for _, u := range flagStrings {
		msg.Append(u)
	}
	err = client.Send(msg)
	return
}
