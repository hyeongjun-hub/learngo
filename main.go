package main // compile 하고싶으면 무조건 main

import (
	"fmt"
	"github.com/hyeongjun-hub/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}

	baseWord := "hello"
	//definition := "Greeting"

	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}

	//definition2, _ := dictionary.Search(baseWord)
	//fmt.Println(definition2)

	dictionary.Delete(baseWord)
	//err = dictionary.Delete("baseWord")
	if err != nil {
		fmt.Println(err)
	}

	definition3, err := dictionary.Search(baseWord)
	fmt.Println(definition3)
	if err != nil {
		fmt.Println(err)
	}
}
