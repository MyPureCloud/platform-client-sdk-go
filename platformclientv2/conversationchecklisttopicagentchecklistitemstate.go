package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchecklisttopicagentchecklistitemstate
type Conversationchecklisttopicagentchecklistitemstate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// StateFromModel
	StateFromModel *string `json:"stateFromModel,omitempty"`

	// StateFromAgent
	StateFromAgent *string `json:"stateFromAgent,omitempty"`

	// DateLastModifiedByModel
	DateLastModifiedByModel *time.Time `json:"dateLastModifiedByModel,omitempty"`

	// DateLastModifiedByAgent
	DateLastModifiedByAgent *time.Time `json:"dateLastModifiedByAgent,omitempty"`

	// AutomatedCheckEnabled
	AutomatedCheckEnabled *bool `json:"automatedCheckEnabled,omitempty"`

	// Important
	Important *bool `json:"important,omitempty"`

	// LastModifiedInAcw
	LastModifiedInAcw *bool `json:"lastModifiedInAcw,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationchecklisttopicagentchecklistitemstate) SetField(field string, fieldValue interface{}) {
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

func (o Conversationchecklisttopicagentchecklistitemstate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLastModifiedByModel","DateLastModifiedByAgent", }
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
	type Alias Conversationchecklisttopicagentchecklistitemstate
	
	DateLastModifiedByModel := new(string)
	if o.DateLastModifiedByModel != nil {
		
		*DateLastModifiedByModel = timeutil.Strftime(o.DateLastModifiedByModel, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastModifiedByModel = nil
	}
	
	DateLastModifiedByAgent := new(string)
	if o.DateLastModifiedByAgent != nil {
		
		*DateLastModifiedByAgent = timeutil.Strftime(o.DateLastModifiedByAgent, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLastModifiedByAgent = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		StateFromModel *string `json:"stateFromModel,omitempty"`
		
		StateFromAgent *string `json:"stateFromAgent,omitempty"`
		
		DateLastModifiedByModel *string `json:"dateLastModifiedByModel,omitempty"`
		
		DateLastModifiedByAgent *string `json:"dateLastModifiedByAgent,omitempty"`
		
		AutomatedCheckEnabled *bool `json:"automatedCheckEnabled,omitempty"`
		
		Important *bool `json:"important,omitempty"`
		
		LastModifiedInAcw *bool `json:"lastModifiedInAcw,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		StateFromModel: o.StateFromModel,
		
		StateFromAgent: o.StateFromAgent,
		
		DateLastModifiedByModel: DateLastModifiedByModel,
		
		DateLastModifiedByAgent: DateLastModifiedByAgent,
		
		AutomatedCheckEnabled: o.AutomatedCheckEnabled,
		
		Important: o.Important,
		
		LastModifiedInAcw: o.LastModifiedInAcw,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationchecklisttopicagentchecklistitemstate) UnmarshalJSON(b []byte) error {
	var ConversationchecklisttopicagentchecklistitemstateMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationchecklisttopicagentchecklistitemstateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationchecklisttopicagentchecklistitemstateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationchecklisttopicagentchecklistitemstateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ConversationchecklisttopicagentchecklistitemstateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if StateFromModel, ok := ConversationchecklisttopicagentchecklistitemstateMap["stateFromModel"].(string); ok {
		o.StateFromModel = &StateFromModel
	}
    
	if StateFromAgent, ok := ConversationchecklisttopicagentchecklistitemstateMap["stateFromAgent"].(string); ok {
		o.StateFromAgent = &StateFromAgent
	}
    
	if dateLastModifiedByModelString, ok := ConversationchecklisttopicagentchecklistitemstateMap["dateLastModifiedByModel"].(string); ok {
		DateLastModifiedByModel, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastModifiedByModelString)
		o.DateLastModifiedByModel = &DateLastModifiedByModel
	}
	
	if dateLastModifiedByAgentString, ok := ConversationchecklisttopicagentchecklistitemstateMap["dateLastModifiedByAgent"].(string); ok {
		DateLastModifiedByAgent, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLastModifiedByAgentString)
		o.DateLastModifiedByAgent = &DateLastModifiedByAgent
	}
	
	if AutomatedCheckEnabled, ok := ConversationchecklisttopicagentchecklistitemstateMap["automatedCheckEnabled"].(bool); ok {
		o.AutomatedCheckEnabled = &AutomatedCheckEnabled
	}
    
	if Important, ok := ConversationchecklisttopicagentchecklistitemstateMap["important"].(bool); ok {
		o.Important = &Important
	}
    
	if LastModifiedInAcw, ok := ConversationchecklisttopicagentchecklistitemstateMap["lastModifiedInAcw"].(bool); ok {
		o.LastModifiedInAcw = &LastModifiedInAcw
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationchecklisttopicagentchecklistitemstate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
