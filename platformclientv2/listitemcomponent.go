package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Listitemcomponent - An entry in a List template.
type Listitemcomponent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - An ID assigned to this list item.
	Id *string `json:"id,omitempty"`

	// Rmid - An ID of the rich message instance.
	Rmid *string `json:"rmid,omitempty"`

	// VarType - The type of list item to render.
	VarType *string `json:"type,omitempty"`

	// Image - URL of an image.
	Image *string `json:"image,omitempty"`

	// Title - The main headline of the list item.
	Title *string `json:"title,omitempty"`

	// Description - Text to show in the list item description.
	Description *string `json:"description,omitempty"`

	// Actions - The list item actions (Deprecated).
	Actions *Contentactions `json:"actions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Listitemcomponent) SetField(field string, fieldValue interface{}) {
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

func (o Listitemcomponent) MarshalJSON() ([]byte, error) {
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
	type Alias Listitemcomponent
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Rmid *string `json:"rmid,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Image *string `json:"image,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Actions *Contentactions `json:"actions,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Rmid: o.Rmid,
		
		VarType: o.VarType,
		
		Image: o.Image,
		
		Title: o.Title,
		
		Description: o.Description,
		
		Actions: o.Actions,
		Alias:    (Alias)(o),
	})
}

func (o *Listitemcomponent) UnmarshalJSON(b []byte) error {
	var ListitemcomponentMap map[string]interface{}
	err := json.Unmarshal(b, &ListitemcomponentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ListitemcomponentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Rmid, ok := ListitemcomponentMap["rmid"].(string); ok {
		o.Rmid = &Rmid
	}
    
	if VarType, ok := ListitemcomponentMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Image, ok := ListitemcomponentMap["image"].(string); ok {
		o.Image = &Image
	}
    
	if Title, ok := ListitemcomponentMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ListitemcomponentMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Actions, ok := ListitemcomponentMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Listitemcomponent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
