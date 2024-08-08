/*
Package loglite provides a simple logging solution with configurable log levels.
It supports logging messages at various levels (Debug, Info, Warning, Error, Fatal)
and writes logs to a specified file.

Usage:

import "path/to/your/loglite"

	func main() {
	    loglite.Loger.Mods = loglite.Mods{
	        Debug:   true,
	        Info:    true,
	        Warning: true,
	        Error:   true,
	        Fatal:   true,
	    }
	    loglite.Start("path/to/your/logfile.log")
	    loglite.Debug("This is a debug message")
	}
*/
package loglite

import (
	"fmt"
	"os"
	"time"
)

type loger struct {
	Mods     Mods
	Queue    chan log
	FileName string
	Works    bool
}

type Mods struct {
	Debug   bool
	Info    bool
	Warning bool
	Error   bool
	Fatal   bool
}

type log struct {
	id    int
	lvl   string
	msg   string
	ptime time.Time
}

var (
	Loger  loger
	file   *os.File
	lastId int
)

func (l *loger) Start(fileName string) {
	var err error
	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		println("The Loger couldn't open/create the .log file, err: ", err.Error())
		return
	}
}

func Debug(msg string, args ...any) {
	if !Loger.Mods.Debug {
		return
	}

	lg := newLog("debug", msg, args...)
	save(lg)
}

func Info(msg string, args ...any) {
	if !Loger.Mods.Info {
		return
	}

	lg := newLog("info", msg, args...)
	save(lg)
}

func Warning(msg string, args ...any) {
	if !Loger.Mods.Warning {
		return
	}

	lg := newLog("warning", msg, args...)
	save(lg)
}

func Error(msg string, args ...any) {
	if !Loger.Mods.Error {
		return
	}

	lg := newLog("error", msg, args...)
	save(lg)
}

func Fatal(msg string, args ...any) {
	if !Loger.Mods.Fatal {
		return
	}

	lg := newLog("fatal", msg, args...)
	save(lg)
}

func newLog(lvl, msg string, args ...any) log {
	lg := log{
		id:    lastId,
		lvl:   lvl,
		msg:   fmt.Sprintf(msg, args...),
		ptime: time.Now(),
	}
	lastId++

	return lg
}

func save(lg log) {
	_, err := file.WriteString(fmt.Sprintln(lg.id, lg.lvl, lg.msg, lg.ptime.String()))
	if err != nil {
		println("The Loger couldn't insert the data of the log to the log's file, err: ", err.Error())
		return
	}
}
