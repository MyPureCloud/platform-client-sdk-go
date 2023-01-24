package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthclientlisting
type Oauthclientlisting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the OAuth client.
	Name *string `json:"name,omitempty"`

	// AccessTokenValiditySeconds - The number of seconds, between 5mins and 48hrs, until tokens created with this client expire. If this field is omitted, a default of 24 hours will be applied.
	AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// RegisteredRedirectUri - List of allowed callbacks for this client. For example: https://myap.example.com/auth/callback
	RegisteredRedirectUri *[]string `json:"registeredRedirectUri,omitempty"`

	// Secret - System created secret assigned to this client. Secrets are required for code authorization and client credential grants.
	Secret *string `json:"secret,omitempty"`

	// RoleIds - Deprecated. Use roleDivisions instead.
	RoleIds *[]string `json:"roleIds,omitempty"`

	// DateCreated - Date this client was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date this client was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// CreatedBy - User that created this client
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`

	// ModifiedBy - User that last modified this client
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`

	// Scope - The scope requested by this client. Scopes only apply to clients not using the client_credential grant
	Scope *[]string `json:"scope,omitempty"`

	// RoleDivisions - Set of roles and their corresponding divisions associated with this client. Roles and divisions only apply to clients using the client_credential grant
	RoleDivisions *[]Roledivision `json:"roleDivisions,omitempty"`

	// State - The state of the OAuth client. Active: The OAuth client can be used to create access tokens. This is the default state. Disabled: Access tokens created by the client are invalid and new ones cannot be created. Inactive: Access tokens cannot be created with this OAuth client and it will be deleted.
	State *string `json:"state,omitempty"`

	// DateToDelete - The time at which this client will be deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateToDelete *time.Time `json:"dateToDelete,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Oauthclientlisting) SetField(field string, fieldValue interface{}) {
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

func (o Oauthclientlisting) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateToDelete", }
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
	type Alias Oauthclientlisting
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateToDelete := new(string)
	if o.DateToDelete != nil {
		
		*DateToDelete = timeutil.Strftime(o.DateToDelete, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateToDelete = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AccessTokenValiditySeconds *int `json:"accessTokenValiditySeconds,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		RegisteredRedirectUri *[]string `json:"registeredRedirectUri,omitempty"`
		
		Secret *string `json:"secret,omitempty"`
		
		RoleIds *[]string `json:"roleIds,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Scope *[]string `json:"scope,omitempty"`
		
		RoleDivisions *[]Roledivision `json:"roleDivisions,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateToDelete *string `json:"dateToDelete,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		AccessTokenValiditySeconds: o.AccessTokenValiditySeconds,
		
		Description: o.Description,
		
		RegisteredRedirectUri: o.RegisteredRedirectUri,
		
		Secret: o.Secret,
		
		RoleIds: o.RoleIds,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		Scope: o.Scope,
		
		RoleDivisions: o.RoleDivisions,
		
		State: o.State,
		
		DateToDelete: DateToDelete,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Oauthclientlisting) UnmarshalJSON(b []byte) error {
	var OauthclientlistingMap map[string]interface{}
	err := json.Unmarshal(b, &OauthclientlistingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OauthclientlistingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OauthclientlistingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if AccessTokenValiditySeconds, ok := OauthclientlistingMap["accessTokenValiditySeconds"].(float64); ok {
		AccessTokenValiditySecondsInt := int(AccessTokenValiditySeconds)
		o.AccessTokenValiditySeconds = &AccessTokenValiditySecondsInt
	}
	
	if Description, ok := OauthclientlistingMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if RegisteredRedirectUri, ok := OauthclientlistingMap["registeredRedirectUri"].([]interface{}); ok {
		RegisteredRedirectUriString, _ := json.Marshal(RegisteredRedirectUri)
		json.Unmarshal(RegisteredRedirectUriString, &o.RegisteredRedirectUri)
	}
	
	if Secret, ok := OauthclientlistingMap["secret"].(string); ok {
		o.Secret = &Secret
	}
    
	if RoleIds, ok := OauthclientlistingMap["roleIds"].([]interface{}); ok {
		RoleIdsString, _ := json.Marshal(RoleIds)
		json.Unmarshal(RoleIdsString, &o.RoleIds)
	}
	
	if dateCreatedString, ok := OauthclientlistingMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := OauthclientlistingMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := OauthclientlistingMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := OauthclientlistingMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Scope, ok := OauthclientlistingMap["scope"].([]interface{}); ok {
		ScopeString, _ := json.Marshal(Scope)
		json.Unmarshal(ScopeString, &o.Scope)
	}
	
	if RoleDivisions, ok := OauthclientlistingMap["roleDivisions"].([]interface{}); ok {
		RoleDivisionsString, _ := json.Marshal(RoleDivisions)
		json.Unmarshal(RoleDivisionsString, &o.RoleDivisions)
	}
	
	if State, ok := OauthclientlistingMap["state"].(string); ok {
		o.State = &State
	}
    
	if dateToDeleteString, ok := OauthclientlistingMap["dateToDelete"].(string); ok {
		DateToDelete, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateToDeleteString)
		o.DateToDelete = &DateToDelete
	}
	
	if SelfUri, ok := OauthclientlistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Oauthclientlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
