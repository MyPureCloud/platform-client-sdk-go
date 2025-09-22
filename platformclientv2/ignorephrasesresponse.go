package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Ignorephrasesresponse
type Ignorephrasesresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TotalPhrases - Total number of phrases in current org
	TotalPhrases *int `json:"totalPhrases,omitempty"`

	// AddedPhrases - Number of phrases added in current request
	AddedPhrases *int `json:"addedPhrases,omitempty"`

	// UpdatedPhrases - Number of phrases updated in current request
	UpdatedPhrases *int `json:"updatedPhrases,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Ignorephrasesresponse) SetField(field string, fieldValue interface{}) {
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

func (o Ignorephrasesresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Ignorephrasesresponse
	
	return json.Marshal(&struct { 
		TotalPhrases *int `json:"totalPhrases,omitempty"`
		
		AddedPhrases *int `json:"addedPhrases,omitempty"`
		
		UpdatedPhrases *int `json:"updatedPhrases,omitempty"`
		Alias
	}{ 
		TotalPhrases: o.TotalPhrases,
		
		AddedPhrases: o.AddedPhrases,
		
		UpdatedPhrases: o.UpdatedPhrases,
		Alias:    (Alias)(o),
	})
}

func (o *Ignorephrasesresponse) UnmarshalJSON(b []byte) error {
	var IgnorephrasesresponseMap map[string]interface{}
	err := json.Unmarshal(b, &IgnorephrasesresponseMap)
	if err != nil {
		return err
	}
	
	if TotalPhrases, ok := IgnorephrasesresponseMap["totalPhrases"].(float64); ok {
		TotalPhrasesInt := int(TotalPhrases)
		o.TotalPhrases = &TotalPhrasesInt
	}
	
	if AddedPhrases, ok := IgnorephrasesresponseMap["addedPhrases"].(float64); ok {
		AddedPhrasesInt := int(AddedPhrases)
		o.AddedPhrases = &AddedPhrasesInt
	}
	
	if UpdatedPhrases, ok := IgnorephrasesresponseMap["updatedPhrases"].(float64); ok {
		UpdatedPhrasesInt := int(UpdatedPhrases)
		o.UpdatedPhrases = &UpdatedPhrasesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Ignorephrasesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
