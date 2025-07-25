package workflow

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/philborlin/zsuite/internal/zsuite"
)

func Workflow() error {
	// Do nothing on error because in production we deploy with ENV variables directly
	_ = godotenv.Load(".env")

	token := getOrDefault("TOKEN", "")
	id := getOrDefault("ID", "")
	httpPrefix := getOrDefault("HTTP_PREFIX", "https://4h.zsuite.org")
	year, err := getOrDefaultInt("YEAR", 2023)
	if err != nil {
		return err
	}

	zsuite := zsuite.New(token, id, httpPrefix, year)

	es, err := zsuite.Enrollments()
	if err != nil {
		return err
	}

	e, err := chooseEnrollment(es)
	if err != nil {
		return err
	}

	fmt.Printf("Chose enrollment: %v\n", e)

	rs, err := zsuite.RecordBooks(e.MemberID)
	if err != nil {
		return err
	}

	fmt.Printf("Found record books: %v\n", rs)

	from, err := chooseRecordBook(rs, []string{})
	if err != nil {
		return err
	}

	fmt.Printf("chooseRecordBook (1): %v\n", rs)

	to, err := chooseRecordBook(rs, []string{from.Name})
	if err != nil {
		return err
	}

	fmt.Printf("chooseRecordBook (2): %v\n", rs)

	return copyActivities(zsuite, from, to)
}

func getOrDefault(envVar string, def string) string {
	if value, ok := os.LookupEnv(envVar); ok {
		return value
	}
	return def
}

func getOrDefaultInt(envVar string, def int) (int, error) {
	if value, ok := os.LookupEnv(envVar); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return def, nil
}
