package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontact
type Externalcontact struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Writablestarrabledivision `json:"division,omitempty"`

	// FirstName - The first name of the contact.
	FirstName *string `json:"firstName,omitempty"`

	// MiddleName
	MiddleName *string `json:"middleName,omitempty"`

	// LastName - The last name of the contact.
	LastName *string `json:"lastName,omitempty"`

	// Salutation
	Salutation *string `json:"salutation,omitempty"`

	// Title
	Title *string `json:"title,omitempty"`

	// WorkPhone
	WorkPhone *Phonenumber `json:"workPhone,omitempty"`

	// CellPhone
	CellPhone *Phonenumber `json:"cellPhone,omitempty"`

	// HomePhone
	HomePhone *Phonenumber `json:"homePhone,omitempty"`

	// OtherPhone
	OtherPhone *Phonenumber `json:"otherPhone,omitempty"`

	// WorkEmail
	WorkEmail *string `json:"workEmail,omitempty"`

	// PersonalEmail
	PersonalEmail *string `json:"personalEmail,omitempty"`

	// OtherEmail
	OtherEmail *string `json:"otherEmail,omitempty"`

	// Address
	Address *Contactaddress `json:"address,omitempty"`

	// TwitterId
	TwitterId *Twitterid `json:"twitterId,omitempty"`

	// LineId
	LineId *Lineid `json:"lineId,omitempty"`

	// WhatsAppId
	WhatsAppId *Whatsappid `json:"whatsAppId,omitempty"`

	// FacebookId
	FacebookId *Facebookid `json:"facebookId,omitempty"`

	// InstagramId - User information for an Instagram account
	InstagramId *Instagramid `json:"instagramId,omitempty"`

	// AppleOpaqueIds - User information for an Apple account
	AppleOpaqueIds *[]Appleopaqueid `json:"appleOpaqueIds,omitempty"`

	// ExternalIds - A list of external identifiers that identify this contact in an external system
	ExternalIds *[]Externalid `json:"externalIds,omitempty"`

	// Identifiers - Identifiers claimed by this contact
	Identifiers *[]Contactidentifier `json:"identifiers,omitempty"`

	// ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ExternalOrganization
	ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`

	// SurveyOptOut
	SurveyOptOut *bool `json:"surveyOptOut,omitempty"`

	// ExternalSystemUrl - A string that identifies an external system-of-record resource that may have more detailed information on the contact. It should be a valid URL (including the http/https protocol, port, and path [if any]). The value is automatically trimmed of any leading and trailing whitespace.
	ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`

	// Schema - The schema defining custom fields for this contact
	Schema *Dataschema `json:"schema,omitempty"`

	// CustomFields - Custom fields defined in the schema referenced by schemaId and schemaVersion.
	CustomFields *map[string]interface{} `json:"customFields,omitempty"`

	// ExternalDataSources - Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`

	// VarType - The type of contact
	VarType *string `json:"type,omitempty"`

	// CanonicalContact - The contact at the head of the merge tree. If null, this contact is not a part of any merge.
	CanonicalContact *Contactaddressableentityref `json:"canonicalContact,omitempty"`

	// MergeSet - The set of all contacts that are a part of the merge tree. If null, this contact is not a part of any merge.
	MergeSet *[]Contactaddressableentityref `json:"mergeSet,omitempty"`

	// MergedFrom - The input contacts from the merge operation.
	MergedFrom *[]Contactaddressableentityref `json:"mergedFrom,omitempty"`

	// MergedTo - The output contact from the merge operation.
	MergedTo *Contactaddressableentityref `json:"mergedTo,omitempty"`

	// MergeOperation - (Deprecated: use mergedTo and mergedFrom instead) Information about the merge history of this contact. If null, this contact is not a part of any merge.
	MergeOperation *Mergeoperation `json:"mergeOperation,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalcontact) SetField(field string, fieldValue interface{}) {
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

func (o Externalcontact) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifyDate","CreateDate", }
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
	type Alias Externalcontact
	
	ModifyDate := new(string)
	if o.ModifyDate != nil {
		
		*ModifyDate = timeutil.Strftime(o.ModifyDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifyDate = nil
	}
	
	CreateDate := new(string)
	if o.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(o.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreateDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Division *Writablestarrabledivision `json:"division,omitempty"`
		
		FirstName *string `json:"firstName,omitempty"`
		
		MiddleName *string `json:"middleName,omitempty"`
		
		LastName *string `json:"lastName,omitempty"`
		
		Salutation *string `json:"salutation,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		WorkPhone *Phonenumber `json:"workPhone,omitempty"`
		
		CellPhone *Phonenumber `json:"cellPhone,omitempty"`
		
		HomePhone *Phonenumber `json:"homePhone,omitempty"`
		
		OtherPhone *Phonenumber `json:"otherPhone,omitempty"`
		
		WorkEmail *string `json:"workEmail,omitempty"`
		
		PersonalEmail *string `json:"personalEmail,omitempty"`
		
		OtherEmail *string `json:"otherEmail,omitempty"`
		
		Address *Contactaddress `json:"address,omitempty"`
		
		TwitterId *Twitterid `json:"twitterId,omitempty"`
		
		LineId *Lineid `json:"lineId,omitempty"`
		
		WhatsAppId *Whatsappid `json:"whatsAppId,omitempty"`
		
		FacebookId *Facebookid `json:"facebookId,omitempty"`
		
		InstagramId *Instagramid `json:"instagramId,omitempty"`
		
		AppleOpaqueIds *[]Appleopaqueid `json:"appleOpaqueIds,omitempty"`
		
		ExternalIds *[]Externalid `json:"externalIds,omitempty"`
		
		Identifiers *[]Contactidentifier `json:"identifiers,omitempty"`
		
		ModifyDate *string `json:"modifyDate,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`
		
		SurveyOptOut *bool `json:"surveyOptOut,omitempty"`
		
		ExternalSystemUrl *string `json:"externalSystemUrl,omitempty"`
		
		Schema *Dataschema `json:"schema,omitempty"`
		
		CustomFields *map[string]interface{} `json:"customFields,omitempty"`
		
		ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		CanonicalContact *Contactaddressableentityref `json:"canonicalContact,omitempty"`
		
		MergeSet *[]Contactaddressableentityref `json:"mergeSet,omitempty"`
		
		MergedFrom *[]Contactaddressableentityref `json:"mergedFrom,omitempty"`
		
		MergedTo *Contactaddressableentityref `json:"mergedTo,omitempty"`
		
		MergeOperation *Mergeoperation `json:"mergeOperation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
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
		
		TwitterId: o.TwitterId,
		
		LineId: o.LineId,
		
		WhatsAppId: o.WhatsAppId,
		
		FacebookId: o.FacebookId,
		
		InstagramId: o.InstagramId,
		
		AppleOpaqueIds: o.AppleOpaqueIds,
		
		ExternalIds: o.ExternalIds,
		
		Identifiers: o.Identifiers,
		
		ModifyDate: ModifyDate,
		
		CreateDate: CreateDate,
		
		ExternalOrganization: o.ExternalOrganization,
		
		SurveyOptOut: o.SurveyOptOut,
		
		ExternalSystemUrl: o.ExternalSystemUrl,
		
		Schema: o.Schema,
		
		CustomFields: o.CustomFields,
		
		ExternalDataSources: o.ExternalDataSources,
		
		VarType: o.VarType,
		
		CanonicalContact: o.CanonicalContact,
		
		MergeSet: o.MergeSet,
		
		MergedFrom: o.MergedFrom,
		
		MergedTo: o.MergedTo,
		
		MergeOperation: o.MergeOperation,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Externalcontact) UnmarshalJSON(b []byte) error {
	var ExternalcontactMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := ExternalcontactMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if FirstName, ok := ExternalcontactMap["firstName"].(string); ok {
		o.FirstName = &FirstName
	}
    
	if MiddleName, ok := ExternalcontactMap["middleName"].(string); ok {
		o.MiddleName = &MiddleName
	}
    
	if LastName, ok := ExternalcontactMap["lastName"].(string); ok {
		o.LastName = &LastName
	}
    
	if Salutation, ok := ExternalcontactMap["salutation"].(string); ok {
		o.Salutation = &Salutation
	}
    
	if Title, ok := ExternalcontactMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if WorkPhone, ok := ExternalcontactMap["workPhone"].(map[string]interface{}); ok {
		WorkPhoneString, _ := json.Marshal(WorkPhone)
		json.Unmarshal(WorkPhoneString, &o.WorkPhone)
	}
	
	if CellPhone, ok := ExternalcontactMap["cellPhone"].(map[string]interface{}); ok {
		CellPhoneString, _ := json.Marshal(CellPhone)
		json.Unmarshal(CellPhoneString, &o.CellPhone)
	}
	
	if HomePhone, ok := ExternalcontactMap["homePhone"].(map[string]interface{}); ok {
		HomePhoneString, _ := json.Marshal(HomePhone)
		json.Unmarshal(HomePhoneString, &o.HomePhone)
	}
	
	if OtherPhone, ok := ExternalcontactMap["otherPhone"].(map[string]interface{}); ok {
		OtherPhoneString, _ := json.Marshal(OtherPhone)
		json.Unmarshal(OtherPhoneString, &o.OtherPhone)
	}
	
	if WorkEmail, ok := ExternalcontactMap["workEmail"].(string); ok {
		o.WorkEmail = &WorkEmail
	}
    
	if PersonalEmail, ok := ExternalcontactMap["personalEmail"].(string); ok {
		o.PersonalEmail = &PersonalEmail
	}
    
	if OtherEmail, ok := ExternalcontactMap["otherEmail"].(string); ok {
		o.OtherEmail = &OtherEmail
	}
    
	if Address, ok := ExternalcontactMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	
	if TwitterId, ok := ExternalcontactMap["twitterId"].(map[string]interface{}); ok {
		TwitterIdString, _ := json.Marshal(TwitterId)
		json.Unmarshal(TwitterIdString, &o.TwitterId)
	}
	
	if LineId, ok := ExternalcontactMap["lineId"].(map[string]interface{}); ok {
		LineIdString, _ := json.Marshal(LineId)
		json.Unmarshal(LineIdString, &o.LineId)
	}
	
	if WhatsAppId, ok := ExternalcontactMap["whatsAppId"].(map[string]interface{}); ok {
		WhatsAppIdString, _ := json.Marshal(WhatsAppId)
		json.Unmarshal(WhatsAppIdString, &o.WhatsAppId)
	}
	
	if FacebookId, ok := ExternalcontactMap["facebookId"].(map[string]interface{}); ok {
		FacebookIdString, _ := json.Marshal(FacebookId)
		json.Unmarshal(FacebookIdString, &o.FacebookId)
	}
	
	if InstagramId, ok := ExternalcontactMap["instagramId"].(map[string]interface{}); ok {
		InstagramIdString, _ := json.Marshal(InstagramId)
		json.Unmarshal(InstagramIdString, &o.InstagramId)
	}
	
	if AppleOpaqueIds, ok := ExternalcontactMap["appleOpaqueIds"].([]interface{}); ok {
		AppleOpaqueIdsString, _ := json.Marshal(AppleOpaqueIds)
		json.Unmarshal(AppleOpaqueIdsString, &o.AppleOpaqueIds)
	}
	
	if ExternalIds, ok := ExternalcontactMap["externalIds"].([]interface{}); ok {
		ExternalIdsString, _ := json.Marshal(ExternalIds)
		json.Unmarshal(ExternalIdsString, &o.ExternalIds)
	}
	
	if Identifiers, ok := ExternalcontactMap["identifiers"].([]interface{}); ok {
		IdentifiersString, _ := json.Marshal(Identifiers)
		json.Unmarshal(IdentifiersString, &o.Identifiers)
	}
	
	if modifyDateString, ok := ExternalcontactMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	
	if createDateString, ok := ExternalcontactMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if ExternalOrganization, ok := ExternalcontactMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if SurveyOptOut, ok := ExternalcontactMap["surveyOptOut"].(bool); ok {
		o.SurveyOptOut = &SurveyOptOut
	}
    
	if ExternalSystemUrl, ok := ExternalcontactMap["externalSystemUrl"].(string); ok {
		o.ExternalSystemUrl = &ExternalSystemUrl
	}
    
	if Schema, ok := ExternalcontactMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if CustomFields, ok := ExternalcontactMap["customFields"].(map[string]interface{}); ok {
		CustomFieldsString, _ := json.Marshal(CustomFields)
		json.Unmarshal(CustomFieldsString, &o.CustomFields)
	}
	
	if ExternalDataSources, ok := ExternalcontactMap["externalDataSources"].([]interface{}); ok {
		ExternalDataSourcesString, _ := json.Marshal(ExternalDataSources)
		json.Unmarshal(ExternalDataSourcesString, &o.ExternalDataSources)
	}
	
	if VarType, ok := ExternalcontactMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if CanonicalContact, ok := ExternalcontactMap["canonicalContact"].(map[string]interface{}); ok {
		CanonicalContactString, _ := json.Marshal(CanonicalContact)
		json.Unmarshal(CanonicalContactString, &o.CanonicalContact)
	}
	
	if MergeSet, ok := ExternalcontactMap["mergeSet"].([]interface{}); ok {
		MergeSetString, _ := json.Marshal(MergeSet)
		json.Unmarshal(MergeSetString, &o.MergeSet)
	}
	
	if MergedFrom, ok := ExternalcontactMap["mergedFrom"].([]interface{}); ok {
		MergedFromString, _ := json.Marshal(MergedFrom)
		json.Unmarshal(MergedFromString, &o.MergedFrom)
	}
	
	if MergedTo, ok := ExternalcontactMap["mergedTo"].(map[string]interface{}); ok {
		MergedToString, _ := json.Marshal(MergedTo)
		json.Unmarshal(MergedToString, &o.MergedTo)
	}
	
	if MergeOperation, ok := ExternalcontactMap["mergeOperation"].(map[string]interface{}); ok {
		MergeOperationString, _ := json.Marshal(MergeOperation)
		json.Unmarshal(MergeOperationString, &o.MergeOperation)
	}
	
	if SelfUri, ok := ExternalcontactMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
