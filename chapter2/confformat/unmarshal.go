package confformat

import "fmt"

const (
	exampleTOML = `name="Example1"
age=99
    `

	exampleJSON = `{"name":"Example2","age":98}`

	exampleYAML = `name: Example3
age: 97    
    `
)

// UnmarshalAll takes data in various formats
// and converts them into structs
func UnmarshalAll() error {
	t := TOMLData{}
	j := JSONData{}
	y := YAMLData{}

	if _, err := t.Decode([]byte(exampleTOML)); err != nil {
		return err
	}
	fmt.Println("TOML Unmarshal =", t)

	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal =", j)

	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}
	fmt.Println("Yaml Unmarshal =", y)
	return nil
}
