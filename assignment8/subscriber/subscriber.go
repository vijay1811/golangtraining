package subscriber

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type subscriptionContainer struct {
	filter  func(string) bool
	subChan chan string
}

type subscriberImpl struct {
	subscriptions []*subscriptionContainer
}

// NewSubscriber creates and returns a slice and keeps on
// reading(from connetcion ) into it through a goroutine
func NewSubscriber(c net.Conn) *subscriberImpl {
	subImpl := &subscriberImpl{
		subscriptions: make([]*subscriptionContainer, 0, 0),
	}
	go func() {
		subImpl.readFromConnection(c)
	}()
	return subImpl
}

func (s *subscriberImpl) readFromConnection(conn net.Conn) {
	rdr := bufio.NewReader(conn)
	for {
		time.Sleep(1 * time.Second)
		line, err := rdr.ReadString('\n')
		// if EOF close all sub channels and return - Done
		if err != nil {
			fmt.Println(err)
			for _, sub := range s.subscriptions {
				close(sub.subChan)
			}
			return
		}
		fmt.Println(line)
		line = strings.TrimSpace(line)
		// use mutex here - Done
		mx := &sync.Mutex{}

		mx.Lock()
		timeOutChan := make(chan struct{})
		wg := new(sync.WaitGroup)
		for _, sub := range s.subscriptions {
			if sub.filter(line) {
				//	send line to channel - Done
				wg.Add(1)
				go func() {
					defer wg.Done()

					select {
					case sub.subChan <- line:
					case <-timeOutChan:
						fmt.Println("dropping message ")
					}
				}()
				// implement "send unblocked after 1 second logic" - Done
			}
		}

		mx.Unlock()
		subDoneChan := make(chan struct{})
		go func() {
			wg.Wait()
			close(subDoneChan)
		}()
		timer := time.NewTimer(2 * time.Second)
		select {
		case <-timer.C:
			close(timeOutChan)
			<-subDoneChan
		case <-subDoneChan:
		}

	}
}

func (s *subscriberImpl) Subscribe(filter func(string) bool) <-chan string {
	m := &sync.Mutex{}
	subContainer := &subscriptionContainer{
		filter:  filter,
		subChan: make(chan string),
	}
	m.Lock()
	s.subscriptions = append(s.subscriptions, subContainer)
	m.Unlock()
	return subContainer.subChan
}

func (s *subscriberImpl) Unsubscribe(ch <-chan string) {
	m := &sync.Mutex{}
	m.Lock()
	for idx, subConainer := range s.subscriptions {
		if subConainer.subChan == ch {
			s.subscriptions[idx] = s.subscriptions[len(s.subscriptions)-1]
			s.subscriptions[len(s.subscriptions)-1] = nil
			s.subscriptions = s.subscriptions[:len(s.subscriptions)-1]
			m.Unlock()
			return
		}
	}
}

func (s *subscriberImpl) Close() error {
	m := &sync.Mutex{}
	m.Lock()
	for _, subContainer := range s.subscriptions {
		close(subContainer.subChan)
	}
	m.Unlock()
	return nil
}
