package platformclientv2
import (
	"encoding/json"
)

// Outboundmessagingmessagingcampaignconfigchangecontactsort
type Outboundmessagingmessagingcampaignconfigchangecontactsort struct { 
	// FieldName
	FieldName *string `json:"fieldName,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Numeric
	Numeric *bool `json:"numeric,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Outboundmessagingmessagingcampaignconfigchangecontactsort) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
