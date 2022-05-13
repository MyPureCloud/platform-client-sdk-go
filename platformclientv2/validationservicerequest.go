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


	// UploadKey - S3 key for the uploaded file
	UploadKey *string `json:"uploadKey,omitempty"`

}

func (o *Validationservicerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validationservicerequest
	
	DateImportEnded := new(string)
	if o.DateImportEnded != nil {
		
		*DateImportEnded = timeutil.Strftime(o.DateImportEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImportEnded = nil
	}
	
	return json.Marshal(&struct { 
		DateImportEnded *string `json:"dateImportEnded,omitempty"`
		
		UploadKey *string `json:"uploadKey,omitempty"`
		*Alias
	}{ 
		DateImportEnded: DateImportEnded,
		
		UploadKey: o.UploadKey,
		Alias:    (*Alias)(o),
	})
}

func (o *Validationservicerequest) UnmarshalJSON(b []byte) error {
	var ValidationservicerequestMap map[string]interface{}
	err := json.Unmarshal(b, &ValidationservicerequestMap)
	if err != nil {
		return err
	}
	
	if dateImportEndedString, ok := ValidationservicerequestMap["dateImportEnded"].(string); ok {
		DateImportEnded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateImportEndedString)
		o.DateImportEnded = &DateImportEnded
	}
	
	if UploadKey, ok := ValidationservicerequestMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Validationservicerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
