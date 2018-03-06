package co

import (
	"sync"
)

// Runner to run and manage life-cycle of several async tasks.
type Runner struct {
	wg sync.WaitGroup
}

// Go run f in go routine.
func (r *Runner) Go(f func()) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		f()
	}()
}

// Wait wait for all go routines started by 'Go' done.
func (r *Runner) Wait() {
	r.wg.Wait()
}

// Done return the done channel for exiting of all go routines.
func (r *Runner) Done() chan struct{} {
	done := make(chan struct{})
	go func() {
		r.Wait()
		close(done)
	}()
	return done
}
