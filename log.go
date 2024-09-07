package loglite

import (
	"fmt"
	"os"
	"time"
)

// loger struct: Handles the logger's state, including log level, file path, file handler, and internal counters.
type loger struct {
	Lvl       int
	FilePath  string
	file      *os.File
	lastId    int
	waitCount int
	listen    bool
}

// Log Levels: Constants Debug, Info, Warning, and Error control log severity.
const (
	Debug   = -4
	Info    = 0
	Warning = 4
	Error   = 8
)

// log struct: Represents a log entry with a unique ID, log level, message, and timestamp.
type log struct {
	id    int
	lvl   string
	msg   string
	ptime time.Time
}

// NewLogger(lvl int, filePath string) *loger: Initializes a new logger, sets the log level, and opens/creates a log file.
func NewLogger(lvl int, filePath string) *loger {
	l := &loger{
		Lvl:      lvl,
		FilePath: filePath,
		listen:   true,
	}

	var err error
	l.file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic("The Loger couldn't open/create the .log file, err: " + err.Error())
	}

	return l
}

// Debug(msg string, args ...any): Logs debug messages.
func (l *loger) Debug(msg string, args ...any) {
	if l.Lvl > Debug || !l.listen {
		return
	}

	l.waitCount++

	lg := l.newLog("debug", msg, args...)
	go l.save(lg)
}

// Info(msg string, args ...any): Logs informational messages.
func (l *loger) Info(msg string, args ...any) {
	if l.Lvl > Info || !l.listen {
		return
	}

	l.waitCount++

	lg := l.newLog("info", msg, args...)
	go l.save(lg)
}

// Warning(msg string, args ...any): Logs warning messages.
func (l *loger) Warning(msg string, args ...any) {
	if l.Lvl > Warning || !l.listen {
		return
	}

	l.waitCount++

	lg := l.newLog("warning", msg, args...)
	go l.save(lg)
}

// Error(msg string, args ...any): Logs error messages.
func (l *loger) Error(msg string, args ...any) {
	if l.Lvl > Error || !l.listen {
		return
	}

	l.waitCount++

	lg := l.newLog("error", msg, args...)
	go l.save(lg)
}

// Fatal(msg string, args ...any): Logs a fatal error and stops the logger.
func (l *loger) Fatal(msg string, args ...any) {
	lg := l.newLog("fatal", msg, args...)

	l.listen = false

	for {
		if l.waitCount == 0 {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}

	panic(l.save(lg))

}

// newLog(lvl, msg string, args ...any) log: Creates a new log entry with a unique ID, log level, and formatted message.
func (l *loger) newLog(lvl, msg string, args ...any) log {
	lg := log{
		id:    l.lastId,
		lvl:   lvl,
		msg:   fmt.Sprintf(msg, args...),
		ptime: time.Now(),
	}

	l.lastId++

	return lg
}

// save(lg log) string: Writes the log entry to the file.
func (l *loger) save(lg log) string {

	lgText := fmt.Sprintln("{", lg.id, ",\n", lg.lvl, ",\n", lg.msg, "\n", lg.ptime.Format("2001-01-01 01:01:01.001"), "}")

	_, err := l.file.WriteString(lgText)

	if err != nil {
		println("The Loger couldn't insert the data of the log to the log's file, err: ", err.Error())
	}

	l.waitCount--

	return lgText
}
