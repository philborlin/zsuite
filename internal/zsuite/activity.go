package zsuite

import (
	"fmt"
)

type Activity struct {
	Date     string `json:"date"`
	Learn    string `json:"learn"`
	Activity string `json:"activity"`
}

func (z *Zsuite) Activities(recordBook int) ([]*Activity, error) {
	var activities []*Activity

	err := z.get(fmt.Sprintf("/api/record-books/%d/entries/activity", recordBook), &activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (z *Zsuite) UpdateActivity(recordBook int, index int, a *Activity) error {
	fmt.Printf("PUT /api/record-books/%d/entries/activity/%d - %v\n", recordBook, index, a)

	err := z.put(fmt.Sprintf("/api/record-books/%d/entries/activity/%d", recordBook, index), a)
	if err != nil {
		return err
	}

	return nil
}
