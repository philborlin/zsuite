package workflow

import (
	"fmt"
	"slices"

	"github.com/manifoldco/promptui"
	"github.com/philborlin/zsuite/internal/zsuite"
)

func chooseRecordBook(rs []*zsuite.RecordBook, ignore []string) (*zsuite.RecordBook, error) {
	var rs2 []*zsuite.RecordBook
	var ss []string
	for _, r := range rs {
		if !slices.Contains(ignore, r.Name) {
			ss = append(ss, fmt.Sprintf("%s (%d)", r.Name, len(r.Activities)))
			rs2 = append(rs2, r)
		}
	}

	prompt := promptui.Select{
		Label: "Select Record Book to copy activities to other Record Books",
		Items: ss,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	return rs2[i], nil
}

func copyActivities(z *zsuite.Zsuite, from, to *zsuite.RecordBook) error {
	for i, a := range from.Activities {
		z.UpdateActivity(to.ID, i, a)
	}

	as, err := z.Activities(to.ID)
	if err != nil {
		return err
	}

	to.Activities = as

	return nil
}
