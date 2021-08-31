package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createbenefitassessmentjobrequest
type Createbenefitassessmentjobrequest struct { 
	// DivisionIds - The list of division ids for routing queues that are to be assessed for Predictive Routing benefit.
	DivisionIds *[]string `json:"divisionIds,omitempty"`

}

func (o *Createbenefitassessmentjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createbenefitassessmentjobrequest
	
	return json.Marshal(&struct { 
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		*Alias
	}{ 
		DivisionIds: o.DivisionIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Createbenefitassessmentjobrequest) UnmarshalJSON(b []byte) error {
	var CreatebenefitassessmentjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatebenefitassessmentjobrequestMap)
	if err != nil {
		return err
	}
	
	if DivisionIds, ok := CreatebenefitassessmentjobrequestMap["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createbenefitassessmentjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
