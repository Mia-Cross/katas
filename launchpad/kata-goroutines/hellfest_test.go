package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TESTS FOR findReservationNumber()

func TestFindAvailableSlotWhenNameNotInReservations(t *testing.T) {
	localResa := Reservations{Ticket{0, "arthur"}, Ticket{1, "billy"}}
	_, err := localResa.findAvailableSlot("lol")
	if err.Error() != "Name was not in reservations" {
		t.Fatal("Should have returned an error because client is not in reservation")
	}
}
func TestFindAvailableSlotWhenNameIsInReservations(t *testing.T) {
	localResa := Reservations{Ticket{0, "arthur"}, Ticket{1, "billy"}}
	i, err := localResa.findAvailableSlot("billy")
	assert.Equal(t, i, 1, "Should have returned index 1 for client billy")
	assert.Equal(t, err, nil, "Should have returned nil as error because client billy is in reservations")
}

// TESTS FOR Queue.isFull()

func TestReturnFalseBecauseQueueIsEmpty(t*testing.T) {
	localQueue := make(Queue, 1)
	defer close(localQueue)
	isFull := localQueue.isFull()
	assert.Equal(t, isFull, false, "Should have returned false because queue is empty")
}
func TestReturnFalseBecauseQueueHasPlace(t*testing.T) {
	localQueue := make(Queue, 5)
	defer close(localQueue)
	localQueue <- "jake"
	isFull := localQueue.isFull()
	assert.Equal(t, isFull, false, "Should have returned false because queue still has some space")
}
func TestReturnTrueBecauseQueueIsFull(t*testing.T) {
	localQueue := make(Queue, 5)
	defer close(localQueue)
	localQueue <- "jake"
	localQueue <- "esther"
	localQueue <- "wanda"
	localQueue <- "belinda"
	localQueue <- "sydney"
	isFull := localQueue.isFull()
	assert.Equal(t, isFull, true, "Should have returned true because queue is full")
}

// TESTS FOR Queue.clientAlreadyInQueue()
