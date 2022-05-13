package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercampaignconfigchangecontactsort - information determining the order in which the contacts will be dialed
type Dialercampaignconfigchangecontactsort struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Numeric - Whether that column contains numeric data
	Numeric *bool `json:"numeric,omitempty"`

}

func (o *Dialercampaignconfigchangecontactsort) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercampaignconfigchangecontactsort
	
	return json.Marshal(&struct { 
		FieldName *string `json:"fieldName,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Numeric *bool `json:"numeric,omitempty"`
		*Alias
	}{ 
		FieldName: o.FieldName,
		
		Direction: o.Direction,
		
		Numeric: o.Numeric,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercampaignconfigchangecontactsort) UnmarshalJSON(b []byte) error {
	var DialercampaignconfigchangecontactsortMap map[string]interface{}
	err := json.Unmarshal(b, &DialercampaignconfigchangecontactsortMap)
	if err != nil {
		return err
	}
	
	if FieldName, ok := DialercampaignconfigchangecontactsortMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
    
	if Direction, ok := DialercampaignconfigchangecontactsortMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Numeric, ok := DialercampaignconfigchangecontactsortMap["numeric"].(bool); ok {
		o.Numeric = &Numeric
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercampaignconfigchangecontactsort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
