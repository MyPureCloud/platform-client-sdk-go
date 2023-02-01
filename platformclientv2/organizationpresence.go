package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Organizationpresence
type Organizationpresence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// LanguageLabels - The label used for the system presence in each specified language
	LanguageLabels *map[string]string `json:"languageLabels,omitempty"`

	// SystemPresence
	SystemPresence *string `json:"systemPresence,omitempty"`

	// Deactivated
	Deactivated *bool `json:"deactivated,omitempty"`

	// Primary
	Primary *bool `json:"primary,omitempty"`

	// CreatedBy
	CreatedBy *User `json:"createdBy,omitempty"`

	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// ModifiedBy
	ModifiedBy *User `json:"modifiedBy,omitempty"`

	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Organizationpresence) SetField(field string, fieldValue interface{}) {
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

func (o Organizationpresence) MarshalJSON() ([]byte, error) {
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
	type Alias Organizationpresence
	
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
		
		LanguageLabels *map[string]string `json:"languageLabels,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		Deactivated *bool `json:"deactivated,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedBy *User `json:"modifiedBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		LanguageLabels: o.LanguageLabels,
		
		SystemPresence: o.SystemPresence,
		
		Deactivated: o.Deactivated,
		
		Primary: o.Primary,
		
		CreatedBy: o.CreatedBy,
		
		CreatedDate: CreatedDate,
		
		ModifiedBy: o.ModifiedBy,
		
		ModifiedDate: ModifiedDate,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Organizationpresence) UnmarshalJSON(b []byte) error {
	var OrganizationpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &OrganizationpresenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OrganizationpresenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := OrganizationpresenceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if LanguageLabels, ok := OrganizationpresenceMap["languageLabels"].(map[string]interface{}); ok {
		LanguageLabelsString, _ := json.Marshal(LanguageLabels)
		json.Unmarshal(LanguageLabelsString, &o.LanguageLabels)
	}
	
	if SystemPresence, ok := OrganizationpresenceMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if Deactivated, ok := OrganizationpresenceMap["deactivated"].(bool); ok {
		o.Deactivated = &Deactivated
	}
    
	if Primary, ok := OrganizationpresenceMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
    
	if CreatedBy, ok := OrganizationpresenceMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if createdDateString, ok := OrganizationpresenceMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if ModifiedBy, ok := OrganizationpresenceMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if modifiedDateString, ok := OrganizationpresenceMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if SelfUri, ok := OrganizationpresenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Organizationpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
