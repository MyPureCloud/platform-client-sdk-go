package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactidentifier
type Contactidentifier struct { 
	// VarType - The type of this identifier
	VarType *string `json:"type,omitempty"`


	// Value - The string value of the identifier. Will vary in syntax by type.
	Value *string `json:"value,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

}

func (o *Contactidentifier) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactidentifier
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Value: o.Value,
		
		DateCreated: DateCreated,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactidentifier) UnmarshalJSON(b []byte) error {
	var ContactidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &ContactidentifierMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ContactidentifierMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := ContactidentifierMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if dateCreatedString, ok := ContactidentifierMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
