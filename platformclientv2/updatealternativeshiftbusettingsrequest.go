package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatealternativeshiftbusettingsrequest
type Updatealternativeshiftbusettingsrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EnabledGranularities - The granularity at which alternative shifts is allowed. An empty list as the wrapped value will indicate alternative shifts is disabled
	EnabledGranularities *Listwrapperalternativeshiftbusettingsgranularity `json:"enabledGranularities,omitempty"`

	// MinMinutesBeforeStartTime - The minimum number of minutes before the start of a shift that an alternative shift can be automatically approved
	MinMinutesBeforeStartTime *int `json:"minMinutesBeforeStartTime,omitempty"`

	// RetainedActivityCategories - Categories of activities that are required to remain at the same time slot for the alternative shifts offered. An empty list indicates no retained activities
	RetainedActivityCategories *Listwrapperalternativeshiftbusettingsactivitycategory `json:"retainedActivityCategories,omitempty"`

	// Metadata - Version metadata for this business unit's alternative shift settings
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updatealternativeshiftbusettingsrequest) SetField(field string, fieldValue interface{}) {
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

func (o Updatealternativeshiftbusettingsrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Updatealternativeshiftbusettingsrequest
	
	return json.Marshal(&struct { 
		EnabledGranularities *Listwrapperalternativeshiftbusettingsgranularity `json:"enabledGranularities,omitempty"`
		
		MinMinutesBeforeStartTime *int `json:"minMinutesBeforeStartTime,omitempty"`
		
		RetainedActivityCategories *Listwrapperalternativeshiftbusettingsactivitycategory `json:"retainedActivityCategories,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		Alias
	}{ 
		EnabledGranularities: o.EnabledGranularities,
		
		MinMinutesBeforeStartTime: o.MinMinutesBeforeStartTime,
		
		RetainedActivityCategories: o.RetainedActivityCategories,
		
		Metadata: o.Metadata,
		Alias:    (Alias)(o),
	})
}

func (o *Updatealternativeshiftbusettingsrequest) UnmarshalJSON(b []byte) error {
	var UpdatealternativeshiftbusettingsrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatealternativeshiftbusettingsrequestMap)
	if err != nil {
		return err
	}
	
	if EnabledGranularities, ok := UpdatealternativeshiftbusettingsrequestMap["enabledGranularities"].(map[string]interface{}); ok {
		EnabledGranularitiesString, _ := json.Marshal(EnabledGranularities)
		json.Unmarshal(EnabledGranularitiesString, &o.EnabledGranularities)
	}
	
	if MinMinutesBeforeStartTime, ok := UpdatealternativeshiftbusettingsrequestMap["minMinutesBeforeStartTime"].(float64); ok {
		MinMinutesBeforeStartTimeInt := int(MinMinutesBeforeStartTime)
		o.MinMinutesBeforeStartTime = &MinMinutesBeforeStartTimeInt
	}
	
	if RetainedActivityCategories, ok := UpdatealternativeshiftbusettingsrequestMap["retainedActivityCategories"].(map[string]interface{}); ok {
		RetainedActivityCategoriesString, _ := json.Marshal(RetainedActivityCategories)
		json.Unmarshal(RetainedActivityCategoriesString, &o.RetainedActivityCategories)
	}
	
	if Metadata, ok := UpdatealternativeshiftbusettingsrequestMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatealternativeshiftbusettingsrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
