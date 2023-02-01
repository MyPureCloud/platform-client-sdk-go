package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestsession
type Knowledgeguestsession struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Session ID.
	Id *string `json:"id,omitempty"`

	// App - The app where the session is started.
	App *Knowledgeguestsessionapp `json:"app,omitempty"`

	// CustomerId - An arbitrary ID for the customer starting the session. Used to track multiple sessions started by the same customer.
	CustomerId *string `json:"customerId,omitempty"`

	// PageUrl - URL of the page where the session is started.
	PageUrl *string `json:"pageUrl,omitempty"`

	// Contexts - The session contexts.
	Contexts *[]Knowledgeguestsessioncontext `json:"contexts,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgeguestsession) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgeguestsession) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgeguestsession
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		App *Knowledgeguestsessionapp `json:"app,omitempty"`
		
		CustomerId *string `json:"customerId,omitempty"`
		
		PageUrl *string `json:"pageUrl,omitempty"`
		
		Contexts *[]Knowledgeguestsessioncontext `json:"contexts,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		App: o.App,
		
		CustomerId: o.CustomerId,
		
		PageUrl: o.PageUrl,
		
		Contexts: o.Contexts,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgeguestsession) UnmarshalJSON(b []byte) error {
	var KnowledgeguestsessionMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestsessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeguestsessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if App, ok := KnowledgeguestsessionMap["app"].(map[string]interface{}); ok {
		AppString, _ := json.Marshal(App)
		json.Unmarshal(AppString, &o.App)
	}
	
	if CustomerId, ok := KnowledgeguestsessionMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if PageUrl, ok := KnowledgeguestsessionMap["pageUrl"].(string); ok {
		o.PageUrl = &PageUrl
	}
    
	if Contexts, ok := KnowledgeguestsessionMap["contexts"].([]interface{}); ok {
		ContextsString, _ := json.Marshal(Contexts)
		json.Unmarshal(ContextsString, &o.Contexts)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestsession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
