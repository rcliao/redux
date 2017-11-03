# Go Redux

[![GoDoc](https://godoc.org/github.com/rcliao/redux?status.svg)](https://godoc.org/github.com/rcliao/redux)[![Build Status](https://travis-ci.org/rcliao/redux.svg?branch=master)](https://travis-ci.org/rcliao/redux)

Redux implementation in GoLang.

## Motivation

Follow Redux pattern for sane state management.

## Example

```go
package main

import (
	"fmt"

	"github.com/rcliao/redux"
)

func main() {
	initialState := map[string]interface{}{
		"counter": 0,
	}
	reducers := []redux.Reducer{
		func(state redux.State, action redux.Action) redux.State {
			switch action.Type {
			case "INCREMENT":
				state["counter"] = state["counter"].(int) + 1
				return state
			case "DECREMENT":
				state["counter"] = state["counter"].(int) - 1
				return state
			default:
				return state
			}
		},
	}
	store := redux.NewStore(initialState, reducers)
	store.Subscribe(func(state redux.State) {
		fmt.Println("pubsub:", state)
	})
	store.Dispatch(redux.Action{Type: "INCREMENT"}) // 1
	store.Dispatch(redux.Action{Type: "INCREMENT"}) // 2
}
```
