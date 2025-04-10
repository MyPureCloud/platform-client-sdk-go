package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contestscoresagenttrend
type Contestscoresagenttrend struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContestScore - The Contest score
	ContestScore *Contestscoreranked `json:"contestScore,omitempty"`

	// MetricScores - The Contest metric scores
	MetricScores *[]Contestmetricscoreranked `json:"metricScores,omitempty"`

	// Disqualified - Indicates whether an agent is disqualified or not
	Disqualified *bool `json:"disqualified,omitempty"`

	// DateWorkday - Workday of the contest scores leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateWorkday *time.Time `json:"dateWorkday,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contestscoresagenttrend) SetField(field string, fieldValue interface{}) {
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

func (o Contestscoresagenttrend) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateWorkday", }

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
	type Alias Contestscoresagenttrend
	
	DateWorkday := new(string)
	if o.DateWorkday != nil {
		*DateWorkday = timeutil.Strftime(o.DateWorkday, "%Y-%m-%d")
	} else {
		DateWorkday = nil
	}
	
	return json.Marshal(&struct { 
		ContestScore *Contestscoreranked `json:"contestScore,omitempty"`
		
		MetricScores *[]Contestmetricscoreranked `json:"metricScores,omitempty"`
		
		Disqualified *bool `json:"disqualified,omitempty"`
		
		DateWorkday *string `json:"dateWorkday,omitempty"`
		Alias
	}{ 
		ContestScore: o.ContestScore,
		
		MetricScores: o.MetricScores,
		
		Disqualified: o.Disqualified,
		
		DateWorkday: DateWorkday,
		Alias:    (Alias)(o),
	})
}

func (o *Contestscoresagenttrend) UnmarshalJSON(b []byte) error {
	var ContestscoresagenttrendMap map[string]interface{}
	err := json.Unmarshal(b, &ContestscoresagenttrendMap)
	if err != nil {
		return err
	}
	
	if ContestScore, ok := ContestscoresagenttrendMap["contestScore"].(map[string]interface{}); ok {
		ContestScoreString, _ := json.Marshal(ContestScore)
		json.Unmarshal(ContestScoreString, &o.ContestScore)
	}
	
	if MetricScores, ok := ContestscoresagenttrendMap["metricScores"].([]interface{}); ok {
		MetricScoresString, _ := json.Marshal(MetricScores)
		json.Unmarshal(MetricScoresString, &o.MetricScores)
	}
	
	if Disqualified, ok := ContestscoresagenttrendMap["disqualified"].(bool); ok {
		o.Disqualified = &Disqualified
	}
    
	if dateWorkdayString, ok := ContestscoresagenttrendMap["dateWorkday"].(string); ok {
		DateWorkday, _ := time.Parse("2006-01-02", dateWorkdayString)
		o.DateWorkday = &DateWorkday
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contestscoresagenttrend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
