package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Healthinfo
type Healthinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status - Status of health computation for this intent.
	Status *string `json:"status,omitempty"`

	// ErrorInfo - Error details for the intent, if any.
	ErrorInfo *Flowhealtherrorinfo `json:"errorInfo,omitempty"`

	// OverallScore - Overall health score for the intent ranged between 0 and 100 as 100 is the perfect health score.
	OverallScore *float32 `json:"overallScore,omitempty"`

	// IssueCount - Number of issues found in the intent.
	IssueCount *int `json:"issueCount,omitempty"`

	// StaticValidationResults - Validation results for the intent.
	StaticValidationResults *[]string `json:"staticValidationResults,omitempty"`

	// Utterances - Utterances for this intent.
	Utterances *[]Flowhealthintentutterance `json:"utterances,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Healthinfo) SetField(field string, fieldValue interface{}) {
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

func (o Healthinfo) MarshalJSON() ([]byte, error) {
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
	type Alias Healthinfo
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		ErrorInfo *Flowhealtherrorinfo `json:"errorInfo,omitempty"`
		
		OverallScore *float32 `json:"overallScore,omitempty"`
		
		IssueCount *int `json:"issueCount,omitempty"`
		
		StaticValidationResults *[]string `json:"staticValidationResults,omitempty"`
		
		Utterances *[]Flowhealthintentutterance `json:"utterances,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		ErrorInfo: o.ErrorInfo,
		
		OverallScore: o.OverallScore,
		
		IssueCount: o.IssueCount,
		
		StaticValidationResults: o.StaticValidationResults,
		
		Utterances: o.Utterances,
		Alias:    (Alias)(o),
	})
}

func (o *Healthinfo) UnmarshalJSON(b []byte) error {
	var HealthinfoMap map[string]interface{}
	err := json.Unmarshal(b, &HealthinfoMap)
	if err != nil {
		return err
	}
	
	if Status, ok := HealthinfoMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ErrorInfo, ok := HealthinfoMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if OverallScore, ok := HealthinfoMap["overallScore"].(float64); ok {
		OverallScoreFloat32 := float32(OverallScore)
		o.OverallScore = &OverallScoreFloat32
	}
	
	if IssueCount, ok := HealthinfoMap["issueCount"].(float64); ok {
		IssueCountInt := int(IssueCount)
		o.IssueCount = &IssueCountInt
	}
	
	if StaticValidationResults, ok := HealthinfoMap["staticValidationResults"].([]interface{}); ok {
		StaticValidationResultsString, _ := json.Marshal(StaticValidationResults)
		json.Unmarshal(StaticValidationResultsString, &o.StaticValidationResults)
	}
	
	if Utterances, ok := HealthinfoMap["utterances"].([]interface{}); ok {
		UtterancesString, _ := json.Marshal(Utterances)
		json.Unmarshal(UtterancesString, &o.Utterances)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Healthinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
