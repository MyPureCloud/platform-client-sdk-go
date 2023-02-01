package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmapestimateoutcomecriteria
type Actionmapestimateoutcomecriteria struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OutcomeId - ID of outcome.
	OutcomeId *string `json:"outcomeId,omitempty"`

	// MaxProbability - Probability value for the selected outcome at or above which the action map will trigger.
	MaxProbability *float32 `json:"maxProbability,omitempty"`

	// Probability - Additional probability condition, where if set, the action map will trigger if the current outcome probability is lower or equal to the value.
	Probability *float32 `json:"probability,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Actionmapestimateoutcomecriteria) SetField(field string, fieldValue interface{}) {
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

func (o Actionmapestimateoutcomecriteria) MarshalJSON() ([]byte, error) {
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
	type Alias Actionmapestimateoutcomecriteria
	
	return json.Marshal(&struct { 
		OutcomeId *string `json:"outcomeId,omitempty"`
		
		MaxProbability *float32 `json:"maxProbability,omitempty"`
		
		Probability *float32 `json:"probability,omitempty"`
		Alias
	}{ 
		OutcomeId: o.OutcomeId,
		
		MaxProbability: o.MaxProbability,
		
		Probability: o.Probability,
		Alias:    (Alias)(o),
	})
}

func (o *Actionmapestimateoutcomecriteria) UnmarshalJSON(b []byte) error {
	var ActionmapestimateoutcomecriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &ActionmapestimateoutcomecriteriaMap)
	if err != nil {
		return err
	}
	
	if OutcomeId, ok := ActionmapestimateoutcomecriteriaMap["outcomeId"].(string); ok {
		o.OutcomeId = &OutcomeId
	}
    
	if MaxProbability, ok := ActionmapestimateoutcomecriteriaMap["maxProbability"].(float64); ok {
		MaxProbabilityFloat32 := float32(MaxProbability)
		o.MaxProbability = &MaxProbabilityFloat32
	}
	
	if Probability, ok := ActionmapestimateoutcomecriteriaMap["probability"].(float64); ok {
		ProbabilityFloat32 := float32(Probability)
		o.Probability = &ProbabilityFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionmapestimateoutcomecriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
