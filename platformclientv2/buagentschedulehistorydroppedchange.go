package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentschedulehistorydroppedchange
type Buagentschedulehistorydroppedchange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Metadata - The metadata of the change, including who and when the change was made
	Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`

	// ShiftIds - The IDs of deleted shifts
	ShiftIds *[]string `json:"shiftIds,omitempty"`

	// FullDayTimeOffMarkerDates - The dates of any deleted full day time off markers
	FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`

	// Deletes - The deleted shifts, full day time off markers, or the entire agent schedule
	Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buagentschedulehistorydroppedchange) SetField(field string, fieldValue interface{}) {
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

func (o Buagentschedulehistorydroppedchange) MarshalJSON() ([]byte, error) {
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
	type Alias Buagentschedulehistorydroppedchange
	
	return json.Marshal(&struct { 
		Metadata *Buagentschedulehistorychangemetadata `json:"metadata,omitempty"`
		
		ShiftIds *[]string `json:"shiftIds,omitempty"`
		
		FullDayTimeOffMarkerDates *[]time.Time `json:"fullDayTimeOffMarkerDates,omitempty"`
		
		Deletes *Buagentschedulehistorydeletedchange `json:"deletes,omitempty"`
		Alias
	}{ 
		Metadata: o.Metadata,
		
		ShiftIds: o.ShiftIds,
		
		FullDayTimeOffMarkerDates: o.FullDayTimeOffMarkerDates,
		
		Deletes: o.Deletes,
		Alias:    (Alias)(o),
	})
}

func (o *Buagentschedulehistorydroppedchange) UnmarshalJSON(b []byte) error {
	var BuagentschedulehistorydroppedchangeMap map[string]interface{}
	err := json.Unmarshal(b, &BuagentschedulehistorydroppedchangeMap)
	if err != nil {
		return err
	}
	
	if Metadata, ok := BuagentschedulehistorydroppedchangeMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if ShiftIds, ok := BuagentschedulehistorydroppedchangeMap["shiftIds"].([]interface{}); ok {
		ShiftIdsString, _ := json.Marshal(ShiftIds)
		json.Unmarshal(ShiftIdsString, &o.ShiftIds)
	}
	
	if FullDayTimeOffMarkerDates, ok := BuagentschedulehistorydroppedchangeMap["fullDayTimeOffMarkerDates"].([]interface{}); ok {
		FullDayTimeOffMarkerDatesString, _ := json.Marshal(FullDayTimeOffMarkerDates)
		json.Unmarshal(FullDayTimeOffMarkerDatesString, &o.FullDayTimeOffMarkerDates)
	}
	
	if Deletes, ok := BuagentschedulehistorydroppedchangeMap["deletes"].(map[string]interface{}); ok {
		DeletesString, _ := json.Marshal(Deletes)
		json.Unmarshal(DeletesString, &o.Deletes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buagentschedulehistorydroppedchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
