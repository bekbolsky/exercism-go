package parsinglogfiles

import (
	"fmt"
	"regexp"
)

// IsValidLine returns true if the line is a valid log line.
// To be considered valid a line should begin with one of the following strings:
// [TRC]
// [DBG]
// [INF]
// [WRN]
// [ERR]
// [FTL]
func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

// SplitLogLine splits the lines of a log file into a slice of strings.
// Each line is a string that is terminated by a newline character.
// Any string that has a first character of "<" and a last character of ">" and
// any combination of the following characters "~", "*", "=" and "-" in between.
func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

// CountQuotedPasswords returns the number of references to word "password" in quoted text.
// It identifies log lines where the string "password",
// which may be in any combination of upper or lower case, is surrounded by quotation marks.
// It accounts for the possibility of additional content between
// the quotation marks before and after "password".
// Each line will contain at most two quotation marks.
func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`(?i)\"+[^"]*\bpassword\b\"+`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText takes a string and remove the end-of-line text and returns a clean string.
// Lines not containing end-of-line text should be returned unmodified.
// Just removes the end of line string. Do not attempts to adjust the whitespaces.
func RemoveEndOfLineText(text string) string {
	// re, _ := regexp.Compile(`[end-of-line]/d`)
	re, _ := regexp.Compile(`end-of-line[\d]*`)
	return re.ReplaceAllString(text, "")
}

// TagWithUserName takes a slice of string and returns a slice of string with the user name tagged.
// It processes log lines:
// - Lines that do not contain the string "User " remain unchanged.
// - For lines that contain the string "User ", prefix the line with [USR] followed by the user name.
func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`(User[\s]+){1}([^\s]\w+)`)
	for i, line := range lines {
		sb := re.FindStringSubmatch(line)
		if len(sb) > 0 {
			lines[i] = fmt.Sprintf("[USR] %s %s", sb[2], line)
		}
	}
	return lines
}
