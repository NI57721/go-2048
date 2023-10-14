package main

import (
	"fmt"
	"testing"
)

func TestCountOfCells(t *testing.T) {
	var target uint32
	var got, want int
	g := Game{cells: [16]uint32{0, 0, 4, 64, 0, 2, 4, 0, 0, 0, 4, 32, 0, 2, 4, 32}}

	target = 0
	want = 7
	got = g.countOfCells(target)
	if got != want {
		t.Errorf("After Game{cells: %v}.countOfCells(%d), cells is %d, but wants %d", g.cells, target, got, want)
	}

	target = 2
	want = 2
	got = g.countOfCells(target)
	if got != want {
		t.Errorf("After Game{cells: %v}.countOfCells(%d), cells is %d, but wants %d", g.cells, target, got, want)
	}

	target = 16
	want = 0
	got = g.countOfCells(target)
	if got != want {
		t.Errorf("After Game{cells: %v}.countOfCells(%d), cells is %d, but wants %d", g.cells, target, got, want)
	}
}

func TestLean(t *testing.T) {
	var target, got, want [4]uint32
	g := Game{}

	target = [4]uint32{0, 0, 0, 0}
	want = [4]uint32{0, 0, 0, 0}
	got = g.lean(target)
	if got != want {
		t.Errorf("g.lean(%v) == %v, but wants %v", target, got, want)
	}

	target = [4]uint32{0, 0, 0, 2}
	want = [4]uint32{2, 0, 0, 0}
	got = g.lean(target)
	if got != want {
		t.Errorf("g.lean(%v) == %v, but wants %v", target, got, want)
	}

	target = [4]uint32{0, 4, 0, 2}
	want = [4]uint32{4, 2, 0, 0}
	got = g.lean(target)
	if got != want {
		t.Errorf("g.lean(%v) == %v, but wants %v", target, got, want)
	}

	target = [4]uint32{0, 2, 0, 2}
	want = [4]uint32{4, 0, 0, 0}
	got = g.lean(target)
	if got != want {
		t.Errorf("g.lean(%v) == %v, but wants %v", target, got, want)
	}

	target = [4]uint32{0, 2, 8, 2}
	want = [4]uint32{2, 8, 2, 0}
	got = g.lean(target)
	if got != want {
		t.Errorf("g.lean(%v) == %v, but wants %v", target, got, want)
	}

	target = [4]uint32{4, 4, 4, 4}
	want = [4]uint32{8, 8, 0, 0}
	got = g.lean(target)
	if got != want {
		t.Errorf("g.lean(%v) == %v, but wants %v", target, got, want)
	}

	target = [4]uint32{64, 32, 32, 32}
	want = [4]uint32{64, 64, 32, 0}
	got = g.lean(target)
	if got != want {
		t.Errorf("g.lean(%v) == %v, but wants %v", target, got, want)
	}
}

func TestLeanToUp(t *testing.T) {
	target := [16]uint32{0, 0, 4, 64, 0, 2, 4, 0, 0, 0, 4, 32, 0, 2, 4, 32}
	want := [16]uint32{0, 4, 8, 64, 0, 0, 8, 64, 0, 0, 0, 0, 0, 0, 0, 0}
	g := Game{cells: target}
	g.leanToUp()
	got := g.cells
	if got != want {
		t.Errorf("After Game{cells: %v}.leanToUp(), cells is %v, but wants %v", target, got, want)
	}
}

func TestLeanToDown(t *testing.T) {
	target := [16]uint32{0, 2, 4, 32, 0, 0, 4, 32, 0, 2, 4, 0, 0, 0, 4, 64}
	want := [16]uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 64, 0, 4, 8, 64}
	g := Game{cells: target}
	g.leanToDown()
	got := g.cells
	if got != want {
		t.Errorf("After Game{cells: %v}.leanToDown(), cells is %v, but wants %v", target, got, want)
	}
}

func TestLeanToRight(t *testing.T) {
	target := [16]uint32{0, 0, 0, 0, 2, 0, 2, 0, 4, 4, 4, 4, 32, 32, 0, 64}
	want := [16]uint32{0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 8, 8, 0, 0, 64, 64}
	g := Game{cells: target}
	g.leanToRight()
	got := g.cells
	if got != want {
		t.Errorf("After Game{cells: %v}.leanToRight(), cells is %v, but wants %v", target, got, want)
	}
}

