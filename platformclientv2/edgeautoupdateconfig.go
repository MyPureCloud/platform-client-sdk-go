package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeautoupdateconfig
type Edgeautoupdateconfig struct { 
	// TimeZone - The timezone of the window in which any updates to the edges assigned to the site can be applied. The minimum size of the window is 2 hours.
	TimeZone *string `json:"timeZone,omitempty"`


	// Rrule - The recurrence rule for updating the Edges assigned to the site. The only supported frequencies are daily and weekly. Weekly frequencies require a day list with at least oneday specified. All other configurations are not supported.
	Rrule *string `json:"rrule,omitempty"`


	// Start - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	Start *time.Time `json:"start,omitempty"`


	// End - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	End *time.Time `json:"end,omitempty"`

}

func (o *Edgeautoupdateconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeautoupdateconfig
	
	Start := new(string)
	if o.Start != nil {
		*Start = timeutil.Strftime(o.Start, "%Y-%m-%dT%H:%M:%S.%f")
		
	} else {
		Start = nil
	}
	
	End := new(string)
	if o.End != nil {
		*End = timeutil.Strftime(o.End, "%Y-%m-%dT%H:%M:%S.%f")
		
	} else {
		End = nil
	}
	
	return json.Marshal(&struct { 
		TimeZone *string `json:"timeZone,omitempty"`
		
		Rrule *string `json:"rrule,omitempty"`
		
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		*Alias
	}{ 
		TimeZone: o.TimeZone,
		
		Rrule: o.Rrule,
		
		Start: Start,
		
		End: End,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgeautoupdateconfig) UnmarshalJSON(b []byte) error {
	var EdgeautoupdateconfigMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeautoupdateconfigMap)
	if err != nil {
		return err
	}
	
	if TimeZone, ok := EdgeautoupdateconfigMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	
	if Rrule, ok := EdgeautoupdateconfigMap["rrule"].(string); ok {
		o.Rrule = &Rrule
	}
	
	if startString, ok := EdgeautoupdateconfigMap["start"].(string); ok {
		Start, _ := time.Parse("2006-01-02T15:04:05.999999", startString)
		o.Start = &Start
	}
	
	if endString, ok := EdgeautoupdateconfigMap["end"].(string); ok {
		End, _ := time.Parse("2006-01-02T15:04:05.999999", endString)
		o.End = &End
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeautoupdateconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
