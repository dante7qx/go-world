package _audio

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

func ReadMap3() {
	var mp3Path = "/Users/dante/Desktop/差不多先生.mp3"

	f, err := os.Open(mp3Path)
	if err != nil {
		fmt.Printf("Failed to open MP3 file: %v\n", err)
		return
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		fmt.Printf("Failed to create decoder: %v\n", err)
		return
	}
	fmt.Printf("Playback duration: %v\n", int(d.Length())/d.SampleRate())
}
