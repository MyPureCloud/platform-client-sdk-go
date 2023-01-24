package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Linestatus
type Linestatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The id of this line
	Id *string `json:"id,omitempty"`

	// Reachable - Indicates whether the edge can reach the line.
	Reachable *bool `json:"reachable,omitempty"`

	// AddressOfRecord - The line's address of record.
	AddressOfRecord *string `json:"addressOfRecord,omitempty"`

	// ContactAddresses - The addresses used to contact the line.
	ContactAddresses *[]string `json:"contactAddresses,omitempty"`

	// ReachableStateTime - The time the line entered its current reachable state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ReachableStateTime *time.Time `json:"reachableStateTime,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Linestatus) SetField(field string, fieldValue interface{}) {
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

func (o Linestatus) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ReachableStateTime", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Linestatus
	
	ReachableStateTime := new(string)
	if o.ReachableStateTime != nil {
		
		*ReachableStateTime = timeutil.Strftime(o.ReachableStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReachableStateTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Reachable *bool `json:"reachable,omitempty"`
		
		AddressOfRecord *string `json:"addressOfRecord,omitempty"`
		
		ContactAddresses *[]string `json:"contactAddresses,omitempty"`
		
		ReachableStateTime *string `json:"reachableStateTime,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Reachable: o.Reachable,
		
		AddressOfRecord: o.AddressOfRecord,
		
		ContactAddresses: o.ContactAddresses,
		
		ReachableStateTime: ReachableStateTime,
		Alias:    (Alias)(o),
	})
}

func (o *Linestatus) UnmarshalJSON(b []byte) error {
	var LinestatusMap map[string]interface{}
	err := json.Unmarshal(b, &LinestatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LinestatusMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Reachable, ok := LinestatusMap["reachable"].(bool); ok {
		o.Reachable = &Reachable
	}
    
	if AddressOfRecord, ok := LinestatusMap["addressOfRecord"].(string); ok {
		o.AddressOfRecord = &AddressOfRecord
	}
    
	if ContactAddresses, ok := LinestatusMap["contactAddresses"].([]interface{}); ok {
		ContactAddressesString, _ := json.Marshal(ContactAddresses)
		json.Unmarshal(ContactAddressesString, &o.ContactAddresses)
	}
	
	if reachableStateTimeString, ok := LinestatusMap["reachableStateTime"].(string); ok {
		ReachableStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", reachableStateTimeString)
		o.ReachableStateTime = &ReachableStateTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Linestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
