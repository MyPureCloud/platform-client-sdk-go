package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Usersrulesdependent - Users rule dependent response
type Usersrulesdependent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Rule - The rule associated with this dependent
	Rule *Usersrulesdependentrule `json:"rule,omitempty"`

	// TypeId - The type id of the dependent
	TypeId *string `json:"typeId,omitempty"`

	// VarType - The type of the dependent
	VarType *string `json:"type,omitempty"`

	// CreatedDate - The date/time the dependent was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// CreatedBy - The user who created the dependent
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// LastRun - Information on the last run of the dependent
	LastRun *Usersruleslastrunmetadata `json:"lastRun,omitempty"`

	// RecentRunCount - The number of times the rule has run
	RecentRunCount *int `json:"recentRunCount,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Usersrulesdependent) SetField(field string, fieldValue interface{}) {
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

func (o Usersrulesdependent) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Usersrulesdependent
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Rule *Usersrulesdependentrule `json:"rule,omitempty"`
		
		TypeId *string `json:"typeId,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		LastRun *Usersruleslastrunmetadata `json:"lastRun,omitempty"`
		
		RecentRunCount *int `json:"recentRunCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Rule: o.Rule,
		
		TypeId: o.TypeId,
		
		VarType: o.VarType,
		
		CreatedDate: CreatedDate,
		
		CreatedBy: o.CreatedBy,
		
		LastRun: o.LastRun,
		
		RecentRunCount: o.RecentRunCount,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Usersrulesdependent) UnmarshalJSON(b []byte) error {
	var UsersrulesdependentMap map[string]interface{}
	err := json.Unmarshal(b, &UsersrulesdependentMap)
	if err != nil {
		return err
	}
	
	if Rule, ok := UsersrulesdependentMap["rule"].(map[string]interface{}); ok {
		RuleString, _ := json.Marshal(Rule)
		json.Unmarshal(RuleString, &o.Rule)
	}
	
	if TypeId, ok := UsersrulesdependentMap["typeId"].(string); ok {
		o.TypeId = &TypeId
	}
    
	if VarType, ok := UsersrulesdependentMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if createdDateString, ok := UsersrulesdependentMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if CreatedBy, ok := UsersrulesdependentMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if LastRun, ok := UsersrulesdependentMap["lastRun"].(map[string]interface{}); ok {
		LastRunString, _ := json.Marshal(LastRun)
		json.Unmarshal(LastRunString, &o.LastRun)
	}
	
	if RecentRunCount, ok := UsersrulesdependentMap["recentRunCount"].(float64); ok {
		RecentRunCountInt := int(RecentRunCount)
		o.RecentRunCount = &RecentRunCountInt
	}
	
	if SelfUri, ok := UsersrulesdependentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Usersrulesdependent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
