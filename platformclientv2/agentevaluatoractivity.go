package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentevaluatoractivity
type Agentevaluatoractivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Agent
	Agent *User `json:"agent,omitempty"`

	// Evaluator
	Evaluator *User `json:"evaluator,omitempty"`

	// NumEvaluations
	NumEvaluations *int `json:"numEvaluations,omitempty"`

	// AverageEvaluationScore
	AverageEvaluationScore *int `json:"averageEvaluationScore,omitempty"`

	// NumEvaluationsWithoutViewPermission
	NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentevaluatoractivity) SetField(field string, fieldValue interface{}) {
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

func (o Agentevaluatoractivity) MarshalJSON() ([]byte, error) {
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
	type Alias Agentevaluatoractivity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		Evaluator *User `json:"evaluator,omitempty"`
		
		NumEvaluations *int `json:"numEvaluations,omitempty"`
		
		AverageEvaluationScore *int `json:"averageEvaluationScore,omitempty"`
		
		NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Agent: o.Agent,
		
		Evaluator: o.Evaluator,
		
		NumEvaluations: o.NumEvaluations,
		
		AverageEvaluationScore: o.AverageEvaluationScore,
		
		NumEvaluationsWithoutViewPermission: o.NumEvaluationsWithoutViewPermission,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Agentevaluatoractivity) UnmarshalJSON(b []byte) error {
	var AgentevaluatoractivityMap map[string]interface{}
	err := json.Unmarshal(b, &AgentevaluatoractivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentevaluatoractivityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AgentevaluatoractivityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Agent, ok := AgentevaluatoractivityMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Evaluator, ok := AgentevaluatoractivityMap["evaluator"].(map[string]interface{}); ok {
		EvaluatorString, _ := json.Marshal(Evaluator)
		json.Unmarshal(EvaluatorString, &o.Evaluator)
	}
	
	if NumEvaluations, ok := AgentevaluatoractivityMap["numEvaluations"].(float64); ok {
		NumEvaluationsInt := int(NumEvaluations)
		o.NumEvaluations = &NumEvaluationsInt
	}
	
	if AverageEvaluationScore, ok := AgentevaluatoractivityMap["averageEvaluationScore"].(float64); ok {
		AverageEvaluationScoreInt := int(AverageEvaluationScore)
		o.AverageEvaluationScore = &AverageEvaluationScoreInt
	}
	
	if NumEvaluationsWithoutViewPermission, ok := AgentevaluatoractivityMap["numEvaluationsWithoutViewPermission"].(float64); ok {
		NumEvaluationsWithoutViewPermissionInt := int(NumEvaluationsWithoutViewPermission)
		o.NumEvaluationsWithoutViewPermission = &NumEvaluationsWithoutViewPermissionInt
	}
	
	if SelfUri, ok := AgentevaluatoractivityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentevaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
