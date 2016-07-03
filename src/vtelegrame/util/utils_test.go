package util

import "testing"
import "github.com/stretchr/testify/assert"

func TestGetCommaSeparatedString(t *testing.T) {
	ints := []int{1, 2, 3}
	strings := []string{"1", "2", "3"}
	expectedString := "1,2,3"

	intArr := IntArray(ints)
	stringArr := StringArray(strings)
	actualIntsString := GetCommaSeparatedString(&intArr)
	actualStringsString := GetCommaSeparatedString(&stringArr)

	assert.Equal(t, expectedString, actualIntsString)
	assert.Equal(t, expectedString, actualStringsString)
}
