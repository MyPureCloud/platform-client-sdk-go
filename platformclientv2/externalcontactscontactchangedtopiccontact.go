package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopiccontact
type Externalcontactscontactchangedtopiccontact struct { 
	// Id
	Id *string `json:"id,omitempty"`


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

}

func (o *Externalcontactscontactchangedtopiccontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactscontactchangedtopiccontact
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
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
		*Alias
	}{ 
		Id: o.Id,
		
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
		Alias:    (*Alias)(o),
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopiccontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
