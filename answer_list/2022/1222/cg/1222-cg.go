// ## Weekly Coding Problem

// ### Week 06, 2022 - (Easy)

// This problem was asked by **Microsoft**.

// Given a clock time in 12 hour ```hh:mm``` format, determine, to the nearest degree, the angle between the hour and the minute hands.

// Bonus: When, during the course of a day, will the angle be zero?

// ```time = [01:45, 06:27, 10:18, 11:11, 07:30, 09:00, 09:59, 12:00, 12:11, 02:19, 04:20, 08:04, 01:08]```

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	times := []string{"01:45", "06:27", "10:18", "11:11", "07:30", "09:00", "09:59", "12:00", "12:11", "02:19", "04:20", "08:04", "01:08"}
	for _, t := range times {
		angle := GetAngle(t)
		outStr := fmt.Sprintf("%s:%5d", t, angle)

		if angle == 360 {
			outStr = outStr + " (angle == 0)"
		}

		fmt.Println(outStr)
	}
}

func GetAngle(t string) int {

	hourC := 360 / 12
	minC := 360 / 60

	tSplit := strings.Split(t, ":")
	hour, _ := strconv.Atoi(tSplit[0])
	minute, _ := strconv.Atoi(tSplit[1])

	hourProgression := math.Floor((float64(minute) / 60) * 5)

	hourD := hour*hourC + int(hourProgression)
	minD := minute * minC

	return int(math.Abs(float64(minD) - float64(hourD)))
}
