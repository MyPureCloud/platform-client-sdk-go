package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopiccontact
type Externalcontactscontactchangedtopiccontact struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Division
	Division *Externalcontactscontactchangedtopicdivision `json:"division,omitempty"`

	// ExternalOrganization
	ExternalOrganization *Externalcontactscontactchangedtopicexternalorganization `json:"externalOrganization,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// FirstName
	FirstName *string `json:"firstName,omitempty"`

	// MiddleName
	MiddleName *string `json:"middleName,omitempty"`

	// LastName
	LastName *string `json:"lastName,omitempty"`

	// Salutation
	Salutation *string `json:"salutation,omitempty"`

	// Title
	Title *string `json:"title,omitempty"`

	// WorkPhone
	WorkPhone *Externalcontactscontactchangedtopicphonenumber `json:"workPhone,omitempty"`

	// CellPhone
	CellPhone *Externalcontactscontactchangedtopicphonenumber `json:"cellPhone,omitempty"`

	// HomePhone
	HomePhone *Externalcontactscontactchangedtopicphonenumber `json:"homePhone,omitempty"`

	// OtherPhone
	OtherPhone *Externalcontactscontactchangedtopicphonenumber `json:"otherPhone,omitempty"`

	// WorkEmail
	WorkEmail *string `json:"workEmail,omitempty"`

	// PersonalEmail
	PersonalEmail *string `json:"personalEmail,omitempty"`

	// OtherEmail
	OtherEmail *string `json:"otherEmail,omitempty"`

	// Address
	Address *Externalcontactscontactchangedtopiccontactaddress `json:"address,omitempty"`

	// SurveyOptOut
	SurveyOptOut *bool `json:"surveyOptOut,omitempty"`

	// ExternalSystemUrl
	ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`

	// TwitterId
	TwitterId *Externalcontactscontactchangedtopictwitterid `json:"twitterId,omitempty"`

	// LineId
	LineId *Externalcontactscontactchangedtopiclineid `json:"lineId,omitempty"`

	// WhatsAppId
	WhatsAppId *Externalcontactscontactchangedtopicwhatsappid `json:"whatsAppId,omitempty"`

	// FacebookId
	FacebookId *Externalcontactscontactchangedtopicfacebookid `json:"facebookId,omitempty"`

	// InstagramId
	InstagramId *Externalcontactscontactchangedtopicinstagramid `json:"instagramId,omitempty"`

	// Schema
	Schema *Externalcontactscontactchangedtopicdataschema `json:"schema,omitempty"`

	// CustomFields
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`

	// CreateDate
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalcontactscontactchangedtopiccontact) SetField(field string, fieldValue interface{}) {
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

func (o Externalcontactscontactchangedtopiccontact) MarshalJSON() ([]byte, error) {
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
	type Alias Externalcontactscontactchangedtopiccontact
	
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
		
		Division *Externalcontactscontactchangedtopicdivision `json:"division,omitempty"`
		
		ExternalOrganization *Externalcontactscontactchangedtopicexternalorganization `json:"externalOrganization,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		MiddleName *string `json:"middleName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Salutation *string `json:"salutation,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		WorkPhone *Externalcontactscontactchangedtopicphonenumber `json:"workPhone,omitempty"`
		
		CellPhone *Externalcontactscontactchangedtopicphonenumber `json:"cellPhone,omitempty"`
		
		HomePhone *Externalcontactscontactchangedtopicphonenumber `json:"homePhone,omitempty"`
		
		OtherPhone *Externalcontactscontactchangedtopicphonenumber `json:"otherPhone,omitempty"`
		
		WorkEmail *string `json:"workEmail,omitempty"`
		
		PersonalEmail *string `json:"personalEmail,omitempty"`
		
		OtherEmail *string `json:"otherEmail,omitempty"`
		
		Address *Externalcontactscontactchangedtopiccontactaddress `json:"address,omitempty"`
		
		SurveyOptOut *bool `json:"surveyOptOut,omitempty"`
		
		ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`
		
		TwitterId *Externalcontactscontactchangedtopictwitterid `json:"twitterId,omitempty"`
		
		LineId *Externalcontactscontactchangedtopiclineid `json:"lineId,omitempty"`
		
		WhatsAppId *Externalcontactscontactchangedtopicwhatsappid `json:"whatsAppId,omitempty"`
		
		FacebookId *Externalcontactscontactchangedtopicfacebookid `json:"facebookId,omitempty"`
		
		InstagramId *Externalcontactscontactchangedtopicinstagramid `json:"instagramId,omitempty"`
		
		Schema *Externalcontactscontactchangedtopicdataschema `json:"schema,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		ModifyDate *string `json:"modifyDate,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
		ExternalOrganization: o.ExternalOrganization,
		
		VarType: o.VarType,
		
		FirstName: o.FirstName,
		
		MiddleName: o.MiddleName,
		
		LastName: o.LastName,
		
		Salutation: o.Salutation,
		
		Title: o.Title,
		
		WorkPhone: o.WorkPhone,
		
		CellPhone: o.CellPhone,
		
		HomePhone: o.HomePhone,
		
		OtherPhone: o.OtherPhone,
		
		WorkEmail: o.WorkEmail,
		
		PersonalEmail: o.PersonalEmail,
		
		OtherEmail: o.OtherEmail,
		
		Address: o.Address,
		
		SurveyOptOut: o.SurveyOptOut,
		
		ExternalSystemUrl: o.ExternalSystemUrl,
		
		TwitterId: o.TwitterId,
		
		LineId: o.LineId,
		
		WhatsAppId: o.WhatsAppId,
		
		FacebookId: o.FacebookId,
		
		InstagramId: o.InstagramId,
		
		Schema: o.Schema,
		
		CustomFields: o.CustomFields,
		
		CreateDate: CreateDate,
		
		ModifyDate: ModifyDate,
		Alias:    (Alias)(o),
	})
}

