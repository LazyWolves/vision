package main

import (
	"vision/core/models"
	"vision/core/fileHandler"
	"fmt"
)

func main() {
	test := &models.QueryHolder{"/var/log/apache2/access.log", "", "hea" ,1, "", "", ""}
	isClean, err := test.Sanitise()
	fmt.Println(err)
	fmt.Println(isClean)
	line, err := fileHandler.ReadFromHead("/home/deep/grep", 2)
	fmt.Println(line)
}
