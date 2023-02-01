package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Batchdownloadjobstatusresult
type Batchdownloadjobstatusresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// JobId - JobId returned when job was initially submitted
	JobId *string `json:"jobId,omitempty"`

	// ExpectedResultCount - Number of results expected when job is completed
	ExpectedResultCount *int `json:"expectedResultCount,omitempty"`

	// ResultCount - Current number of results available
	ResultCount *int `json:"resultCount,omitempty"`

	// ErrorCount - Number of error results produced so far
	ErrorCount *int `json:"errorCount,omitempty"`

	// Results - Current set of results for the job
	Results *[]Batchdownloadjobresult `json:"results,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Batchdownloadjobstatusresult) SetField(field string, fieldValue interface{}) {
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

func (o Batchdownloadjobstatusresult) MarshalJSON() ([]byte, error) {
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
	type Alias Batchdownloadjobstatusresult
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		JobId *string `json:"jobId,omitempty"`
		
		ExpectedResultCount *int `json:"expectedResultCount,omitempty"`
		
		ResultCount *int `json:"resultCount,omitempty"`
		
		ErrorCount *int `json:"errorCount,omitempty"`
		
		Results *[]Batchdownloadjobresult `json:"results,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		JobId: o.JobId,
		
		ExpectedResultCount: o.ExpectedResultCount,
		
		ResultCount: o.ResultCount,
		
		ErrorCount: o.ErrorCount,
		
		Results: o.Results,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Batchdownloadjobstatusresult) UnmarshalJSON(b []byte) error {
	var BatchdownloadjobstatusresultMap map[string]interface{}
	err := json.Unmarshal(b, &BatchdownloadjobstatusresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BatchdownloadjobstatusresultMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if JobId, ok := BatchdownloadjobstatusresultMap["jobId"].(string); ok {
		o.JobId = &JobId
	}
    
	if ExpectedResultCount, ok := BatchdownloadjobstatusresultMap["expectedResultCount"].(float64); ok {
		ExpectedResultCountInt := int(ExpectedResultCount)
		o.ExpectedResultCount = &ExpectedResultCountInt
	}
	
	if ResultCount, ok := BatchdownloadjobstatusresultMap["resultCount"].(float64); ok {
		ResultCountInt := int(ResultCount)
		o.ResultCount = &ResultCountInt
	}
	
	if ErrorCount, ok := BatchdownloadjobstatusresultMap["errorCount"].(float64); ok {
		ErrorCountInt := int(ErrorCount)
		o.ErrorCount = &ErrorCountInt
	}
	
	if Results, ok := BatchdownloadjobstatusresultMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if SelfUri, ok := BatchdownloadjobstatusresultMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Batchdownloadjobstatusresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
