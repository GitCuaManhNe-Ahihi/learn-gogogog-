package main
import (
    "fmt"
    "log"
    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    log.SetPrefix("greetings:")
    log.SetFlags(0)
    messages,err := greetings.Hellos([]string{"Gladys","Mạnh"})
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(messages)
}