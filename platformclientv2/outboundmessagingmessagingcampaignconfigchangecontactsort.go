package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundmessagingmessagingcampaignconfigchangecontactsort
type Outboundmessagingmessagingcampaignconfigchangecontactsort struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Numeric - Whether that column contains numeric data
	Numeric *bool `json:"numeric,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Outboundmessagingmessagingcampaignconfigchangecontactsort) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundmessagingmessagingcampaignconfigchangecontactsort
	
	return json.Marshal(&struct { 
		FieldName *string `json:"fieldName,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Numeric *bool `json:"numeric,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		FieldName: o.FieldName,
		
		Direction: o.Direction,
		
		Numeric: o.Numeric,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundmessagingmessagingcampaignconfigchangecontactsort) UnmarshalJSON(b []byte) error {
	var OutboundmessagingmessagingcampaignconfigchangecontactsortMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundmessagingmessagingcampaignconfigchangecontactsortMap)
	if err != nil {
		return err
	}
	
	if FieldName, ok := OutboundmessagingmessagingcampaignconfigchangecontactsortMap["fieldName"].(string); ok {
		o.FieldName = &FieldName
	}
    
	if Direction, ok := OutboundmessagingmessagingcampaignconfigchangecontactsortMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Numeric, ok := OutboundmessagingmessagingcampaignconfigchangecontactsortMap["numeric"].(bool); ok {
		o.Numeric = &Numeric
	}
    
	if AdditionalProperties, ok := OutboundmessagingmessagingcampaignconfigchangecontactsortMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangecontactsort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
