package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricalshrinkageresponse
type Wfmhistoricalshrinkageresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OperationId - The operationId for which to listen
	OperationId *string `json:"operationId,omitempty"`

	// DownloadUrls - The url list to GET the results of the Historical Shrinkage query. This field is populated only if query state is Complete
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`

	// DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
	DownloadResult *Historicalshrinkageresultlisting `json:"downloadResult,omitempty"`

	// State - The state of the shrinkage query
	State *string `json:"state,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmhistoricalshrinkageresponse) SetField(field string, fieldValue interface{}) {
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

func (o Wfmhistoricalshrinkageresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmhistoricalshrinkageresponse
	
	return json.Marshal(&struct { 
		OperationId *string `json:"operationId,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		DownloadResult *Historicalshrinkageresultlisting `json:"downloadResult,omitempty"`
		
		State *string `json:"state,omitempty"`
		Alias
	}{ 
		OperationId: o.OperationId,
		
		DownloadUrls: o.DownloadUrls,
		
		DownloadResult: o.DownloadResult,
		
		State: o.State,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmhistoricalshrinkageresponse) UnmarshalJSON(b []byte) error {
	var WfmhistoricalshrinkageresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricalshrinkageresponseMap)
	if err != nil {
		return err
	}
	
	if OperationId, ok := WfmhistoricalshrinkageresponseMap["operationId"].(string); ok {
		o.OperationId = &OperationId
	}
    
	if DownloadUrls, ok := WfmhistoricalshrinkageresponseMap["downloadUrls"].([]interface{}); ok {
		DownloadUrlsString, _ := json.Marshal(DownloadUrls)
		json.Unmarshal(DownloadUrlsString, &o.DownloadUrls)
	}
	
	if DownloadResult, ok := WfmhistoricalshrinkageresponseMap["downloadResult"].(map[string]interface{}); ok {
		DownloadResultString, _ := json.Marshal(DownloadResult)
		json.Unmarshal(DownloadResultString, &o.DownloadResult)
	}
	
	if State, ok := WfmhistoricalshrinkageresponseMap["state"].(string); ok {
		o.State = &State
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricalshrinkageresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
