package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainresourceconditionvalue
type Domainresourceconditionvalue struct { 
	// User
	User *User `json:"user,omitempty"`


	// Queue
	Queue *Queue `json:"queue,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Domainresourceconditionvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainresourceconditionvalue
	
	return json.Marshal(&struct { 
		User *User `json:"user,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		Queue: o.Queue,
		
		Value: o.Value,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainresourceconditionvalue) UnmarshalJSON(b []byte) error {
	var DomainresourceconditionvalueMap map[string]interface{}
	err := json.Unmarshal(b, &DomainresourceconditionvalueMap)
	if err != nil {
		return err
	}
	
	if User, ok := DomainresourceconditionvalueMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := DomainresourceconditionvalueMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Value, ok := DomainresourceconditionvalueMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if VarType, ok := DomainresourceconditionvalueMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainresourceconditionvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
