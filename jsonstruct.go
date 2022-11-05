package main

import (
	"encoding/json"
	"fmt"
)

type BodyInfo struct {
	Height uint64 `json:"height"`
	Weight uint64 `json:"weight"`
}

type JsonGroup struct {
	Address     uint64 `json:"myAddress"`
	Phonenumber uint64 `json:"myPhone"`

	//Bdi BodyInfo `json:"BodyInfo"`
}

func testJsonStruct() {
	jsgroup := JsonGroup{Address: 123098129038190, Phonenumber: 19238019209}
	jsString, err := json.Marshal(jsgroup)
	fmt.Println(string(jsString))
	jsGroupUnmarshalled := &JsonGroup{}
	errUnmarshal := json.Unmarshal(jsString, jsGroupUnmarshalled)
	if errUnmarshal != nil {
		fmt.Println("unMarshal error happens")
	}
	fmt.Println(jsGroupUnmarshalled)
	fmt.Println(err)
}

func main() {
	testJsonStruct()
}
