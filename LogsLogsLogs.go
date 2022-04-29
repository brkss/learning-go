package logs

import "fmt"
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
    for _, c := range log {
        switch c {
            case 10071:
        		return "recommendation";
            case 128269:
        		return "search";
            case 9728:
        		return "weather"
        }
    }
	return "default";
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string;
    for _, c := range log {
        if c == oldRune {
			c = newRune;
        }
    	newLog = fmt.Sprintf("%s%c", newLog, c);
    }
	return newLog;
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    if utf8.RuneCountInString(log) > limit {
        return false;
    }
	return true;
}

