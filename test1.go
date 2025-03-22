package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"playsee.co/interview/api/request"
	"playsee.co/interview/domain"
)

func Test1(w http.ResponseWriter, r *http.Request) {
	var req request.TestRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var head *domain.Node
	var current *domain.Node

	for _, val := range req.Array {
		newNode := &domain.Node{Value: val}
		if head == nil {
			head = newNode
			current = head
		} else {
			current.Next = newNode
			current = current.Next
		}
	}

	current = head
	index := 0
	for current != nil {
		switch {
		case index == 0:
			fmt.Fprintf(w, "head -> %v\n", current.Value)
		case current.Next == nil:
			fmt.Fprintf(w, "tail -> %v\n", current.Value)
		default:
			fmt.Fprintf(w, "node%d -> %v\n", index, current.Value)
		}
		current = current.Next
		index++
	}
}
