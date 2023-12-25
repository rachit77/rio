package main

import (
	"fmt"
	"log"
)

func main() {
	cfg := &Config{
		ListenAddr: ":3000",
		StoreProducerFunc: func() Storer {
			return NewMemoryStore()
		},
	}
	s, err := NewServer(cfg)
	fmt.Println(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	// s.Store.Push([]byte("foobar"))
	// data, err := s.Store.Fetch(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))

	s.Start()
}
