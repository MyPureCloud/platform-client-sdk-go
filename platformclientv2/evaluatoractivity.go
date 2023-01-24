package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluatoractivity
type Evaluatoractivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Evaluator
	Evaluator *User `json:"evaluator,omitempty"`

	// NumEvaluationsAssigned
	NumEvaluationsAssigned *int `json:"numEvaluationsAssigned,omitempty"`

	// NumEvaluationsStarted
	NumEvaluationsStarted *int `json:"numEvaluationsStarted,omitempty"`

	// NumEvaluationsCompleted
	NumEvaluationsCompleted *int `json:"numEvaluationsCompleted,omitempty"`

	// NumCalibrationsAssigned
	NumCalibrationsAssigned *int `json:"numCalibrationsAssigned,omitempty"`

	// NumCalibrationsStarted
	NumCalibrationsStarted *int `json:"numCalibrationsStarted,omitempty"`

	// NumCalibrationsCompleted
	NumCalibrationsCompleted *int `json:"numCalibrationsCompleted,omitempty"`

	// NumEvaluationsWithoutViewPermission
	NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluatoractivity) SetField(field string, fieldValue interface{}) {
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

func (o Evaluatoractivity) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluatoractivity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Evaluator *User `json:"evaluator,omitempty"`
		
		NumEvaluationsAssigned *int `json:"numEvaluationsAssigned,omitempty"`
		
		NumEvaluationsStarted *int `json:"numEvaluationsStarted,omitempty"`
		
		NumEvaluationsCompleted *int `json:"numEvaluationsCompleted,omitempty"`
		
		NumCalibrationsAssigned *int `json:"numCalibrationsAssigned,omitempty"`
		
		NumCalibrationsStarted *int `json:"numCalibrationsStarted,omitempty"`
		
		NumCalibrationsCompleted *int `json:"numCalibrationsCompleted,omitempty"`
		
		NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Evaluator: o.Evaluator,
		
		NumEvaluationsAssigned: o.NumEvaluationsAssigned,
		
		NumEvaluationsStarted: o.NumEvaluationsStarted,
		
		NumEvaluationsCompleted: o.NumEvaluationsCompleted,
		
		NumCalibrationsAssigned: o.NumCalibrationsAssigned,
		
		NumCalibrationsStarted: o.NumCalibrationsStarted,
		
		NumCalibrationsCompleted: o.NumCalibrationsCompleted,
		
		NumEvaluationsWithoutViewPermission: o.NumEvaluationsWithoutViewPermission,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluatoractivity) UnmarshalJSON(b []byte) error {
	var EvaluatoractivityMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluatoractivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EvaluatoractivityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EvaluatoractivityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Evaluator, ok := EvaluatoractivityMap["evaluator"].(map[string]interface{}); ok {
		EvaluatorString, _ := json.Marshal(Evaluator)
		json.Unmarshal(EvaluatorString, &o.Evaluator)
	}
	
	if NumEvaluationsAssigned, ok := EvaluatoractivityMap["numEvaluationsAssigned"].(float64); ok {
		NumEvaluationsAssignedInt := int(NumEvaluationsAssigned)
		o.NumEvaluationsAssigned = &NumEvaluationsAssignedInt
	}
	
	if NumEvaluationsStarted, ok := EvaluatoractivityMap["numEvaluationsStarted"].(float64); ok {
		NumEvaluationsStartedInt := int(NumEvaluationsStarted)
		o.NumEvaluationsStarted = &NumEvaluationsStartedInt
	}
	
	if NumEvaluationsCompleted, ok := EvaluatoractivityMap["numEvaluationsCompleted"].(float64); ok {
		NumEvaluationsCompletedInt := int(NumEvaluationsCompleted)
		o.NumEvaluationsCompleted = &NumEvaluationsCompletedInt
	}
	
	if NumCalibrationsAssigned, ok := EvaluatoractivityMap["numCalibrationsAssigned"].(float64); ok {
		NumCalibrationsAssignedInt := int(NumCalibrationsAssigned)
		o.NumCalibrationsAssigned = &NumCalibrationsAssignedInt
	}
	
	if NumCalibrationsStarted, ok := EvaluatoractivityMap["numCalibrationsStarted"].(float64); ok {
		NumCalibrationsStartedInt := int(NumCalibrationsStarted)
		o.NumCalibrationsStarted = &NumCalibrationsStartedInt
	}
	
	if NumCalibrationsCompleted, ok := EvaluatoractivityMap["numCalibrationsCompleted"].(float64); ok {
		NumCalibrationsCompletedInt := int(NumCalibrationsCompleted)
		o.NumCalibrationsCompleted = &NumCalibrationsCompletedInt
	}
	
	if NumEvaluationsWithoutViewPermission, ok := EvaluatoractivityMap["numEvaluationsWithoutViewPermission"].(float64); ok {
		NumEvaluationsWithoutViewPermissionInt := int(NumEvaluationsWithoutViewPermission)
		o.NumEvaluationsWithoutViewPermission = &NumEvaluationsWithoutViewPermissionInt
	}
	
	if SelfUri, ok := EvaluatoractivityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
