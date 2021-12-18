package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	targetArea := GetPuzzleInput()
	maxY := 0
	count := 0
	for i := 1; i < targetArea["x"][1] + 1; i++ {
		for j := targetArea["y"][0] - 1; j < 500; j++ {
			p := Probe{}
			p.Init(i, j, targetArea)
			if p.IsInTargetArea {
				count +=1
				for _, coor := range p.Trajectory {
					if maxY < coor[1] {
						maxY = coor[1]
					}
				}
			}
		}
	}
	fmt.Println("Puzzle 1 =", maxY)

	fmt.Println("Puzzle 2 =", count)

}

type Probe struct {
	InitialVelocityX int
	InitialVelocityY int
	VelocityX  int
	VelocityY  int
	Trajectory [][]int
	IsInTargetArea	 bool
}

func (p *Probe) Init(x int, y int, targetArea map[string][]int) {
	p.InitialVelocityX = x
	p.InitialVelocityY = y
	p.VelocityX = x
	p.VelocityY = y
	p.Trajectory = append(p.Trajectory, []int{0, 0})
	p.IsInTargetArea = false
	p.CalculateTrajectory(targetArea)
}

func (p *Probe) CalculateTrajectory(area map[string][]int) {
	for {
		p.Step()
		coor := p.Trajectory[len(p.Trajectory) - 1]
		x, y := coor[0], coor[1]
		if x >= area["x"][0] && x <= area["x"][1] && y >= area["y"][0] && y <= area["y"][1] {
			p.IsInTargetArea = true
			break
		}
		if x > area["x"][1] || y < area["y"][0] {
			break
		}
	}
}

func (p *Probe) Step() {
	coor := p.Trajectory[len(p.Trajectory) - 1]
	x, y := coor[0], coor[1]
	newX := x + p.VelocityX
	newY := y + p.VelocityY
	newCoor := []int{newX, newY}
	p.Trajectory = append(p.Trajectory, newCoor)
	switch {
	case p.VelocityX > 0:
		p.VelocityX -= 1
	case p.VelocityX < 0:
		p.VelocityX += 1
	}
	p.VelocityY -= 1
}

func GetPuzzleInput() map[string][]int {
	raw := "target area: x=48..70, y=-189..-148"
	//raw = "target area: x=20..30, y=-10..-5"
	raw = strings.Replace(raw, "target area: x=", "", 1)
	raw = strings.Replace(raw, ", y=", "..", 1)
	coor := strings.Split(raw, "..")
	x0, _ := strconv.Atoi(coor[0])
	x1, _ := strconv.Atoi(coor[1])
	y0, _ := strconv.Atoi(coor[2])
	y1, _ := strconv.Atoi(coor[3])
	targetArea := make(map[string][]int)
	targetArea["x"] = []int{x0, x1}
	targetArea["y"] = []int{y0, y1}

	return targetArea
}