package app

import "time"

type Grinch struct {
	Age       int // In hours since creation
	Hunger    int
	Happiness int
	Energy    int
	Stage     StageType
	Created   time.Time
	LastFed   time.Time
	LastPlay  time.Time
	LastSleep time.Time
	IsAlive   bool
}

// Seconds
const HungerDecayInterval = 3
const EnergyDecayInterval = 8

type EventType int

const (
	EventFeed  EventType = iota // 1
	EventSleep                  // 2
	EventPlay                   // 3
	EventGrow                   // 4
)

type GameEvent struct {
	Timestamp time.Time
	Type      EventType
	Message   string
}

type StageType int

const (
	StageBaby     StageType = iota // 1
	StageJuvenile                  // 2
	StageAdult                     // 3
)

type Food struct {
	Name      string
	Happiness int
	Energy    int
}
