package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Schedulermessageseveritycount
type Schedulermessageseveritycount struct { 
	// Severity - The schedule message severity
	Severity *string `json:"severity,omitempty"`


	// Count - The number of schedule messages with the given severity
	Count *int `json:"count,omitempty"`

}

func (o *Schedulermessageseveritycount) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Schedulermessageseveritycount
	
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

func (o *Schedulermessageseveritycount) UnmarshalJSON(b []byte) error {
	var SchedulermessageseveritycountMap map[string]interface{}
	err := json.Unmarshal(b, &SchedulermessageseveritycountMap)
	if err != nil {
		return err
	}
	
	if Severity, ok := SchedulermessageseveritycountMap["severity"].(string); ok {
		o.Severity = &Severity
	}
	
	if Count, ok := SchedulermessageseveritycountMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Schedulermessageseveritycount) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
