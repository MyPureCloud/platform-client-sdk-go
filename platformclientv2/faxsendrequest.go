package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Faxsendrequest
type Faxsendrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Addresses - A list of outbound fax dialing addresses. E.g. +13175555555 or 3175555555
	Addresses *[]string `json:"addresses,omitempty"`

	// DocumentId - DocumentId of Content Management artifact. If Content Management document is not used for faxing, documentId should be null
	DocumentId *string `json:"documentId,omitempty"`

	// ContentType - The content type that is going to be uploaded. If Content Management document is used for faxing, contentType will be ignored
	ContentType *string `json:"contentType,omitempty"`

	// Workspace - Workspace in which the document should be stored. If Content Management document is used for faxing, workspace will be ignored
	Workspace *Workspace `json:"workspace,omitempty"`

	// CoverSheet - Data for coversheet generation.
	CoverSheet *Coversheet `json:"coverSheet,omitempty"`

	// TimeZoneOffsetMinutes - Time zone offset minutes from GMT
	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Faxsendrequest) SetField(field string, fieldValue interface{}) {
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

func (o Faxsendrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Faxsendrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Addresses *[]string `json:"addresses,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		Workspace *Workspace `json:"workspace,omitempty"`
		
		CoverSheet *Coversheet `json:"coverSheet,omitempty"`
		
		TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Addresses: o.Addresses,
		
		DocumentId: o.DocumentId,
		
		ContentType: o.ContentType,
		
		Workspace: o.Workspace,
		
		CoverSheet: o.CoverSheet,
		
		TimeZoneOffsetMinutes: o.TimeZoneOffsetMinutes,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Faxsendrequest) UnmarshalJSON(b []byte) error {
	var FaxsendrequestMap map[string]interface{}
	err := json.Unmarshal(b, &FaxsendrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FaxsendrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FaxsendrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Addresses, ok := FaxsendrequestMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if DocumentId, ok := FaxsendrequestMap["documentId"].(string); ok {
		o.DocumentId = &DocumentId
	}
    
	if ContentType, ok := FaxsendrequestMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Workspace, ok := FaxsendrequestMap["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if CoverSheet, ok := FaxsendrequestMap["coverSheet"].(map[string]interface{}); ok {
		CoverSheetString, _ := json.Marshal(CoverSheet)
		json.Unmarshal(CoverSheetString, &o.CoverSheet)
	}
	
	if TimeZoneOffsetMinutes, ok := FaxsendrequestMap["timeZoneOffsetMinutes"].(float64); ok {
		TimeZoneOffsetMinutesInt := int(TimeZoneOffsetMinutes)
		o.TimeZoneOffsetMinutes = &TimeZoneOffsetMinutesInt
	}
	
	if SelfUri, ok := FaxsendrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Faxsendrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
