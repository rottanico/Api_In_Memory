package methods

import (
	stre "DB_In_Memory/structures"
	"encoding/json"
	"io/ioutil"
)

func BinarySearch(s []stre.Course, k int) int {
	lo, hi := 0, len(s)-1
	for lo <= hi {
		m := (lo + hi) >> 1
		if s[m].Id < k {
			lo = m + 1
		} else if s[m].Id > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

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
	key := array[len(array)-1].Id + 1
	if c.Id == 0 || c.Id > len(array) {
		if i := BinarySearch(array, key); i == -1 {
			c.Id = key
		} else {
			c.Id = key + 1
		}

	}

	for _, value := range array {
		if (value.Name == c.Name) && (value.Duration == c.Duration) {
			return false
		}
	}
	array = append(array, c)
	result, err := json.Marshal(array)
	if err != nil {
		panic(err)

	}
	err = ioutil.WriteFile("data.json", result, 0644)
	if err != nil {
		panic(err)

	}

	return true
}
func DeleteCourseDB(key int) bool {
	array := GetAllDB()

	if key > 0 && key < len(array)+1 {
		i := BinarySearch(array, key)
		if i != -1 {
			array = append(array[:i], array[i+1:]...)

			result, err := json.Marshal(array)
			if err != nil {
				panic(err)

			}
			err = ioutil.WriteFile("data.json", result, 0644)
			if err != nil {
				panic(err)

			}
			return true
		}

	}

	return false
}
func GetCourseDB(key int) (stre.Course, bool) {
	array := GetAllDB()
	if key > 0 && key < len(array)+1 {
		i := BinarySearch(array, key)
		if i != -1 {
			return array[i], true
		}
	}

	ret := stre.Course{}
	return ret, false
}
