//package main
//
//import (
	//"vision/core/models"
	//"fmt"
//)
//
//func main() {
	//test := &models.QueryHolder{"/var/log/apache2/access.log", "", "hea" ,1, "", "", ""}
	//isClean, err := test.Sanitise()
	//fmt.Println(err)
	//fmt.Println(isClean)
//}


package main

import (
    "errors"
    "flag"
    "fmt"
    "io"
    "os"
)

var (
    //DEFINE FLAG DEFAULTS
    filename = ""
    numLines = 10

    //ErrNoFilename is thrown when the path the the file to tail was not given
    ErrNoFilename = errors.New("You must provide the path to a file in the \"-file\" flag.")

    //ErrInvalidLineCount is thrown when the user provided 0 (zero) as the value for number of lines to tail
    ErrInvalidLineCount = errors.New("You cannot tail zero lines.")
)

//GET FLAGS FROM CLI
func init() {
    flag.StringVar(&filename, "file", "", "Filename to get last lines of text from.")
    flag.IntVar(&numLines, "n", numLines, "Number of lines to get from end of file.")
    flag.Parse()
}

//MAIN FUNCTIONALITY OF APP
//make sure filename (path to file) was given
//run it through the tailing function
//print the output to stdout
func main() {
    //TAIL
    text, err := GoTail(filename, numLines)
    if err != nil {
        fmt.Println(err)
        return
    }

    //DONE
    fmt.Print(text)
    return
}

//GoTail IS THE FUNCTION THAT ACTUALLY DOES THE "TAILING"
//this can be used this package is imported into other golang projects
func GoTail(filename string, numLines int) (string, error) {
    //MAKE SURE FILENAME IS GIVEN
    //actually, a path to the file
    if len(filename) == 0 {
        return "", ErrNoFilename
    }

    //MAKE SURE USER WANTS TO GET AT LEAST ONE LINE
    if numLines == 0 {
        return "", ErrInvalidLineCount
    }

    //OPEN FILE
    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()

    //SEEK BACKWARD CHARACTER BY CHARACTER ADDING UP NEW LINES
    //offset must start at "-1" otherwise we are already at the EOF
    //"-1" from numLines since we ignore "last" newline in a file
    numNewLines := 0
    var offset int64 = -1
    var finalReadStartPos int64
    for numNewLines <= numLines-1 {
        //seek to new position in file
        startPos, err := file.Seek(offset, 2)
        if err != nil {
            return "", err
        }

        //make sure start position can never be less than 0
        //aka, you cannot read from before the file starts
        if startPos == 0 {
            //set to -1 since we +1 to this below
            //the position will then start from the first character
            finalReadStartPos = -1
            break
        }

        //read the character at this position
        b := make([]byte, 1)
        _, err = file.ReadAt(b, startPos)
        if err != nil {
            return "", err
        }

        //ignore if first character being read is a newline
        if offset == int64(-1) && string(b) == "\n" {
            offset--
            continue
        }

        //if the character is a newline
        //add this to the number of lines read
        //and remember position in case we have reached our target number of lines
        if string(b) == "\n" {
            numNewLines++
            finalReadStartPos = startPos
        }

        //decrease offset for reading next character
        //remember, we are reading backward!
        offset--
    }

    //READ TO END OF FILE
    //add "1" here to move offset from the newline position to first character in line of text
    //this position should be the first character in the "first" line of data we want
    endPos, err := file.Seek(int64(-1), 2)
    if err != nil {
        return "", err
    }
    b := make([]byte, (endPos+1)-finalReadStartPos)
    _, err = file.ReadAt(b, finalReadStartPos+1)
    if err == io.EOF {
        return string(b), nil
    } else if err != nil {
        return "", err
    }

    //special case
    //if text is read, then err == io.EOF should hit
    //there should *never* not be an error above
    //so this line should never return
    return "**No error but no text read.**", nil
}
