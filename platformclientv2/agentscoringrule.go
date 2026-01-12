package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentscoringrule
type Agentscoringrule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ProgramId - The program ID that this rule belongs to. When provided in request body, this value is ignored in favor of the path parameter.
	ProgramId *string `json:"programId,omitempty"`

	// SamplingType - Sampling type setting. Valid values: All, Percentage
	SamplingType *string `json:"samplingType,omitempty"`

	// SamplingPercentage - Percentage of interactions to evaluate (0.01-100.00). Required when samplingType is Percentage, null when All.
	SamplingPercentage *float32 `json:"samplingPercentage,omitempty"`

	// SubmissionType - Submission type for evaluations. Valid values: Automated, Manual
	SubmissionType *string `json:"submissionType,omitempty"`

	// EvaluationFormContextId - The evaluation form contextID to use for scoring.
	EvaluationFormContextId *string `json:"evaluationFormContextId,omitempty"`

	// Enabled - Whether the rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Published - Whether the rule is published.
	Published *bool `json:"published,omitempty"`

	// Evaluator - The evaluator for evaluations created by this rule.
	Evaluator *Addressableentityref `json:"evaluator,omitempty"`

	// DateCreated - Date when the rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date when the rule was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentscoringrule) SetField(field string, fieldValue interface{}) {
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

func (o Agentscoringrule) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Agentscoringrule
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ProgramId *string `json:"programId,omitempty"`
		
		SamplingType *string `json:"samplingType,omitempty"`
		
		SamplingPercentage *float32 `json:"samplingPercentage,omitempty"`
		
		SubmissionType *string `json:"submissionType,omitempty"`
		
		EvaluationFormContextId *string `json:"evaluationFormContextId,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		Evaluator *Addressableentityref `json:"evaluator,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ProgramId: o.ProgramId,
		
		SamplingType: o.SamplingType,
		
		SamplingPercentage: o.SamplingPercentage,
		
		SubmissionType: o.SubmissionType,
		
		EvaluationFormContextId: o.EvaluationFormContextId,
		
		Enabled: o.Enabled,
		
		Published: o.Published,
		
		Evaluator: o.Evaluator,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Agentscoringrule) UnmarshalJSON(b []byte) error {
	var AgentscoringruleMap map[string]interface{}
	err := json.Unmarshal(b, &AgentscoringruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentscoringruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ProgramId, ok := AgentscoringruleMap["programId"].(string); ok {
		o.ProgramId = &ProgramId
	}
    
	if SamplingType, ok := AgentscoringruleMap["samplingType"].(string); ok {
		o.SamplingType = &SamplingType
	}
    
	if SamplingPercentage, ok := AgentscoringruleMap["samplingPercentage"].(float64); ok {
		SamplingPercentageFloat32 := float32(SamplingPercentage)
		o.SamplingPercentage = &SamplingPercentageFloat32
	}
    
	if SubmissionType, ok := AgentscoringruleMap["submissionType"].(string); ok {
		o.SubmissionType = &SubmissionType
	}
    
	if EvaluationFormContextId, ok := AgentscoringruleMap["evaluationFormContextId"].(string); ok {
		o.EvaluationFormContextId = &EvaluationFormContextId
	}
    
	if Enabled, ok := AgentscoringruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Published, ok := AgentscoringruleMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if Evaluator, ok := AgentscoringruleMap["evaluator"].(map[string]interface{}); ok {
		EvaluatorString, _ := json.Marshal(Evaluator)
		json.Unmarshal(EvaluatorString, &o.Evaluator)
	}
	
	if dateCreatedString, ok := AgentscoringruleMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := AgentscoringruleMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := AgentscoringruleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentscoringrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
