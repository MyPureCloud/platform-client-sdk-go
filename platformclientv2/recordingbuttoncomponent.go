package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordingbuttoncomponent - Structured template button object.
type Recordingbuttoncomponent struct { 
	// Title
	Title *string `json:"title,omitempty"`


	// Actions
	Actions *Recordingcontentactions `json:"actions,omitempty"`


	// IsSelected
	IsSelected *bool `json:"isSelected,omitempty"`

}

func (o *Recordingbuttoncomponent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordingbuttoncomponent
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Actions *Recordingcontentactions `json:"actions,omitempty"`
		
		IsSelected *bool `json:"isSelected,omitempty"`
		*Alias
	}{ 
		Title: o.Title,
		
		Actions: o.Actions,
		
		IsSelected: o.IsSelected,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordingbuttoncomponent) UnmarshalJSON(b []byte) error {
	var RecordingbuttoncomponentMap map[string]interface{}
	err := json.Unmarshal(b, &RecordingbuttoncomponentMap)
	if err != nil {
		return err
	}
	
	if Title, ok := RecordingbuttoncomponentMap["title"].(string); ok {
		o.Title = &Title
	}
	
	if Actions, ok := RecordingbuttoncomponentMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if IsSelected, ok := RecordingbuttoncomponentMap["isSelected"].(bool); ok {
		o.IsSelected = &IsSelected
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Recordingbuttoncomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
