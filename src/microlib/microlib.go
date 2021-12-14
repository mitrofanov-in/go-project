package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aleitner/microphone"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gen2brain/malgo"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "start reading microphone input",
				Action: func(c *cli.Context) error {

					ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
						fmt.Printf("LOG <%v>\n", message)
					})
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					defer func() {
						_ = ctx.Uninit()
						ctx.Free()
					}()

					f, err := os.Create("1.wav")
					if err != nil {
						log.Fatal(err)
					}

					deviceConfig := malgo.DefaultDeviceConfig(malgo.Capture)
					deviceConfig.Capture.Format = malgo.FormatS24
					deviceConfig.Capture.Channels = 2
					deviceConfig.SampleRate = 44100

					stream, format, err := microphone.OpenStream(ctx, deviceConfig)
					if err != nil {
						log.Fatal(err)
					}

					speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

					// Stop the stream when the user tries to quit the program.

					stream.Start()

					// Encode the stream. This is a blocking operation because
					// wav.Encode will try to drain the stream. However, this
					// doesn't happen until stream.Close() is called.

					err = wav.Encode(f, stream, format)
					if err != nil {
						log.Fatal(err)
					}

					ctrlc := make(chan os.Signal)
					signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)

					go func() {
						<-ctrlc
						fmt.Println("\r- Turning off microphone...")
						stream.Close()
						os.Exit(0)
					}()

					done := make(chan bool)
					speaker.Play(beep.Seq(stream, beep.Callback(func() {
						done <- true
					})))

					<-done
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
