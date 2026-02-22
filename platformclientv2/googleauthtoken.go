package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Googleauthtoken
type Googleauthtoken struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - ID of the Google OAuth 2 access token. The token cannot be accessed via Genesys API, only referenced by this property. When the token is not referenced by any integration, it is deleted after 24 hours.
	Id *string `json:"id,omitempty"`

	// ClientId - ID of the Genesys-owned Google API client
	ClientId *string `json:"clientId,omitempty"`

	// Scopes - Google API authorization scopes that have been granted to the Genesys-owned Google API client
	Scopes *[]string `json:"scopes,omitempty"`

	// DateCreated - Date this token was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Googleauthtoken) SetField(field string, fieldValue interface{}) {
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

func (o Googleauthtoken) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
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
	type Alias Googleauthtoken
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ClientId *string `json:"clientId,omitempty"`
		
		Scopes *[]string `json:"scopes,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ClientId: o.ClientId,
		
		Scopes: o.Scopes,
		
		DateCreated: DateCreated,
		
		CreatedBy: o.CreatedBy,
		Alias:    (Alias)(o),
	})
}

func (o *Googleauthtoken) UnmarshalJSON(b []byte) error {
	var GoogleauthtokenMap map[string]interface{}
	err := json.Unmarshal(b, &GoogleauthtokenMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GoogleauthtokenMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ClientId, ok := GoogleauthtokenMap["clientId"].(string); ok {
		o.ClientId = &ClientId
	}
    
	if Scopes, ok := GoogleauthtokenMap["scopes"].([]interface{}); ok {
		ScopesString, _ := json.Marshal(Scopes)
		json.Unmarshal(ScopesString, &o.Scopes)
	}
	
	if dateCreatedString, ok := GoogleauthtokenMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if CreatedBy, ok := GoogleauthtokenMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Googleauthtoken) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
