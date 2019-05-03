package main

import (
	"vision/core/models"
	"vision/core/fileDriver"
	"fmt"
)

func main() {
	request := &models.QueryHolder{"/home/deep/grep", "", "tail" ,2, "", "", ""}
	lines, err := fileDriver.FileDriver(request)
	fmt.Println(err)
	fmt.Println(lines)
}
