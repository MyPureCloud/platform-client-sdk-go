package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Defaultobjective
type Defaultobjective struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// TemplateId - The id of this objective's base template
	TemplateId *string `json:"templateId,omitempty"`

	// Zones - Objective zone specifies min,max points and values for the associated metric
	Zones *[]Objectivezone `json:"zones,omitempty"`

	// Enabled - A flag for whether this objective is enabled for the related metric
	Enabled *bool `json:"enabled,omitempty"`

	// MediaTypes - A list of media types for the metric
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

	// Queues - A list of queues for the metric
	Queues *[]Addressableentityref `json:"queues,omitempty"`

	// Topics - A list of topic ids for detected topic metrics
	Topics *[]Addressableentityref `json:"topics,omitempty"`

	// TopicIdsFilterType - A filter type for topic Ids. It's only used for objectives with topicIds. Default filter behavior is \"or\".
	TopicIdsFilterType *string `json:"topicIdsFilterType,omitempty"`

	// EvaluationFormContextIds - The ids of associated evaluation form context, for Quality Evaluation Score metrics
	EvaluationFormContextIds *[]string `json:"evaluationFormContextIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Defaultobjective) SetField(field string, fieldValue interface{}) {
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

func (o Defaultobjective) MarshalJSON() ([]byte, error) {
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
	type Alias Defaultobjective
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		TemplateId *string `json:"templateId,omitempty"`
		
		Zones *[]Objectivezone `json:"zones,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		Queues *[]Addressableentityref `json:"queues,omitempty"`
		
		Topics *[]Addressableentityref `json:"topics,omitempty"`
		
		TopicIdsFilterType *string `json:"topicIdsFilterType,omitempty"`
		
		EvaluationFormContextIds *[]string `json:"evaluationFormContextIds,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		TemplateId: o.TemplateId,
		
		Zones: o.Zones,
		
		Enabled: o.Enabled,
		
		MediaTypes: o.MediaTypes,
		
		Queues: o.Queues,
		
		Topics: o.Topics,
		
		TopicIdsFilterType: o.TopicIdsFilterType,
		
		EvaluationFormContextIds: o.EvaluationFormContextIds,
		Alias:    (Alias)(o),
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
    
	if MediaTypes, ok := DefaultobjectiveMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if Queues, ok := DefaultobjectiveMap["queues"].([]interface{}); ok {
		QueuesString, _ := json.Marshal(Queues)
		json.Unmarshal(QueuesString, &o.Queues)
	}
	
	if Topics, ok := DefaultobjectiveMap["topics"].([]interface{}); ok {
		TopicsString, _ := json.Marshal(Topics)
		json.Unmarshal(TopicsString, &o.Topics)
	}
	
	if TopicIdsFilterType, ok := DefaultobjectiveMap["topicIdsFilterType"].(string); ok {
		o.TopicIdsFilterType = &TopicIdsFilterType
	}
    
	if EvaluationFormContextIds, ok := DefaultobjectiveMap["evaluationFormContextIds"].([]interface{}); ok {
		EvaluationFormContextIdsString, _ := json.Marshal(EvaluationFormContextIds)
		json.Unmarshal(EvaluationFormContextIdsString, &o.EvaluationFormContextIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Defaultobjective) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
