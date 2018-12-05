package day3

import (
	"fmt"
	"github.com/heyimblake/AdventOfCode-2018/util"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
	Claims []*Claim
}

type Claim struct {
	Id int
	FromLeft int
	FromTop int
	Width int
	Height int
	Points []*Point
}

func CreateClaim(data string) *Claim {
	splitData := strings.Split(data, " ")
	fromNums := strings.Split(splitData[2], ",") // Ex: "45" "64:"
	dimensions := strings.Split(splitData[3], "x")

	id, _ := strconv.Atoi(splitData[0][1:]) // Remove the # symbol
	fromLeft, _ :=  strconv.Atoi(fromNums[0])
	fromTop, _ := strconv.Atoi(fromNums[1][:len(fromNums[1]) - 1]) // Remove the : symbol
	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	claim := Claim{id, fromLeft, fromTop, width, height, nil}

	return &claim
}

func ClaimToString(claim *Claim) string {
	return fmt.Sprintf("#%d @ %d,%d: %dx%d", claim.Id, claim.FromLeft, claim.FromTop, claim.Width, claim.Height)
}

func LoadClaimsFromFile(filePath string) []*Claim {
	claims := make([]*Claim, 0)

	content, _ := util.GetFileContents(filePath)

	for _, line := range *content {
		claims = append(claims, CreateClaim(line))
	}

	return claims
}

func PrepareClaims(filePath string) []*Claim {
	// Read file and make a Claim struct for each one
	claims := LoadClaimsFromFile(filePath)

	// Map string to point pointer, so we can ensure we are using the same reference
	stringPointMap := make(map[string]*Point)


	// Calculate the points of each claim and append it to the array in the struct
	for _, claim := range claims {
		claim.Points = make([]*Point, 0)

		for x := claim.FromLeft; x < claim.FromLeft + claim.Width; x++ {
			for y := claim.FromTop; y < claim.FromTop + claim.Height; y++ {
				point := stringPointMap[fmt.Sprintf("{%d %d}", x, y)]

				if point == nil {
					point = &Point{x, y, make([]*Claim, 0)}
					stringPointMap[fmt.Sprintf("{%d %d}", x, y)] = point
				}

				claim.Points = append(claim.Points, point)
				point.Claims = append(point.Claims, claim)
			}
		}
	}

	return claims
}

/*
	Part 1
 */
func FindOverlappingSquareInches(claims []*Claim) int {
	// Create Map to keep track of how many occurrences of each point there are
	pointOccurrences := make(map[*Point]int)

	// Go through all points of each Claim and add to map as needed
	for _, claim := range claims {
		for _, point := range claim.Points{
			pointOccurrences[point]++
		}
	}

	sqInchCount := 0

	// Find which occurred more than once
	for _, count := range pointOccurrences {
		if count > 1 {
			sqInchCount++
		}
	}

	return sqInchCount
}

/*
	Part 2
 */
func FindNonOverlappingClaim(claims []*Claim) *Claim {
	// Find the Claim where all its points has 1 claim to it
	for _, claim := range claims {
		passed := true

		pointsLoop:
		for _, point := range claim.Points {
			if len(point.Claims) != 1 {
				passed = false
				break pointsLoop
			}
		}

		// We found it!!! Woo!!
		if passed {
			return claim
		}
	}

	// Or if we couldn't find anything :(
	return nil
}
