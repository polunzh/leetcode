package swappingnodesinalinkedlist

import (
	. "leetcode/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, []int{1}, LinklistToArray(swapNodes(MakeLinkList([]int{1}), 2)))
}

func Test2(t *testing.T) {
	// assert.Equal(t, []int{1, 4, 3, 2, 5}, swapNodes(MakeLinkList([]int{1, 2, 3, 4, 5}), 2))
	assert.Equal(t, []int{1}, LinklistToArray(swapNodes(MakeLinkList([]int{1}), 1)))
}

func Test3(t *testing.T) {
	assert.Equal(t, []int{1, 4, 3, 2, 5}, LinklistToArray(swapNodes(MakeLinkList([]int{1, 2, 3, 4, 5}), 2)))
}

func Test4(t *testing.T) {
	l := MakeLinkList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	assert.Equal(t, []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5}, LinklistToArray(swapNodes(l, 5)))
}

func Test5(t *testing.T) {
	assert.Equal(t, []int{2, 1}, LinklistToArray(swapNodes(MakeLinkList([]int{1, 2}), 1)))
}

func Test6(t *testing.T) {
	assert.Equal(t, []int{100, 90}, LinklistToArray(swapNodes(MakeLinkList([]int{90, 100}), 2)))
}

func Test7(t *testing.T) {
	assert.Equal(t, []int{64, 46, 66, 35, 80}, LinklistToArray(swapNodes(MakeLinkList([]int{80, 46, 66, 35, 64}), 1)))
}

func Test8(t *testing.T) {
	assert.Equal(t, []int{26, 24, 24, 36, 18, 52, 95, 61, 54, 88, 86, 79, 11, 1, 31, 100}, LinklistToArray(swapNodes(MakeLinkList([]int{100, 24, 24, 36, 18, 52, 95, 61, 54, 88, 86, 79, 11, 1, 31, 26}), 16)))
}

func Test9(t *testing.T) {
	l := MakeLinkList([]int{55, 60, 78, 53, 93, 37, 31, 4, 61, 11, 13, 51, 34, 83, 24, 96, 5, 77, 1, 67})
	assert.Equal(t, []int{55, 60, 78, 53, 93, 37, 31, 4, 61, 13, 11, 51, 34, 83, 24, 96, 5, 77, 1, 67}, LinklistToArray(swapNodes(l, 11)))
}
