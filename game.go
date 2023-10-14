package main

import (
	"fmt"
	"math/rand"
	"time"
)

var colors = map[uint32]string{
	0:      "60;40;30",
	2:      "110;50;30",
	4:      "160;50;30",
	8:      "210;50;30",
	16:     "255;50;30",
	32:     "255;50;70",
	64:     "245;40;100",
	128:    "235;30;130",
	256:    "225;20;160",
	512:    "215;10;190",
	1024:   "205;0;220",
	2048:   "195;0;255",
	4096:   "185;0;245",
	8192:   "175;0;235",
	16384:  "165;0;225",
	32768:  "155;0;215",
	65536:  "145;0;205",
	131072: "135;0;195",
}

type Game struct {
	cells [16]uint32
	score uint32
}

func (g *Game) countOfCells(n uint32) int {
	cnt := 0
	for _, c := range g.cells {
		if c == n {
			cnt++
		}
	}
	return cnt
}

// Make a row lean to the left
// e.g. [0, 2, 2, 4] -> [4, 4, 0, 0]
func (g *Game) lean(row [4]uint32) [4]uint32 {
	for i := 0; i < 3; i++ {
		if row[i] == 0 {
			continue
		}
		for j := i + 1; j < 4; j++ {
			if row[j] == 0 {
				continue
			}
			if row[i] == row[j] {
				row[i] = row[i] * 2
				row[j] = 0
				g.score += row[i]
			}
			break
		}
	}
	for i := 0; i < 3; i++ {
		if row[i] != 0 {
			continue
		}
		for j := i + 1; j < 4; j++ {
			if row[j] == 0 {
				continue
			}
			row[i] = row[j]
			row[j] = 0
			break
		}
	}
	return row
}

func (g *Game) leanToUp() {
	for i := 0; i < 4; i++ {
		var row [4]uint32
		for j := 0; j < 4; j++ {
			row[j] = g.getCellAt(i, j)
		}
		leanedRow := g.lean(row)
		for j, c := range leanedRow {
			g.setCellAt(i, j, c)
		}
	}
}

func (g *Game) leanToDown() {
	for i := 0; i < 4; i++ {
		var row [4]uint32
		for j := 0; j < 4; j++ {
			row[j] = g.getCellAt(i, 3-j)
		}
		leanedRow := g.lean(row)
		for j, c := range leanedRow {
			g.setCellAt(i, 3-j, c)
		}
	}
}

func (g *Game) leanToRight() {
	for i := 0; i < 4; i++ {
		var row [4]uint32
		for j := 0; j < 4; j++ {
			row[j] = g.getCellAt(3-j, i)
		}
		leanedRow := g.lean(row)
		for j, c := range leanedRow {
			g.setCellAt(3-j, i, c)
		}
	}
}

func (g *Game) leanToLeft() {
	for i := 0; i < 4; i++ {
		var row [4]uint32
		for j := 0; j < 4; j++ {
			row[j] = g.getCellAt(j, i)
		}
		leanedRow := g.lean(row)
		for j, c := range leanedRow {
			g.setCellAt(j, i, c)
		}
	}
}

func (g *Game) drawIfAbleToLean(dir string) {
	preCells := g.cells
	switch dir {
	case "up":
		g.leanToUp()
	case "down":
		g.leanToDown()
	case "right":
		g.leanToRight()
	case "left":
		g.leanToLeft()
	default:
		panic("Invalid direction")
	}
	if g.cells != preCells {
		g.draw()
	}
}

func (g *Game) getCellAt(x, y int) uint32 {
	return g.cells[y*4+x]
}

func (g *Game) setCellAt(x, y int, c uint32) {
	g.cells[y*4+x] = c
}

func (g *Game) draw() {
	var emptyIndices []int
	for i, c := range g.cells {
		if c == 0 {
			emptyIndices = append(emptyIndices, i)
		}
	}
	if len(emptyIndices) == 0 {
		panic("There is not any more room for new drawing.")
	}

	rand.Seed(time.Now().UnixNano())
	var nextNum uint32
	if rand.Float32() < .9 {
		nextNum = 2
	} else {
		nextNum = 4
	}
	rand.Seed(time.Now().UnixNano())
	index := emptyIndices[rand.Intn(len(emptyIndices))]
	g.cells[index] = nextNum
}

func uint32ToCell(i uint32) string {
	switch i {
	case 0:
		return fmt.Sprintf("\x1b[48;2;%sm      \x1b[m", colors[i])
	case 2, 4, 8:
		return fmt.Sprintf("\x1b[48;2;%sm   %d  \x1b[m", colors[i], i)
	case 16, 32, 64:
		return fmt.Sprintf("\x1b[48;2;%sm  %d  \x1b[m", colors[i], i)
	case 128, 256, 512:
		return fmt.Sprintf("\x1b[48;2;%sm %d  \x1b[m", colors[i], i)
	case 1024, 2048, 4096, 8192:
		return fmt.Sprintf("\x1b[48;2;%sm %d \x1b[m", colors[i], i)
	case 16384, 32768, 65536:
		return fmt.Sprintf("\x1b[48;2;%sm %d\x1b[m", colors[i], i)
	case 131072:
		return fmt.Sprintf("\x1b[48;2;%sm%d\x1b[m", colors[i], i)
	default:
		panic(fmt.Sprintf("Range Error: %d is not between 0 and 131072", i))
	}
}

func (g *Game) String() string {
	const indent = "  "
	str := fmt.Sprintf("%s      Score: %06d\n", indent, g.score)
	for i, c := range g.cells {
		if i%4 == 0 {
			str += indent
		}
		str += uint32ToCell(c)
		if i%4 == 3 {
			str += "\n"
		}
	}
	return str
}
