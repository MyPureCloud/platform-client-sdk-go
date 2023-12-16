package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Performancepredictionrecalculationresponse
type Performancepredictionrecalculationresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OperationId - The operationId for which to listen
	OperationId *string `json:"operationId,omitempty"`

	// DownloadUrl - The url to GET the results of the performance prediction. This field is populated only if query state is 'Complete'
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
	DownloadResult *Performancepredictionoutputs `json:"downloadResult,omitempty"`

	// State - The state of the performance prediction
	State *string `json:"state,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Performancepredictionrecalculationresponse) SetField(field string, fieldValue interface{}) {
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

func (o Performancepredictionrecalculationresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Performancepredictionrecalculationresponse
	
	return json.Marshal(&struct { 
		OperationId *string `json:"operationId,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadResult *Performancepredictionoutputs `json:"downloadResult,omitempty"`
		
		State *string `json:"state,omitempty"`
		Alias
	}{ 
		OperationId: o.OperationId,
		
		DownloadUrl: o.DownloadUrl,
		
		DownloadResult: o.DownloadResult,
		
		State: o.State,
		Alias:    (Alias)(o),
	})
}

func (o *Performancepredictionrecalculationresponse) UnmarshalJSON(b []byte) error {
	var PerformancepredictionrecalculationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &PerformancepredictionrecalculationresponseMap)
	if err != nil {
		return err
	}
	
	if OperationId, ok := PerformancepredictionrecalculationresponseMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if DownloadUrl, ok := PerformancepredictionrecalculationresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if DownloadResult, ok := PerformancepredictionrecalculationresponseMap["downloadResult"].(map[string]interface{}); ok {
		DownloadResultString, _ := json.Marshal(DownloadResult)
		json.Unmarshal(DownloadResultString, &o.DownloadResult)
	}
	
	if State, ok := PerformancepredictionrecalculationresponseMap["state"].(string); ok {
		o.State = &State
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Performancepredictionrecalculationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
