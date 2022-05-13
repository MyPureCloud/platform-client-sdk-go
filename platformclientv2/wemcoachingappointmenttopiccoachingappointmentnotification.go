package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wemcoachingappointmenttopiccoachingappointmentnotification
type Wemcoachingappointmenttopiccoachingappointmentnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateStart
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Facilitator
	Facilitator *Wemcoachingappointmenttopicuserreference `json:"facilitator,omitempty"`


	// Attendees
	Attendees *[]Wemcoachingappointmenttopicuserreference `json:"attendees,omitempty"`


	// CreatedBy
	CreatedBy *Wemcoachingappointmenttopicuserreference `json:"createdBy,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy
	ModifiedBy *Wemcoachingappointmenttopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Conversations
	Conversations *[]Wemcoachingappointmenttopiccoachingappointmentconversation `json:"conversations,omitempty"`


	// Documents
	Documents *[]Wemcoachingappointmenttopiccoachingappointmentdocument `json:"documents,omitempty"`


	// ChangeType
	ChangeType *string `json:"changeType,omitempty"`


	// DateCompleted
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`


	// ExternalLinks
	ExternalLinks *[]Wemcoachingappointmenttopiccoachingappointmentexternallink `json:"externalLinks,omitempty"`

}

func (o *Wemcoachingappointmenttopiccoachingappointmentnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wemcoachingappointmenttopiccoachingappointmentnotification
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Facilitator *Wemcoachingappointmenttopicuserreference `json:"facilitator,omitempty"`
		
		Attendees *[]Wemcoachingappointmenttopicuserreference `json:"attendees,omitempty"`
		
		CreatedBy *Wemcoachingappointmenttopicuserreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Wemcoachingappointmenttopicuserreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Conversations *[]Wemcoachingappointmenttopiccoachingappointmentconversation `json:"conversations,omitempty"`
		
		Documents *[]Wemcoachingappointmenttopiccoachingappointmentdocument `json:"documents,omitempty"`
		
		ChangeType *string `json:"changeType,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		ExternalLinks *[]Wemcoachingappointmenttopiccoachingappointmentexternallink `json:"externalLinks,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Status: o.Status,
		
		Facilitator: o.Facilitator,
		
		Attendees: o.Attendees,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		Conversations: o.Conversations,
		
		Documents: o.Documents,
		
		ChangeType: o.ChangeType,
		
		DateCompleted: DateCompleted,
		
		ExternalLinks: o.ExternalLinks,
		Alias:    (*Alias)(o),
	})
}

func (o *Wemcoachingappointmenttopiccoachingappointmentnotification) UnmarshalJSON(b []byte) error {
	var WemcoachingappointmenttopiccoachingappointmentnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WemcoachingappointmenttopiccoachingappointmentnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateStartString, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Status, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Facilitator, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["facilitator"].(map[string]interface{}); ok {
		FacilitatorString, _ := json.Marshal(Facilitator)
		json.Unmarshal(FacilitatorString, &o.Facilitator)
	}
	
	if Attendees, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["attendees"].([]interface{}); ok {
		AttendeesString, _ := json.Marshal(Attendees)
		json.Unmarshal(AttendeesString, &o.Attendees)
	}
	
	if CreatedBy, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ModifiedBy, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Conversations, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["conversations"].([]interface{}); ok {
		ConversationsString, _ := json.Marshal(Conversations)
		json.Unmarshal(ConversationsString, &o.Conversations)
	}
	
	if Documents, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["documents"].([]interface{}); ok {
		DocumentsString, _ := json.Marshal(Documents)
		json.Unmarshal(DocumentsString, &o.Documents)
	}
	
	if ChangeType, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["changeType"].(string); ok {
		o.ChangeType = &ChangeType
	}
    
	if dateCompletedString, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if ExternalLinks, ok := WemcoachingappointmenttopiccoachingappointmentnotificationMap["externalLinks"].([]interface{}); ok {
		ExternalLinksString, _ := json.Marshal(ExternalLinks)
		json.Unmarshal(ExternalLinksString, &o.ExternalLinks)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wemcoachingappointmenttopiccoachingappointmentnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
