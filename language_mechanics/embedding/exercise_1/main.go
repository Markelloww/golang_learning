package main

import (
	"fmt"
	"log"
	"time"
)

type Document struct {
	Key   string
	Title string
}

type Feed struct{}

func (f *Feed) Count() int {
	return 42
}

func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)

	doc := Document{
		Key:   key,
		Title: "Title for " + key,
	}
	return doc, nil
}

type CachingFeed struct {
	docs map[string]Document
	*Feed
}

func NewCachingFeed(f *Feed) *CachingFeed {
	return &CachingFeed{
		docs: make(map[string]Document),
		Feed: f,
	}
}

func (cf *CachingFeed) Fetch(key string) (Document, error) {
	if doc, ok := cf.docs[key]; ok {
		return doc, nil
	}

	doc, err := cf.Feed.Fetch(key)
	if err != nil {
		return Document{}, err
	}

	cf.docs[key] = doc
	return doc, nil
}

type FetchCounter interface {
	Fetch(key string) (Document, error)
	Count() int
}

func process(fc FetchCounter) {
	fmt.Printf("There are %d documents\n", fc.Count())

	keys := []string{"a", "a", "a", "b", "b", "b"}

	for _, key := range keys {
		doc, err := fc.Fetch(key)
		if err != nil {
			log.Printf("Could not fetch %s : %v", key, err)
			return
		}

		fmt.Printf("%s : %v\n", key, doc)
	}
}

func main() {
	fmt.Println("Using Feed directly")
	process(&Feed{})

	fmt.Println("Using CachingFeed")
	c := NewCachingFeed(&Feed{})
	process(c)
}
