package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactsexport
type Contactsexport struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// DivisionIds - Division IDs of entities
	DivisionIds *[]string `json:"divisionIds,omitempty"`

	// QueryConditions - Query conditions to apply on export
	QueryConditions *Contactsexportqueryconditions `json:"queryConditions,omitempty"`

	// CreatedBy - The user that created this request
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`

	// DateCreated - When the request was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// Status - The status of the request
	Status *string `json:"status,omitempty"`

	// DownloadUrl - The location where the results of the request can be retrieved
	DownloadUrl *string `json:"downloadUrl,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactsexport) SetField(field string, fieldValue interface{}) {
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

func (o Contactsexport) MarshalJSON() ([]byte, error) {
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
	type Alias Contactsexport
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		
		QueryConditions *Contactsexportqueryconditions `json:"queryConditions,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DivisionIds: o.DivisionIds,
		
		QueryConditions: o.QueryConditions,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		Status: o.Status,
		
		DownloadUrl: o.DownloadUrl,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Contactsexport) UnmarshalJSON(b []byte) error {
	var ContactsexportMap map[string]interface{}
	err := json.Unmarshal(b, &ContactsexportMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContactsexportMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DivisionIds, ok := ContactsexportMap["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	
	if QueryConditions, ok := ContactsexportMap["queryConditions"].(map[string]interface{}); ok {
		QueryConditionsString, _ := json.Marshal(QueryConditions)
		json.Unmarshal(QueryConditionsString, &o.QueryConditions)
	}
	
	if CreatedBy, ok := ContactsexportMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := ContactsexportMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if Status, ok := ContactsexportMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if DownloadUrl, ok := ContactsexportMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if SelfUri, ok := ContactsexportMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactsexport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
