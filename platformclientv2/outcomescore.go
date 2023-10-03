package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outcomescore
type Outcomescore struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Outcome - The outcome that the score was calculated for.
	Outcome *Addressableentityref `json:"outcome,omitempty"`

	// SessionMaxProbability - Represents the max probability reached in the session.
	SessionMaxProbability *float32 `json:"sessionMaxProbability,omitempty"`

	// Probability - Represents the likelihood of a customer reaching or achieving a given outcome.
	Probability *float32 `json:"probability,omitempty"`

	// Percentile - (Deprecated: use the 'quantile' field instead) Represents the predicted probability's percentile score when compared with all other generated probabilities for a given outcome.
	Percentile *int `json:"percentile,omitempty"`

	// SessionMaxPercentile - (Deprecated: use the 'quantile' field instead) Represents the maximum likelihood percentile score reached for a given outcome by the current session.
	SessionMaxPercentile *int `json:"sessionMaxPercentile,omitempty"`

	// Quantile - Represents the quantity of sessions that have a maximum probability less than the predicted probability.
	Quantile *float32 `json:"quantile,omitempty"`

	// SessionMaxQuantile - Represents the quantity of sessions that have a maximum probability less than the predicted session max probability.
	SessionMaxQuantile *float32 `json:"sessionMaxQuantile,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outcomescore) SetField(field string, fieldValue interface{}) {
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

func (o Outcomescore) MarshalJSON() ([]byte, error) {
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
	type Alias Outcomescore
	
	return json.Marshal(&struct { 
		Outcome *Addressableentityref `json:"outcome,omitempty"`
		
		SessionMaxProbability *float32 `json:"sessionMaxProbability,omitempty"`
		
		Probability *float32 `json:"probability,omitempty"`
		
		Percentile *int `json:"percentile,omitempty"`
		
		SessionMaxPercentile *int `json:"sessionMaxPercentile,omitempty"`
		
		Quantile *float32 `json:"quantile,omitempty"`
		
		SessionMaxQuantile *float32 `json:"sessionMaxQuantile,omitempty"`
		Alias
	}{ 
		Outcome: o.Outcome,
		
		SessionMaxProbability: o.SessionMaxProbability,
		
		Probability: o.Probability,
		
		Percentile: o.Percentile,
		
		SessionMaxPercentile: o.SessionMaxPercentile,
		
		Quantile: o.Quantile,
		
		SessionMaxQuantile: o.SessionMaxQuantile,
		Alias:    (Alias)(o),
	})
}

func (o *Outcomescore) UnmarshalJSON(b []byte) error {
	var OutcomescoreMap map[string]interface{}
	err := json.Unmarshal(b, &OutcomescoreMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := OutcomescoreMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if SessionMaxProbability, ok := OutcomescoreMap["sessionMaxProbability"].(float64); ok {
		SessionMaxProbabilityFloat32 := float32(SessionMaxProbability)
		o.SessionMaxProbability = &SessionMaxProbabilityFloat32
	}
	
	if Probability, ok := OutcomescoreMap["probability"].(float64); ok {
		ProbabilityFloat32 := float32(Probability)
		o.Probability = &ProbabilityFloat32
	}
	
	if Percentile, ok := OutcomescoreMap["percentile"].(float64); ok {
		PercentileInt := int(Percentile)
		o.Percentile = &PercentileInt
	}
	
	if SessionMaxPercentile, ok := OutcomescoreMap["sessionMaxPercentile"].(float64); ok {
		SessionMaxPercentileInt := int(SessionMaxPercentile)
		o.SessionMaxPercentile = &SessionMaxPercentileInt
	}
	
	if Quantile, ok := OutcomescoreMap["quantile"].(float64); ok {
		QuantileFloat32 := float32(Quantile)
		o.Quantile = &QuantileFloat32
	}
	
	if SessionMaxQuantile, ok := OutcomescoreMap["sessionMaxQuantile"].(float64); ok {
		SessionMaxQuantileFloat32 := float32(SessionMaxQuantile)
		o.SessionMaxQuantile = &SessionMaxQuantileFloat32
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outcomescore) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
