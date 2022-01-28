package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Defaultobjective
type Defaultobjective struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// TemplateId - The id of this objective's base template
	TemplateId *string `json:"templateId,omitempty"`


	// Zones - Objective zone specifies min,max points and values for the associated metric
	Zones *[]Objectivezone `json:"zones,omitempty"`


	// Enabled - A flag for whether this objective is enabled for the related metric
	Enabled *bool `json:"enabled,omitempty"`


	// Topics - A list of topic ids for detected topic metrics
	Topics *[]Addressableentityref `json:"topics,omitempty"`


	// TopicIdsFilterType - A filter type for topic Ids. It's only used for objectives with topicIds. Default filter behavior is \"or\".
	TopicIdsFilterType *string `json:"topicIdsFilterType,omitempty"`

}

func (o *Defaultobjective) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Defaultobjective
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		TemplateId *string `json:"templateId,omitempty"`
		
		Zones *[]Objectivezone `json:"zones,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Topics *[]Addressableentityref `json:"topics,omitempty"`
		
		TopicIdsFilterType *string `json:"topicIdsFilterType,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		TemplateId: o.TemplateId,
		
		Zones: o.Zones,
		
		Enabled: o.Enabled,
		
		Topics: o.Topics,
		
		TopicIdsFilterType: o.TopicIdsFilterType,
		Alias:    (*Alias)(o),
	})
}

func (o *Defaultobjective) UnmarshalJSON(b []byte) error {
	var DefaultobjectiveMap map[string]interface{}
	err := json.Unmarshal(b, &DefaultobjectiveMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DefaultobjectiveMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if TemplateId, ok := DefaultobjectiveMap["templateId"].(string); ok {
		o.TemplateId = &TemplateId
	}
	
	if Zones, ok := DefaultobjectiveMap["zones"].([]interface{}); ok {
		ZonesString, _ := json.Marshal(Zones)
		json.Unmarshal(ZonesString, &o.Zones)
	}
	
	if Enabled, ok := DefaultobjectiveMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if Topics, ok := DefaultobjectiveMap["topics"].([]interface{}); ok {
		TopicsString, _ := json.Marshal(Topics)
		json.Unmarshal(TopicsString, &o.Topics)
	}
	
	if TopicIdsFilterType, ok := DefaultobjectiveMap["topicIdsFilterType"].(string); ok {
		o.TopicIdsFilterType = &TopicIdsFilterType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Defaultobjective) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
