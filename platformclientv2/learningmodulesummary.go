package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulesummary - Learning module summary data
type Learningmodulesummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AssignedCount - The total number of assignments assigned for a learning module
	AssignedCount *int `json:"assignedCount,omitempty"`

	// CompletedCount - The number of assignments completed for a learning module
	CompletedCount *int `json:"completedCount,omitempty"`

	// PassedCount - The number of assignments passed for a learning module
	PassedCount *int `json:"passedCount,omitempty"`

	// CompletedSum - The sum of assignment scores for a learning module
	CompletedSum *float32 `json:"completedSum,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodulesummary) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodulesummary) MarshalJSON() ([]byte, error) {
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
	type Alias Learningmodulesummary
	
	return json.Marshal(&struct { 
		AssignedCount *int `json:"assignedCount,omitempty"`
		
		CompletedCount *int `json:"completedCount,omitempty"`
		
		PassedCount *int `json:"passedCount,omitempty"`
		
		CompletedSum *float32 `json:"completedSum,omitempty"`
		Alias
	}{ 
		AssignedCount: o.AssignedCount,
		
		CompletedCount: o.CompletedCount,
		
		PassedCount: o.PassedCount,
		
		CompletedSum: o.CompletedSum,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodulesummary) UnmarshalJSON(b []byte) error {
	var LearningmodulesummaryMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulesummaryMap)
	if err != nil {
		return err
	}
	
	if AssignedCount, ok := LearningmodulesummaryMap["assignedCount"].(float64); ok {
		AssignedCountInt := int(AssignedCount)
		o.AssignedCount = &AssignedCountInt
	}
	
	if CompletedCount, ok := LearningmodulesummaryMap["completedCount"].(float64); ok {
		CompletedCountInt := int(CompletedCount)
		o.CompletedCount = &CompletedCountInt
	}
	
	if PassedCount, ok := LearningmodulesummaryMap["passedCount"].(float64); ok {
		PassedCountInt := int(PassedCount)
		o.PassedCount = &PassedCountInt
	}
	
	if CompletedSum, ok := LearningmodulesummaryMap["completedSum"].(float64); ok {
		CompletedSumFloat32 := float32(CompletedSum)
		o.CompletedSum = &CompletedSumFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulesummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
