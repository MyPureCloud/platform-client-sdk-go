package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userschedule
type Userschedule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Shifts - The shifts that belong to this schedule
	Shifts *[]Userscheduleshift `json:"shifts,omitempty"`

	// FullDayTimeOffMarkers - Markers to indicate a full day time off request, relative to the management unit time zone
	FullDayTimeOffMarkers *[]Userschedulefulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`

	// Delete - If marked true for updating an existing user schedule, it will be deleted
	Delete *bool `json:"delete,omitempty"`

	// Metadata - Version metadata for this schedule
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

	// WorkPlanId - ID of the work plan associated with the user during schedule creation
	WorkPlanId *string `json:"workPlanId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userschedule) SetField(field string, fieldValue interface{}) {
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

func (o Userschedule) MarshalJSON() ([]byte, error) {
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
	type Alias Userschedule
	
	return json.Marshal(&struct { 
		Shifts *[]Userscheduleshift `json:"shifts,omitempty"`
		
		FullDayTimeOffMarkers *[]Userschedulefulldaytimeoffmarker `json:"fullDayTimeOffMarkers,omitempty"`
		
		Delete *bool `json:"delete,omitempty"`
		
		Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`
		
		WorkPlanId *string `json:"workPlanId,omitempty"`
		Alias
	}{ 
		Shifts: o.Shifts,
		
		FullDayTimeOffMarkers: o.FullDayTimeOffMarkers,
		
		Delete: o.Delete,
		
		Metadata: o.Metadata,
		
		WorkPlanId: o.WorkPlanId,
		Alias:    (Alias)(o),
	})
}

func (o *Userschedule) UnmarshalJSON(b []byte) error {
	var UserscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &UserscheduleMap)
	if err != nil {
		return err
	}
	
	if Shifts, ok := UserscheduleMap["shifts"].([]interface{}); ok {
		ShiftsString, _ := json.Marshal(Shifts)
		json.Unmarshal(ShiftsString, &o.Shifts)
	}
	
	if FullDayTimeOffMarkers, ok := UserscheduleMap["fullDayTimeOffMarkers"].([]interface{}); ok {
		FullDayTimeOffMarkersString, _ := json.Marshal(FullDayTimeOffMarkers)
		json.Unmarshal(FullDayTimeOffMarkersString, &o.FullDayTimeOffMarkers)
	}
	
	if Delete, ok := UserscheduleMap["delete"].(bool); ok {
		o.Delete = &Delete
	}
    
	if Metadata, ok := UserscheduleMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if WorkPlanId, ok := UserscheduleMap["workPlanId"].(string); ok {
		o.WorkPlanId = &WorkPlanId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
