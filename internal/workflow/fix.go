package workflow

import (
	"github.com/joho/godotenv"
	"github.com/philborlin/zsuite/internal/zsuite"
)

func Fix(activityID int) error {
	_ = godotenv.Load(".env")

	token := getOrDefault("TOKEN", "")
	id := getOrDefault("ID", "")
	httpPrefix := getOrDefault("HTTP_PREFIX", "https://4h.zsuite.org")
	year, err := getOrDefaultInt("YEAR", 2023)
	if err != nil {
		return err
	}

	z := zsuite.New(token, id, httpPrefix, year)
	var as []*zsuite.Activity

	z.PutActivities(activityID, as)

	return nil
}
