package exercises

import "fmt"

func DailyTemperatures(temperatures []int) []int {
	answer := []int{}
	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] && j > i {
				answer = append(answer, j-i)
				break
			} else if j == len(temperatures)-1 && temperatures[i] >= temperatures[j] {
				answer = append(answer, 0)
				continue
			} else {
				continue
			}
		}
	}
	answer = append(answer, 0)
	return answer
}

func DailyTemperatures2(temperatures []int) []int {
	answer := []int{}

	i := 0
	j := 1

	for i < len(temperatures)-1 {
		if temperatures[j] > temperatures[i] && j > i {
			answer = append(answer, j-i)
			i++
			j = i + 1

		} else if j == len(temperatures)-1 && temperatures[i] >= temperatures[j] {
			answer = append(answer, 0)
			i++
			j = i + 1
		} else {
			j++
		}
	}
	answer = append(answer, 0)
	return answer
}

func TestTemps() {
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	t2 := []int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}
	fmt.Println(DailyTemperatures2(t))
	fmt.Println(DailyTemperatures2(t2))
	fmt.Println("[3,1,1,2,1,1,0,1,1,0]")

}
