package pubsub

type PubSub struct {
	Chans map[string][]chan bool
}

func NewQueue() *PubSub {
  return &PubSub{
    Chans: make(map[string][]chan bool),
  }
}

func (ps *PubSub) Sub(topic string) chan bool {
  ch := make(chan bool)
  ps.Chans[topic] = append(ps.Chans[topic], ch)

  return ch
}

func (ps *PubSub) Unsub(topic string, ch chan bool) {
  for i, ch2 := range ps.Chans[topic] {
    if ch2 == ch {
      ps.Chans[topic][i] = ps.Chans[topic][len(ps.Chans[topic])-1]
      ps.Chans[topic][len(ps.Chans[topic])-1] = nil
      ps.Chans[topic] = ps.Chans[topic][:len(ps.Chans[topic])-1]

      close(ch)
      break
    }
  }
}

func (ps *PubSub) Pub(topic string) {
	for _, ch := range ps.Chans[topic] {
		ch <- true
	}
}
