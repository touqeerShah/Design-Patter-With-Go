package main

import "fmt"

type SAdapter struct {
	sum *Sumsung
}

func (s *SAdapter) volumeUP() {
	vol := s.sum.getVolume()
	vol++
	s.sum.setVolume(vol)
	fmt.Println("volume", vol)
}

func (s *SAdapter) volumeDown() {
	vol := s.sum.getVolume()
	vol--
	s.sum.setVolume(vol)
	fmt.Println("volume", vol)
}

func (s *SAdapter) channelUP() {
	ch := s.sum.getChanne()
	ch++
	s.sum.setChannel(ch)
	fmt.Println("volume", ch)
}

func (s *SAdapter) channelDown() {
	ch := s.sum.getChanne()
	ch--
	s.sum.setChannel(ch)
	fmt.Println("volume", ch)

}
func (s *SAdapter) goChannel(channel int) {
	s.sum.setChannel(channel)
	fmt.Println("goChannel", channel)

}
func (s *SAdapter) trunOn() {
	s.sum.trunTrun(true)
	fmt.Println("turn on")
}

func (s *SAdapter) trunOff() {
	s.sum.trunTrun(false)
	fmt.Println("turn on")
}
