package xdelta

import (
	"fmt"
	"sync"
	"time"

	lib "github.com/Giulio2002/go-xdelta/xdelta-lib"
	"github.com/dustin/go-humanize"
)

type Stats struct {
	states    map[lib.XdeltaState]time.Duration
	dataBytes map[lib.XdeltaState]uint64
	lock      sync.Mutex
}

func newStats() *Stats {
	return &Stats{
		states:    make(map[lib.XdeltaState]time.Duration),
		dataBytes: make(map[lib.XdeltaState]uint64),
	}
}

func (s *Stats) DumpToStdout() {
	fmt.Println("STATS:")

	s.lock.Lock()

	for k, v := range s.states {
		dataBytes, _ := s.dataBytes[k]

		fmt.Printf("  State %10v lastet %v and processed %v bytes (%v)\n", k, v, dataBytes, humanize.Bytes(dataBytes))
	}

	s.lock.Unlock()
}

func (s *Stats) addStateTime(state lib.XdeltaState, duration time.Duration) {
	s.lock.Lock()

	prev, _ := s.states[state]

	s.states[state] = prev + duration

	s.lock.Unlock()
}

func (s *Stats) addDataBytes(state lib.XdeltaState, amount int) {
	if amount <= 0 {
		return
	}

	s.lock.Lock()

	prev, _ := s.dataBytes[state]

	s.dataBytes[state] = prev + uint64(amount)

	s.lock.Unlock()
}
