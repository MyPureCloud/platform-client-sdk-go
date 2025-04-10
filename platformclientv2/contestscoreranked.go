package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contestscoreranked
type Contestscoreranked struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Score - The Contest score
	Score *float64 `json:"score,omitempty"`

	// TotalPoints - The Contest totalPoints
	TotalPoints *float64 `json:"totalPoints,omitempty"`

	// PercentOfGoal - The Contest percentOfGoal
	PercentOfGoal *float64 `json:"percentOfGoal,omitempty"`

	// Rank - The Contest Score rank
	Rank *int `json:"rank,omitempty"`

	// Tier - The Contest Score tier
	Tier *int `json:"tier,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contestscoreranked) SetField(field string, fieldValue interface{}) {
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

func (o Contestscoreranked) MarshalJSON() ([]byte, error) {
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
	type Alias Contestscoreranked
	
	return json.Marshal(&struct { 
		Score *float64 `json:"score,omitempty"`
		
		TotalPoints *float64 `json:"totalPoints,omitempty"`
		
		PercentOfGoal *float64 `json:"percentOfGoal,omitempty"`
		
		Rank *int `json:"rank,omitempty"`
		
		Tier *int `json:"tier,omitempty"`
		Alias
	}{ 
		Score: o.Score,
		
		TotalPoints: o.TotalPoints,
		
		PercentOfGoal: o.PercentOfGoal,
		
		Rank: o.Rank,
		
		Tier: o.Tier,
		Alias:    (Alias)(o),
	})
}

func (o *Contestscoreranked) UnmarshalJSON(b []byte) error {
	var ContestscorerankedMap map[string]interface{}
	err := json.Unmarshal(b, &ContestscorerankedMap)
	if err != nil {
		return err
	}
	
	if Score, ok := ContestscorerankedMap["score"].(float64); ok {
		o.Score = &Score
	}
    
	if TotalPoints, ok := ContestscorerankedMap["totalPoints"].(float64); ok {
		o.TotalPoints = &TotalPoints
	}
    
	if PercentOfGoal, ok := ContestscorerankedMap["percentOfGoal"].(float64); ok {
		o.PercentOfGoal = &PercentOfGoal
	}
    
	if Rank, ok := ContestscorerankedMap["rank"].(float64); ok {
		RankInt := int(Rank)
		o.Rank = &RankInt
	}
	
	if Tier, ok := ContestscorerankedMap["tier"].(float64); ok {
		TierInt := int(Tier)
		o.Tier = &TierInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contestscoreranked) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
