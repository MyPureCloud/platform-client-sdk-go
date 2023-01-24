package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustee
type Trustee struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Organization Id for this trust.
	Id *string `json:"id,omitempty"`

	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`

	// UsesDefaultRole - Denotes if trustee uses admin role by default.
	UsesDefaultRole *bool `json:"usesDefaultRole,omitempty"`

	// DateCreated - Date Trust was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpired *time.Time `json:"dateExpired,omitempty"`

	// CreatedBy - User that created trust.
	CreatedBy *Orguser `json:"createdBy,omitempty"`

	// Organization - Organization associated with this trust.
	Organization *Organization `json:"organization,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Trustee) SetField(field string, fieldValue interface{}) {
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

func (o Trustee) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateExpired", }
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
	type Alias Trustee
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateExpired := new(string)
	if o.DateExpired != nil {
		
		*DateExpired = timeutil.Strftime(o.DateExpired, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpired = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		UsesDefaultRole *bool `json:"usesDefaultRole,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateExpired *string `json:"dateExpired,omitempty"`
		
		CreatedBy *Orguser `json:"createdBy,omitempty"`
		
		Organization *Organization `json:"organization,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Enabled: o.Enabled,
		
		UsesDefaultRole: o.UsesDefaultRole,
		
		DateCreated: DateCreated,
		
		DateExpired: DateExpired,
		
		CreatedBy: o.CreatedBy,
		
		Organization: o.Organization,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Trustee) UnmarshalJSON(b []byte) error {
	var TrusteeMap map[string]interface{}
	err := json.Unmarshal(b, &TrusteeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TrusteeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Enabled, ok := TrusteeMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if UsesDefaultRole, ok := TrusteeMap["usesDefaultRole"].(bool); ok {
		o.UsesDefaultRole = &UsesDefaultRole
	}
    
	if dateCreatedString, ok := TrusteeMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateExpiredString, ok := TrusteeMap["dateExpired"].(string); ok {
		DateExpired, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiredString)
		o.DateExpired = &DateExpired
	}
	
	if CreatedBy, ok := TrusteeMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if Organization, ok := TrusteeMap["organization"].(map[string]interface{}); ok {
		OrganizationString, _ := json.Marshal(Organization)
		json.Unmarshal(OrganizationString, &o.Organization)
	}
	
	if SelfUri, ok := TrusteeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Trustee) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
