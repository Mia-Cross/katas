package main

// add thread so only 5 people at a time can take places
// 10 sec per places
// erreur file pleine
// requete get /queue qui renvoie les 5 personne en train de prendre leur place

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var MAX_TICKETS = 50
var MAX_CLI_IN_QUEUE = 5
var TIME_TO_RESERVE = 10

type Ticket struct {
	Id   int
	Name string
}

type Reservations []Ticket

var r Reservations

type Queue chan string

var q Queue

func (r Reservations) findAvailableSlot(name string) (int, error) {
	for i, ticket := range r {
		fmt.Println(fmt.Sprintf("scanning index %d : %s", i, ticket.Name))
		if ticket.Name == name {
			return i, nil
		}
	}
	return len(r), errors.New("Name was not in reservations")
}

func (q Queue) isFull() bool {
	if len(q) == MAX_CLI_IN_QUEUE {
		return true
	}
	return false
}

func (q Queue) clientAlreadyInQueue(name string) bool {
	for clientName := range q {
		if clientName == name {
			return true
		}
	}
	return false
}

func (r Reservations) buyTicketIfAvailable(name string) Ticket {
	if len(r) >= MAX_TICKETS {
		log.Printf("Sorry, all %d tickets were sold :(", MAX_TICKETS)
		return Ticket{}
	}
	r = append(r, Ticket{len(r), name})
	return r[len(r)-1]
}

func putClientInQueue(name string) (Ticket, error) {
	id, err := r.findAvailableSlot(name)
	if err != nil {
		return Ticket{}, errors.New(fmt.Sprintf("There is already a place for %s (number %d).", name, id))
	}
	if q.isFull() {
		return Ticket{}, fmt.Errorf("Queue is full, please retry later...")
	}
	if q.clientAlreadyInQueue(name) {
		return Ticket{}, fmt.Errorf("Client is already queuing")
	}
	q <- name
	time.Sleep(time.Duration(TIME_TO_RESERVE) * time.Second)
	// time.Sleep(TIME_TO_RESERVE * time.Second)
	ticket := r.buyTicketIfAvailable(name)
	<-q
	r[id] = ticket
	return ticket, nil
}

func makeReservation(w http.ResponseWriter, req *http.Request) {
	names, ok := req.URL.Query()["name"]
	name := names[0]
	if !ok {
		log.Fatal("Need a name to reserve a place\n")
	}
	ticket, err := putClientInQueue(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		// data := SomeStruct{}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ticket)
		// enc := json.NewEncoder(w)
		// enc.Encode(r)
		// err := json.NewEncoder(w).Encode(r)
		// formatted, err := json.Marshal(r)
		// if err != nil {
		// w.Write([]byte(ticket.Name)
		// }
	}
}

func main() {
	http.HandleFunc("/place", makeReservation)
	http.ListenAndServe(":8080", nil)
}
