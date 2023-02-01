package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Tokeninfo
type Tokeninfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Organization - The current organization
	Organization *Namedentity `json:"organization,omitempty"`

	// HomeOrganization - The token's home organization
	HomeOrganization *Namedentity `json:"homeOrganization,omitempty"`

	// AuthorizedScope - The list of scopes authorized for the OAuth client
	AuthorizedScope *[]string `json:"authorizedScope,omitempty"`

	// ClonedUser - Only present when a user is a clone of trustee user in the trustor org.
	ClonedUser *Tokeninfocloneduser `json:"clonedUser,omitempty"`

	// OAuthClient
	OAuthClient *Orgoauthclient `json:"OAuthClient,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Tokeninfo) SetField(field string, fieldValue interface{}) {
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

func (o Tokeninfo) MarshalJSON() ([]byte, error) {
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
	type Alias Tokeninfo
	
	return json.Marshal(&struct { 
		Organization *Namedentity `json:"organization,omitempty"`
		
		HomeOrganization *Namedentity `json:"homeOrganization,omitempty"`
		
		AuthorizedScope *[]string `json:"authorizedScope,omitempty"`
		
		ClonedUser *Tokeninfocloneduser `json:"clonedUser,omitempty"`
		
		OAuthClient *Orgoauthclient `json:"OAuthClient,omitempty"`
		Alias
	}{ 
		Organization: o.Organization,
		
		HomeOrganization: o.HomeOrganization,
		
		AuthorizedScope: o.AuthorizedScope,
		
		ClonedUser: o.ClonedUser,
		
		OAuthClient: o.OAuthClient,
		Alias:    (Alias)(o),
	})
}

func (o *Tokeninfo) UnmarshalJSON(b []byte) error {
	var TokeninfoMap map[string]interface{}
	err := json.Unmarshal(b, &TokeninfoMap)
	if err != nil {
		return err
	}
	
	if Organization, ok := TokeninfoMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	
	if HomeOrganization, ok := TokeninfoMap["homeOrganization"].(map[string]interface{}); ok {
		HomeOrganizationString, _ := json.Marshal(HomeOrganization)
		json.Unmarshal(HomeOrganizationString, &o.HomeOrganization)
	}
	
	if AuthorizedScope, ok := TokeninfoMap["authorizedScope"].([]interface{}); ok {
		AuthorizedScopeString, _ := json.Marshal(AuthorizedScope)
		json.Unmarshal(AuthorizedScopeString, &o.AuthorizedScope)
	}
	
	if ClonedUser, ok := TokeninfoMap["clonedUser"].(map[string]interface{}); ok {
		ClonedUserString, _ := json.Marshal(ClonedUser)
		json.Unmarshal(ClonedUserString, &o.ClonedUser)
	}
	
	if OAuthClient, ok := TokeninfoMap["OAuthClient"].(map[string]interface{}); ok {
		OAuthClientString, _ := json.Marshal(OAuthClient)
		json.Unmarshal(OAuthClientString, &o.OAuthClient)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Tokeninfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
