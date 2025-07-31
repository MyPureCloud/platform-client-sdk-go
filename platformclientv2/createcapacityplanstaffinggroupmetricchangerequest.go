package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcapacityplanstaffinggroupmetricchangerequest
type Createcapacityplanstaffinggroupmetricchangerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// NumberOfWeeks - The number of weeks to which the metric change applies
	NumberOfWeeks *int `json:"numberOfWeeks,omitempty"`

	// WeekStartNumber - The start number of the week (starting from 1) to which the metric change applies, related to numberOfWeeks
	WeekStartNumber *int `json:"weekStartNumber,omitempty"`

	// Value - The value of the metric
	Value *float64 `json:"value,omitempty"`

	// Metric - The metric which is going to be modified for the selected staffing groups
	Metric *string `json:"metric,omitempty"`

	// Notes - Notes about the staffing groups metric changes
	Notes *string `json:"notes,omitempty"`

	// StaffingGroupIds - The IDs of the staffing groups affected by the metric change
	StaffingGroupIds *[]string `json:"staffingGroupIds,omitempty"`

	// Version - The version of the capacity plan
	Version *int `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createcapacityplanstaffinggroupmetricchangerequest) SetField(field string, fieldValue interface{}) {
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

func (o Createcapacityplanstaffinggroupmetricchangerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createcapacityplanstaffinggroupmetricchangerequest
	
	return json.Marshal(&struct { 
		NumberOfWeeks *int `json:"numberOfWeeks,omitempty"`
		
		WeekStartNumber *int `json:"weekStartNumber,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		StaffingGroupIds *[]string `json:"staffingGroupIds,omitempty"`
		
		Version *int `json:"version,omitempty"`
		Alias
	}{ 
		NumberOfWeeks: o.NumberOfWeeks,
		
		WeekStartNumber: o.WeekStartNumber,
		
		Value: o.Value,
		
		Metric: o.Metric,
		
		Notes: o.Notes,
		
		StaffingGroupIds: o.StaffingGroupIds,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Createcapacityplanstaffinggroupmetricchangerequest) UnmarshalJSON(b []byte) error {
	var CreatecapacityplanstaffinggroupmetricchangerequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatecapacityplanstaffinggroupmetricchangerequestMap)
	if err != nil {
		return err
	}
	
	if NumberOfWeeks, ok := CreatecapacityplanstaffinggroupmetricchangerequestMap["numberOfWeeks"].(float64); ok {
		NumberOfWeeksInt := int(NumberOfWeeks)
		o.NumberOfWeeks = &NumberOfWeeksInt
	}
	
	if WeekStartNumber, ok := CreatecapacityplanstaffinggroupmetricchangerequestMap["weekStartNumber"].(float64); ok {
		WeekStartNumberInt := int(WeekStartNumber)
		o.WeekStartNumber = &WeekStartNumberInt
	}
	
	if Value, ok := CreatecapacityplanstaffinggroupmetricchangerequestMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Metric, ok := CreatecapacityplanstaffinggroupmetricchangerequestMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Notes, ok := CreatecapacityplanstaffinggroupmetricchangerequestMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if StaffingGroupIds, ok := CreatecapacityplanstaffinggroupmetricchangerequestMap["staffingGroupIds"].([]interface{}); ok {
		StaffingGroupIdsString, _ := json.Marshal(StaffingGroupIds)
		json.Unmarshal(StaffingGroupIdsString, &o.StaffingGroupIds)
	}
	
	if Version, ok := CreatecapacityplanstaffinggroupmetricchangerequestMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createcapacityplanstaffinggroupmetricchangerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
