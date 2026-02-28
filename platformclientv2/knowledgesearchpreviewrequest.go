package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesearchpreviewrequest
type Knowledgesearchpreviewrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Query - Query to search content in the knowledge sources.
	Query *string `json:"query,omitempty"`

	// Sources - Source information to search upon.
	Sources *[]V3sourceref `json:"sources,omitempty"`

	// GenerationSetting - Setting for answer generation.
	GenerationSetting *Knowledgegenerationsetting `json:"generationSetting,omitempty"`

	// Stateful - Indicates if stateful search and generation is enabled for the knowledge setting.
	Stateful *bool `json:"stateful,omitempty"`

	// ConversationTurns - List of conversation turns to use for stateful search.
	ConversationTurns *[]Knowledgeconversationturn `json:"conversationTurns,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgesearchpreviewrequest) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgesearchpreviewrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgesearchpreviewrequest
	
	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		Sources *[]V3sourceref `json:"sources,omitempty"`
		
		GenerationSetting *Knowledgegenerationsetting `json:"generationSetting,omitempty"`
		
		Stateful *bool `json:"stateful,omitempty"`
		
		ConversationTurns *[]Knowledgeconversationturn `json:"conversationTurns,omitempty"`
		Alias
	}{ 
		Query: o.Query,
		
		Sources: o.Sources,
		
		GenerationSetting: o.GenerationSetting,
		
		Stateful: o.Stateful,
		
		ConversationTurns: o.ConversationTurns,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgesearchpreviewrequest) UnmarshalJSON(b []byte) error {
	var KnowledgesearchpreviewrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesearchpreviewrequestMap)
	if err != nil {
		return err
	}
	
	if Query, ok := KnowledgesearchpreviewrequestMap["query"].(string); ok {
		o.Query = &Query
	}
    
	if Sources, ok := KnowledgesearchpreviewrequestMap["sources"].([]interface{}); ok {
		SourcesString, _ := json.Marshal(Sources)
		json.Unmarshal(SourcesString, &o.Sources)
	}
	
	if GenerationSetting, ok := KnowledgesearchpreviewrequestMap["generationSetting"].(map[string]interface{}); ok {
		GenerationSettingString, _ := json.Marshal(GenerationSetting)
		json.Unmarshal(GenerationSettingString, &o.GenerationSetting)
	}
	
	if Stateful, ok := KnowledgesearchpreviewrequestMap["stateful"].(bool); ok {
		o.Stateful = &Stateful
	}
    
	if ConversationTurns, ok := KnowledgesearchpreviewrequestMap["conversationTurns"].([]interface{}); ok {
		ConversationTurnsString, _ := json.Marshal(ConversationTurns)
		json.Unmarshal(ConversationTurnsString, &o.ConversationTurns)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesearchpreviewrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
