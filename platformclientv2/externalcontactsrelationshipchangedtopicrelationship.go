package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsrelationshipchangedtopicrelationship
type Externalcontactsrelationshipchangedtopicrelationship struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Division
	Division *Externalcontactsrelationshipchangedtopicdivision `json:"division,omitempty"`

	// User
	User *Externalcontactsrelationshipchangedtopicuser `json:"user,omitempty"`

	// ExternalOrganization
	ExternalOrganization *Externalcontactsrelationshipchangedtopicexternalorganization `json:"externalOrganization,omitempty"`

	// Relationship
	Relationship *string `json:"relationship,omitempty"`

	// CreateDate
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalcontactsrelationshipchangedtopicrelationship) SetField(field string, fieldValue interface{}) {
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

func (o Externalcontactsrelationshipchangedtopicrelationship) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreateDate","ModifyDate", }
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
	type Alias Externalcontactsrelationshipchangedtopicrelationship
	
	CreateDate := new(string)
	if o.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(o.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreateDate = nil
	}
	
	ModifyDate := new(string)
	if o.ModifyDate != nil {
		
		*ModifyDate = timeutil.Strftime(o.ModifyDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifyDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Division *Externalcontactsrelationshipchangedtopicdivision `json:"division,omitempty"`
		
		User *Externalcontactsrelationshipchangedtopicuser `json:"user,omitempty"`
		
		ExternalOrganization *Externalcontactsrelationshipchangedtopicexternalorganization `json:"externalOrganization,omitempty"`
		
		Relationship *string `json:"relationship,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		ModifyDate *string `json:"modifyDate,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
		User: o.User,
		
		ExternalOrganization: o.ExternalOrganization,
		
		Relationship: o.Relationship,
		
		CreateDate: CreateDate,
		
		ModifyDate: ModifyDate,
		Alias:    (Alias)(o),
	})
}

func (o *Externalcontactsrelationshipchangedtopicrelationship) UnmarshalJSON(b []byte) error {
	var ExternalcontactsrelationshipchangedtopicrelationshipMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsrelationshipchangedtopicrelationshipMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactsrelationshipchangedtopicrelationshipMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ExternalcontactsrelationshipchangedtopicrelationshipMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if User, ok := ExternalcontactsrelationshipchangedtopicrelationshipMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if ExternalOrganization, ok := ExternalcontactsrelationshipchangedtopicrelationshipMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Relationship, ok := ExternalcontactsrelationshipchangedtopicrelationshipMap["relationship"].(string); ok {
		o.Relationship = &Relationship
	}
    
	if createDateString, ok := ExternalcontactsrelationshipchangedtopicrelationshipMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if modifyDateString, ok := ExternalcontactsrelationshipchangedtopicrelationshipMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsrelationshipchangedtopicrelationship) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
