package main

import (
	"fmt"
	"math/rand"
	"time"
)

type color []string

type cubedesc struct {
	color  string
	areaid int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	colors := color{
		"red",
		"gre",
		"blu"}

	var columns = 5
	var rows = 5
	var areaid = 1

	areas := makeareas(columns, rows)
	box := makebox(columns, rows)

	fillbox(box, colors)
	showbox(box)

	for i := range box {
		for j := range box[i] {

			up := i - 1
			left := j - 1

			cur := box[i][j]

			if left >= 0 && cur == box[i][left] {
				areas[i][j] = areas[i][left]
			} else if up >= 0 && cur == box[up][j] {
				areas[i][j] = areas[up][j]
			} else {
				areas[i][j] = areaid
				areaid++
			}

			if left >= 0 && up >= 0 &&
				box[i][left] == box[up][j] &&
				box[i][left] == box[i][j] &&
				areas[i][left] != areas[up][j] {
				changeAreas(areas[i][left], areas[up][j], j, i, areas)
			}
		}
	}

	showareas(areas)

	fmt.Println("")
	fmt.Println(findmax(box, areas))
}

func changeAreas(oldAreaId int, newAreaId int, colEnd int, rowEnd int, areas [][]int) {
	for i := 0; i <= rowEnd; i++ {
		for j := 0; j <= colEnd; j++ {
			if areas[i][j] == oldAreaId {
				areas[i][j] = newAreaId
			}
		}
	}
}

func randColor(colors []string) string {
	return colors[rand.Intn(len(colors))]
}

func showbox(box [][]string) {
	for i := range box {
		for j := range box[i] {
			fmt.Print(box[i][j] + " ")
		}
		fmt.Println("")
	}
}

func showareas(areas [][]int) {
	fmt.Println("")
	for i := range areas {
		for j := range areas[i] {
			val := fmt.Sprintf("%03d", areas[i][j])
			fmt.Print(val + " ")
		}
		fmt.Println("")
	}
}

func makebox(columns int, rows int) [][]string {
	box := make([][]string, rows)
	for i := range box {
		box[i] = make([]string, columns)
	}
	return box
}

func makeareas(columns int, rows int) [][]int {
	areas := make([][]int, rows)
	for i := range areas {
		areas[i] = make([]int, columns)
	}
	return areas
}

func fillbox(box [][]string, colors []string) {
	for i := range box {
		for j := range box[i] {
			box[i][j] = randColor(colors)
		}
	}
}

func findmax(box [][]string, areas [][]int) (cubedesc, int) {
	var cubes = make(map[cubedesc]int)
	for i := range areas {
		for j := range areas[i] {
			cube := cubedesc{
				color:  box[i][j],
				areaid: areas[i][j],
			}

			if _, ok := cubes[cube]; ok {
				cubes[cube]++
			} else {
				cubes[cube] = 1
			}
		}
	}

	var retcube cubedesc
	var max = 0
	for k, v := range cubes {
		if v > max {
			retcube = k
			max = v
		}
	}
	return retcube, max
}
