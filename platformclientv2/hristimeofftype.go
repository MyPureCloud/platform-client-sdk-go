package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Hristimeofftype
type Hristimeofftype struct { 
	// Id - The ID of the time off type configured in integration.
	Id *string `json:"id,omitempty"`


	// Name - The name of the time off type configured in integration.
	Name *string `json:"name,omitempty"`


	// HrisIntegrationId - The ID of the integration.
	HrisIntegrationId *string `json:"hrisIntegrationId,omitempty"`


	// SecondaryId - Secondary ID of the time off type, if configured in integration.
	SecondaryId *string `json:"secondaryId,omitempty"`

}

func (o *Hristimeofftype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Hristimeofftype
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		HrisIntegrationId *string `json:"hrisIntegrationId,omitempty"`
		
		SecondaryId *string `json:"secondaryId,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		HrisIntegrationId: o.HrisIntegrationId,
		
		SecondaryId: o.SecondaryId,
		Alias:    (*Alias)(o),
	})
}

func (o *Hristimeofftype) UnmarshalJSON(b []byte) error {
	var HristimeofftypeMap map[string]interface{}
	err := json.Unmarshal(b, &HristimeofftypeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := HristimeofftypeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := HristimeofftypeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if HrisIntegrationId, ok := HristimeofftypeMap["hrisIntegrationId"].(string); ok {
		o.HrisIntegrationId = &HrisIntegrationId
	}
    
	if SecondaryId, ok := HristimeofftypeMap["secondaryId"].(string); ok {
		o.SecondaryId = &SecondaryId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Hristimeofftype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
