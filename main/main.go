package main

import (
	"vision/core/models"
	"vision/core/fileDriver"
	"vision/core/models"
	"fmt"
	"vision/api"
)

func main() {
	request := &models.QueryHolder{"/home/deep/hhtt", "", "head" ,2, "", "", ""}
	lines, err := fileDriver.FileDriver(request)
	fmt.Println(err)
	fmt.Println(lines)

	api.Api()
}
