package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Snippetrecordingsettings
type Snippetrecordingsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DaysToArchive - The number of days after a snippet recording has ended before it is archived. Use -1 to never archive, or a positive number
	DaysToArchive *int `json:"daysToArchive,omitempty"`

	// DaysToDelete - The number of days after a snippet recording has ended before it is deleted. Use -1 to retain indefinitely, or a positive number
	DaysToDelete *int `json:"daysToDelete,omitempty"`

	// ExportIntegration - The aws-s3-recording-bulk-actions-integration used for snippet recording exports.
	ExportIntegration *Domainentityref `json:"exportIntegration,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Snippetrecordingsettings) SetField(field string, fieldValue interface{}) {
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

func (o Snippetrecordingsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Snippetrecordingsettings
	
	return json.Marshal(&struct { 
		DaysToArchive *int `json:"daysToArchive,omitempty"`
		
		DaysToDelete *int `json:"daysToDelete,omitempty"`
		
		ExportIntegration *Domainentityref `json:"exportIntegration,omitempty"`
		Alias
	}{ 
		DaysToArchive: o.DaysToArchive,
		
		DaysToDelete: o.DaysToDelete,
		
		ExportIntegration: o.ExportIntegration,
		Alias:    (Alias)(o),
	})
}

func (o *Snippetrecordingsettings) UnmarshalJSON(b []byte) error {
	var SnippetrecordingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &SnippetrecordingsettingsMap)
	if err != nil {
		return err
	}
	
	if DaysToArchive, ok := SnippetrecordingsettingsMap["daysToArchive"].(float64); ok {
		DaysToArchiveInt := int(DaysToArchive)
		o.DaysToArchive = &DaysToArchiveInt
	}
	
	if DaysToDelete, ok := SnippetrecordingsettingsMap["daysToDelete"].(float64); ok {
		DaysToDeleteInt := int(DaysToDelete)
		o.DaysToDelete = &DaysToDeleteInt
	}
	
	if ExportIntegration, ok := SnippetrecordingsettingsMap["exportIntegration"].(map[string]interface{}); ok {
		ExportIntegrationString, _ := json.Marshal(ExportIntegration)
		json.Unmarshal(ExportIntegrationString, &o.ExportIntegration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Snippetrecordingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
