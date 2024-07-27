package app

import (
	"os"
	"waiting_room/pkg/stream"
)

const TWITCH_LIVE_URL = "rtmp://live-ber.twitch.tv/app/"

func Run() error {
	loadEnv()

	liveUrl := os.Getenv("LIVE_URL")

	if liveUrl == "" {
		liveUrl = TWITCH_LIVE_URL
	}

	// @todo download random vod and stream if the streamer is ofline, otherwise show default video loop

	return stream.StreamVideo("./vods/stream.mp4", liveUrl+os.Getenv("LIVE_STREAM_ID"), 0)
}
