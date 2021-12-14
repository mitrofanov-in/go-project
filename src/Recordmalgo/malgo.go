// This example simply captures data from your default microphone until you press Enter, after which it plays back the captured audio.
package main

import (
	"fmt"
	"os"

	"github.com/aleitner/microphone"
	"github.com/faiface/beep/wav"
	"github.com/gen2brain/malgo"
)

const (
	bits = 32
	rate = 44100
)

func main() {
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

	wavOut, err := os.Create("Test.wav")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		defer wavOut.Close()
	}

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Duplex)
	deviceConfig.Capture.Format = malgo.FormatS24
	deviceConfig.Capture.Channels = 1
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1

	//var playbackSampleCount uint32
	var capturedSampleCount uint32
	pCapturedSamples := make([]byte, 0)

	stream, format, err := microphone.OpenStream(ctx, deviceConfig)
	if err != nil {
		fmt.Println(err)
	}

	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {

		sampleCount := framecount * deviceConfig.Capture.Channels * sizeInBytes

		newCapturedSampleCount := capturedSampleCount + sampleCount

		pCapturedSamples = append(pCapturedSamples, pSample...)

		/*
			_, err = wavOut.Write(pSample)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		*/

		capturedSampleCount = newCapturedSampleCount

	}

	fmt.Println("Recording...")
	captureCallbacks := malgo.DeviceCallbacks{
		Data: onRecvFrames,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, captureCallbacks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stream.Start()

	fmt.Println("Press Enter to stop recording...")
	fmt.Scanln()

	err = wav.Encode(wavOut, stream, format)
	if err != nil {
		fmt.Println(err)
	}

	/*
		ctrlc := make(chan os.Signal)
		signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)

		go func() {
			<-ctrlc
			fmt.Println("\r- Turning off microphone...")
			stream.Close()
			os.Exit(1)
		}()
	*/
	device.Uninit()
	stream.Close()
	/*

		onSendFrames := func(pSample, nil []byte, framecount uint32) {

			samplesToRead := framecount * deviceConfig.Playback.Channels * sizeInBytes
			if samplesToRead > capturedSampleCount-playbackSampleCount {
				samplesToRead = capturedSampleCount - playbackSampleCount

			}

			copy(pSample, pCapturedSamples[playbackSampleCount:playbackSampleCount+samplesToRead])

			playbackSampleCount += samplesToRead

			if playbackSampleCount == uint32(len(pCapturedSamples)) {
				playbackSampleCount = 0
			}

		}

		fmt.Println("Playing...")
		playbackCallbacks := malgo.DeviceCallbacks{
			Data: onSendFrames,
		}

		device, err = malgo.InitDevice(ctx.Context, deviceConfig, playbackCallbacks)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = device.Start()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Press Enter to quit...")
		fmt.Scanln()

		device.Uninit()
	*/
}
