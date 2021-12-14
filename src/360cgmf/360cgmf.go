package main

import (
	"log"

	. "github.com/3d0c/gmf"
)

func fatal(err error) {
	log.Fatal(err)
}

func main() {
	/// input
	mic, _ := NewInputCtxWithFormatName("default", "alsa")
	mic.Dump()

	ast, err := mic.GetBestStream(AVMEDIA_TYPE_AUDIO)
	if err != nil {
		log.Fatal("failed to find audio stream")
	}
	cc := ast.CodecCtx()

	/// fifo
	fifo := NewAVAudioFifo(cc.SampleFmt(), cc.Channels(), 1024)
	if fifo == nil {
		log.Fatal("failed to create audio fifo")
	}

	/// output
	codec, err := FindEncoder("libmp3lame")
	if err != nil {
		log.Fatal("find encoder error:", err.Error())
	}

	occ := NewCodecCtx(codec)
	if occ == nil {
		log.Fatal("new output codec context error:", err.Error())
	}
	defer Release(occ)

	occ.SetSampleFmt(AV_SAMPLE_FMT_S16P).
		SetSampleRate(cc.SampleRate()).
		SetChannels(cc.Channels()).
		SetBitRate(128e3)
	channelLayout := occ.SelectChannelLayout()
	occ.SetChannelLayout(channelLayout)

	if err := occ.Open(nil); err != nil {
		log.Fatal("can't open output codec context", err.Error())
		return
	}

	/// resample
	options := []*Option{
		{"in_channel_count", cc.Channels()},
		{"out_channel_count", cc.Channels()},
		{"in_sample_rate", cc.SampleRate()},
		{"out_sample_rate", cc.SampleRate()},
		{"in_sample_fmt", SampleFmt(cc.SampleFmt())},
		{"out_sample_fmt", SampleFmt(AV_SAMPLE_FMT_S16P)},
	}

	swrCtx := NewSwrCtx(options, occ)
	if swrCtx == nil {
		log.Fatal("unable to create Swr Context")
	}

	/// mp3 file
	outputCtx, err := NewOutputCtx("test.mp3")
	if err != nil {
		log.Fatal("new output fail", err.Error())
		return
	}

	ost := outputCtx.NewStream(codec)
	if ost == nil {
		log.Fatal("Unable to create stream for [%s]\n", codec.LongName())
	}
	defer func() {
		Release(ost)
	}()

	ost.SetCodecCtx(occ)

	if err := outputCtx.WriteHeader(); err != nil {
		log.Fatal(err.Error())
	}

	count := 0
	for packet := range mic.GetNewPackets() {
		srcFrame, err := packet.Frames(ast.CodecCtx())
		Release(packet)
		if err != nil {
			log.Println("capture audio error:", err)
			continue
		}

		wrote := fifo.Write(srcFrame)
		count += wrote
		exit := false

		for fifo.SamplesToRead() >= 1152 {
			winFrame := fifo.Read(1152)
			dstFrame := swrCtx.Convert(winFrame)
			Release(winFrame)

			if dstFrame == nil {
				continue
			}

			writePacket, err := dstFrame.Encode(occ)
			if err == nil {
				if err := outputCtx.WritePacket(writePacket); err != nil {
					log.Println("write packet err", err.Error())
				}

				Release(writePacket)

				if count < int(cc.SampleRate())*10 {
					break
				} else { //exit
					exit = true
					writePacket, err = dstFrame.Encode(occ)
				}
			} else {
				fatal(err)
			}
			Release(dstFrame)
		}
		Release(srcFrame)
		if exit {
			break
		}

	}
}
