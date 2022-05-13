package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventcondition
type Eventcondition struct { 
	// Key - The event key.
	Key *string `json:"key,omitempty"`


	// Values - The event values.
	Values *[]string `json:"values,omitempty"`


	// Operator - The comparison operator.
	Operator *string `json:"operator,omitempty"`


	// StreamType - The stream type for which this condition can be satisfied.
	StreamType *string `json:"streamType,omitempty"`


	// SessionType - The session type for which this condition can be satisfied.
	SessionType *string `json:"sessionType,omitempty"`


	// EventName - The name of the event for which this condition can be satisfied.
	EventName *string `json:"eventName,omitempty"`

}

func (o *Eventcondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventcondition
	
	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		StreamType *string `json:"streamType,omitempty"`
		
		SessionType *string `json:"sessionType,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		*Alias
	}{ 
		Key: o.Key,
		
		Values: o.Values,
		
		Operator: o.Operator,
		
		StreamType: o.StreamType,
		
		SessionType: o.SessionType,
		
		EventName: o.EventName,
		Alias:    (*Alias)(o),
	})
}

func (o *Eventcondition) UnmarshalJSON(b []byte) error {
	var EventconditionMap map[string]interface{}
	err := json.Unmarshal(b, &EventconditionMap)
	if err != nil {
		return err
	}
	
	if Key, ok := EventconditionMap["key"].(string); ok {
		o.Key = &Key
	}
    
	if Values, ok := EventconditionMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if Operator, ok := EventconditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if StreamType, ok := EventconditionMap["streamType"].(string); ok {
		o.StreamType = &StreamType
	}
    
	if SessionType, ok := EventconditionMap["sessionType"].(string); ok {
		o.SessionType = &SessionType
	}
    
	if EventName, ok := EventconditionMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Eventcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
