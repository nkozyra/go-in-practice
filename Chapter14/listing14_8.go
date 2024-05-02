package main
import (
     "encoding/json"
     "fmt"
)
type Name struct {
     First string `json:"firstName"` 
     Last  string `json:"lastName "` 
}
func main() {
     n := &Name{"Inigo", "Montoya"}
     data, _ := json.Marshal(n) 
     fmt.Printf("%s\n", data)B 
}
