package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerresponsesetconfigchangereaction
type Dialerresponsesetconfigchangereaction struct { 
	// Data
	Data *string `json:"data,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ReactionType
	ReactionType *string `json:"reactionType,omitempty"`

}

func (o *Dialerresponsesetconfigchangereaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerresponsesetconfigchangereaction
	
	return json.Marshal(&struct { 
		Data *string `json:"data,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ReactionType *string `json:"reactionType,omitempty"`
		*Alias
	}{ 
		Data: o.Data,
		
		Name: o.Name,
		
		ReactionType: o.ReactionType,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerresponsesetconfigchangereaction) UnmarshalJSON(b []byte) error {
	var DialerresponsesetconfigchangereactionMap map[string]interface{}
	err := json.Unmarshal(b, &DialerresponsesetconfigchangereactionMap)
	if err != nil {
		return err
	}
	
	if Data, ok := DialerresponsesetconfigchangereactionMap["data"].(string); ok {
		o.Data = &Data
	}
	
	if Name, ok := DialerresponsesetconfigchangereactionMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ReactionType, ok := DialerresponsesetconfigchangereactionMap["reactionType"].(string); ok {
		o.ReactionType = &ReactionType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerresponsesetconfigchangereaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
