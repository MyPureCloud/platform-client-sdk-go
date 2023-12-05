package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalorganizationtrustorlink
type Externalorganizationtrustorlink struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// ExternalOrganizationId - The id of a PureCloud External Organization entity in the External Contacts system that will be used to represent the trustor org
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`

	// TrustorOrgId - The id of a PureCloud organization that has granted trust to this PureCloud organization
	TrustorOrgId *string `json:"trustorOrgId,omitempty"`

	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// ExternalOrganizationUri - The URI for the External Organization that is linked to the trustor org
	ExternalOrganizationUri *string `json:"externalOrganizationUri,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalorganizationtrustorlink) SetField(field string, fieldValue interface{}) {
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

func (o Externalorganizationtrustorlink) MarshalJSON() ([]byte, error) {
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
	type Alias Externalorganizationtrustorlink
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		TrustorOrgId *string `json:"trustorOrgId,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ExternalOrganizationUri *string `json:"externalOrganizationUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		TrustorOrgId: o.TrustorOrgId,
		
		DateCreated: DateCreated,
		
		ExternalOrganizationUri: o.ExternalOrganizationUri,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Externalorganizationtrustorlink) UnmarshalJSON(b []byte) error {
	var ExternalorganizationtrustorlinkMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalorganizationtrustorlinkMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalorganizationtrustorlinkMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ExternalOrganizationId, ok := ExternalorganizationtrustorlinkMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if TrustorOrgId, ok := ExternalorganizationtrustorlinkMap["trustorOrgId"].(string); ok {
		o.TrustorOrgId = &TrustorOrgId
	}
    
	if dateCreatedString, ok := ExternalorganizationtrustorlinkMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ExternalOrganizationUri, ok := ExternalorganizationtrustorlinkMap["externalOrganizationUri"].(string); ok {
		o.ExternalOrganizationUri = &ExternalOrganizationUri
	}
    
	if SelfUri, ok := ExternalorganizationtrustorlinkMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalorganizationtrustorlink) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
