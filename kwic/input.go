package kwic

import "os"
import "log"
import "bufio"

func ReadFile(file_name string) ([]string, error) {
	var lines []string

	file, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
