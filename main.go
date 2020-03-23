package main

import "fmt"

type position struct {
	i   int
	j   int
	val int
}

func printTable(s [][]int) {
	for i := 0; i <= 8; i++ {
		for k := 0; k <= 8; k++ {
			if k%3 == 0 && k != 0 {
				fmt.Print(" ")
			}
			fmt.Print(s[i][k])
		}
		if i%3 == 2 && i != 0 {
			fmt.Println("")
		}
		fmt.Println("")

	}
}

func findEmpty(s [][]int) []position {
	tempPos := []position{}
	for i := 0; i < 9; i++ {
		for k := 0; k < 9; k++ {
			if s[i][k] == 0 {
				tempInd := position{
					i:   i,
					j:   k,
					val: 0,
				}
				tempPos = append(tempPos, tempInd)
				//fmt.Println(pos)
				//now in pos there are values of empty cell positions
			}
		}
	}
	return tempPos
}

func main() {
	row1 := []int{0, 0, 3, 0, 6, 1, 2, 9, 0}
	row2 := []int{8, 0, 6, 0, 0, 9, 7, 0, 1}
	row3 := []int{0, 4, 9, 0, 7, 5, 0, 3, 0}
	row4 := []int{0, 5, 2, 0, 0, 0, 1, 7, 3}
	row5 := []int{3, 0, 0, 9, 0, 2, 0, 4, 6}
	row6 := []int{6, 1, 0, 3, 5, 7, 0, 0, 0}
	row7 := []int{7, 6, 1, 0, 9, 0, 4, 0, 0}
	row8 := []int{0, 3, 0, 7, 0, 6, 0, 1, 5}
	row9 := []int{2, 0, 0, 1, 4, 0, 3, 0, 7}

	table := [][]int{row1, row2, row3, row4, row5, row6, row7, row8, row9}
	for {
		//Print table to screen
		printTable(table)

		zeros := findEmpty(table)
		fmt.Println(len(findEmpty(table)))

		fmt.Println(findValues(table, zeros))
		temp := findValues(table, zeros)
		for _, v := range temp {
			table[v.i][v.j] = v.val
		}
		if len(temp) == 0 {
			break
		}
	}
	printTable(table)
}

func findValues(s [][]int, zeros []position) []position {
	found := []position{}
	ret := []position{}
	for _, v := range zeros {
		found = nil
		for val := 1; val < 10; val++ {
			isFound := false
			for i := 0; i < 9; i++ {
				if s[i][v.j] == val {
					isFound = true
				}
			}
			if !isFound {
				for i := 0; i < 9; i++ {
					if s[v.i][i] == val {
						isFound = true
					}
				}
				if !isFound {
					//buraya kucuk kare içerisinde arama
					smallSquarei := []int{}
					smallSquarej := []int{}

					switch v.i {
					case 0, 1, 2:
						smallSquarei = []int{0, 1, 2}
						switch v.j {
						case 0, 1, 2:
							smallSquarej = []int{0, 1, 2}
						case 3, 4, 5:
							smallSquarej = []int{3, 4, 5}
						case 6, 7, 8:
							smallSquarej = []int{6, 7, 8}
						}
					case 3, 4, 5:
						smallSquarei = []int{3, 4, 5}
						switch v.j {
						case 0, 1, 2:
							smallSquarej = []int{0, 1, 2}
						case 3, 4, 5:
							smallSquarej = []int{3, 4, 5}
						case 6, 7, 8:
							smallSquarej = []int{6, 7, 8}
						}
					case 6, 7, 8:
						smallSquarei = []int{6, 7, 8}
						switch v.j {
						case 0, 1, 2:
							smallSquarej = []int{0, 1, 2}
						case 3, 4, 5:
							smallSquarej = []int{3, 4, 5}
						case 6, 7, 8:
							smallSquarej = []int{6, 7, 8}
						}
					}

					//fmt.Println(v.i, v.j, val, smallSquarei, smallSquarej)

					for _, l := range smallSquarei {
						for _, m := range smallSquarej {
							if s[l][m] == val {
								isFound = true
							}
						}
					}
					if !isFound {
						temp := position{
							i:   v.i,
							j:   v.j,
							val: val,
						}
						found = append(found, temp)
						//fmt.Println(v.i, v.j, val)
					}
				}
			}
		}
		if len(found) == 1 {
			ret = append(ret, found[0])
			//fmt.Println(found)
		}
	}
	return ret
}
