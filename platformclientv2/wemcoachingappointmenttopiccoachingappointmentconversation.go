package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingappointmenttopiccoachingappointmentconversation
type Wemcoachingappointmenttopiccoachingappointmentconversation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

func (o *Wemcoachingappointmenttopiccoachingappointmentconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wemcoachingappointmenttopiccoachingappointmentconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Action: o.Action,
		Alias:    (*Alias)(o),
	})
}

func (o *Wemcoachingappointmenttopiccoachingappointmentconversation) UnmarshalJSON(b []byte) error {
	var WemcoachingappointmenttopiccoachingappointmentconversationMap map[string]interface{}
	err := json.Unmarshal(b, &WemcoachingappointmenttopiccoachingappointmentconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WemcoachingappointmenttopiccoachingappointmentconversationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Action, ok := WemcoachingappointmenttopiccoachingappointmentconversationMap["action"].(string); ok {
		o.Action = &Action
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
