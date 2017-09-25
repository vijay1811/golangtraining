package main

import (
	"fmt"
	"golangtraining/assignment8/conn"
	"golangtraining/assignment8/subscriber"
	"sync"
)

func mny(str string) bool {
	if str == "Money" {
		return true
	}
	return false
}
func pari(str string) bool {
	if str == "Paritosh" {
		return true
	}
	return false
}
func vij(str string) bool {
	if str == "Vijay" {
		return true
	}
	return false
}
func shi(str string) bool {
	if str == "Shivam" {
		return true
	}
	return false
}

func main() {
	// make a connection here
	conn := conn.NewConnection()
	// pass that to subscriber
	var wg sync.WaitGroup
	s := subscriber.NewSubscriber(conn)
	mx := &sync.Mutex{}
	//use waitgroup here
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func() {
			listenOnSubscription(s, "money", subscribe(s, mny, mx), mx)
			wg.Done()
		}()
	}
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func() {
			listenOnSubscription(s, "shivam", subscribe(s, shi, mx), mx)
			wg.Done()
		}()
	}
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func() {
			listenOnSubscription(s, "vijay", subscribe(s, vij, mx), mx)
			wg.Done()
		}()
	}
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func() {
			listenOnSubscription(s, "paritosh", subscribe(s, pari, mx), mx)
			wg.Done()
		}()
	}

	wg.Wait()
	s.Close()
}

func listenOnSubscription(s subscriber.Subscriber, name string, ch <-chan string, mx *sync.Mutex) {
	for {
		str, ok := <-ch
		if !ok {
			unsubscribe(s, ch, mx)
			fmt.Println(name + " subscription closed")
			return
		}
		fmt.Println(name + " got: " + str)
	}
}

func subscribe(s subscriber.Subscriber, filter func(string) bool, mx *sync.Mutex) <-chan string {
	return s.Subscribe(filter)
}
func unsubscribe(s subscriber.Subscriber, ch <-chan string, mx *sync.Mutex) {
	s.Unsubscribe(ch)
	return
}
