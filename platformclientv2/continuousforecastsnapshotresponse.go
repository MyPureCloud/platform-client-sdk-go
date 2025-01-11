package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Continuousforecastsnapshotresponse
type Continuousforecastsnapshotresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SessionId - Session Id of the continuous forecast
	SessionId *string `json:"sessionId,omitempty"`

	// SnapshotId - Snapshot Id of the continuous forecast session
	SnapshotId *string `json:"snapshotId,omitempty"`

	// Files - Link to the files containing data for requested snapshot
	Files *Snapshotfiles `json:"files,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Continuousforecastsnapshotresponse) SetField(field string, fieldValue interface{}) {
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

func (o Continuousforecastsnapshotresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Continuousforecastsnapshotresponse
	
	return json.Marshal(&struct { 
		SessionId *string `json:"sessionId,omitempty"`
		
		SnapshotId *string `json:"snapshotId,omitempty"`
		
		Files *Snapshotfiles `json:"files,omitempty"`
		Alias
	}{ 
		SessionId: o.SessionId,
		
		SnapshotId: o.SnapshotId,
		
		Files: o.Files,
		Alias:    (Alias)(o),
	})
}

func (o *Continuousforecastsnapshotresponse) UnmarshalJSON(b []byte) error {
	var ContinuousforecastsnapshotresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ContinuousforecastsnapshotresponseMap)
	if err != nil {
		return err
	}
	
	if SessionId, ok := ContinuousforecastsnapshotresponseMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if SnapshotId, ok := ContinuousforecastsnapshotresponseMap["snapshotId"].(string); ok {
		o.SnapshotId = &SnapshotId
	}
    
	if Files, ok := ContinuousforecastsnapshotresponseMap["files"].(map[string]interface{}); ok {
		FilesString, _ := json.Marshal(Files)
		json.Unmarshal(FilesString, &o.Files)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Continuousforecastsnapshotresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
