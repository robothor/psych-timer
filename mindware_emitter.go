package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// We need to emit files in Mindware format
// https://support.mindwaretech.com/knowledge-base/kb0080/

// The first 2 rows of a MindWare event file must contain specific values
// for these columns:
//
// Header – This row contains the labels for each of the columns specified
//          above
// Start Event – This event must correspond with the start of the data file,
//               as it is the event from which all subsequent events will
//               calculate their relative offset from the beginning of the
//               data file

const (
	DATE_FORMAT = "01/02/2006"      //  Spec: MM/DD/YYYY
	TIME_FORMAT = "03:04:05.000 PM" // Spec: HH:MM:SS.fff AM/PM
)

type mindwareEvent struct {
	EventType string
	Name      string
	Timestamp time.Time
}

type MindwareFile struct {
	file   *os.File
	events []*mindwareEvent
	mutex  *sync.Mutex
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewMindwareFile(name string) *MindwareFile {
	f, err := os.Create(name)
	check(err)

	m := &MindwareFile{
		file:  f,
		mutex: &sync.Mutex{},
	}
	m.writeHeader()

	return m
}

func (f *MindwareFile) writeHeader() {
	fmt.Fprintln(f.file, "Event Type\tName\tDate\tTime")
	f.WriteEvent("Start Event", "")
}

func (f *MindwareFile) WriteEvent(eventType, name string) {
	e := &mindwareEvent{
		EventType: eventType,
		Name:      name,
		Timestamp: time.Now().UTC(),
	}
	f.mutex.Lock()
	f.events = append(f.events, e)
	fmt.Fprintf(f.file, "%s\t%s\t%s\t%s\n",
		e.EventType,
		e.Name,
		e.Timestamp.Format(DATE_FORMAT),
		e.Timestamp.Format(TIME_FORMAT))
	f.mutex.Unlock()
}

func (f *MindwareFile) Close() error {

	return f.file.Close()
}
