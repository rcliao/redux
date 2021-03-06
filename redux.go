package redux

import "github.com/mohae/deepcopy"

// State is simply a map of name to anything
type State map[string]interface{}

// Action is sent to dispatcher to create new state
type Action struct {
	Type  string
	Value interface{}
}

// Reducer takes an action and return updated state
type Reducer func(State, Action) State

// Listener function defines subscriber to store
type Listener func(State)

// Store is the central glue between everything – takes actions dispatch such
// action to all reducers and modify state accordingly
type Store struct {
	currentState State
	reducers     []Reducer
	listeners    []Listener
}

// NewStore is a constructor pattern for creating a new store
func NewStore(initialState State, reducers []Reducer) *Store {
	return &Store{
		currentState: initialState,
		reducers:     reducers,
		listeners:    []Listener{},
	}
}

// Dispatch sends action to store and returns updated state
func (s *Store) Dispatch(action Action) State {
	for _, reducer := range s.reducers {
		copy := deepcopy.Copy(s.currentState).(State)

		s.currentState = reducer(copy, action)
		for _, listener := range s.listeners {
			listener(s.currentState)
		}
	}
	return s.currentState
}

// GetState returns the copy of the current state
func (s *Store) GetState() State {
	return deepcopy.Copy(s.currentState).(State)
}

// SetReducers updates the internal reducers
func (s *Store) SetReducers(reducers []Reducer) {
	s.reducers = reducers
}

// Subscribe allows consumer to listen for state changes
func (s *Store) Subscribe(listener Listener) {
	s.listeners = append(s.listeners, listener)
}