func (o *Externalcontactscontactchangedtopiccontact) UnmarshalJSON(b []byte) error {
	var ExternalcontactscontactchangedtopiccontactMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactscontactchangedtopiccontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactscontactchangedtopiccontactMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ExternalcontactscontactchangedtopiccontactMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if ExternalOrganization, ok := ExternalcontactscontactchangedtopiccontactMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if VarType, ok := ExternalcontactscontactchangedtopiccontactMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if FirstName, ok := ExternalcontactscontactchangedtopiccontactMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if MiddleName, ok := ExternalcontactscontactchangedtopiccontactMap["middleName"].(string); ok {
		o.MiddleName = &MiddleName
	}
    
	if LastName, ok := ExternalcontactscontactchangedtopiccontactMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Salutation, ok := ExternalcontactscontactchangedtopiccontactMap["salutation"].(string); ok {
		o.Salutation = &Salutation
	}
    
	if Title, ok := ExternalcontactscontactchangedtopiccontactMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if WorkPhone, ok := ExternalcontactscontactchangedtopiccontactMap["workPhone"].(map[string]interface{}); ok {
		WorkPhoneString, _ := json.Marshal(WorkPhone)
		json.Unmarshal(WorkPhoneString, &o.WorkPhone)
	}
	
	if CellPhone, ok := ExternalcontactscontactchangedtopiccontactMap["cellPhone"].(map[string]interface{}); ok {
		CellPhoneString, _ := json.Marshal(CellPhone)
		json.Unmarshal(CellPhoneString, &o.CellPhone)
	}
	
	if HomePhone, ok := ExternalcontactscontactchangedtopiccontactMap["homePhone"].(map[string]interface{}); ok {
		HomePhoneString, _ := json.Marshal(HomePhone)
		json.Unmarshal(HomePhoneString, &o.HomePhone)
	}
	
	if OtherPhone, ok := ExternalcontactscontactchangedtopiccontactMap["otherPhone"].(map[string]interface{}); ok {
		OtherPhoneString, _ := json.Marshal(OtherPhone)
		json.Unmarshal(OtherPhoneString, &o.OtherPhone)
	}
	
	if WorkEmail, ok := ExternalcontactscontactchangedtopiccontactMap["workEmail"].(string); ok {
		o.WorkEmail = &WorkEmail
	}
    
	if PersonalEmail, ok := ExternalcontactscontactchangedtopiccontactMap["personalEmail"].(string); ok {
		o.PersonalEmail = &PersonalEmail
	}
    
	if OtherEmail, ok := ExternalcontactscontactchangedtopiccontactMap["otherEmail"].(string); ok {
		o.OtherEmail = &OtherEmail
	}
    
	if Address, ok := ExternalcontactscontactchangedtopiccontactMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if SurveyOptOut, ok := ExternalcontactscontactchangedtopiccontactMap["surveyOptOut"].(bool); ok {
		o.SurveyOptOut = &SurveyOptOut
	}
    
	if ExternalSystemUrl, ok := ExternalcontactscontactchangedtopiccontactMap["externalSystemUrl"].(string); ok {
		o.ExternalSystemUrl = &ExternalSystemUrl
	}
    
	if TwitterId, ok := ExternalcontactscontactchangedtopiccontactMap["twitterId"].(map[string]interface{}); ok {
		TwitterIdString, _ := json.Marshal(TwitterId)
		json.Unmarshal(TwitterIdString, &o.TwitterId)
	}
	
	if LineId, ok := ExternalcontactscontactchangedtopiccontactMap["lineId"].(map[string]interface{}); ok {
		LineIdString, _ := json.Marshal(LineId)
		json.Unmarshal(LineIdString, &o.LineId)
	}
	
	if WhatsAppId, ok := ExternalcontactscontactchangedtopiccontactMap["whatsAppId"].(map[string]interface{}); ok {
		WhatsAppIdString, _ := json.Marshal(WhatsAppId)
		json.Unmarshal(WhatsAppIdString, &o.WhatsAppId)
	}
	
	if FacebookId, ok := ExternalcontactscontactchangedtopiccontactMap["facebookId"].(map[string]interface{}); ok {
		FacebookIdString, _ := json.Marshal(FacebookId)
		json.Unmarshal(FacebookIdString, &o.FacebookId)
	}
	
	if InstagramId, ok := ExternalcontactscontactchangedtopiccontactMap["instagramId"].(map[string]interface{}); ok {
		InstagramIdString, _ := json.Marshal(InstagramId)
		json.Unmarshal(InstagramIdString, &o.InstagramId)
	}
	
	if Schema, ok := ExternalcontactscontactchangedtopiccontactMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if CustomFields, ok := ExternalcontactscontactchangedtopiccontactMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if createDateString, ok := ExternalcontactscontactchangedtopiccontactMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if modifyDateString, ok := ExternalcontactscontactchangedtopiccontactMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopiccontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
