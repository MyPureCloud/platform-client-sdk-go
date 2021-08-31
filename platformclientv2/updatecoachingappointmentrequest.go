package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatecoachingappointmentrequest - Update coaching appointment request
type Updatecoachingappointmentrequest struct { 
	// Name - The name of coaching appointment.
	Name *string `json:"name,omitempty"`


	// Description - The description of coaching appointment.
	Description *string `json:"description,omitempty"`


	// DateStart - The date/time the coaching appointment starts. Times will be rounded down to the minute. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// LengthInMinutes - The duration of coaching appointment in minutes.
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// ConversationIds - IDs of conversations associated with this coaching appointment.
	ConversationIds *[]string `json:"conversationIds,omitempty"`


	// DocumentIds - IDs of documents associated with this coaching appointment.
	DocumentIds *[]string `json:"documentIds,omitempty"`


	// Status - The status of the coaching appointment.
	Status *string `json:"status,omitempty"`

}

func (o *Updatecoachingappointmentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatecoachingappointmentrequest
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		
		DocumentIds *[]string `json:"documentIds,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		DateStart: DateStart,
		
		LengthInMinutes: o.LengthInMinutes,
		
		ConversationIds: o.ConversationIds,
		
		DocumentIds: o.DocumentIds,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatecoachingappointmentrequest) UnmarshalJSON(b []byte) error {
	var UpdatecoachingappointmentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatecoachingappointmentrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdatecoachingappointmentrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := UpdatecoachingappointmentrequestMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if dateStartString, ok := UpdatecoachingappointmentrequestMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if LengthInMinutes, ok := UpdatecoachingappointmentrequestMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if ConversationIds, ok := UpdatecoachingappointmentrequestMap["conversationIds"].([]interface{}); ok {
		ConversationIdsString, _ := json.Marshal(ConversationIds)
		json.Unmarshal(ConversationIdsString, &o.ConversationIds)
	}
	
	if DocumentIds, ok := UpdatecoachingappointmentrequestMap["documentIds"].([]interface{}); ok {
		DocumentIdsString, _ := json.Marshal(DocumentIds)
		json.Unmarshal(DocumentIdsString, &o.DocumentIds)
	}
	
	if Status, ok := UpdatecoachingappointmentrequestMap["status"].(string); ok {
		o.Status = &Status
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatecoachingappointmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
