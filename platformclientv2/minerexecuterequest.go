package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Minerexecuterequest
type Minerexecuterequest struct { 
	// DateStart - Start date for the date range to mine. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`


	// DateEnd - End date for the date range to mine. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEnd *time.Time `json:"dateEnd,omitempty"`


	// UploadKey - Location of input conversations.
	UploadKey *string `json:"uploadKey,omitempty"`


	// MediaType - Media type for filtering conversations.
	MediaType *string `json:"mediaType,omitempty"`


	// ParticipantType - Type of the participant, either agent, customer or both.
	ParticipantType *string `json:"participantType,omitempty"`


	// QueueIds - List of queue IDs for filtering conversations.
	QueueIds *[]string `json:"queueIds,omitempty"`

}

func (o *Minerexecuterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Minerexecuterequest
	
	DateStart := new(string)
	if o.DateStart != nil {
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%d")
	} else {
		DateEnd = nil
	}
	
	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ParticipantType *string `json:"participantType,omitempty"`
		
		QueueIds *[]string `json:"queueIds,omitempty"`
		*Alias
	}{ 
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		UploadKey: o.UploadKey,
		
		MediaType: o.MediaType,
		
		ParticipantType: o.ParticipantType,
		
		QueueIds: o.QueueIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Minerexecuterequest) UnmarshalJSON(b []byte) error {
	var MinerexecuterequestMap map[string]interface{}
	err := json.Unmarshal(b, &MinerexecuterequestMap)
	if err != nil {
		return err
	}
	
	if dateStartString, ok := MinerexecuterequestMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := MinerexecuterequestMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if UploadKey, ok := MinerexecuterequestMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if MediaType, ok := MinerexecuterequestMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ParticipantType, ok := MinerexecuterequestMap["participantType"].(string); ok {
		o.ParticipantType = &ParticipantType
	}
    
	if QueueIds, ok := MinerexecuterequestMap["queueIds"].([]interface{}); ok {
		QueueIdsString, _ := json.Marshal(QueueIds)
		json.Unmarshal(QueueIdsString, &o.QueueIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Minerexecuterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
