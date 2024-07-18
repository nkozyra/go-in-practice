package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{ 
"firstName": "Jean", 
"lastName": "Bartik", 
"age": 86, 
"education": [ 
     { 
            "institution": "Northwest Missouri State Teachers College", 
            "degree": "Bachelor of Science in Mathematics" 
     }, 
     {  
            "institution": "University of Pennsylvania", 
            "degree": "Masters in English" 
     } 
], 
"spouse": "William Bartik", 
"children": [ 
     "Timothy John Bartik", 
     "Jane Helen Bartik", 
     "Mary Ruth Bartik" 
]  
}`)

func main() {
	var f interface{}             // # B
	err := json.Unmarshal(ks, &f) // # C
	if err != nil {               // # D
		fmt.Println(err) // # D
		os.Exit(1)       // # D
	}
	fmt.Println(f) // # E
}
