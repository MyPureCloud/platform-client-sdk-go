package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkmetricsoptions
type Trunkinstancetopictrunkmetricsoptions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ProxyAddress
	ProxyAddress *string `json:"proxyAddress,omitempty"`

	// OptionState
	OptionState *bool `json:"optionState,omitempty"`

	// OptionStateTime
	OptionStateTime *time.Time `json:"optionStateTime,omitempty"`

	// ErrorInfo
	ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Trunkinstancetopictrunkmetricsoptions) SetField(field string, fieldValue interface{}) {
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

func (o Trunkinstancetopictrunkmetricsoptions) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "OptionStateTime", }
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
	type Alias Trunkinstancetopictrunkmetricsoptions
	
	OptionStateTime := new(string)
	if o.OptionStateTime != nil {
		
		*OptionStateTime = timeutil.Strftime(o.OptionStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		OptionStateTime = nil
	}
	
	return json.Marshal(&struct { 
		ProxyAddress *string `json:"proxyAddress,omitempty"`
		
		OptionState *bool `json:"optionState,omitempty"`
		
		OptionStateTime *string `json:"optionStateTime,omitempty"`
		
		ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`
		Alias
	}{ 
		ProxyAddress: o.ProxyAddress,
		
		OptionState: o.OptionState,
		
		OptionStateTime: OptionStateTime,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (Alias)(o),
	})
}

func (o *Trunkinstancetopictrunkmetricsoptions) UnmarshalJSON(b []byte) error {
	var TrunkinstancetopictrunkmetricsoptionsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkinstancetopictrunkmetricsoptionsMap)
	if err != nil {
		return err
	}
	
	if ProxyAddress, ok := TrunkinstancetopictrunkmetricsoptionsMap["proxyAddress"].(string); ok {
		o.ProxyAddress = &ProxyAddress
	}
    
	if OptionState, ok := TrunkinstancetopictrunkmetricsoptionsMap["optionState"].(bool); ok {
		o.OptionState = &OptionState
	}
    
	if optionStateTimeString, ok := TrunkinstancetopictrunkmetricsoptionsMap["optionStateTime"].(string); ok {
		OptionStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", optionStateTimeString)
		o.OptionStateTime = &OptionStateTime
	}
	
	if ErrorInfo, ok := TrunkinstancetopictrunkmetricsoptionsMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsoptions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
