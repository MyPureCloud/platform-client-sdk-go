package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Leaderboarditem
type Leaderboarditem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// User - The user object for this leaderboard rank
	User *Userreference `json:"user,omitempty"`

	// Rank - The rank of the user
	Rank *int `json:"rank,omitempty"`

	// Points - The points collected by the user
	Points *int `json:"points,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Leaderboarditem) SetField(field string, fieldValue interface{}) {
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

func (o Leaderboarditem) MarshalJSON() ([]byte, error) {
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
	type Alias Leaderboarditem
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		Rank *int `json:"rank,omitempty"`
		
		Points *int `json:"points,omitempty"`
		Alias
	}{ 
		User: o.User,
		
		Rank: o.Rank,
		
		Points: o.Points,
		Alias:    (Alias)(o),
	})
}

func (o *Leaderboarditem) UnmarshalJSON(b []byte) error {
	var LeaderboarditemMap map[string]interface{}
	err := json.Unmarshal(b, &LeaderboarditemMap)
	if err != nil {
		return err
	}
	
	if User, ok := LeaderboarditemMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Rank, ok := LeaderboarditemMap["rank"].(float64); ok {
		RankInt := int(Rank)
		o.Rank = &RankInt
	}
	
	if Points, ok := LeaderboarditemMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Leaderboarditem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
