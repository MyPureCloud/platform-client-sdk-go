package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingdata
type Routingdata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueueId - The identifier of the routing queue
	QueueId *string `json:"queueId,omitempty"`

	// LanguageId - The identifier of a language to be considered in routing
	LanguageId *string `json:"languageId,omitempty"`

	// Label - An optional label that categorizes the conversation.  Max-utilization settings can be configured at a per-label level
	Label *string `json:"label,omitempty"`

	// Priority - The priority for routing
	Priority *int `json:"priority,omitempty"`

	// SkillIds - A list of skill identifiers to be considered in routing
	SkillIds *[]string `json:"skillIds,omitempty"`

	// PreferredAgentIds - A list of agents to be preferred in routing
	PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`

	// ScoredAgents - A list of scored agents for routing decisions. For Agent Owned Callbacks use one scored agent with a score of 100.
	ScoredAgents *[]Scoredagent `json:"scoredAgents,omitempty"`

	// RoutingFlags - An array of flags indicating how the conversation should be routed. Use \"AGENT_OWNED_CALLBACK\" when creating an Agent Owned Callback.
	RoutingFlags *[]string `json:"routingFlags,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Routingdata) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Routingdata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingdata
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		PreferredAgentIds *[]string `json:"preferredAgentIds,omitempty"`
		
		ScoredAgents *[]Scoredagent `json:"scoredAgents,omitempty"`
		
		RoutingFlags *[]string `json:"routingFlags,omitempty"`
		Alias
	}{ 
		QueueId: o.QueueId,
		
		LanguageId: o.LanguageId,
		
		Label: o.Label,
		
		Priority: o.Priority,
		
		SkillIds: o.SkillIds,
		
		PreferredAgentIds: o.PreferredAgentIds,
		
		ScoredAgents: o.ScoredAgents,
		
		RoutingFlags: o.RoutingFlags,
		Alias:    (Alias)(o),
	})
}

func (o *Routingdata) UnmarshalJSON(b []byte) error {
	var RoutingdataMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingdataMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := RoutingdataMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if LanguageId, ok := RoutingdataMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if Label, ok := RoutingdataMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if Priority, ok := RoutingdataMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if SkillIds, ok := RoutingdataMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if PreferredAgentIds, ok := RoutingdataMap["preferredAgentIds"].([]interface{}); ok {
		PreferredAgentIdsString, _ := json.Marshal(PreferredAgentIds)
		json.Unmarshal(PreferredAgentIdsString, &o.PreferredAgentIds)
	}
	
	if ScoredAgents, ok := RoutingdataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	
	if RoutingFlags, ok := RoutingdataMap["routingFlags"].([]interface{}); ok {
		RoutingFlagsString, _ := json.Marshal(RoutingFlags)
		json.Unmarshal(RoutingFlagsString, &o.RoutingFlags)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
