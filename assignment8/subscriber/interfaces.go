package subscriber

type Subscriber interface {
	Subscribe(filter func(string) bool) <-chan string
	Unsubscribe(<-chan string)
	Close() error
}
