package main

import "fmt"

type Sumsung struct {
	volume  int
	channel int
	isOn    bool
}

func (s *Sumsung) getVolume() int {
	fmt.Println("volume", s.volume)
	return s.volume
}

func (s *Sumsung) setVolume(vol int) {
	s.volume = vol
	fmt.Println("volume", s.volume)

}
func (s *Sumsung) getChanne() int {
	fmt.Println("channel", s.channel)
	return s.channel
}

func (s *Sumsung) setChannel(vol int) {
	s.channel = vol
	fmt.Println("channel", s.channel)

}

func (s *Sumsung) trunTrun(tvon bool) {
	s.isOn = tvon
	if tvon == true {
		fmt.Println("turn on")
	} else {
		fmt.Println("turn off")

	}
}
