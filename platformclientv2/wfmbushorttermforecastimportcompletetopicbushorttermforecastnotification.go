package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification
type Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status
	Status *string `json:"status,omitempty"`

	// Result
	Result *Wfmbushorttermforecastimportcompletetopicbushorttermforecast `json:"result,omitempty"`

	// OperationId
	OperationId *string `json:"operationId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		Result *Wfmbushorttermforecastimportcompletetopicbushorttermforecast `json:"result,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		Result: o.Result,
		
		OperationId: o.OperationId,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastimportcompletetopicbushorttermforecastnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastimportcompletetopicbushorttermforecastnotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Result, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastnotificationMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if OperationId, ok := WfmbushorttermforecastimportcompletetopicbushorttermforecastnotificationMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicbushorttermforecastnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
