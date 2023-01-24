package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludomainversionqualityreport
type Nludomainversionqualityreport struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Version - The domain and version details of the quality report
	Version *Nludomainversion `json:"version,omitempty"`

	// ConfusionMatrix - The confusion matrix for the Domain Version
	ConfusionMatrix *[]Nluconfusionmatrixrow `json:"confusionMatrix,omitempty"`

	// Summary - The quality report summary for the Domain Version
	Summary *Nluqualityreportsummary `json:"summary,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Nludomainversionqualityreport) SetField(field string, fieldValue interface{}) {
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

func (o Nludomainversionqualityreport) MarshalJSON() ([]byte, error) {
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
	type Alias Nludomainversionqualityreport
	
	return json.Marshal(&struct { 
		Version *Nludomainversion `json:"version,omitempty"`
		
		ConfusionMatrix *[]Nluconfusionmatrixrow `json:"confusionMatrix,omitempty"`
		
		Summary *Nluqualityreportsummary `json:"summary,omitempty"`
		Alias
	}{ 
		Version: o.Version,
		
		ConfusionMatrix: o.ConfusionMatrix,
		
		Summary: o.Summary,
		Alias:    (Alias)(o),
	})
}

func (o *Nludomainversionqualityreport) UnmarshalJSON(b []byte) error {
	var NludomainversionqualityreportMap map[string]interface{}
	err := json.Unmarshal(b, &NludomainversionqualityreportMap)
	if err != nil {
		return err
	}
	
	if Version, ok := NludomainversionqualityreportMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if ConfusionMatrix, ok := NludomainversionqualityreportMap["confusionMatrix"].([]interface{}); ok {
		ConfusionMatrixString, _ := json.Marshal(ConfusionMatrix)
		json.Unmarshal(ConfusionMatrixString, &o.ConfusionMatrix)
	}
	
	if Summary, ok := NludomainversionqualityreportMap["summary"].(map[string]interface{}); ok {
		SummaryString, _ := json.Marshal(Summary)
		json.Unmarshal(SummaryString, &o.Summary)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nludomainversionqualityreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
