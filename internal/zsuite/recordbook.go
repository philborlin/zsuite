package zsuite

import (
	"fmt"
	"strings"
)

type RecordBook struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Activities []*Activity
}

func (z *Zsuite) RecordBooks(memberID int) ([]*RecordBook, error) {
	var recordBooks []*RecordBook

	err := z.get(
		fmt.Sprintf("/api/record-books?accountId=%s&householdMemberId=%d&orgId=4h", z.id, memberID),
		&recordBooks,
	)
	if err != nil {
		return nil, err
	}

	var filtered []*RecordBook
	for _, r := range recordBooks {
		n := strings.TrimSpace(r.Name)
		if n != "Involvement Report" && n != "Short Term Record Books" {
			a, err := z.Activities(r.ID)
			if err != nil {
				return nil, err
			}
			r.Activities = a

			filtered = append(filtered, r)
		}
	}

	return filtered, nil
}
