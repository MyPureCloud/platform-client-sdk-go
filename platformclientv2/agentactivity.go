package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentactivity
type Agentactivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Agent
	Agent *User `json:"agent,omitempty"`

	// NumEvaluations
	NumEvaluations *int `json:"numEvaluations,omitempty"`

	// AverageEvaluationScore
	AverageEvaluationScore *int `json:"averageEvaluationScore,omitempty"`

	// NumCriticalEvaluations
	NumCriticalEvaluations *int `json:"numCriticalEvaluations,omitempty"`

	// AverageCriticalScore
	AverageCriticalScore *float32 `json:"averageCriticalScore,omitempty"`

	// HighestEvaluationScore
	HighestEvaluationScore *float32 `json:"highestEvaluationScore,omitempty"`

	// LowestEvaluationScore
	LowestEvaluationScore *float32 `json:"lowestEvaluationScore,omitempty"`

	// HighestCriticalScore
	HighestCriticalScore *float32 `json:"highestCriticalScore,omitempty"`

	// LowestCriticalScore
	LowestCriticalScore *float32 `json:"lowestCriticalScore,omitempty"`

	// AgentEvaluatorActivityList
	AgentEvaluatorActivityList *[]Agentevaluatoractivity `json:"agentEvaluatorActivityList,omitempty"`

	// NumEvaluationsWithoutViewPermission
	NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentactivity) SetField(field string, fieldValue interface{}) {
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

func (o Agentactivity) MarshalJSON() ([]byte, error) {
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
	type Alias Agentactivity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		NumEvaluations *int `json:"numEvaluations,omitempty"`
		
		AverageEvaluationScore *int `json:"averageEvaluationScore,omitempty"`
		
		NumCriticalEvaluations *int `json:"numCriticalEvaluations,omitempty"`
		
		AverageCriticalScore *float32 `json:"averageCriticalScore,omitempty"`
		
		HighestEvaluationScore *float32 `json:"highestEvaluationScore,omitempty"`
		
		LowestEvaluationScore *float32 `json:"lowestEvaluationScore,omitempty"`
		
		HighestCriticalScore *float32 `json:"highestCriticalScore,omitempty"`
		
		LowestCriticalScore *float32 `json:"lowestCriticalScore,omitempty"`
		
		AgentEvaluatorActivityList *[]Agentevaluatoractivity `json:"agentEvaluatorActivityList,omitempty"`
		
		NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Agent: o.Agent,
		
		NumEvaluations: o.NumEvaluations,
		
		AverageEvaluationScore: o.AverageEvaluationScore,
		
		NumCriticalEvaluations: o.NumCriticalEvaluations,
		
		AverageCriticalScore: o.AverageCriticalScore,
		
		HighestEvaluationScore: o.HighestEvaluationScore,
		
		LowestEvaluationScore: o.LowestEvaluationScore,
		
		HighestCriticalScore: o.HighestCriticalScore,
		
		LowestCriticalScore: o.LowestCriticalScore,
		
		AgentEvaluatorActivityList: o.AgentEvaluatorActivityList,
		
		NumEvaluationsWithoutViewPermission: o.NumEvaluationsWithoutViewPermission,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Agentactivity) UnmarshalJSON(b []byte) error {
	var AgentactivityMap map[string]interface{}
	err := json.Unmarshal(b, &AgentactivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentactivityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AgentactivityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Agent, ok := AgentactivityMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if NumEvaluations, ok := AgentactivityMap["numEvaluations"].(float64); ok {
		NumEvaluationsInt := int(NumEvaluations)
		o.NumEvaluations = &NumEvaluationsInt
	}
	
	if AverageEvaluationScore, ok := AgentactivityMap["averageEvaluationScore"].(float64); ok {
		AverageEvaluationScoreInt := int(AverageEvaluationScore)
		o.AverageEvaluationScore = &AverageEvaluationScoreInt
	}
	
	if NumCriticalEvaluations, ok := AgentactivityMap["numCriticalEvaluations"].(float64); ok {
		NumCriticalEvaluationsInt := int(NumCriticalEvaluations)
		o.NumCriticalEvaluations = &NumCriticalEvaluationsInt
	}
	
	if AverageCriticalScore, ok := AgentactivityMap["averageCriticalScore"].(float64); ok {
		AverageCriticalScoreFloat32 := float32(AverageCriticalScore)
		o.AverageCriticalScore = &AverageCriticalScoreFloat32
	}
	
	if HighestEvaluationScore, ok := AgentactivityMap["highestEvaluationScore"].(float64); ok {
		HighestEvaluationScoreFloat32 := float32(HighestEvaluationScore)
		o.HighestEvaluationScore = &HighestEvaluationScoreFloat32
	}
	
	if LowestEvaluationScore, ok := AgentactivityMap["lowestEvaluationScore"].(float64); ok {
		LowestEvaluationScoreFloat32 := float32(LowestEvaluationScore)
		o.LowestEvaluationScore = &LowestEvaluationScoreFloat32
	}
	
	if HighestCriticalScore, ok := AgentactivityMap["highestCriticalScore"].(float64); ok {
		HighestCriticalScoreFloat32 := float32(HighestCriticalScore)
		o.HighestCriticalScore = &HighestCriticalScoreFloat32
	}
	
	if LowestCriticalScore, ok := AgentactivityMap["lowestCriticalScore"].(float64); ok {
		LowestCriticalScoreFloat32 := float32(LowestCriticalScore)
		o.LowestCriticalScore = &LowestCriticalScoreFloat32
	}
	
	if AgentEvaluatorActivityList, ok := AgentactivityMap["agentEvaluatorActivityList"].([]interface{}); ok {
		AgentEvaluatorActivityListString, _ := json.Marshal(AgentEvaluatorActivityList)
		json.Unmarshal(AgentEvaluatorActivityListString, &o.AgentEvaluatorActivityList)
	}
	
	if NumEvaluationsWithoutViewPermission, ok := AgentactivityMap["numEvaluationsWithoutViewPermission"].(float64); ok {
		NumEvaluationsWithoutViewPermissionInt := int(NumEvaluationsWithoutViewPermission)
		o.NumEvaluationsWithoutViewPermission = &NumEvaluationsWithoutViewPermissionInt
	}
	
	if SelfUri, ok := AgentactivityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
