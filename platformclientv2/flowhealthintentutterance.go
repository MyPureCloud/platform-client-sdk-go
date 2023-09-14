package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowhealthintentutterance
type Flowhealthintentutterance struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Text - Utterance Text.
	Text *string `json:"text,omitempty"`

	// IssueCount - Number of issues found for this utterance.
	IssueCount *int `json:"issueCount,omitempty"`

	// StaticValidationResults - Validation results for this utterance.
	StaticValidationResults *[]string `json:"staticValidationResults,omitempty"`

	// OutlierInfo - Details about this utterance being an outlier or not.
	OutlierInfo *Outlierinfo `json:"outlierInfo,omitempty"`

	// ConfusionInfo - Confusion details with other utterances.
	ConfusionInfo *Confusioninfo `json:"confusionInfo,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowhealthintentutterance) SetField(field string, fieldValue interface{}) {
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

func (o Flowhealthintentutterance) MarshalJSON() ([]byte, error) {
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
	type Alias Flowhealthintentutterance
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		IssueCount *int `json:"issueCount,omitempty"`
		
		StaticValidationResults *[]string `json:"staticValidationResults,omitempty"`
		
		OutlierInfo *Outlierinfo `json:"outlierInfo,omitempty"`
		
		ConfusionInfo *Confusioninfo `json:"confusionInfo,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		IssueCount: o.IssueCount,
		
		StaticValidationResults: o.StaticValidationResults,
		
		OutlierInfo: o.OutlierInfo,
		
		ConfusionInfo: o.ConfusionInfo,
		Alias:    (Alias)(o),
	})
}

func (o *Flowhealthintentutterance) UnmarshalJSON(b []byte) error {
	var FlowhealthintentutteranceMap map[string]interface{}
	err := json.Unmarshal(b, &FlowhealthintentutteranceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowhealthintentutteranceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Text, ok := FlowhealthintentutteranceMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if IssueCount, ok := FlowhealthintentutteranceMap["issueCount"].(float64); ok {
		IssueCountInt := int(IssueCount)
		o.IssueCount = &IssueCountInt
	}
	
	if StaticValidationResults, ok := FlowhealthintentutteranceMap["staticValidationResults"].([]interface{}); ok {
		StaticValidationResultsString, _ := json.Marshal(StaticValidationResults)
		json.Unmarshal(StaticValidationResultsString, &o.StaticValidationResults)
	}
	
	if OutlierInfo, ok := FlowhealthintentutteranceMap["outlierInfo"].(map[string]interface{}); ok {
		OutlierInfoString, _ := json.Marshal(OutlierInfo)
		json.Unmarshal(OutlierInfoString, &o.OutlierInfo)
	}
	
	if ConfusionInfo, ok := FlowhealthintentutteranceMap["confusionInfo"].(map[string]interface{}); ok {
		ConfusionInfoString, _ := json.Marshal(ConfusionInfo)
		json.Unmarshal(ConfusionInfoString, &o.ConfusionInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowhealthintentutterance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
