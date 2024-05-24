package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopiccontact
type Externalcontactsunresolvedcontactchangedtopiccontact struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Division
	Division *Externalcontactsunresolvedcontactchangedtopicdivision `json:"division,omitempty"`

	// ExternalOrganization
	ExternalOrganization *Externalcontactsunresolvedcontactchangedtopicexternalorganization `json:"externalOrganization,omitempty"`

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
	WorkPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"workPhone,omitempty"`

	// CellPhone
	CellPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"cellPhone,omitempty"`

	// HomePhone
	HomePhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"homePhone,omitempty"`

	// OtherPhone
	OtherPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"otherPhone,omitempty"`

	// WorkEmail
	WorkEmail *string `json:"workEmail,omitempty"`

	// PersonalEmail
	PersonalEmail *string `json:"personalEmail,omitempty"`

	// OtherEmail
	OtherEmail *string `json:"otherEmail,omitempty"`

	// Address
	Address *Externalcontactsunresolvedcontactchangedtopiccontactaddress `json:"address,omitempty"`

	// SurveyOptOut
	SurveyOptOut *bool `json:"surveyOptOut,omitempty"`

	// ExternalSystemUrl
	ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`

	// TwitterId
	TwitterId *Externalcontactsunresolvedcontactchangedtopictwitterid `json:"twitterId,omitempty"`

	// LineId
	LineId *Externalcontactsunresolvedcontactchangedtopiclineid `json:"lineId,omitempty"`

	// WhatsAppId
	WhatsAppId *Externalcontactsunresolvedcontactchangedtopicwhatsappid `json:"whatsAppId,omitempty"`

	// FacebookId
	FacebookId *Externalcontactsunresolvedcontactchangedtopicfacebookid `json:"facebookId,omitempty"`

	// InstagramId
	InstagramId *Externalcontactsunresolvedcontactchangedtopicinstagramid `json:"instagramId,omitempty"`

	// ExternalIds
	ExternalIds *[]Externalcontactsunresolvedcontactchangedtopicexternalid `json:"externalIds,omitempty"`

	// Schema
	Schema *Externalcontactsunresolvedcontactchangedtopicdataschema `json:"schema,omitempty"`

	// CustomFields
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`

	// CreateDate
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalcontactsunresolvedcontactchangedtopiccontact) SetField(field string, fieldValue interface{}) {
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

func (o Externalcontactsunresolvedcontactchangedtopiccontact) MarshalJSON() ([]byte, error) {
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
	type Alias Externalcontactsunresolvedcontactchangedtopiccontact
	
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
		
		Division *Externalcontactsunresolvedcontactchangedtopicdivision `json:"division,omitempty"`
		
		ExternalOrganization *Externalcontactsunresolvedcontactchangedtopicexternalorganization `json:"externalOrganization,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		MiddleName *string `json:"middleName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Salutation *string `json:"salutation,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		WorkPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"workPhone,omitempty"`
		
		CellPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"cellPhone,omitempty"`
		
		HomePhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"homePhone,omitempty"`
		
		OtherPhone *Externalcontactsunresolvedcontactchangedtopicphonenumber `json:"otherPhone,omitempty"`
		
		WorkEmail *string `json:"workEmail,omitempty"`
		
		PersonalEmail *string `json:"personalEmail,omitempty"`
		
		OtherEmail *string `json:"otherEmail,omitempty"`
		
		Address *Externalcontactsunresolvedcontactchangedtopiccontactaddress `json:"address,omitempty"`
		
		SurveyOptOut *bool `json:"surveyOptOut,omitempty"`
		
		ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`
		
		TwitterId *Externalcontactsunresolvedcontactchangedtopictwitterid `json:"twitterId,omitempty"`
		
		LineId *Externalcontactsunresolvedcontactchangedtopiclineid `json:"lineId,omitempty"`
		
		WhatsAppId *Externalcontactsunresolvedcontactchangedtopicwhatsappid `json:"whatsAppId,omitempty"`
		
		FacebookId *Externalcontactsunresolvedcontactchangedtopicfacebookid `json:"facebookId,omitempty"`
		
		InstagramId *Externalcontactsunresolvedcontactchangedtopicinstagramid `json:"instagramId,omitempty"`
		
		ExternalIds *[]Externalcontactsunresolvedcontactchangedtopicexternalid `json:"externalIds,omitempty"`
		
		Schema *Externalcontactsunresolvedcontactchangedtopicdataschema `json:"schema,omitempty"`
		
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
		
		ExternalIds: o.ExternalIds,
		
		Schema: o.Schema,
		
		CustomFields: o.CustomFields,
		
		CreateDate: CreateDate,
		
		ModifyDate: ModifyDate,
		Alias:    (Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopiccontact) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopiccontactMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopiccontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if ExternalOrganization, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if VarType, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if FirstName, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if MiddleName, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["middleName"].(string); ok {
		o.MiddleName = &MiddleName
	}
    
	if LastName, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Salutation, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["salutation"].(string); ok {
		o.Salutation = &Salutation
	}
    
	if Title, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if WorkPhone, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["workPhone"].(map[string]interface{}); ok {
		WorkPhoneString, _ := json.Marshal(WorkPhone)
		json.Unmarshal(WorkPhoneString, &o.WorkPhone)
	}
	
	if CellPhone, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["cellPhone"].(map[string]interface{}); ok {
		CellPhoneString, _ := json.Marshal(CellPhone)
		json.Unmarshal(CellPhoneString, &o.CellPhone)
	}
	
	if HomePhone, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["homePhone"].(map[string]interface{}); ok {
		HomePhoneString, _ := json.Marshal(HomePhone)
		json.Unmarshal(HomePhoneString, &o.HomePhone)
	}
	
	if OtherPhone, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["otherPhone"].(map[string]interface{}); ok {
		OtherPhoneString, _ := json.Marshal(OtherPhone)
		json.Unmarshal(OtherPhoneString, &o.OtherPhone)
	}
	
	if WorkEmail, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["workEmail"].(string); ok {
		o.WorkEmail = &WorkEmail
	}
    
	if PersonalEmail, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["personalEmail"].(string); ok {
		o.PersonalEmail = &PersonalEmail
	}
    
	if OtherEmail, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["otherEmail"].(string); ok {
		o.OtherEmail = &OtherEmail
	}
    
	if Address, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if SurveyOptOut, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["surveyOptOut"].(bool); ok {
		o.SurveyOptOut = &SurveyOptOut
	}
    
	if ExternalSystemUrl, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["externalSystemUrl"].(string); ok {
		o.ExternalSystemUrl = &ExternalSystemUrl
	}
    
	if TwitterId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["twitterId"].(map[string]interface{}); ok {
		TwitterIdString, _ := json.Marshal(TwitterId)
		json.Unmarshal(TwitterIdString, &o.TwitterId)
	}
	
	if LineId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["lineId"].(map[string]interface{}); ok {
		LineIdString, _ := json.Marshal(LineId)
		json.Unmarshal(LineIdString, &o.LineId)
	}
	
	if WhatsAppId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["whatsAppId"].(map[string]interface{}); ok {
		WhatsAppIdString, _ := json.Marshal(WhatsAppId)
		json.Unmarshal(WhatsAppIdString, &o.WhatsAppId)
	}
	
	if FacebookId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["facebookId"].(map[string]interface{}); ok {
		FacebookIdString, _ := json.Marshal(FacebookId)
		json.Unmarshal(FacebookIdString, &o.FacebookId)
	}
	
	if InstagramId, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["instagramId"].(map[string]interface{}); ok {
		InstagramIdString, _ := json.Marshal(InstagramId)
		json.Unmarshal(InstagramIdString, &o.InstagramId)
	}
	
	if ExternalIds, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["externalIds"].([]interface{}); ok {
		ExternalIdsString, _ := json.Marshal(ExternalIds)
		json.Unmarshal(ExternalIdsString, &o.ExternalIds)
	}
	
	if Schema, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if CustomFields, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if createDateString, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if modifyDateString, ok := ExternalcontactsunresolvedcontactchangedtopiccontactMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopiccontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
