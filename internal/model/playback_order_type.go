package model

type PlaybackOrderType string

const (
	PlaybackInOrder PlaybackOrderType = "IN_ORDER"
	PlaybackReplay  PlaybackOrderType = "REPLAY"
	PlaybackRandom  PlaybackOrderType = "RANDOM"
)
