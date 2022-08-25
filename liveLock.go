package main

import (
	"fmt"
	"time"
)

type hallway struct {
	gopherCrossing *gopher
}

func (s hallway) crossing() {
	fmt.Printf("%s has crossing!\n", s.gopherCrossing.name)
}

type gopher struct {
	name          string
	isNeedCrosing bool
}

func (g *gopher) walk(hallway *hallway, otherGopher *gopher) {
	for g.isNeedCrosing {
		if hallway.gopherCrossing != g {
			time.Sleep((1 * time.Millisecond * 100))
			continue
		}
		/* if otherGopher.isNeedCrosing {
			fmt.Printf("%s: please, You crossing first %s!\n", g.name, otherGopher.name)
			hallway.gopherCrossing = otherGopher
			continue
		} */
		hallway.crossing()
		g.isNeedCrosing = false
		fmt.Printf("%s I have crossed the hallway, thank you %s!\n", g.name, otherGopher.name)
		hallway.gopherCrossing = otherGopher
		return
	}
}

func main() {
	gopherPurple := &gopher{
		name:          "Gopher Purple",
		isNeedCrosing: true,
	}

	gopherGreen := &gopher{
		name:          "Gopher Green",
		isNeedCrosing: true,
	}

	halfway := &hallway{gopherCrossing: gopherPurple}

	go func() {
		gopherPurple.walk(halfway, gopherGreen)
	}()
	go func() {
		gopherGreen.walk(halfway, gopherPurple)
	}()
	time.Sleep(time.Second)
}
