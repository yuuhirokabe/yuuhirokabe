package main

import (
	"encoding/json"
	"fmt"
	"log"
)


type Linux struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Narukoshin struct {
	Name        string   `json:"name"`
	Age         int      `json:"age"`
	Hobbies     []string `json:"hobbies"`
	Languages   []string `json:"languages"`
	Programming []string `json:"programming_languages"`
	Linux       []Linux  `json:"linux"`
}

func main() {
	naru := Narukoshin{
		Name: "narukoshin",
		Age:  20,
		Hobbies: []string{
			"👨‍💻 writting the code",
			"😾 hacking the systems",
			"🤳 photographying",
			"🎹 music production",
		},
		Languages: []string{
			"latvian",
			"russian",
			"english",
			"japanese",
		},
		Programming: []string{
			"PHP",
			"JavaScript",
			"Python",
			"C++",
			"Go",
		},
		Linux: []Linux{
			{Name: "BlackArch", Desc: "for penetration testing"},
			{Name: "Kali", Desc: "alternative for penetration testing"},
			{Name: "Ubuntu", Desc: "for penetration testing"},
		},
	}
	json, err := json.Marshal(&naru);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json));
}