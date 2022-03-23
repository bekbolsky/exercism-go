package logs

/*
Application identifies the application emitting the given log.
Takes a log line and returns the application that emitted the log line.

To identify which application emitted a given log line,
search the log line for a specific character as specified by the following table:

| Application      | Character | Unicode Code Point |
|------------------|-----------|--------------------|
| `recommendation` | ❗        | `U+2757`           |
| `search`         | 🔍        | `U+1F50D`          |
| `weather`        | ☀         | `U+2600`           |

If a log line does not contain one of the characters from the above table, return `default` to the caller.
If a log line contains more than one character in the above table,
return the application corresponding to the first character
found in the log line starting from left to right. */
func Application(log string) string {
	applications := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
	result := "default"

	var foundCharacters []rune
	for _, r := range log {
		if r == '❗' || r == '🔍' || r == '☀' {
			foundCharacters = append(foundCharacters, r)
		}
	}

	if len(foundCharacters) != 0 {
		result = applications[foundCharacters[0]]
	}
	return result
}

// Replace replaces all occurrences of old with new, returning the modified log to the caller.
// Takes a log line, a corrupted character, and the original value.
// Returns a modified log line that has all occurrences of the corrupted character
// replaced with the original value.
func Replace(log string, oldRune, newRune rune) string {
	var newLog []rune
	for _, r := range log {
		if r != oldRune {
			newLog = append(newLog, r)
		} else {
			newLog = append(newLog, newRune)
		}
	}
	return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is within the limit.
// Takes a log line and character limit.
// Returns whether or not the log line is within the character limit.
func WithinLimit(log string, limit int) bool {
	var numberOfCharacters int
	for range log {
		numberOfCharacters++
	}
	return numberOfCharacters <= limit
}
