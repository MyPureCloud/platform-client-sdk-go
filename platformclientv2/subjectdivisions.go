package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Subjectdivisions
type Subjectdivisions struct { 
	// SubjectIds - A collection of subject IDs to associate with the given divisions
	SubjectIds *[]string `json:"subjectIds,omitempty"`


	// DivisionIds - A collection of division IDs to associate with the given subjects
	DivisionIds *[]string `json:"divisionIds,omitempty"`

}

func (o *Subjectdivisions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Subjectdivisions
	
	return json.Marshal(&struct { 
		SubjectIds *[]string `json:"subjectIds,omitempty"`
		
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		*Alias
	}{ 
		SubjectIds: o.SubjectIds,
		
		DivisionIds: o.DivisionIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Subjectdivisions) UnmarshalJSON(b []byte) error {
	var SubjectdivisionsMap map[string]interface{}
	err := json.Unmarshal(b, &SubjectdivisionsMap)
	if err != nil {
		return err
	}
	
	if SubjectIds, ok := SubjectdivisionsMap["subjectIds"].([]interface{}); ok {
		SubjectIdsString, _ := json.Marshal(SubjectIds)
		json.Unmarshal(SubjectIdsString, &o.SubjectIds)
	}
	
	if DivisionIds, ok := SubjectdivisionsMap["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Subjectdivisions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
