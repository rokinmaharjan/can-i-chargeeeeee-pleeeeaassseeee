package notify

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func Notify() {
	// Open the WAV file
	f, err := os.Open("./resources/alert.wav")
	if err != nil {
		log.Fatal(err)
	}

	// Decode the WAV file
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize the speaker
	sr := format.SampleRate * 1
	speaker.Init(sr, sr.N(time.Second/10))

	// Play the sound
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
