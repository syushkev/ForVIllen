package main

import (
	"fmt"
	"regexp"
	"errors"
)

const defaultREconstant  = `(\d+)\.(.+)`

func FindVersion123(version , expression string) (r []string, err error) {
	f := flag.NewFlagSet("flag", flag.ExitOnError)
	versio := f.String("version", version, "Some description")
	re := f.String("re", defaultREconstant, "Some description")

	if expression == "" {
	f.Parse([]string{defaultREconstant})
	}

	if expression != ""{
	f.Parse([]string{"-re",expression})
	}
	req, err := regexp.Compile(*re)


	if err != nil {
		return nil, err
	}
	r = req.FindStringSubmatch(*versio)
	if len(r) == 0 {
		return nil, errors.New("invalid version")
	}
	r = r[1:]
	return r, err

}

func main() {

	one, err := FindVersion123("10.20.3.5", "")
	fmt.Println(one, "", err)
	fmt.Println()
	four, err := FindVersion123("10.20.4.5", `(\d+\.\d+)\.(.+)`)
	fmt.Println(four, "", err)
	five, err := FindVersion123("10.20.4.5.54.65.5.5.5.5.5.5.5.5.5", `(\d+\.\d+)\.(.+)`)
	fmt.Println(five, "", err)

}

