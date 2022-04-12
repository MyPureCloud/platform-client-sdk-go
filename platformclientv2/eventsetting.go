package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventsetting
type Eventsetting struct { 
	// Typing - Settings regarding typing events
	Typing *Typingsetting `json:"typing,omitempty"`

}

func (o *Eventsetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventsetting
	
	return json.Marshal(&struct { 
		Typing *Typingsetting `json:"typing,omitempty"`
		*Alias
	}{ 
		Typing: o.Typing,
		Alias:    (*Alias)(o),
	})
}

func (o *Eventsetting) UnmarshalJSON(b []byte) error {
	var EventsettingMap map[string]interface{}
	err := json.Unmarshal(b, &EventsettingMap)
	if err != nil {
		return err
	}
	
	if Typing, ok := EventsettingMap["typing"].(map[string]interface{}); ok {
		TypingString, _ := json.Marshal(Typing)
		json.Unmarshal(TypingString, &o.Typing)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Eventsetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
