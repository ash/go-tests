package main

import "os"
import "flag"

var omitNewline = flag.Bool("n", false, "do not print final newline")

const Space = " "
const Newline = "\n"

func main() {
    flag.Parse()
    
    var s string = ""
    for i := 0; i < flag.NArg(); i++ {
	if i > 0 {
	    s += Space
	}
	s += flag.Arg(i)
    }
    
    if !*omitNewline {
	s += Newline
    }
    
    os.Stdout.WriteString(s)
}
