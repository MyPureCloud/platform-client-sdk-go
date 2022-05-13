package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactablestatus
type Contactablestatus struct { 
	// Contactable - Indicates whether or not the entire contact is contactable for the associated media type.
	Contactable *bool `json:"contactable,omitempty"`


	// ColumnStatus - A map of individual contact method columns to whether the individual column is contactable for the associated media type.
	ColumnStatus *map[string]Columnstatus `json:"columnStatus,omitempty"`

}

func (o *Contactablestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactablestatus
	
	return json.Marshal(&struct { 
		Contactable *bool `json:"contactable,omitempty"`
		
		ColumnStatus *map[string]Columnstatus `json:"columnStatus,omitempty"`
		*Alias
	}{ 
		Contactable: o.Contactable,
		
		ColumnStatus: o.ColumnStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactablestatus) UnmarshalJSON(b []byte) error {
	var ContactablestatusMap map[string]interface{}
	err := json.Unmarshal(b, &ContactablestatusMap)
	if err != nil {
		return err
	}
	
	if Contactable, ok := ContactablestatusMap["contactable"].(bool); ok {
		o.Contactable = &Contactable
	}
    
	if ColumnStatus, ok := ContactablestatusMap["columnStatus"].(map[string]interface{}); ok {
		ColumnStatusString, _ := json.Marshal(ColumnStatus)
		json.Unmarshal(ColumnStatusString, &o.ColumnStatus)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactablestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
