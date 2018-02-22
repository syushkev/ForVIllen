package main

import (
	"fmt"
	"regexp"
	"errors"
)

const defaultREconstant  = `(\d+)\.(.+)`

func FindVersion123(version , expression string) (r []string, err error) {

	req, err := regexp.Compile(expression)

	if err != nil {
		return nil, err
	}
	if version == ""{
		return nil, errors.New("empty field version")
	}

	r = req.FindStringSubmatch(version)
	if len(r) == 0 {
		return nil, errors.New("invalid version")
	}
	r = r[1:]
	return r, err

}

func main() {
	versio := flag.String("version", "", "Some description")
	re := flag.String("re", defaultREconstant, "Some description")
	flag.Parse()
	one, err := FindVersion123(*versio, *re)
	fmt.Println(one, "", err)
}
