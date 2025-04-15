package main

import (
	"sync"
	"time"
)

const (
	Follower  = "Follower"
	Candidate = "Candidate"
	Leader    = "leader"
)

type Raft struct {
	mu          sync.Mutex
	id          int
	peers       []int
	status      string
	currentTime int
	voteFor     int
	hearbeat    chan bool
	voteGranted chan bool
	leaderChan  chan bool
}

func NewRaft(id int, peers []int) *Raft {
	r := &Raft{
		id:          id,
		peers:       peers,
		status:      Follower,
		currentTime: 0,
		voteFor:     -1,
		hearbeat:    make(chan bool),
		voteGranted: make(chan bool),
		leaderChan:  make(chan bool),
	}

	go r.run()
	return r
}

func (r *Raft) run() {
	duration := 10 * time.Second
	times := time.NewTimer(duration)
	for {
		times.Reset(duration)
		switch r.status {
		case Follower:
			select {
			case <-r.hearbeat:
			case <-times.C:
				r.status = Candidate
			}
		case Candidate:
			r.startElection()
		case Leader:
			r.sendHearbeat()
		}
	}
}

func (r *Raft) startElection() {

}

func (r *Raft) sendHearbeat() {

}
