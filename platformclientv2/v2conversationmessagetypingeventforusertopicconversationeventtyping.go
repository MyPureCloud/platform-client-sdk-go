package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationeventtyping
type V2conversationmessagetypingeventforusertopicconversationeventtyping struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Duration
	Duration *int `json:"duration,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationeventtyping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationeventtyping
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Duration *int `json:"duration,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Duration: o.Duration,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforusertopicconversationeventtyping) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationeventtypingMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationeventtypingMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := V2conversationmessagetypingeventforusertopicconversationeventtypingMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Duration, ok := V2conversationmessagetypingeventforusertopicconversationeventtypingMap["duration"].(float64); ok {
		DurationInt := int(Duration)
		o.Duration = &DurationInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationeventtyping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
