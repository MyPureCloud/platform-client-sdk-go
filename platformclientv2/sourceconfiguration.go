package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sourceconfiguration
type Sourceconfiguration struct { 
	// SourceId - Identifies the external platform that is the source of the conversation.
	SourceId *string `json:"sourceId,omitempty"`


	// InteractionId - The customer's unique external identifier associated with the conversation that comes from the external platform.
	InteractionId *string `json:"interactionId,omitempty"`


	// TagId - The customer's external identifier or tag associated with the conversation. If set, it will be used to tag the conversation.
	TagId *string `json:"tagId,omitempty"`

}

func (o *Sourceconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sourceconfiguration
	
	return json.Marshal(&struct { 
		SourceId *string `json:"sourceId,omitempty"`
		
		InteractionId *string `json:"interactionId,omitempty"`
		
		TagId *string `json:"tagId,omitempty"`
		*Alias
	}{ 
		SourceId: o.SourceId,
		
		InteractionId: o.InteractionId,
		
		TagId: o.TagId,
		Alias:    (*Alias)(o),
	})
}

func (o *Sourceconfiguration) UnmarshalJSON(b []byte) error {
	var SourceconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &SourceconfigurationMap)
	if err != nil {
		return err
	}
	
	if SourceId, ok := SourceconfigurationMap["sourceId"].(string); ok {
		o.SourceId = &SourceId
	}
    
	if InteractionId, ok := SourceconfigurationMap["interactionId"].(string); ok {
		o.InteractionId = &InteractionId
	}
    
	if TagId, ok := SourceconfigurationMap["tagId"].(string); ok {
		o.TagId = &TagId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sourceconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
