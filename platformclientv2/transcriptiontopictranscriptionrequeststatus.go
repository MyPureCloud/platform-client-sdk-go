package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptionrequeststatus
type Transcriptiontopictranscriptionrequeststatus struct { 
	// OffsetMs
	OffsetMs *int `json:"offsetMs,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

func (o *Transcriptiontopictranscriptionrequeststatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptiontopictranscriptionrequeststatus
	
	return json.Marshal(&struct { 
		OffsetMs *int `json:"offsetMs,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		OffsetMs: o.OffsetMs,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptiontopictranscriptionrequeststatus) UnmarshalJSON(b []byte) error {
	var TranscriptiontopictranscriptionrequeststatusMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptiontopictranscriptionrequeststatusMap)
	if err != nil {
		return err
	}
	
	if OffsetMs, ok := TranscriptiontopictranscriptionrequeststatusMap["offsetMs"].(float64); ok {
		OffsetMsInt := int(OffsetMs)
		o.OffsetMs = &OffsetMsInt
	}
	
	if Status, ok := TranscriptiontopictranscriptionrequeststatusMap["status"].(string); ok {
		o.Status = &Status
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptionrequeststatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
