package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createmanagementunitapirequest
type Createmanagementunitapirequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the management unit
	Name *string `json:"name,omitempty"`

	// TimeZone - The default time zone to use for this management unit.  Moving to Business Unit
	TimeZone *string `json:"timeZone,omitempty"`

	// StartDayOfWeek - The configured first day of the week for scheduling and forecasting purposes. Moving to Business Unit
	StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`

	// Settings - The configuration for the management unit.  If omitted, reasonable defaults will be assigned
	Settings *Createmanagementunitsettingsrequest `json:"settings,omitempty"`

	// DivisionId - The id of the division to which this management unit belongs.  Defaults to home division ID
	DivisionId *string `json:"divisionId,omitempty"`

	// BusinessUnitId - The id of the business unit to which this management unit belongs
	BusinessUnitId *string `json:"businessUnitId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createmanagementunitapirequest) SetField(field string, fieldValue interface{}) {
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

func (o Createmanagementunitapirequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createmanagementunitapirequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		StartDayOfWeek *string `json:"startDayOfWeek,omitempty"`
		
		Settings *Createmanagementunitsettingsrequest `json:"settings,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		TimeZone: o.TimeZone,
		
		StartDayOfWeek: o.StartDayOfWeek,
		
		Settings: o.Settings,
		
		DivisionId: o.DivisionId,
		
		BusinessUnitId: o.BusinessUnitId,
		Alias:    (Alias)(o),
	})
}

func (o *Createmanagementunitapirequest) UnmarshalJSON(b []byte) error {
	var CreatemanagementunitapirequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatemanagementunitapirequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreatemanagementunitapirequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if TimeZone, ok := CreatemanagementunitapirequestMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if StartDayOfWeek, ok := CreatemanagementunitapirequestMap["startDayOfWeek"].(string); ok {
		o.StartDayOfWeek = &StartDayOfWeek
	}
    
	if Settings, ok := CreatemanagementunitapirequestMap["settings"].(map[string]interface{}); ok {
		SettingsString, _ := json.Marshal(Settings)
		json.Unmarshal(SettingsString, &o.Settings)
	}
	
	if DivisionId, ok := CreatemanagementunitapirequestMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if BusinessUnitId, ok := CreatemanagementunitapirequestMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createmanagementunitapirequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
