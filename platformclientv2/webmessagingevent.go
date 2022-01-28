package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webmessagingevent - Message event element.  Examples include: system messages, typing indicators, cobrowse offerings.
type Webmessagingevent struct { 
	// EventType - Type of this event element
	EventType *string `json:"eventType,omitempty"`


	// CoBrowse - Cobrowse event.
	CoBrowse *Webmessagingeventcobrowse `json:"coBrowse,omitempty"`

}

func (o *Webmessagingevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webmessagingevent
	
	return json.Marshal(&struct { 
		EventType *string `json:"eventType,omitempty"`
		
		CoBrowse *Webmessagingeventcobrowse `json:"coBrowse,omitempty"`
		*Alias
	}{ 
		EventType: o.EventType,
		
		CoBrowse: o.CoBrowse,
		Alias:    (*Alias)(o),
	})
}

func (o *Webmessagingevent) UnmarshalJSON(b []byte) error {
	var WebmessagingeventMap map[string]interface{}
	err := json.Unmarshal(b, &WebmessagingeventMap)
	if err != nil {
		return err
	}
	
	if EventType, ok := WebmessagingeventMap["eventType"].(string); ok {
		o.EventType = &EventType
	}
	
	if CoBrowse, ok := WebmessagingeventMap["coBrowse"].(map[string]interface{}); ok {
		CoBrowseString, _ := json.Marshal(CoBrowse)
		json.Unmarshal(CoBrowseString, &o.CoBrowse)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webmessagingevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
