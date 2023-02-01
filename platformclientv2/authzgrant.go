package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Authzgrant
type Authzgrant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SubjectId
	SubjectId *string `json:"subjectId,omitempty"`

	// Division
	Division *Authzdivision `json:"division,omitempty"`

	// Role
	Role *Authzgrantrole `json:"role,omitempty"`

	// GrantMadeAt - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	GrantMadeAt *time.Time `json:"grantMadeAt,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Authzgrant) SetField(field string, fieldValue interface{}) {
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

func (o Authzgrant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "GrantMadeAt", }
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
	type Alias Authzgrant
	
	GrantMadeAt := new(string)
	if o.GrantMadeAt != nil {
		
		*GrantMadeAt = timeutil.Strftime(o.GrantMadeAt, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		GrantMadeAt = nil
	}
	
	return json.Marshal(&struct { 
		SubjectId *string `json:"subjectId,omitempty"`
		
		Division *Authzdivision `json:"division,omitempty"`
		
		Role *Authzgrantrole `json:"role,omitempty"`
		
		GrantMadeAt *string `json:"grantMadeAt,omitempty"`
		Alias
	}{ 
		SubjectId: o.SubjectId,
		
		Division: o.Division,
		
		Role: o.Role,
		
		GrantMadeAt: GrantMadeAt,
		Alias:    (Alias)(o),
	})
}

func (o *Authzgrant) UnmarshalJSON(b []byte) error {
	var AuthzgrantMap map[string]interface{}
	err := json.Unmarshal(b, &AuthzgrantMap)
	if err != nil {
		return err
	}
	
	if SubjectId, ok := AuthzgrantMap["subjectId"].(string); ok {
		o.SubjectId = &SubjectId
	}
    
	if Division, ok := AuthzgrantMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Role, ok := AuthzgrantMap["role"].(map[string]interface{}); ok {
		RoleString, _ := json.Marshal(Role)
		json.Unmarshal(RoleString, &o.Role)
	}
	
	if grantMadeAtString, ok := AuthzgrantMap["grantMadeAt"].(string); ok {
		GrantMadeAt, _ := time.Parse("2006-01-02T15:04:05.999999Z", grantMadeAtString)
		o.GrantMadeAt = &GrantMadeAt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Authzgrant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
