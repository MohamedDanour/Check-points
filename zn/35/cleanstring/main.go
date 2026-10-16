// IMPORTANT: 
// A function gets its input from whoever calls it in code. 
// A program gets its input from the command line.
// Everything after go run . is an argument

// ALMOST ALMOST LIKE FIRSTWORD, BUT DONT BREAK WHEN A WORD IS completed and collected

import "fmt"

func main() {
    if len(os.Args) != 2 {
        fmt.Println() // ny rad
        return
    }

    fmt.Println(CleanString(os.Args[1])) // kör funktionen CleanString
}

func CleanString(s string) string {
    // gå igenom s, samla ord avgränsade av ' ' och '\t',
    // sätt ihop dem med exakt ett mellanslag
}


