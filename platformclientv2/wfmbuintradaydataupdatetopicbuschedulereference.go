package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbuschedulereference
type Wfmbuintradaydataupdatetopicbuschedulereference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// WeekDate
	WeekDate *time.Time `json:"weekDate,omitempty"`

}

func (o *Wfmbuintradaydataupdatetopicbuschedulereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuschedulereference
	
	WeekDate := new(string)
	if o.WeekDate != nil {
		
		*WeekDate = timeutil.Strftime(o.WeekDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		WeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		WeekDate: WeekDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicbuschedulereference) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicbuschedulereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicbuschedulereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuintradaydataupdatetopicbuschedulereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if weekDateString, ok := WfmbuintradaydataupdatetopicbuschedulereferenceMap["weekDate"].(string); ok {
		WeekDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", weekDateString)
		o.WeekDate = &WeekDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuschedulereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
