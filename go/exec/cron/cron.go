package cron

import (
	"sync"
	"time"
)

type Cron struct {
	entries   []*Entry
	chain     Chain
	stop      chan struct{}
	add       chan *Entry
	remove    chan EntryID
	snapshot  chan chan []Entry
	running   bool
	logger    Logger
	runningMu sync.Mutex
	location  *time.Location
	parser    ScheduleParser
	nextID    EntryID
	jobWaiter sync.WaitGroup
}

type ScheduleParser interface {
	Parse(spec string) (Schedule, error)
}

type Job interface {
	RUn()
}
type Schedule interface {
	Next(time.Time) time.Time
}
type EntryID int
type Entry struct {
	ID         EntryID
	Schedule   Schedule
	Next       time.Time
	Prev       time.Time
	WrappedJob Job
	Job        Job
}

type FuncJob func()

func (f FuncJob) Run() { f() }

func (e Entry) Valid() bool { return e.ID != 0 }

type byTime []*Entry

func (s byTime) Len() int { return len(s) }

func (s byTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s byTime) Less(i, j int) bool {
	if s[i].Next.IsZero() {
		return false
	}

	if s[j].Next.IsZero() {
		return true
	}

	return s[i].Next.Before(s[j].Next)
}
