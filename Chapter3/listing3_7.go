package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Shuffleable interface {
	contents() string
	shuffle()
}

type shuffleString string

func (s *shuffleString) shuffle() {
	tmp := strings.Split(string(*s), "")
	rand.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})
	*s = shuffleString(strings.Join(tmp, ""))
}

func (s *shuffleString) contents() string {
	return string(*s)
}

func NewShuffleString(init string) *shuffleString {
	var s shuffleString = shuffleString(init)
	return &s
}

type shuffleSlice []interface{}

func (sl shuffleSlice) contents() string {
	data, _ := json.Marshal(sl)
	return fmt.Sprintf("%v", string(data))
}

func (sl shuffleSlice) shuffle() {
	rand.Shuffle(len(sl), func(i, j int) {
		sl[i], sl[j] = sl[j], sl[i]
	})
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var myShuffle Shuffleable

	myShuffle = NewShuffleString("my name is inigo montoya")
	myShuffle.shuffle()
	fmt.Println(myShuffle.contents())

	myShuffle = &shuffleSlice{1, 2, 3, 4, 5}
	myShuffle.shuffle()
	fmt.Println(myShuffle.contents())
}
