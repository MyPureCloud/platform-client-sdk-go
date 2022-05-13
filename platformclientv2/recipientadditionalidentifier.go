package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recipientadditionalidentifier - Additional identifiers for describing messaging recipient.
type Recipientadditionalidentifier struct { 
	// VarType - Type of the Identifier
	VarType *string `json:"type,omitempty"`


	// Value - The Identifier value.
	Value *string `json:"value,omitempty"`

}

func (o *Recipientadditionalidentifier) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recipientadditionalidentifier
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Value: o.Value,
		Alias:    (*Alias)(o),
	})
}

func (o *Recipientadditionalidentifier) UnmarshalJSON(b []byte) error {
	var RecipientadditionalidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &RecipientadditionalidentifierMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := RecipientadditionalidentifierMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := RecipientadditionalidentifierMap["value"].(string); ok {
		o.Value = &Value
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recipientadditionalidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
