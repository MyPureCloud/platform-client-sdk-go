package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scrollpercentageeventtrigger - Details about a scroll percentage event trigger
type Scrollpercentageeventtrigger struct { 
	// Percentage - Percentage of a webpage at which an event is triggered.
	Percentage *int `json:"percentage,omitempty"`


	// EventName - Name of event triggered after scrolling to the specified percentage.
	EventName *string `json:"eventName,omitempty"`

}

func (o *Scrollpercentageeventtrigger) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scrollpercentageeventtrigger
	
	return json.Marshal(&struct { 
		Percentage *int `json:"percentage,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		*Alias
	}{ 
		Percentage: o.Percentage,
		
		EventName: o.EventName,
		Alias:    (*Alias)(o),
	})
}

func (o *Scrollpercentageeventtrigger) UnmarshalJSON(b []byte) error {
	var ScrollpercentageeventtriggerMap map[string]interface{}
	err := json.Unmarshal(b, &ScrollpercentageeventtriggerMap)
	if err != nil {
		return err
	}
	
	if Percentage, ok := ScrollpercentageeventtriggerMap["percentage"].(float64); ok {
		PercentageInt := int(Percentage)
		o.Percentage = &PercentageInt
	}
	
	if EventName, ok := ScrollpercentageeventtriggerMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scrollpercentageeventtrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
