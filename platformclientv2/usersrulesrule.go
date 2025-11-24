package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Usersrulesrule - Users rule response
type Usersrulesrule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description - The description of the rule
	Description *string `json:"description,omitempty"`

	// VarType - The type of the rule
	VarType *string `json:"type,omitempty"`

	// Criteria - The criteria of the rule
	Criteria *[]Usersrulescriteria `json:"criteria,omitempty"`

	// CreatedDate - The date/time the rule was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// CreatedBy - The user who created the rule
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// ModifiedDate - The date/time the rule was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// ModifiedBy - The last user to modify the rule
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// LastRun - Information on the last run of the rule
	LastRun *Usersruleslastrunmetadata `json:"lastRun,omitempty"`

	// RecentRunCount - The number of times the rule has run
	RecentRunCount *int `json:"recentRunCount,omitempty"`

	// DependentCount - The number of dependents this rule has
	DependentCount *int `json:"dependentCount,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Usersrulesrule) SetField(field string, fieldValue interface{}) {
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

func (o Usersrulesrule) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate","ModifiedDate", }
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
	type Alias Usersrulesrule
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Criteria *[]Usersrulescriteria `json:"criteria,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		LastRun *Usersruleslastrunmetadata `json:"lastRun,omitempty"`
		
		RecentRunCount *int `json:"recentRunCount,omitempty"`
		
		DependentCount *int `json:"dependentCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		Criteria: o.Criteria,
		
		CreatedDate: CreatedDate,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedDate: ModifiedDate,
		
		ModifiedBy: o.ModifiedBy,
		
		LastRun: o.LastRun,
		
		RecentRunCount: o.RecentRunCount,
		
		DependentCount: o.DependentCount,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Usersrulesrule) UnmarshalJSON(b []byte) error {
	var UsersrulesruleMap map[string]interface{}
	err := json.Unmarshal(b, &UsersrulesruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UsersrulesruleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UsersrulesruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := UsersrulesruleMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if VarType, ok := UsersrulesruleMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Criteria, ok := UsersrulesruleMap["criteria"].([]interface{}); ok {
		CriteriaString, _ := json.Marshal(Criteria)
		json.Unmarshal(CriteriaString, &o.Criteria)
	}
	
	if createdDateString, ok := UsersrulesruleMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if CreatedBy, ok := UsersrulesruleMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if modifiedDateString, ok := UsersrulesruleMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if ModifiedBy, ok := UsersrulesruleMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if LastRun, ok := UsersrulesruleMap["lastRun"].(map[string]interface{}); ok {
		LastRunString, _ := json.Marshal(LastRun)
		json.Unmarshal(LastRunString, &o.LastRun)
	}
	
	if RecentRunCount, ok := UsersrulesruleMap["recentRunCount"].(float64); ok {
		RecentRunCountInt := int(RecentRunCount)
		o.RecentRunCount = &RecentRunCountInt
	}
	
	if DependentCount, ok := UsersrulesruleMap["dependentCount"].(float64); ok {
		DependentCountInt := int(DependentCount)
		o.DependentCount = &DependentCountInt
	}
	
	if SelfUri, ok := UsersrulesruleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Usersrulesrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
