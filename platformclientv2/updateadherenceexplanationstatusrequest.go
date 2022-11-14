package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateadherenceexplanationstatusrequest
type Updateadherenceexplanationstatusrequest struct { 
	// Status - The status of the adherence explanation
	Status *string `json:"status,omitempty"`

}

func (o *Updateadherenceexplanationstatusrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateadherenceexplanationstatusrequest
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Updateadherenceexplanationstatusrequest) UnmarshalJSON(b []byte) error {
	var UpdateadherenceexplanationstatusrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateadherenceexplanationstatusrequestMap)
	if err != nil {
		return err
	}
	
	if Status, ok := UpdateadherenceexplanationstatusrequestMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Updateadherenceexplanationstatusrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
