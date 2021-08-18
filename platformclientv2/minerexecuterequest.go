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


	// QueueIds - List of queue IDs for filtering conversations.
	QueueIds *[]string `json:"queueIds,omitempty"`

}

func (u *Minerexecuterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Minerexecuterequest

	
	DateStart := new(string)
	if u.DateStart != nil {
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if u.DateEnd != nil {
		*DateEnd = timeutil.Strftime(u.DateEnd, "%Y-%m-%d")
	} else {
		DateEnd = nil
	}
	

	return json.Marshal(&struct { 
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		QueueIds *[]string `json:"queueIds,omitempty"`
		*Alias
	}{ 
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		UploadKey: u.UploadKey,
		
		MediaType: u.MediaType,
		
		QueueIds: u.QueueIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Minerexecuterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
