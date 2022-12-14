package main

import "fmt"

type Lg struct {
	volume  int
	channel int
	isOn    bool
}

func (s *Lg) volumeUP() {
	s.volume--
	fmt.Println("volume", s.volume)
}

func (s *Lg) volumeDown() {
	s.volume++
	fmt.Println("volume", s.volume)

}
func (s *Lg) channelUP() {
	s.channel--
	fmt.Println("channel", s.channel)
}

func (s *Lg) channelDown() {
	s.channel++
	fmt.Println("channel", s.channel)

}
func (s *Lg) goChannel(channel int) {
	s.channel = channel
	fmt.Println("goChannel", s.channel)

}
func (s *Lg) trunOn() {
	s.isOn = true
	fmt.Println("turn on")
}

func (s *Lg) trunOff() {
	s.isOn = false
	fmt.Println("turn on")
}
