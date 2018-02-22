package main

import (
	"fmt"
	"regexp"
	"errors"
)

const defaultREconstant  = `(\d+)\.(.+)`

func parceFlags(version, expression string)  (version1, expression1 string){
	f := flag.NewFlagSet("flag", flag.ExitOnError)
	versio := f.String("version", version, "Some description")
	re := f.String("re", defaultREconstant, "Some description")

	if expression == "" {
		f.Parse([]string{"-re", defaultREconstant, "version", version})
	}

	if expression != ""{
		f.Parse([]string{"-re",expression, "version", version})
	}
	return *versio, *re
}

func FindVersion123(version , expression string) (r []string, err error) {


	req, err := regexp.Compile(expression)


	if err != nil {
		return nil, err
	}
	r = req.FindStringSubmatch(version)
	if len(r) == 0 {
		return nil, errors.New("invalid version")
	}
	r = r[1:]
	return r, err

}

func main() {
	v,re := parceFlags("10.20.3.5", "")
	one, err := FindVersion123(v, re)
	fmt.Println(one, "", err)
	fmt.Println()
	v2,re2 := parceFlags("10.20.4.5", `(\d+\.\d+)\.(.+)`)
	one2, err2 := FindVersion123(v2, re2)
	fmt.Println(one2, "", err2)
	fmt.Println()
	v3,re3 := parceFlags("10.20.4.5.54.65.5.5.5.5.5.5.5.5.5", `(\d+\.\d+)\.(.+)`)
	one3, err3 := FindVersion123(v3, re3)
	fmt.Println(one3, "", err3)

}

