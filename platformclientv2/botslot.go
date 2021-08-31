package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botslot - Description of a data value returned from an intent
type Botslot struct { 
	// Name - The name of the slot. This can be up to 100 characters long and must be comprised of displayable characters without leading or trailing whitespace
	Name *string `json:"name,omitempty"`


	// VarType - The data type of the slot string, integer, decimal, duration, boolean, currency, datetime or the xxxCollection versions of those types
	VarType *string `json:"type,omitempty"`

}

func (o *Botslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botslot
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Botslot) UnmarshalJSON(b []byte) error {
	var BotslotMap map[string]interface{}
	err := json.Unmarshal(b, &BotslotMap)
	if err != nil {
		return err
	}
	
	if Name, ok := BotslotMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VarType, ok := BotslotMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Botslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
