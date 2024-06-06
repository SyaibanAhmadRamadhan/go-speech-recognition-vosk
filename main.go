package main

import (
	"encoding/json"
	"fmt"
	vosk "github.com/alphacep/vosk-api/go"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gordonklaus/portaudio"
)

const (
	// Number of samples per seconds.
	sampleRate = 16_000

	// Number of samples to send at once.
	framesPerBuffer = 3_200
)

type VoskResult struct {
	Text string `json:"text"`
}

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Initialize PortAudio
	err := portaudio.Initialize()
	if err != nil {
		panic(err)
	}
	defer portaudio.Terminate()

	// Open the Vosk model
	modelPath := "vosk-model-small-en-us-0.15" // Path to your Vosk model directory
	model, err := vosk.NewModel(modelPath)
	if err != nil {
		log.Fatalf("Failed to open Vosk model: %v", err)
	}

	recognizer, err := vosk.NewRecognizer(model, float64(sampleRate))
	if err != nil {
		panic(err)
	}

	rec, err := newRecorder(sampleRate, framesPerBuffer)
	if err != nil {
		panic(err)
	}
	if err := rec.Start(); err != nil {
		log.Fatalf("Error starting stream: %v", err)
	}

	for {
		select {
		case <-sigs:
			slog.Info("stopping recording...")

			var err error

			err = rec.Stop()
			if err != nil {
				panic(err)
			}

			os.Exit(0)
		default:
			b, err := rec.Read()
			if err != nil {
				panic(err)
			}

			r := recognizer.AcceptWaveform(b)
			if r != 0 {
				resultStr := recognizer.Result()
				var result VoskResult
				if err := json.Unmarshal([]byte(resultStr), &result); err != nil {
					log.Printf("Error unmarshalling result: %v", err)
					return
				}
				if result.Text != "" {
					fmt.Printf("Transcription: %s\n", result.Text)
				}
			} else {
				partialResultStr := recognizer.PartialResult()
				var partialResult VoskResult
				if err := json.Unmarshal([]byte(partialResultStr), &partialResult); err != nil {
					log.Printf("Error unmarshalling partial result: %v", err)
					return
				}
				if partialResult.Text != "" {
					fmt.Printf("Partial transcription: %s\r", partialResult.Text)
				}
			}
		}
	}
}
