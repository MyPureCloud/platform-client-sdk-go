package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Searchunmatchedshifttradelistjobrequest
type Searchunmatchedshifttradelistjobrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ManagementUnitIds - The IDs of management units from which to query shift trades
	ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`

	// WeekDates - The start week dates in which to query shift trades in the business unit time zone (yyyy-MM-dd format)
	WeekDates *[]time.Time `json:"weekDates,omitempty"`

	// ReceivingSchedule - Associated schedule information for the receiving user
	ReceivingSchedule *Receivingschedulelookup `json:"receivingSchedule,omitempty"`

	// ReceivingShiftIds - The IDs of shifts that the receiving user would potentially be willing to trade. If empty, only returns one-sided trades
	ReceivingShiftIds *[]string `json:"receivingShiftIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Searchunmatchedshifttradelistjobrequest) SetField(field string, fieldValue interface{}) {
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

func (o Searchunmatchedshifttradelistjobrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Searchunmatchedshifttradelistjobrequest
	
	return json.Marshal(&struct { 
		ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`
		
		WeekDates *[]time.Time `json:"weekDates,omitempty"`
		
		ReceivingSchedule *Receivingschedulelookup `json:"receivingSchedule,omitempty"`
		
		ReceivingShiftIds *[]string `json:"receivingShiftIds,omitempty"`
		Alias
	}{ 
		ManagementUnitIds: o.ManagementUnitIds,
		
		WeekDates: o.WeekDates,
		
		ReceivingSchedule: o.ReceivingSchedule,
		
		ReceivingShiftIds: o.ReceivingShiftIds,
		Alias:    (Alias)(o),
	})
}

func (o *Searchunmatchedshifttradelistjobrequest) UnmarshalJSON(b []byte) error {
	var SearchunmatchedshifttradelistjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SearchunmatchedshifttradelistjobrequestMap)
	if err != nil {
		return err
	}
	
	if ManagementUnitIds, ok := SearchunmatchedshifttradelistjobrequestMap["managementUnitIds"].([]interface{}); ok {
		ManagementUnitIdsString, _ := json.Marshal(ManagementUnitIds)
		json.Unmarshal(ManagementUnitIdsString, &o.ManagementUnitIds)
	}
	
	if WeekDates, ok := SearchunmatchedshifttradelistjobrequestMap["weekDates"].([]interface{}); ok {
		WeekDatesString, _ := json.Marshal(WeekDates)
		json.Unmarshal(WeekDatesString, &o.WeekDates)
	}
	
	if ReceivingSchedule, ok := SearchunmatchedshifttradelistjobrequestMap["receivingSchedule"].(map[string]interface{}); ok {
		ReceivingScheduleString, _ := json.Marshal(ReceivingSchedule)
		json.Unmarshal(ReceivingScheduleString, &o.ReceivingSchedule)
	}
	
	if ReceivingShiftIds, ok := SearchunmatchedshifttradelistjobrequestMap["receivingShiftIds"].([]interface{}); ok {
		ReceivingShiftIdsString, _ := json.Marshal(ReceivingShiftIds)
		json.Unmarshal(ReceivingShiftIdsString, &o.ReceivingShiftIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Searchunmatchedshifttradelistjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
