package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicaddress
type Queueconversationeventtopicaddress struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// NameRaw
	NameRaw *string `json:"nameRaw,omitempty"`


	// AddressNormalized
	AddressNormalized *string `json:"addressNormalized,omitempty"`


	// AddressRaw
	AddressRaw *string `json:"addressRaw,omitempty"`


	// AddressDisplayable
	AddressDisplayable *string `json:"addressDisplayable,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationeventtopicaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicaddress
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		NameRaw *string `json:"nameRaw,omitempty"`
		
		AddressNormalized *string `json:"addressNormalized,omitempty"`
		
		AddressRaw *string `json:"addressRaw,omitempty"`
		
		AddressDisplayable *string `json:"addressDisplayable,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		NameRaw: o.NameRaw,
		
		AddressNormalized: o.AddressNormalized,
		
		AddressRaw: o.AddressRaw,
		
		AddressDisplayable: o.AddressDisplayable,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopicaddress) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicaddressMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicaddressMap)
	if err != nil {
		return err
	}
	
	if Name, ok := QueueconversationeventtopicaddressMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if NameRaw, ok := QueueconversationeventtopicaddressMap["nameRaw"].(string); ok {
		o.NameRaw = &NameRaw
	}
	
	if AddressNormalized, ok := QueueconversationeventtopicaddressMap["addressNormalized"].(string); ok {
		o.AddressNormalized = &AddressNormalized
	}
	
	if AddressRaw, ok := QueueconversationeventtopicaddressMap["addressRaw"].(string); ok {
		o.AddressRaw = &AddressRaw
	}
	
	if AddressDisplayable, ok := QueueconversationeventtopicaddressMap["addressDisplayable"].(string); ok {
		o.AddressDisplayable = &AddressDisplayable
	}
	
	if AdditionalProperties, ok := QueueconversationeventtopicaddressMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
