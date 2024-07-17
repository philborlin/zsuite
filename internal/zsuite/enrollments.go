package zsuite

import (
	"fmt"
	"strings"
)

type EnrollmentRequest struct {
	MemberID int                    `json:"householdMemberId"`
	Clubs    []*ClubRequest         `json:"clubEnrollments"`
	Data     *EnrollmentRequestData `json:"data"`
	Year     int                    `json:"year"`
}

type ClubRequest struct {
	ClubID       int `json:"clubId"`
	EnrollmentID int `json:"enrollmentId"`
}

type EnrollmentRequestData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Enrollment struct {
	MemberID  int
	FirstName string
	LastName  string
	Clubs     []*Club
}

type Club struct {
	ClubID       int
	EnrollmentID int
}

func (z *Zsuite) Enrollments() ([]*Enrollment, error) {
	var enrollments []*EnrollmentRequest

	err := z.get("/api/enrollments?accountId="+z.id, &enrollments)
	if err != nil {
		return nil, err
	}

	var es []*Enrollment
	for _, e := range enrollments {
		if e.Year == z.year {
			es = append(es, convertEnrollment(e))
		}
	}

	return es, nil
}

func (e *Enrollment) FullName() string {
	return fmt.Sprintf("%s %s", strings.TrimSpace(e.FirstName), strings.TrimSpace(e.LastName))
}

func convertEnrollment(er *EnrollmentRequest) *Enrollment {
	var cs []*Club
	for _, c := range er.Clubs {
		cs = append(cs, &Club{ClubID: c.ClubID, EnrollmentID: c.EnrollmentID})
	}

	return &Enrollment{
		MemberID:  er.MemberID,
		FirstName: er.Data.FirstName,
		LastName:  er.Data.LastName,
		Clubs:     cs,
	}
}
