package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	//open the .log file, or create one
	logfile, err := os.OpenFile("lg.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening/creating:", err)
		return
	}
	defer logfile.Close()

	//create a writer with a timestamp(already declared outside func main)
	logwriter := NewTimeStampWriter(logfile) //using the func to write in content into the struct

	//Lv2 plus
	twO := NewTimeStampWriter(os.Stdout)
	mulw := io.MultiWriter(twO, logwriter)

	//simulate user operation
	fmt.Fprintf(mulw, "user login. ") //user login
	time.Sleep(2 * time.Second)       //sleep..
	fmt.Fprintf(mulw, "user conducted operation A. ")
	time.Sleep(1 * time.Second)
	fmt.Fprintf(mulw, "User conducted operation B. ")
	time.Sleep(10 * time.Second)
	fmt.Fprintf(mulw, "User logout. ")

}

//code below is the key part(main points of lesson03), try to understand entirely of each.

type TimeStampWriter struct {
	logFile io.Writer //implement the interface 'Writer' in io package
}

func NewTimeStampWriter(lf io.Writer) *TimeStampWriter {
	return &TimeStampWriter{
		logFile: lf, //similar to the func 'NewCommon' in lesson03/LvX
	}
}

// a method towards *TimeStampWriter
func (tw *TimeStampWriter) Write(p []byte) (n int, err error) {
	wt := time.Now()
	ts := wt.Unix()
	t := fmt.Sprintf("at time: %s, timestamp: %d  \n", wt, ts)
	tw.logFile.Write(append(p, []byte(t)...)) //first get in tw, then write in p
	return len(p), nil
}
