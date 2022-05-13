package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Presenceeventuserpresence
type Presenceeventuserpresence struct { 
	// Source
	Source *string `json:"source,omitempty"`


	// PresenceDefinition
	PresenceDefinition *Presenceeventorganizationpresence `json:"presenceDefinition,omitempty"`


	// Primary
	Primary *bool `json:"primary,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// ModifiedDate
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (o *Presenceeventuserpresence) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Presenceeventuserpresence
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Source *string `json:"source,omitempty"`
		
		PresenceDefinition *Presenceeventorganizationpresence `json:"presenceDefinition,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Source: o.Source,
		
		PresenceDefinition: o.PresenceDefinition,
		
		Primary: o.Primary,
		
		Message: o.Message,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Presenceeventuserpresence) UnmarshalJSON(b []byte) error {
	var PresenceeventuserpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &PresenceeventuserpresenceMap)
	if err != nil {
		return err
	}
	
	if Source, ok := PresenceeventuserpresenceMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if PresenceDefinition, ok := PresenceeventuserpresenceMap["presenceDefinition"].(map[string]interface{}); ok {
		PresenceDefinitionString, _ := json.Marshal(PresenceDefinition)
		json.Unmarshal(PresenceDefinitionString, &o.PresenceDefinition)
	}
	
	if Primary, ok := PresenceeventuserpresenceMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
    
	if Message, ok := PresenceeventuserpresenceMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if modifiedDateString, ok := PresenceeventuserpresenceMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Presenceeventuserpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
