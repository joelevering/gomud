package pubsub

import "testing"

func Test_PubSendsToSubs(t *testing.T) {
  q := NewQueue()

  ch := q.Sub("test")
  defer close(ch)

	go q.Pub("test")

	res := <-ch

  if !res {
    t.Error("Expected to get result when publishing to a sub'd channel, but didn't")
  }
}

func Test_UnsubRemovesAndClosesChan(t *testing.T) {
  q := NewQueue()

  ch := q.Sub("test")
  q.Unsub("test", ch)

  q.Pub("test") // shouldn't block because there should be no chans for topic
}
