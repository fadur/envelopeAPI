package main

type Person struct {
	ID   string `json:"id"`
	Name string `json:name"`
	Age  string `json:age"`
}

func all() []Person {
	return []Person{
		Person{"1", "Bob", "33"},
		Person{"2", "Mike", "30"},
		Person{"3", "Ann", "22"},
	}
}

func GetPeople() []Person {
	return all()
}
