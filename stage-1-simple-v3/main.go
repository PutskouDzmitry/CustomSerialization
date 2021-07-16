package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type IntStringMap map[int]string

func (m *IntStringMap) MarshalJSON() ([]byte, error) {
	ss := map[string]string{}
	for k, v := range *m {
		i := strconv.Itoa(k)
		ss[i] = v
	}
	return json.Marshal(ss)
}

func (m *IntStringMap) UnmarshalJSON(data []byte ) error {
	ss := map[string]string{}
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}
	for k, v := range ss {
		i, err := strconv.Atoi(k)
		if err != nil {
			return err
		}
		(*m)[i] = v
	}
	return nil
}

func main () {
	m := IntStringMap{4: "four", 5: "five"}
	data, err := m.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("IntStringMap to JSON: ", string(data))

	m = IntStringMap{}

	jsonString := []byte("{\"1\": \"one\", \"2\": \"two\"}")
	m.UnmarshalJSON(jsonString)

	fmt.Printf("IntStringMap from JSON: %v\n", m)
	fmt.Println("m[1]:", m[1], "m[2]:", m[2])
}
