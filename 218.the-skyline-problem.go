/*
 * @lc app=leetcode id=218 lang=golang
 *
 * [218] The Skyline Problem
 */
package main

import (
	utils "reply2future.com/utils"
	"sort"
)

// @lc code=start
type Point struct {
	X int
	Y int
}

func getSkyline(buildings [][]int) [][]int {
	var leftPoints [][]int
	var rightPoints [][]int
	for _, building := range buildings {
		// right point which x < building[0] to leftpoints
		i := 0
		for _, point := range rightPoints {
			if point[0] >= building[0] {
				break
			}
			i += 1
			leftPoints = append(leftPoints, point)
		}
		rightPoints = rightPoints[i:]

		if len(rightPoints) == 0 {
			leftPoints = append(leftPoints, []int{building[0], building[2]})
			rightPoints = append(rightPoints, []int{building[1], 0})
			continue
		}

		closestLeftPoint := []int{leftPoints[len(leftPoints)-1][0], leftPoints[len(leftPoints)-1][1]} 
		closestRightPoint := []int{rightPoints[0][0], rightPoints[0][1]}

		// cover left point
		// TODO: =?
		coverLeft := closestLeftPoint[1] >= building[2]

		coverIndices := make([]int, 0)
		for i, point := range rightPoints {
			if point[0] > building[1] || point[1] > building[2] {
				break
			}
			coverIndices = append(coverIndices, i)
		}

		cStartIndex := coverIndices[0]
		cEndIndex := coverIndices[len(coverIndices)-1]
		rightPoints = append(rightPoints[:cStartIndex], rightPoints[cEndIndex+1:]...)	

		// left side
		var intersection Point 
		if closestLeftPoint[0] == building[0] {
			if closestLeftPoint[1] < building[2] {
				leftPoints[len(leftPoints)-1][1] = building[2]

			} else {

			}
		} else {
			if closestLeftPoint[1] < building[2] {
				leftPoints = append(leftPoints, []int{building[0], building[2]})

			} else {

			}	
		}

		if closestRightPoint[0] == building[1] {
			rightPoints = append(rightPoints, closestRightPoint)
		} else {
			if len(rightPoints) == 0 {
				rightPoints = append(rightPoints, []int{building[1], 0})
			} else {
				
			}
		}


		coverRight := len(coverIndices) == 0

		if coverLeft {
			if coverRight {
				continue
			} else {
				cStartIndex := coverIndices[0]
				cEndIndex := coverIndices[len(coverIndices)-1]
				xLeft := closestRightPoint[0]
				yLeft := building[2]
				xRight := building[1]
				var yRight int

				if cEndIndex < len(rightPoints)-1 {
					yRight = rightPoints[cEndIndex][1]
				} else {
					yRight = 0
				}

				rightPoints = append(rightPoints[:cStartIndex], rightPoints[cEndIndex+1:]...)
				if yLeft != closestLeftPoint[1] && building[0] != closestLeftPoint[0] {
					rightPoints = append(rightPoints, []int{xLeft, yLeft})
				}
				rightPoints = append(rightPoints, []int{xRight, yRight})
			}
		} else {
			originalClosestY := closestLeftPoint[1]
			if building[0] == closestLeftPoint[0] {
				if building[2] > closestLeftPoint[1] {
					closestLeftPoint[1] = building[2]
				}
			} else {
				leftPoints = append(leftPoints, []int{building[0], building[2]})
			}

			if coverRight {
				rightPoints = append(rightPoints, []int{building[1], originalClosestY})
			} else {
				cStartIndex := coverIndices[0]
				cEndIndex := coverIndices[len(coverIndices)-1]
				xRight := building[1]
				yRight := rightPoints[cEndIndex][1]
				rightPoints = append(rightPoints[:cStartIndex], rightPoints[cEndIndex+1:]...)
				rightPoints = append(rightPoints, []int{xRight, yRight})
			}
		}

		sort.Slice(rightPoints, func(i, j int) bool {
			return rightPoints[i][0] < rightPoints[j][0]
		})
	}

	ret := append(leftPoints, rightPoints...)
	return ret
}

// @lc code=end
func main() {
	utils.AssertEqual1[[][]int, [][]int]("example 1", [][]int{
		{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0},
	}, getSkyline, [][]int{
		{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8},
	})
	utils.AssertEqual1[[][]int, [][]int]("example 2", [][]int{
		{0, 3}, {5, 0},
	}, getSkyline, [][]int{
		{0, 2, 3}, {2, 5, 3},
	})
	utils.AssertEqual1[[][]int, [][]int]("example 3", [][]int{
		{0, 4}, {1, 3}, {2, 0},
	}, getSkyline, [][]int{
		{0, 2, 3}, {0, 1, 4},
	})
	utils.AssertEqual1[[][]int, [][]int]("example 4", [][]int{
		{1, 3}, {2, 0},
	}, getSkyline, [][]int{
		{1, 2, 1}, {1, 2, 2}, {1, 2, 3},
	})
	utils.AssertEqual1[[][]int, [][]int]("example 5", [][]int{
		{4, 15}, {9, 0}, {10, 10}, {12, 0},
	}, getSkyline, [][]int{
		{4, 9, 10}, {4, 9, 15}, {4, 9, 12}, {10, 12, 10}, {10, 12, 8},
	})
	utils.AssertEqual1[[][]int, [][]int]("example 6", [][]int{
		{3,8},{7,7},{8,6},{9,5},{10,4},{11,3},{12,2},{13,1},{14,0},
	}, getSkyline, [][]int{
		{3,7,8},{3,8,7},{3,9,6},{3,10,5},{3,11,4},{3,12,3},{3,13,2},{3,14,1},
	})
}
