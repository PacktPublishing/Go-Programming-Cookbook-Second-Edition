package channels

import "time"

// Sender sends "tick"" on ch until done is
// written to, then it sends "sender done."
// and exits
func Sender(ch chan string, done chan bool) {
	t := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-done:
			ch <- "sender done."
			return
		case <-t:
			ch <- "tick"
		}
	}
}
