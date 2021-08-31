package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicaddress
type Queueconversationsocialexpressioneventtopicaddress struct { 
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

func (o *Queueconversationsocialexpressioneventtopicaddress) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicaddress
	
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

func (o *Queueconversationsocialexpressioneventtopicaddress) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicaddressMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicaddressMap)
	if err != nil {
		return err
	}
	
	if Name, ok := QueueconversationsocialexpressioneventtopicaddressMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if NameRaw, ok := QueueconversationsocialexpressioneventtopicaddressMap["nameRaw"].(string); ok {
		o.NameRaw = &NameRaw
	}
	
	if AddressNormalized, ok := QueueconversationsocialexpressioneventtopicaddressMap["addressNormalized"].(string); ok {
		o.AddressNormalized = &AddressNormalized
	}
	
	if AddressRaw, ok := QueueconversationsocialexpressioneventtopicaddressMap["addressRaw"].(string); ok {
		o.AddressRaw = &AddressRaw
	}
	
	if AddressDisplayable, ok := QueueconversationsocialexpressioneventtopicaddressMap["addressDisplayable"].(string); ok {
		o.AddressDisplayable = &AddressDisplayable
	}
	
	if AdditionalProperties, ok := QueueconversationsocialexpressioneventtopicaddressMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
