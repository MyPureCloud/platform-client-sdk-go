package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentlist - List content object.
type Contentlist struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A unique ID assigned to this rich message content.
	Id *string `json:"id,omitempty"`

	// ListType - The type of list this instance represents.
	ListType *string `json:"listType,omitempty"`

	// Title - Text to show in the title.
	Title *string `json:"title,omitempty"`

	// Description - Text to show in the description.
	Description *string `json:"description,omitempty"`

	// SubmitLabel - Label for Submit button.
	SubmitLabel *string `json:"submitLabel,omitempty"`

	// Actions - The list actions (Deprecated).
	Actions *Contentactions `json:"actions,omitempty"`

	// Components - An array of component objects.
	Components *[]Listitemcomponent `json:"components,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contentlist) SetField(field string, fieldValue interface{}) {
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

func (o Contentlist) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Contentlist
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ListType *string `json:"listType,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SubmitLabel *string `json:"submitLabel,omitempty"`
		
		Actions *Contentactions `json:"actions,omitempty"`
		
		Components *[]Listitemcomponent `json:"components,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ListType: o.ListType,
		
		Title: o.Title,
		
		Description: o.Description,
		
		SubmitLabel: o.SubmitLabel,
		
		Actions: o.Actions,
		
		Components: o.Components,
		Alias:    (Alias)(o),
	})
}

func (o *Contentlist) UnmarshalJSON(b []byte) error {
	var ContentlistMap map[string]interface{}
	err := json.Unmarshal(b, &ContentlistMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentlistMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ListType, ok := ContentlistMap["listType"].(string); ok {
		o.ListType = &ListType
	}
    
	if Title, ok := ContentlistMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ContentlistMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SubmitLabel, ok := ContentlistMap["submitLabel"].(string); ok {
		o.SubmitLabel = &SubmitLabel
	}
    
	if Actions, ok := ContentlistMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Components, ok := ContentlistMap["components"].([]interface{}); ok {
		ComponentsString, _ := json.Marshal(Components)
		json.Unmarshal(ComponentsString, &o.Components)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
