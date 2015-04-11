package main

import (
	"math/rand"
	"time"

	"github.com/voldyman/portaudio-go/portaudio"
)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()
	h, err := portaudio.DefaultHostApi()
	chk(err)
	stream, err := portaudio.OpenStream(portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice), func(out []int32) {
		for i := range out {
			out[i] = int32(rand.Uint32())
		}
	})
	chk(err)
	defer stream.Close()
	chk(stream.Start())
	time.Sleep(time.Second)
	chk(stream.Stop())
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
