package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentupdate
type Documentupdate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`

	// Name - The name of the document
	Name *string `json:"name,omitempty"`

	// Read
	Read *bool `json:"read,omitempty"`

	// AddTags
	AddTags *[]string `json:"addTags,omitempty"`

	// RemoveTags
	RemoveTags *[]string `json:"removeTags,omitempty"`

	// AddTagIds
	AddTagIds *[]string `json:"addTagIds,omitempty"`

	// RemoveTagIds
	RemoveTagIds *[]string `json:"removeTagIds,omitempty"`

	// UpdateAttributes
	UpdateAttributes *[]Documentattribute `json:"updateAttributes,omitempty"`

	// RemoveAttributes
	RemoveAttributes *[]string `json:"removeAttributes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentupdate) SetField(field string, fieldValue interface{}) {
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

func (o Documentupdate) MarshalJSON() ([]byte, error) {
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
	type Alias Documentupdate
	
	return json.Marshal(&struct { 
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		AddTags *[]string `json:"addTags,omitempty"`
		
		RemoveTags *[]string `json:"removeTags,omitempty"`
		
		AddTagIds *[]string `json:"addTagIds,omitempty"`
		
		RemoveTagIds *[]string `json:"removeTagIds,omitempty"`
		
		UpdateAttributes *[]Documentattribute `json:"updateAttributes,omitempty"`
		
		RemoveAttributes *[]string `json:"removeAttributes,omitempty"`
		Alias
	}{ 
		ChangeNumber: o.ChangeNumber,
		
		Name: o.Name,
		
		Read: o.Read,
		
		AddTags: o.AddTags,
		
		RemoveTags: o.RemoveTags,
		
		AddTagIds: o.AddTagIds,
		
		RemoveTagIds: o.RemoveTagIds,
		
		UpdateAttributes: o.UpdateAttributes,
		
		RemoveAttributes: o.RemoveAttributes,
		Alias:    (Alias)(o),
	})
}

func (o *Documentupdate) UnmarshalJSON(b []byte) error {
	var DocumentupdateMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentupdateMap)
	if err != nil {
		return err
	}
	
	if ChangeNumber, ok := DocumentupdateMap["changeNumber"].(float64); ok {
		ChangeNumberInt := int(ChangeNumber)
		o.ChangeNumber = &ChangeNumberInt
	}
	
	if Name, ok := DocumentupdateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Read, ok := DocumentupdateMap["read"].(bool); ok {
		o.Read = &Read
	}
    
	if AddTags, ok := DocumentupdateMap["addTags"].([]interface{}); ok {
		AddTagsString, _ := json.Marshal(AddTags)
		json.Unmarshal(AddTagsString, &o.AddTags)
	}
	
	if RemoveTags, ok := DocumentupdateMap["removeTags"].([]interface{}); ok {
		RemoveTagsString, _ := json.Marshal(RemoveTags)
		json.Unmarshal(RemoveTagsString, &o.RemoveTags)
	}
	
	if AddTagIds, ok := DocumentupdateMap["addTagIds"].([]interface{}); ok {
		AddTagIdsString, _ := json.Marshal(AddTagIds)
		json.Unmarshal(AddTagIdsString, &o.AddTagIds)
	}
	
	if RemoveTagIds, ok := DocumentupdateMap["removeTagIds"].([]interface{}); ok {
		RemoveTagIdsString, _ := json.Marshal(RemoveTagIds)
		json.Unmarshal(RemoveTagIdsString, &o.RemoveTagIds)
	}
	
	if UpdateAttributes, ok := DocumentupdateMap["updateAttributes"].([]interface{}); ok {
		UpdateAttributesString, _ := json.Marshal(UpdateAttributes)
		json.Unmarshal(UpdateAttributesString, &o.UpdateAttributes)
	}
	
	if RemoveAttributes, ok := DocumentupdateMap["removeAttributes"].([]interface{}); ok {
		RemoveAttributesString, _ := json.Marshal(RemoveAttributes)
		json.Unmarshal(RemoveAttributesString, &o.RemoveAttributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
