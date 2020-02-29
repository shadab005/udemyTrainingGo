package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type person struct {
	First       string
	Last        string `json:"-"`     // - to exclude field from marshalling
	Age         int    `json:"myAge"` //creating another alias for the field
	notexported int    //lowercase are not marshalled to jsom
}

//like a class method. the one in the bracket after func specifies the receiver.
func (p person) structMethod() string {
	return "Hello " + p.First + " " + p.Last
}

func main() {

	fmt.Println("--------struct----------")
	p1 := person{"ankit", "dewan", 26, 8}
	//	fmt.Println(p1)
	fmt.Println(p1.structMethod())

	fmt.Println("--------Marshalling----------")
	//marhsalling : converting the object to json byte slice
	x, _ := json.Marshal(p1)
	fmt.Println(x)
	fmt.Println(string(x))

	//unmarshalling : converting back the byte slice to object

	fmt.Println("--------Unmarshalling----------")

	bs := []byte(`{"First" : "Marie", "Last" : "Curie", "Age" : 60, "myAge" : 30}`)
	var p2 person
	json.Unmarshal(bs, &p2)
	fmt.Println(p2)

	fmt.Println("--------Encoding----------")
	//encoding : converting the object to byte slice  and writing to some stream. Writer
	json.NewEncoder(os.Stdout).Encode(p1)

	fmt.Println("--------Decoding----------")
	//decoding : read from stream byte array and convert it into object
	var p3 person
	rdr := strings.NewReader(`{"First":"ankit","myAge":26}`)
	json.NewDecoder(rdr).Decode(&p3)
	fmt.Println(p3)

}
