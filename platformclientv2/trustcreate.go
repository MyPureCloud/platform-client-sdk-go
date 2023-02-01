package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustcreate
type Trustcreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PairingId - The pairing Id created by the trustee. This is required to prove that the trustee agrees to the relationship.  Not required when creating a default pairing with Customer Care.
	PairingId *string `json:"pairingId,omitempty"`

	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`

	// Users - The list of users and their roles to which access will be granted. The users are from the trustee and the roles are from the trustor. If no users are specified, at least one group is required.
	Users *[]Trustmembercreate `json:"users,omitempty"`

	// Groups - The list of groups and their roles to which access will be granted. The groups are from the trustee and the roles are from the trustor. If no groups are specified, at least one user is required.
	Groups *[]Trustmembercreate `json:"groups,omitempty"`

	// DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpired *time.Time `json:"dateExpired,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Trustcreate) SetField(field string, fieldValue interface{}) {
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

func (o Trustcreate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateExpired", }
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
	type Alias Trustcreate
	
	DateExpired := new(string)
	if o.DateExpired != nil {
		
		*DateExpired = timeutil.Strftime(o.DateExpired, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpired = nil
	}
	
	return json.Marshal(&struct { 
		PairingId *string `json:"pairingId,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Users *[]Trustmembercreate `json:"users,omitempty"`
		
		Groups *[]Trustmembercreate `json:"groups,omitempty"`
		
		DateExpired *string `json:"dateExpired,omitempty"`
		Alias
	}{ 
		PairingId: o.PairingId,
		
		Enabled: o.Enabled,
		
		Users: o.Users,
		
		Groups: o.Groups,
		
		DateExpired: DateExpired,
		Alias:    (Alias)(o),
	})
}

func (o *Trustcreate) UnmarshalJSON(b []byte) error {
	var TrustcreateMap map[string]interface{}
	err := json.Unmarshal(b, &TrustcreateMap)
	if err != nil {
		return err
	}
	
	if PairingId, ok := TrustcreateMap["pairingId"].(string); ok {
		o.PairingId = &PairingId
	}
    
	if Enabled, ok := TrustcreateMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Users, ok := TrustcreateMap["users"].([]interface{}); ok {
		UsersString, _ := json.Marshal(Users)
		json.Unmarshal(UsersString, &o.Users)
	}
	
	if Groups, ok := TrustcreateMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}
	
	if dateExpiredString, ok := TrustcreateMap["dateExpired"].(string); ok {
		DateExpired, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiredString)
		o.DateExpired = &DateExpired
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
