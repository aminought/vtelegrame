package util

import "strconv"

// IntArray is alias for int array
type IntArray []int

// StringArray is alias for string array
type StringArray []string

type genArray interface {
	Get(i int) string
	Len() int
}

// Get returns value by index
func (arr *IntArray) Get(i int) string {
	return strconv.Itoa((*arr)[i])
}

// Len returns length of array
func (arr *IntArray) Len() int {
	return len([]int(*arr))
}

// Get returns value by index
func (arr *StringArray) Get(i int) string {
	return (*arr)[i]
}

// Len returns length of array
func (arr *StringArray) Len() int {
	return len([]string(*arr))
}

// GetCommaSeparatedString joins strings with comma separator
func GetCommaSeparatedString(arr genArray) string {
	var str string
	len := arr.Len()
	for i := 0; i < len; i++ {
		str += arr.Get(i)
		if i < len-1 {
			str += ","
		}
	}
	return str
}
