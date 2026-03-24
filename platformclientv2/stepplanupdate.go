package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Stepplanupdate
type Stepplanupdate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the Stepplan. Valid length between 3 and 256 characters.
	Name *string `json:"name,omitempty"`

	// Description - The description of the Stepplan. Maximum length of 512 characters.
	Description *string `json:"description,omitempty"`

	// WorkitemSettings - The workitemSettings of the Stepplan.
	WorkitemSettings *Workitemsettings `json:"workitemSettings,omitempty"`

	// ActivityType - The activityType of the Stepplan.
	ActivityType *string `json:"activityType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Stepplanupdate) SetField(field string, fieldValue interface{}) {
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

func (o Stepplanupdate) MarshalJSON() ([]byte, error) {
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
	type Alias Stepplanupdate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		WorkitemSettings *Workitemsettings `json:"workitemSettings,omitempty"`
		
		ActivityType *string `json:"activityType,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		WorkitemSettings: o.WorkitemSettings,
		
		ActivityType: o.ActivityType,
		Alias:    (Alias)(o),
	})
}

func (o *Stepplanupdate) UnmarshalJSON(b []byte) error {
	var StepplanupdateMap map[string]interface{}
	err := json.Unmarshal(b, &StepplanupdateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := StepplanupdateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := StepplanupdateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if WorkitemSettings, ok := StepplanupdateMap["workitemSettings"].(map[string]interface{}); ok {
		WorkitemSettingsString, _ := json.Marshal(WorkitemSettings)
		json.Unmarshal(WorkitemSettingsString, &o.WorkitemSettings)
	}
	
	if ActivityType, ok := StepplanupdateMap["activityType"].(string); ok {
		o.ActivityType = &ActivityType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Stepplanupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
