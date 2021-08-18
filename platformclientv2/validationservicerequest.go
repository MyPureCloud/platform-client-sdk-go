package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validationservicerequest
type Validationservicerequest struct { 
	// DateImportEnded - The last day of the data you are importing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateImportEnded *time.Time `json:"dateImportEnded,omitempty"`


	// FileUrl - Path to the file in the storage including the file name
	FileUrl *string `json:"fileUrl,omitempty"`

}

func (u *Validationservicerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validationservicerequest

	
	DateImportEnded := new(string)
	if u.DateImportEnded != nil {
		
		*DateImportEnded = timeutil.Strftime(u.DateImportEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImportEnded = nil
	}
	

	return json.Marshal(&struct { 
		DateImportEnded *string `json:"dateImportEnded,omitempty"`
		
		FileUrl *string `json:"fileUrl,omitempty"`
		*Alias
	}{ 
		DateImportEnded: DateImportEnded,
		
		FileUrl: u.FileUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Validationservicerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
