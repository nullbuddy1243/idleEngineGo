package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const exclaim string = "! "

type player struct {
	Id       int   `json:"id"`
	Outposts []int `json:"outposts"` // ids of outposts belonging to player
}

type outpost struct {
	Id   int `json:"id"`
	Land int `json:"land"`
}

type land struct {
	Id        int `json:"id"`
	Xposition int `json:"xposition"`
	Yposition int `json:"yposition"`
	Life      int `json:"life"`  //🟩
	Flow      int `json:"flow"`  // 🟦
	X         int `json:"x"`     // ⬛
	Spark     int `json:"spark"` // 🟥
}

// func newPlayer() (struct, struct, struct) {
func newPlayer() (player, outpost, land) {

	// build player
	player := player{Id: 1}
	// determine starting Land tile
	land := land{
		Id:        1,
		Xposition: rand.Intn(1000),
		Yposition: rand.Intn(1000),
		Life:      rand.Intn(10),
		Flow:      rand.Intn(10),
		X:         rand.Intn(10),
		Spark:     rand.Intn(10),
	}
	// build first outpost on starting Land title
	outpost := outpost{Id: 1, Land: land.Id}
	return player, outpost, land
}

func main() {
	thestring := "starting up"
	fmt.Println(thestring + exclaim)
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	player, land, outpost := newPlayer()

	fmt.Println("====")
	// playerJSON, _ := convertJSON(player)
	playerJSON, _ := json.MarshalIndent(player, "", "	")
	landJSON, _ := json.MarshalIndent(land, "", "	")
	outpostJSON, _ := json.MarshalIndent(outpost, "", "	")
	fmt.Println("Player: ", string(playerJSON))
	fmt.Println("Starting Land: ", string(landJSON))
	fmt.Println("Outpost: ", string(outpostJSON))
	fmt.Println("====")

	runcount := 0
	for true {
		time.Sleep(time.Second)
		tic(runcount)
		runcount++
	}

}

// func convertJSON(thing) (string, error) {
// 	output, err := json.MarshalIndent(thing, "", "	")
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 		return err
// 	} else {
// 		return string(output)
// 	}
// }

func tic(runcount int) {
	outstr := ""
	for i := 0; i < runcount; i++ {
		outstr += "."
	}

	fmt.Println(outstr)
}

func calcLife(life int, localpop int, flow int) {
	var gain float32
	// see note
	gain = math.Round(life/localpop) + flow
	consume := localpop
	gain = gain - consume
	// nah im off wrong thinking here
	// the life of a land / outpost doesnt change tic by tic
	// rather it can be added to with structures or upgrades (?)

}

func popGrowth(int life, int pop) int {
	var newPop int
	if life > 0 {
		// life / spark
	}
}
