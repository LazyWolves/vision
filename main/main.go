package main

import (
	"vision/core/models"
	"fmt"
)

func main() {
	test := &models.QueryHolder{"/var/log/apache2", "", "" ,0, "", "", ""}
	isClean, err := test.Sanitise()
	fmt.Println(err)
	fmt.Println(isClean)
}
