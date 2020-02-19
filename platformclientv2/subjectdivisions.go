package platformclientv2
import (
	"encoding/json"
)

// Subjectdivisions
type Subjectdivisions struct { 
	// SubjectIds - A collection of subject IDs to associate with the given divisions
	SubjectIds *[]string `json:"subjectIds,omitempty"`


	// DivisionIds - A collection of division IDs to associate with the given subjects
	DivisionIds *[]string `json:"divisionIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Subjectdivisions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
