package main

func main() {
	s := &Sumsung{
		volume: 34, channel: 23, isOn: true,
	}
	sd := &SAdapter{
		sum: s,
	}
	sd.trunOn()
	sd.volumeUP()
	sd.channelDown()
	sd.goChannel(69)
	sd.channelUP()

	sd.trunOff()

}
