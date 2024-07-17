package workflow

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/philborlin/zsuite/internal/zsuite"
)

func chooseEnrollment(es []*zsuite.Enrollment) (*zsuite.Enrollment, error) {
	var ss []string
	for _, e := range es {
		ss = append(ss, e.FullName())
	}

	prompt := promptui.Select{
		Label: "Select Enrollment",
		Items: ss,
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	return es[i], nil
}
