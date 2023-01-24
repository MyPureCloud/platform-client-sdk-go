package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricsprocessor
type Edgemetricsprocessor struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActiveTimePct - Percent time processor was active.
	ActiveTimePct *float64 `json:"activeTimePct,omitempty"`

	// CpuId - Machine CPU identifier. 'total' will always be included in the array and is the total of all CPU resources.
	CpuId *string `json:"cpuId,omitempty"`

	// IdleTimePct - Percent time processor was idle.
	IdleTimePct *float64 `json:"idleTimePct,omitempty"`

	// PrivilegedTimePct - Percent time processor spent in privileged mode.
	PrivilegedTimePct *float64 `json:"privilegedTimePct,omitempty"`

	// UserTimePct - Percent time processor spent in user mode.
	UserTimePct *float64 `json:"userTimePct,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgemetricsprocessor) SetField(field string, fieldValue interface{}) {
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

func (o Edgemetricsprocessor) MarshalJSON() ([]byte, error) {
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
	type Alias Edgemetricsprocessor
	
	return json.Marshal(&struct { 
		ActiveTimePct *float64 `json:"activeTimePct,omitempty"`
		
		CpuId *string `json:"cpuId,omitempty"`
		
		IdleTimePct *float64 `json:"idleTimePct,omitempty"`
		
		PrivilegedTimePct *float64 `json:"privilegedTimePct,omitempty"`
		
		UserTimePct *float64 `json:"userTimePct,omitempty"`
		Alias
	}{ 
		ActiveTimePct: o.ActiveTimePct,
		
		CpuId: o.CpuId,
		
		IdleTimePct: o.IdleTimePct,
		
		PrivilegedTimePct: o.PrivilegedTimePct,
		
		UserTimePct: o.UserTimePct,
		Alias:    (Alias)(o),
	})
}

func (o *Edgemetricsprocessor) UnmarshalJSON(b []byte) error {
	var EdgemetricsprocessorMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricsprocessorMap)
	if err != nil {
		return err
	}
	
	if ActiveTimePct, ok := EdgemetricsprocessorMap["activeTimePct"].(float64); ok {
		o.ActiveTimePct = &ActiveTimePct
	}
    
	if CpuId, ok := EdgemetricsprocessorMap["cpuId"].(string); ok {
		o.CpuId = &CpuId
	}
    
	if IdleTimePct, ok := EdgemetricsprocessorMap["idleTimePct"].(float64); ok {
		o.IdleTimePct = &IdleTimePct
	}
    
	if PrivilegedTimePct, ok := EdgemetricsprocessorMap["privilegedTimePct"].(float64); ok {
		o.PrivilegedTimePct = &PrivilegedTimePct
	}
    
	if UserTimePct, ok := EdgemetricsprocessorMap["userTimePct"].(float64); ok {
		o.UserTimePct = &UserTimePct
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricsprocessor) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
