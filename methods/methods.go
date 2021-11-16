package methods

import (
	stre "DB_In_Memory/structures"
	"encoding/json"
	"io/ioutil"
)

func GetAllDB() []stre.Course {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	result := []stre.Course{}

	err = json.Unmarshal(file, &result)
	if err != nil {
		panic(err)
	}

	return result
}
func PostCourseDB(c stre.Course) bool {
	array := GetAllDB()
	var ban bool
	c.Id = len(array) + 1

	array = append(array, c)
	result, err := json.Marshal(array)
	if err != nil {
		panic(err)

	}
	err = ioutil.WriteFile("data.json", result, 0644)
	if err != nil {
		panic(err)

	}
	ban = true
	return ban
}
