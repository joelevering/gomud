package mocks

type MockQueue struct {
  Subs []string
  Unsubs []string
  Pubs []string
}

func (m *MockQueue) Sub(topic string) chan bool {
  m.Subs = append(m.Subs, topic)
  return make(chan bool)
}

func (m *MockQueue) Unsub(topic string, ch chan bool) {
  m.Unsubs = append(m.Unsubs, topic)
  close(ch)
}

func (m *MockQueue) Pub(topic string) {
  m.Pubs = append(m.Pubs, topic)
}
