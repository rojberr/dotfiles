package goscripts

import (
	"fmt"
	"strings"

	"github.com/bitfield/script"
	"github.com/mgutz/ansi"
)

type LogManager struct {
	logEntries []*LogEntry
}

type LogCategory int64

const (
	Todo LogCategory = iota
	Failure
	Success
	End
	Default
	Warning
)

func (lt LogCategory) String() string {
	switch lt {
	case Todo:
		return ansi.ColorCode("cyan+b") + "💠 "
	case Failure:
		return ansi.ColorCode("red+b") + "❌ "
	case Success:
		return ansi.ColorCode("green+b") + "✅ "
	case Warning:
		return ansi.ColorCode("yellow+b") + "⚠️ "
	case End:
		return ansi.ColorCode("magenta+b") + "🎆🎆🎆 "
	case Default:
		return ""
	}
	return "unknown"
}

type LogEntry struct {
	category LogCategory
	message  string
}

func (le LogEntry) String() string {
	return fmt.Sprint(le.category, le.message, ansi.ColorCode("reset"))
}

// AddLog appends a log entry with a given category and message. Optionally suppresses output.
func (lm *LogManager) AddLog(category LogCategory, msg string, silent ...bool) {
	quiet := false
	entry := LogEntry{
		category: category,
		message:  msg,
	}
	lm.logEntries = append(lm.logEntries, &entry)

	if len(silent) > 0 {
		quiet = silent[0]
	}
	if !quiet {
		fmt.Println(entry)
	}
}

// AppendFrom merges logs from another LogManager instance.
func (lm *LogManager) AppendFrom(other LogManager) {
	lm.logEntries = append(lm.logEntries, other.logEntries...)
}

// RunCommand executes a shell command and logs the result.
func (lm *LogManager) RunCommand(cmd string) error {
	scr := script.Exec(cmd)

	if scr.Error() != nil {
		lm.AddLog(Failure, fmt.Sprintf("Command (%s...) failed!", cmd[:10]))
		lm.AddLog(Failure, fmt.Sprint(scr.Error()))
		lastErr := scr.Error()
		scr.SetError(nil)
		lastMsg, _ := scr.String()
		lm.AddLog(Failure, fmt.Sprint(strings.Replace(lastMsg, "\n", "", -1)))
		return lastErr
	}
	lm.AddLog(Default, fmt.Sprint(scr.String()))
	return nil
}
