package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := []string{"01:45", "06:27", "10:18", "11:11", "07:30", "09:00", "09:59", "12:00", "12:11", "02:19", "04:20", "08:04", "01:08", "12:30"}
	timesToAngles(input)

}

func timesToAngles(times []string) {
	for _, v := range times {
		hour := getHour(v)
		minute := getMinute(v)
		angle := getAngle(hour, minute)

		fmt.Printf("%v has an outer angle of %v° and an inner angle of %v°\n", v, angle, (360 - angle))

	}
}

func getHour(strTime string) float64 {
	strHr := strTime[0:2]
	hour, _ := strconv.ParseFloat(strHr, 64)

	return hour
}

func getMinute(strTime string) float64 {
	strMn := strTime[3:]
	minute, _ := strconv.ParseFloat(strMn, 64)

	return minute / 5
}

func getAngle(h, m float64) float64 {
	diff := h - m
	angle := math.Round(math.Abs((diff * 30)))


	return angle
}
