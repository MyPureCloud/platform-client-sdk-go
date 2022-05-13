package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicschedulermessageseveritycount
type Wfmbuscheduletopicschedulermessageseveritycount struct { 
	// Severity
	Severity *string `json:"severity,omitempty"`


	// Count
	Count *int `json:"count,omitempty"`

}

func (o *Wfmbuscheduletopicschedulermessageseveritycount) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicschedulermessageseveritycount
	
	return json.Marshal(&struct { 
		Severity *string `json:"severity,omitempty"`
		
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		Severity: o.Severity,
		
		Count: o.Count,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduletopicschedulermessageseveritycount) UnmarshalJSON(b []byte) error {
	var WfmbuscheduletopicschedulermessageseveritycountMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduletopicschedulermessageseveritycountMap)
	if err != nil {
		return err
	}
	
	if Severity, ok := WfmbuscheduletopicschedulermessageseveritycountMap["severity"].(string); ok {
		o.Severity = &Severity
	}
    
	if Count, ok := WfmbuscheduletopicschedulermessageseveritycountMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicschedulermessageseveritycount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
