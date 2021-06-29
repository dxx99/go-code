package main

import "time"

//func Fetch(url string) (items []Item, next time.Time, err error)

// Item a subset(子集) of Rss fields
type Item struct {
	Title   string
	Channel string
	Guid    string
}

type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

type Subscription interface {
	Updates() <-chan Item // stream of Items
	Close() error         // shuts down the stream
}

func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
		closing: make(chan chan error),
	}
	return s
}

type sub struct {
	fetcher Fetcher
	updates chan Item
	closing chan chan error
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

//func Merge(subs ...Subscription) Subscription {
//
//}

func main() {
	//todo
}
