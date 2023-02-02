package monitor

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestCpuTemperature(t *testing.T) {
	temp, err := getCPUTemperature()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("CPU temperature: %.2f°C\n", temp)
}
func getCPUTemperature() (float64, error) {
	file, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	re := regexp.MustCompile("[0-9]+")
	match := re.FindString(text)
	temp, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, err
	}
	return temp / 1000, nil
}
