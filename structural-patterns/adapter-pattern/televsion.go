package main

type televsion interface {
	volumeUP()
	volumeDown()
	channelUP()
	channelDown()
	goChannel(int)
	trunOn()
	trunOff()
}
