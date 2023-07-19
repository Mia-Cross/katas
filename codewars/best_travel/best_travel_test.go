package best_travel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisitOneTownWithOneElemListShouldReturnElem(t *testing.T) {
	ts := []int{50}
	actual := ChooseBestSum(50, 1, ts)
	assert.Equal(t, 50, actual)
}

func TestVisitOneTownWithOneElemListShouldReturnMinusOne(t *testing.T) {
	ts := []int{57}
	actual := ChooseBestSum(56, 1, ts)
	assert.Equal(t, -1, actual)
}

func TestVisitOneTown(t *testing.T) {
	ts := []int{50, 55, 56, 57, 58}
	actual := ChooseBestSum(65, 1, ts)
	assert.Equal(t, 58, actual)
	actual = ChooseBestSum(56, 1, ts)
	assert.Equal(t, 56, actual)
}

func TestVisitTwoTowns(t *testing.T) {
	ts := []int{50, 55, 56, 57, 58}
	actual := ChooseBestSum(106, 2, ts)
	assert.Equal(t, 106, actual)
	ts = []int{50, 55, 54, 57, 58}
	actual = ChooseBestSum(106, 2, ts)
	assert.Equal(t, 105, actual)
	ts = []int{91, 74, 73, 85, 73, 81, 87}
	actual = ChooseBestSum(331, 2, ts)
	assert.Equal(t, 178, actual)
}

func TestVisitTwoTownsNoDuplicate(t *testing.T) {
	// ts := []int{50, 55, 56, 57, 50}
	// actual := ChooseBestSum(100, 2, ts)
	// assert.Equal(t, 100, actual)
	ts := []int{50, 55, 56, 57, 58}
	actual := ChooseBestSum(100, 2, ts)
	assert.Equal(t, -1, actual)
}

func TestVisitThreeTowns(t *testing.T) {
	ts := []int{50, 55, 56, 57, 50}
	actual := ChooseBestSum(156, 3, ts)
	assert.Equal(t, 156, actual)
	ts = []int{50, 55, 56, 57, 58}
	actual = ChooseBestSum(163, 3, ts)
	assert.Equal(t, 163, actual)
	ts = []int{91, 74, 73, 85, 73, 81, 87}
	actual = ChooseBestSum(230, 3, ts)
	assert.Equal(t, 228, actual)
}

func TestVisitMoreTownsThatThereAreInList(t *testing.T) {
	ts := []int{50}
	actual := ChooseBestSum(163, 3, ts)
	assert.Equal(t, -1, actual)
}

func TestVisitFourTowns(t *testing.T) {
	ts := []int{91, 74, 73, 85, 73, 81, 87}
	actual := ChooseBestSum(331, 4, ts)
	assert.Equal(t, 331, actual)
}

func TestVisitFiveTownsWithMaxDistanceTooSmall(t *testing.T) {
	ts := []int{91, 74, 73, 85, 73, 81, 87}
	actual := ChooseBestSum(331, 5, ts)
	assert.Equal(t, -1, actual)
}

func TestVisitFiveTowns(t *testing.T) {
	ts := []int{91, 74, 73, 85, 73, 81, 87}
	actual := ChooseBestSum(409, 5, ts)
	assert.Equal(t, 409, actual)
}
