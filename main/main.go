package main

import (
	"fmt"
	"regexp"
	"errors"
)

func FindVersion123(version ...string) (r []string, err error){
	if version[1] == "" {
		version[1] = `(\d+)\.(.+)`
	}
	req, err := regexp.Compile(version[1])
	if err != nil {
		return nil, err
	}
	r = req.FindStringSubmatch(version[0])
	if len(r) == 0 {
		return nil, errors.New("invalid version")
	}
	r = r[1:]
	return 	r, err


}

func main() {
	//var validID = regexp.MustCompile(`(\d+\.\d+)\.(.+)`)
	//var validID2 = regexp.MustCompile(`(\d+)\.(.+)`)
	one, err :=FindVersion123("10.20.3.5", "1vjh")
	fmt.Println(one, "", err)
	fmt.Println()
	four, err := FindVersion123("10.20.3.5", `(\d+\.\d+)\.(.+)`)
	fmt.Println(four, "", err)


}
