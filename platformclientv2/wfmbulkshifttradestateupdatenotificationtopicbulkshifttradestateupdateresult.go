package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult
type Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// ReviewedBy
	ReviewedBy *Wfmbulkshifttradestateupdatenotificationtopicuserreference `json:"reviewedBy,omitempty"`


	// ReviewedDate
	ReviewedDate *time.Time `json:"reviewedDate,omitempty"`


	// FailureReason
	FailureReason *string `json:"failureReason,omitempty"`


	// Metadata
	Metadata *Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata `json:"metadata,omitempty"`

}

func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult
	
	ReviewedDate := new(string)
	if o.ReviewedDate != nil {
		
		*ReviewedDate = timeutil.Strftime(o.ReviewedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReviewedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ReviewedBy *Wfmbulkshifttradestateupdatenotificationtopicuserreference `json:"reviewedBy,omitempty"`
		
		ReviewedDate *string `json:"reviewedDate,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		Metadata *Wfmbulkshifttradestateupdatenotificationtopicwfmversionedentitymetadata `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		ReviewedBy: o.ReviewedBy,
		
		ReviewedDate: ReviewedDate,
		
		FailureReason: o.FailureReason,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) UnmarshalJSON(b []byte) error {
	var WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["state"].(string); ok {
		o.State = &State
	}
    
	if ReviewedBy, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["reviewedBy"].(map[string]interface{}); ok {
		ReviewedByString, _ := json.Marshal(ReviewedBy)
		json.Unmarshal(ReviewedByString, &o.ReviewedBy)
	}
	
	if reviewedDateString, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["reviewedDate"].(string); ok {
		ReviewedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", reviewedDateString)
		o.ReviewedDate = &ReviewedDate
	}
	
	if FailureReason, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
    
	if Metadata, ok := WfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
