// say replicates the OS X say command using festival as a backend.
package main

import (
        "bufio"
        "fmt"
        "bitbucket.org/kisom/gofestival"
        "io"
        "path/filepath"
        "os"
        "strings"
)

var text string
var readStdin bool

func main() {
        config()
        if !readStdin {
                err := festival.Speak(text)
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
        } else {
                reader := bufio.NewReader(os.Stdin)
                for {
                        line, _, err := reader.ReadLine()
                        if err == io.EOF {
                                break
                        }
                        text = string(line)
                        festival.Speak(text)
                }
        }
}

func usage() {
        bin := filepath.Base(os.Args[0])
        fmt.Println("speak text from the command line.")
        fmt.Printf("\nusage: %s [-] <text>\n", bin)
        fmt.Printf("    if - is passed, %s will read from the commandline\n",
                bin)
        fmt.Printf("    otherwise, speak the text from the command line\n")
}

func config() {
        if len(os.Args) == 1 {
                usage()
                os.Exit(1)
        }
        if os.Args[1] == "-" {
                readStdin = true
        } else {
                text = strings.Join(os.Args[1:], " ")
        }
}
