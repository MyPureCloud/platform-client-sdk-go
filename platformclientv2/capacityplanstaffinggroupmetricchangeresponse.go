package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanstaffinggroupmetricchangeresponse
type Capacityplanstaffinggroupmetricchangeresponse struct { 
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

	// StaffingGroups - The staffing groups affected by the metric change
	StaffingGroups *[]Staffinggroupreference `json:"staffingGroups,omitempty"`

	// CreatedBy - The user who created the metric change
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// CreatedDate - The date the entity was created, in ISO-8601 format
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// Version - The version of the capacity plan
	Version *int `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanstaffinggroupmetricchangeresponse) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanstaffinggroupmetricchangeresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Capacityplanstaffinggroupmetricchangeresponse
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		NumberOfWeeks *int `json:"numberOfWeeks,omitempty"`
		
		WeekStartNumber *int `json:"weekStartNumber,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		StaffingGroups *[]Staffinggroupreference `json:"staffingGroups,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		Version *int `json:"version,omitempty"`
		Alias
	}{ 
		NumberOfWeeks: o.NumberOfWeeks,
		
		WeekStartNumber: o.WeekStartNumber,
		
		Value: o.Value,
		
		Metric: o.Metric,
		
		Notes: o.Notes,
		
		StaffingGroups: o.StaffingGroups,
		
		CreatedBy: o.CreatedBy,
		
		CreatedDate: CreatedDate,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanstaffinggroupmetricchangeresponse) UnmarshalJSON(b []byte) error {
	var CapacityplanstaffinggroupmetricchangeresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanstaffinggroupmetricchangeresponseMap)
	if err != nil {
		return err
	}
	
	if NumberOfWeeks, ok := CapacityplanstaffinggroupmetricchangeresponseMap["numberOfWeeks"].(float64); ok {
		NumberOfWeeksInt := int(NumberOfWeeks)
		o.NumberOfWeeks = &NumberOfWeeksInt
	}
	
	if WeekStartNumber, ok := CapacityplanstaffinggroupmetricchangeresponseMap["weekStartNumber"].(float64); ok {
		WeekStartNumberInt := int(WeekStartNumber)
		o.WeekStartNumber = &WeekStartNumberInt
	}
	
	if Value, ok := CapacityplanstaffinggroupmetricchangeresponseMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Metric, ok := CapacityplanstaffinggroupmetricchangeresponseMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Notes, ok := CapacityplanstaffinggroupmetricchangeresponseMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if StaffingGroups, ok := CapacityplanstaffinggroupmetricchangeresponseMap["staffingGroups"].([]interface{}); ok {
		StaffingGroupsString, _ := json.Marshal(StaffingGroups)
		json.Unmarshal(StaffingGroupsString, &o.StaffingGroups)
	}
	
	if CreatedBy, ok := CapacityplanstaffinggroupmetricchangeresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if createdDateString, ok := CapacityplanstaffinggroupmetricchangeresponseMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if Version, ok := CapacityplanstaffinggroupmetricchangeresponseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanstaffinggroupmetricchangeresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
