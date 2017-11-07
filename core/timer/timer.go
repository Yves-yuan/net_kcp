package timer

import (
	"math"
	"server/core/stack"
	"sync/atomic"
	"time"
)

const (
	loopForever = -1
)

var (
	// default timer backlog
	backlog = 128

	// manager manages all timers
	manager = &struct {
		incrementId    int64            // auto increment id
		timers         map[int64]*Timer // all timers
		chClosingTimer chan int64       // timer for closing
		chCreatedTimer chan *Timer
	}{}

	// precision indicates the precision of timer, default is time.Second
	precision = time.Second
)

type (
	// TimerFunc represents a function which will be called periodically in main
	// logic gorontine.
	TimerFunc func()

	TimerCondition interface {
		Check(time.Time) bool
	}

	// Timer represents a Cron job
	Timer struct {
		id        int64         // timer id
		fn        TimerFunc     // function that execute
		createAt  int64         // timer create time
		interval  time.Duration // execution interval
		condition TimerCondition
		elapse    int64 // total elapse time
		closed    int32 // is timer closed
		counter   int   // counter
	}
)

func init() {
	manager.timers = map[int64]*Timer{}
	manager.chClosingTimer = make(chan int64, backlog)
	manager.chCreatedTimer = make(chan *Timer, backlog)
}

// ID returns id of current timer
func (t *Timer) ID() int64 {
	return t.id
}

// Stop turns off a timer. After Stop, fn will not be called forever
func (t *Timer) Stop() {
	if atomic.LoadInt32(&t.closed) > 0 {
		return
	}

	// guarantee that logic is not blocked
	if len(manager.chClosingTimer) < backlog {
		manager.chClosingTimer <- t.id
		atomic.StoreInt32(&t.closed, 1)
	} else {
		t.counter = 0 // automatically closed in next Cron
	}
}

// execute job function with protection
func pexec(id int64, fn TimerFunc) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("Call timer function error, TimerID=%d, Error=%v\n%s", id, err, stack.Backtrace(2))
		}
	}()

	fn()
}

// TODO: if closing timers'count in single Cron call more than backlog will case problem.
func Cron() {
	// created/deleted timers
ctrl:
	for {
		select {
		case id := <-manager.chClosingTimer:
			delete(manager.timers, id)
		case t := <-manager.chCreatedTimer:
			manager.timers[t.id] = t
		default:
			break ctrl
		}
	}

	if len(manager.timers) < 1 {
		return
	}

	now := time.Now()
	unn := now.UnixNano()
	for id, t := range manager.timers {
		// prevent chClosingTimer exceed
		if t.counter == 0 {
			if len(manager.chClosingTimer) < backlog {
				t.Stop()
			}
			continue
		}

		// condition timer
		if t.condition != nil {
			if t.condition.Check(now) {
				pexec(t.id, t.fn)
			}
			continue
		}

		// execute job
		if t.createAt+t.elapse <= unn {
			pexec(id, t.fn)
			t.elapse += int64(t.interval)

			// update timer counter
			if t.counter != loopForever && t.counter > 0 {
				t.counter--
			}
		}
	}
}

// NewTimer returns a new Timer containing a function that will be called
// with a period specified by the duration argument. It adjusts the intervals
// for slow receivers.
// The duration d must be greater than zero; if not, NewTimer will panic.
// Stop the timer to release associated resources.
func NewTimer(interval time.Duration, fn TimerFunc) *Timer {
	return NewCountTimer(interval, loopForever, fn)
}

// NewCountTimer returns a new Timer containing a function that will be called
// with a period specified by the duration argument. After count times, timer
// will be stopped automatically, It adjusts the intervals for slow receivers.
// The duration d must be greater than zero; if not, NewTimer will panic.
// Stop the timer to release associated resources.
func NewCountTimer(interval time.Duration, count int, fn TimerFunc) *Timer {
	if fn == nil {
		panic("timer: nil timer function")
	}
	if interval <= 0 {
		panic("non-positive interval for NewTimer")
	}

	id := atomic.AddInt64(&manager.incrementId, 1)
	t := &Timer{
		id:       id,
		fn:       fn,
		createAt: time.Now().UnixNano(),
		interval: interval,
		elapse:   int64(interval), // first execution will be after interval
		counter:  count,
	}

	// add to manager
	manager.chCreatedTimer <- t
	return t
}

func NewAfterTimer(interval time.Duration, fn TimerFunc) *Timer {
	return NewCountTimer(interval, 1, fn)
}

// type TestCondition byte
//
// // Every year
// func(t TestCondition) Check(now time.Time)bool {
// 	return now.Year() == 0
// }
//
// NewCondTimer(TestCondition{}, func(){})
func NewCondTimer(condition TimerCondition, fn TimerFunc) *Timer {
	if condition == nil {
		panic("timer: nil condition")
	}

	t := NewCountTimer(time.Duration(math.MaxInt64), loopForever, fn)
	t.condition = condition

	return t
}

// SetPrecision set the ticker precision, and time precision can not less
// than a Millisecond, and can not change after application running. The default
// precision is time.Second
func SetPrecision(p time.Duration) {
	if p < time.Millisecond {
		panic("time p can not less than a Millisecond")
	}
	precision = p
}

// SetBacklog set the timer created/closing channel backlog, A small backlog
// may cause the logic to be blocked when call NewTimer/NewCountTimer/timer.Stop
// in main logic gorontine.
func SetBacklog(c int) {
	if c < 16 {
		c = 16
	}
	backlog = c
}

func Precision() time.Duration {
	return precision
}