func TestLeanToLeft(t *testing.T) {
	target := [16]uint32{0, 0, 0, 0, 0, 2, 0, 2, 4, 4, 4, 4, 64, 0, 32, 32}
	want := [16]uint32{0, 0, 0, 0, 4, 0, 0, 0, 8, 8, 0, 0, 64, 64, 0, 0}
	g := Game{cells: target}
	g.leanToLeft()
	got := g.cells
	if got != want {
		t.Errorf("After Game{cells: %v}.leanToLeft(), cells is %v, but wants %v", target, got, want)
	}
}

func TestDrawIfAbleToLean(t *testing.T) {
	var target, got, want [16]uint32
	var g Game

	target = [16]uint32{0, 0, 0, 0, 8, 4, 0, 0, 16, 64, 0, 0, 128, 0, 0, 0}
	want = target
	g = Game{cells: target}
	g.drawIfAbleToLean("left")
	got = g.cells
	if got != want {
		t.Errorf("After Game{cells: %v}.drawIfAbleToLean(\"left\"), cells is %v, but wants %v", target, got, want)
	}

	target = [16]uint32{0, 32, 0, 0, 8, 256, 0, 0, 16, 64, 0, 0, 128, 0, 0, 0}
	want = target
	g = Game{cells: target}
	g.drawIfAbleToLean("left")
	got = g.cells
	if g.countOfCells(2) != 1 && g.countOfCells(4) != 1 {
		t.Errorf("After Game{cells: %v}.drawIfAbleToLean(\"left\"), draw() is not excuted", target)
	}
}

func TestDraw(t *testing.T) {
	g := Game{}
	for i := 0; i < len(g.cells); i++ {
		g.draw()
	}

	isAnyEmpty := false
	for _, c := range g.cells {
		if c == 0 {
			isAnyEmpty = true
		}
	}
	if isAnyEmpty != false {
		t.Errorf("There is any room even after Game.draw() 16 times.")
	}
	// g.draw()
}

func TestUint32ToCell(t *testing.T) {
	var target uint32
	var got, want string

	target = 0
	want = fmt.Sprintf("\x1b[48;2;%sm      \x1b[m", colors[target])
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}

	target = 2
	want = fmt.Sprintf("\x1b[48;2;%sm   %d  \x1b[m", colors[target], target)
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}

	target = 16
	want = fmt.Sprintf("\x1b[48;2;%sm  %d  \x1b[m", colors[target], target)
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}

	target = 128
	want = fmt.Sprintf("\x1b[48;2;%sm %d  \x1b[m", colors[target], target)
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}

	target = 1024
	want = fmt.Sprintf("\x1b[48;2;%sm %d \x1b[m", colors[target], target)
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}

	target = 16384
	want = fmt.Sprintf("\x1b[48;2;%sm %d\x1b[m", colors[target], target)
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}

	target = 32768
	want = fmt.Sprintf("\x1b[48;2;%sm %d\x1b[m", colors[target], target)
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}

	target = 65536
	want = fmt.Sprintf("\x1b[48;2;%sm %d\x1b[m", colors[target], target)
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}

	target = 131072
	want = fmt.Sprintf("\x1b[48;2;%sm%d\x1b[m", colors[target], target)
	got = uint32ToCell(target)
	if got != want {
		t.Errorf("uint32ToCell(%d) == \"%s\", but wants \"%s\"", target, got, want)
	}
}

func TestString(t *testing.T) {
	indent := "  "
	target := Game{
		cells: [16]uint32{0, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 65536, 131072},
		score: 128,
	}
	want := fmt.Sprintf(
		"%s      Score: 000128\n"+
			"%s%s%s%s%s\n"+
			"%s%s%s%s%s\n"+
			"%s%s%s%s%s\n"+
			"%s%s%s%s%s\n",
		indent,
		indent, uint32ToCell(0), uint32ToCell(2), uint32ToCell(4), uint32ToCell(8),
		indent, uint32ToCell(16), uint32ToCell(32), uint32ToCell(64), uint32ToCell(128),
		indent, uint32ToCell(256), uint32ToCell(512), uint32ToCell(1024), uint32ToCell(2048),
		indent, uint32ToCell(4096), uint32ToCell(8192), uint32ToCell(65536), uint32ToCell(131072),
	)
	got := target.String()
	if got != want {
		t.Errorf("%#v.String() is \n%s\nbut wants\n%s", target, got, want)
	}
}
