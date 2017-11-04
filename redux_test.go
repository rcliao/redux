package redux

import "testing"

func TestInitialState(t *testing.T) {
	initialState := map[string]interface{}{
		"test": true,
	}
	store := NewStore(initialState, []Reducer{})
	if store.currentState["test"].(bool) != true {
		t.Error("initial state wasn't set correctly", store.currentState)
	}
	if _, okay := store.currentState["hello"]; okay {
		t.Error("initial state shouldn't add any other value yet", store.currentState)
	}
}

func TestCounter(t *testing.T) {
	initialState := map[string]interface{}{
		"counter": 0,
	}
	reducers := []Reducer{
		func(s State, a Action) State {
			switch a.Type {
			case "INCREMENT":
				s["counter"] = s["counter"].(int) + 1
				return s
			case "DECREMENT":
				s["counter"] = s["counter"].(int) - 1
				return s
			default:
				return s
			}
		},
	}
	subCounter := 0
	store := NewStore(initialState, reducers)
	store.Subscribe(func(s State) {
		subCounter++
	})
	store.Dispatch(Action{Type: "INCREMENT"})
	if store.currentState["counter"].(int) != 1 {
		t.Error("should increment state to 1", store.currentState)
	}
	if initialState["counter"].(int) != 0 {
		t.Error("should not mutate initial state", initialState)
	}
	store.Dispatch(Action{Type: "INCREMENT"})
	if store.currentState["counter"].(int) != 2 {
		t.Error("should increment state to 1", store.currentState)
	}
	if initialState["counter"].(int) != 0 {
		t.Error("should not mutate initial state", initialState)
	}
	store.Dispatch(Action{Type: "DECREMENT"})
	if store.currentState["counter"].(int) != 1 {
		t.Error("should increment state to 1", store.currentState)
	}
	if initialState["counter"].(int) != 0 {
		t.Error("should not mutate initial state", initialState)
	}
	if subCounter != 3 {
		t.Error("should notify subscriber three times", subCounter)
	}
	if store.GetState()["counter"] != 1 {
		t.Error("should get the updated state correctly", store.GetState())
	}
	store.SetReducers([]Reducer{})
	store.Dispatch(Action{Type: "INCREMENT"})
	if store.GetState()["counter"] != 1 {
		t.Error("should not update state after removing reducer", store.GetState())
	}
}
