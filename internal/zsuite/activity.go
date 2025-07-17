package zsuite

import (
	"fmt"
)

type Activity struct {
	Date     string `json:"date"`
	Learn    string `json:"learn"`
	Activity string `json:"activity"`
}

func (z *Zsuite) PutActivities(recordBook int, activities []*Activity) error {
	err := z.put(fmt.Sprintf("/api/record-books/%d/entries/activity", recordBook), &activities)
	if err != nil {
		return err
	}

	return nil
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
	s := fmt.Sprintf("/api/record-books/%d/entries/activity/%d", recordBook, index)
	fmt.Printf("PUT %s - %v\n", s, a)

	err := z.put(s, a)
	if err != nil {
		return err
	}

	return nil
}
