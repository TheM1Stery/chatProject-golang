package chat

type ChatBroker struct {
	stopCh    chan struct{}
	publishCh chan string
	subCh     chan chan string
	unsubCh   chan chan string
}

func NewChat() *ChatBroker {
	return &ChatBroker{
		stopCh:    make(chan struct{}),
		publishCh: make(chan string, 1),
		subCh:     make(chan chan string, 1),
		unsubCh:   make(chan chan string, 1),
	}
}

func (b *ChatBroker) Start() {
	subs := map[chan string]struct{}{}
	for {
		select {
		case <-b.stopCh:
			for subCh := range subs {
				close(subCh)
			}
			return
		case msgCh := <-b.subCh:
			subs[msgCh] = struct{}{}
		case msgCh := <-b.unsubCh:
			delete(subs, msgCh)
			close(msgCh)
		case msg := <-b.publishCh:
			for msgCh := range subs {
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

// add the subscriber to the list of clients and return the receiver channel
func (b *ChatBroker) Subscribe() chan string {
	msgCh := make(chan string, 5)
	b.subCh <- msgCh
	return msgCh
}

func (b *ChatBroker) Unsubscribe(msgCh chan string) {
	b.unsubCh <- msgCh
}

func (b *ChatBroker) Publish(msg string) {
	b.publishCh <- msg
}

func (b *ChatBroker) Stop() {
}
