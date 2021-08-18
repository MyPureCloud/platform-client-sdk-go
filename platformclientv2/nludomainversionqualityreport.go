package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludomainversionqualityreport
type Nludomainversionqualityreport struct { 
	// Version - The domain and version details of the quality report
	Version *Nludomainversion `json:"version,omitempty"`


	// ConfusionMatrix - The confusion matrix for the Domain Version
	ConfusionMatrix *[]Nluconfusionmatrixrow `json:"confusionMatrix,omitempty"`


	// Summary - The quality report summary for the Domain Version
	Summary *Nluqualityreportsummary `json:"summary,omitempty"`

}

func (u *Nludomainversionqualityreport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludomainversionqualityreport

	

	return json.Marshal(&struct { 
		Version *Nludomainversion `json:"version,omitempty"`
		
		ConfusionMatrix *[]Nluconfusionmatrixrow `json:"confusionMatrix,omitempty"`
		
		Summary *Nluqualityreportsummary `json:"summary,omitempty"`
		*Alias
	}{ 
		Version: u.Version,
		
		ConfusionMatrix: u.ConfusionMatrix,
		
		Summary: u.Summary,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nludomainversionqualityreport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
