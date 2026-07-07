package main

import (
	"fmt"
)

type Post struct {
	ID int `json:"id"`
}

func MergeFeeds(a, b []Post) []Post {
	result := make([]Post, len(a), len(a)+len(b))
	copy(result, a)
	for _, post := range b {
		result = append(result, post)
	}
	return result
}

func main() {

	a := []Post{
		Post{ID: 12},
		Post{ID: 13},
		Post{ID: 14},
	}

	b := []Post{
		Post{ID: 15},
		Post{ID: 16},
		Post{ID: 17},
	}
	fmt.Println(MergeFeeds(a, b))
}
