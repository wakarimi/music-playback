package model

type PlaybackOrderType string

const (
	InOrder PlaybackOrderType = "IN_ORDER"
	Replay  PlaybackOrderType = "REPLAY"
	Random  PlaybackOrderType = "RANDOM"
)
