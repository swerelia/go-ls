package printcols

import (
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode/utf8"
)

var Output io.Writer = os.Stdout

func PrintColumns(strs *[]string, margin int) {
	maxLength := 0
	marginStr := strings.Repeat(" ", margin)
	lengths := []int{}
	for _, str := range *strs {
		length := utf8.RuneCountInString(str)
		maxLength = max(maxLength, length)
		lengths = append(lengths, length)
	}

	width := getTermWidth()
	numCols, numRows := calculateTableSize(width, margin, maxLength, len(*strs))

	for i := 0; i < numCols*numRows; i++ {
		x, y := rowIndexToTableCoords(i, numCols)
		j := tableCoordsToColIndex(x, y, numRows)

		strLen := 0
		str := ""
		if j < len(lengths) {
			strLen = lengths[j]
			str = (*strs)[j]
		}

		numSpacesRequired := maxLength - strLen
		spaceStr := strings.Repeat(" ", numSpacesRequired)

		fmt.Fprintf(Output, str)

		if x+1 == numCols {
			fmt.Fprintf(Output, "\n")
		} else {
			fmt.Fprintf(Output, spaceStr)
			fmt.Fprintf(Output, marginStr)
		}
	}
}

func getTermWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err1 := cmd.Output()
	check(err1)
	numsStr := strings.Trim(string(out), "\n ")
	width, err2 := strconv.Atoi(strings.Split(numsStr, " ")[1])
	check(err2)
	return width
}

func calculateTableSize(width, margin, maxLength, numCells int) (int, int) {
	numCols := (width + margin) / (maxLength + margin)
	if numCols == 0 {
		numCols = 1
	}
	numRows := int(math.Ceil(float64(numCells) / float64(numCols)))
	return numCols, numRows
}

func rowIndexToTableCoords(i, numCols int) (int, int) {
	x := i % numCols
	y := i / numCols
	return x, y
}

func tableCoordsToColIndex(x, y, numRows int) int {
	return y + numRows*x
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}