package main
import (
     "encoding/json"
     "fmt"
)
type Name struct {
     First string `json:"firstName"` // # A
     Last  string `json:"lastName "` // # A
}
func main() {
     n := &Name{"Inigo", "Montoya"}
     data, _ := json.Marshal(n) // # B
     fmt.Printf("%s\n", data)B // # B
}
