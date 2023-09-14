package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowhealthutterance
type Flowhealthutterance struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Text - Utterance Text.
	Text *string `json:"text,omitempty"`

	// IssueCount - Number of issues found for this utterance.
	IssueCount *int `json:"issueCount,omitempty"`

	// Language - Language provided for this utterance's health.
	Language *string `json:"language,omitempty"`

	// StaticValidationResults - Validation results for the utterance.
	StaticValidationResults *[]string `json:"staticValidationResults,omitempty"`

	// OutlierInfo - Details about this utterance being an outlier or not.
	OutlierInfo *Outlierinfo `json:"outlierInfo,omitempty"`

	// ConfusionInfo - Confusion details with other utterances.
	ConfusionInfo *Confusiondetails `json:"confusionInfo,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowhealthutterance) SetField(field string, fieldValue interface{}) {
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

func (o Flowhealthutterance) MarshalJSON() ([]byte, error) {
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
	type Alias Flowhealthutterance
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		IssueCount *int `json:"issueCount,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		StaticValidationResults *[]string `json:"staticValidationResults,omitempty"`
		
		OutlierInfo *Outlierinfo `json:"outlierInfo,omitempty"`
		
		ConfusionInfo *Confusiondetails `json:"confusionInfo,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Text: o.Text,
		
		IssueCount: o.IssueCount,
		
		Language: o.Language,
		
		StaticValidationResults: o.StaticValidationResults,
		
		OutlierInfo: o.OutlierInfo,
		
		ConfusionInfo: o.ConfusionInfo,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Flowhealthutterance) UnmarshalJSON(b []byte) error {
	var FlowhealthutteranceMap map[string]interface{}
	err := json.Unmarshal(b, &FlowhealthutteranceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowhealthutteranceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Text, ok := FlowhealthutteranceMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if IssueCount, ok := FlowhealthutteranceMap["issueCount"].(float64); ok {
		IssueCountInt := int(IssueCount)
		o.IssueCount = &IssueCountInt
	}
	
	if Language, ok := FlowhealthutteranceMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if StaticValidationResults, ok := FlowhealthutteranceMap["staticValidationResults"].([]interface{}); ok {
		StaticValidationResultsString, _ := json.Marshal(StaticValidationResults)
		json.Unmarshal(StaticValidationResultsString, &o.StaticValidationResults)
	}
	
	if OutlierInfo, ok := FlowhealthutteranceMap["outlierInfo"].(map[string]interface{}); ok {
		OutlierInfoString, _ := json.Marshal(OutlierInfo)
		json.Unmarshal(OutlierInfoString, &o.OutlierInfo)
	}
	
	if ConfusionInfo, ok := FlowhealthutteranceMap["confusionInfo"].(map[string]interface{}); ok {
		ConfusionInfoString, _ := json.Marshal(ConfusionInfo)
		json.Unmarshal(ConfusionInfoString, &o.ConfusionInfo)
	}
	
	if SelfUri, ok := FlowhealthutteranceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowhealthutterance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
