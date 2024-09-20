package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main(){
    flag.Parse()
    f := flag.Arg(0)
    f2 := flag.Arg(1)

    if (f == "" || f2 == ""){
        log.Fatal("Need two files to work")
    }

    // To get all the different lines of each file we use a 2 time loop
    // First loop gets the lines f2 has that f doesn't
    // Second loop gets the lines f has that f2 doesn't
    for i := 0; i < 2; i++ {
        var first string;
        var sec string;
        if (i == 0){
            first = f;
            sec = f2;
        }else {
            first = f2;
            sec = f;
        }

        lines := make(map[string]bool)
        // Open the first file and set all of it's lines as true in the map
        r, err := os.Open(first)
        if (err != nil){
            log.Fatal("Error opening file")
        }

        sc := bufio.NewScanner(r);

        for sc.Scan() {
            lines[sc.Text()] = true;
        }
        r.Close()

        // Open the second file and loop through the lines
        r, err = os.Open(sec)
        if (err != nil){
            log.Fatal("Error opening file")
        }

        sc = bufio.NewScanner(r);

        // If the line isn't true in the map then it is different from the first file
        for sc.Scan() {
            t := sc.Text()
            if (!lines[t]){
                fmt.Println(t)
                lines[t] = true;
            }
        }
        r.Close()
    }
}
