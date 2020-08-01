package leap_seconds

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseFile(path string) LeapSecondsData {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return ParseReader(file)
}

func ParseReader(input io.Reader) LeapSecondsData {
	scanner := bufio.NewScanner(input)
	var result LeapSecondsData
	for scanner.Scan() {
		line := scanner.Text()

		// #$ is the comment that denotes the line with last file update timestamp
		if strings.HasPrefix(line, "#$") {
			result.LastUpdate = extractIntFromLine(line)
			continue
		}

		// #$ is the comment that denotes the line with file expiration timestamp
		if strings.HasPrefix(line, "#@") {
			result.ExpiresOn = extractIntFromLine(line)
			continue
		}

		// # is the general purpose comment, ignore
		if strings.HasPrefix(line, "#") {
			continue
		}

		result.LeapSeconds = append(result.LeapSeconds, extractLeapSecondsFromLine(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func extractIntFromLine(line string) int64 {
	// line is a comment followed by a whitespace separator and value
	fields := strings.Fields(line)
	value, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatal(err)
	}
	return int64(value)
}

func extractLeapSecondsFromLine(line string) LeapSecond {
	// leap second line consists of the timestamp
	// followed by a whitespace separator
	// followed by the total count
	entry := strings.Fields(line)
	timestamp, err := strconv.Atoi(entry[0])
	if err != nil {
		log.Fatal(err)
	}
	count, err := strconv.Atoi(entry[1])
	if err != nil {
		log.Fatal(err)
	}
	return LeapSecond{
		AddedOn:    int64(timestamp),
		TotalCount: count,
	}
}
