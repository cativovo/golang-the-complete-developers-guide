package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg, input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")

	if len(components) != 3 {
		return Time{}, &TimeParseError{msg: "Invalid number of time components", input: input}
	} else {
		hour, err := strconv.Atoi(components[0])

		if err != nil {
			return Time{}, &TimeParseError{msg: fmt.Sprintf("Error parsing hour: %v\n", err), input: input}
		}

		minute, err := strconv.Atoi(components[1])

		if err != nil {
			return Time{}, &TimeParseError{msg: fmt.Sprintf("Error parsing minute: %v\n", err), input: input}
		}

		second, err := strconv.Atoi(components[2])

		if err != nil {
			return Time{}, &TimeParseError{msg: fmt.Sprintf("Error parsing seconds: %v\n", err), input: input}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{msg: "hour out of range: 0 <= hour <= 23", input: fmt.Sprintf("%v", hour)}
		}

		if minute > 59 || minute < 0 {
			return Time{}, &TimeParseError{msg: "minute out of range: 0 <= minute <= 59", input: fmt.Sprintf("%v", minute)}
		}

		if second > 59 || second < 0 {
			return Time{}, &TimeParseError{msg: "seconds out of range: 0 <= seconds <= 59", input: fmt.Sprintf("%v", second)}
		}

		return Time{hour: hour, minute: minute, second: second}, nil
	}
}
