package day3

import (
	"github.com/heyimblake/AdventOfCode-2018/util"
	"testing"
)

func TestCreateClaim(t *testing.T) {
	contents, _ := util.GetFileContents("input.txt")

	claim := CreateClaim((*contents)[3])

	if !claimEquals(&claim, 4, 694, 889, 28, 28) {
		t.Fatalf("Claim does not match. (Case 1)")
	}

	claim = CreateClaim((*contents)[49])

	if !claimEquals(&claim, 50, 310, 86, 11, 27) {
		t.Fatalf("Claim does not match. (Case 2)")
	}

	claim = CreateClaim((*contents)[1300])

	if !claimEquals(&claim, 1301, 130, 498, 26, 10) {
		t.Fatalf("Claim does not match. (Case 3)")
	}
}

func TestClaimToString(t *testing.T) {
	cases := []string{
		"#1234 @ 23,232: 91x53",
		"#129842 @ 12414124132,345345346: 28294x9",
		"#1284 @ 347,704: 27x23",
		"#830 @ 815,683: 16x22",
		"#-1 @ 0,214: 2x23",
	}

	for _, expected := range cases {
		actual := ClaimToString(CreateClaim(expected))

		if actual != expected {
			t.Fatalf("Expected \"%s\" but got \"%s\"\n", expected, actual)
		}
	}
}

func TestLoadClaimsFromFile(t *testing.T)  {
	claims := LoadClaimsFromFile("input.txt")
	actual := claims[493]

	if !claimEquals(&actual, 494, 651, 815, 25, 18) {
		t.Fatalf("Claim does not match. (Case 1)")
	}

	actual = claims[7]

	if !claimEquals(&actual, 8, 612, 291, 21, 17) {
		t.Fatalf("Claim does not match. (Case 2)")
	}

	actual = claims[241]

	if !claimEquals(&actual, 242, 178, 287, 22, 23) {
		t.Fatalf("Claim does not match. (Case 3)")
	}
}

func TestFindOverlappingSquareInches(t *testing.T)  {
	expected := 104712
	actual := FindOverlappingSquareInches(PrepareClaims("input.txt"))

	if actual != expected {
		t.Fatalf("Expected %d but got %d\n", expected, actual)
	}
}

func TestFindNonOverlappingClaim(t *testing.T)  {
	expected := 840
	actual := FindNonOverlappingClaim(PrepareClaims("input.txt")).Id

	if actual != expected {
		t.Fatalf("Expected %d but got %d\n", expected, actual)
	}
}

func claimEquals(claim **Claim, id int, fromLeft int, fromTop int, width int, height int) bool {
	if (*claim).Id != id || (*claim).FromLeft != fromLeft || (*claim).FromTop != fromTop || (*claim).Width != width || (*claim).Height != height {
		return false
	}

	return true
}
