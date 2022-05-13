package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchactiontarget
type Patchactiontarget struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ServiceLevel - Service Level of the action target. Chat offers for the target will be throttled with the aim of achieving this service level.
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`


	// ShortAbandonThreshold - Indicates the non-default short abandon threshold
	ShortAbandonThreshold *int `json:"shortAbandonThreshold,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Patchactiontarget) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchactiontarget
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`
		
		ShortAbandonThreshold *int `json:"shortAbandonThreshold,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ServiceLevel: o.ServiceLevel,
		
		ShortAbandonThreshold: o.ShortAbandonThreshold,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchactiontarget) UnmarshalJSON(b []byte) error {
	var PatchactiontargetMap map[string]interface{}
	err := json.Unmarshal(b, &PatchactiontargetMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PatchactiontargetMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := PatchactiontargetMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ServiceLevel, ok := PatchactiontargetMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if ShortAbandonThreshold, ok := PatchactiontargetMap["shortAbandonThreshold"].(float64); ok {
		ShortAbandonThresholdInt := int(ShortAbandonThreshold)
		o.ShortAbandonThreshold = &ShortAbandonThresholdInt
	}
	
	if SelfUri, ok := PatchactiontargetMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Patchactiontarget) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
