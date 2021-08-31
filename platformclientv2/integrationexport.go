package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationexport
type Integrationexport struct { 
	// Integration - The aws-s3-recording-bulk-actions-integration that the policy uses for exports.
	Integration *Domainentityref `json:"integration,omitempty"`


	// ShouldExportScreenRecordings - True if the policy should export screen recordings in addition to the other conversation media. Default = true
	ShouldExportScreenRecordings *bool `json:"shouldExportScreenRecordings,omitempty"`

}

func (o *Integrationexport) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationexport
	
	return json.Marshal(&struct { 
		Integration *Domainentityref `json:"integration,omitempty"`
		
		ShouldExportScreenRecordings *bool `json:"shouldExportScreenRecordings,omitempty"`
		*Alias
	}{ 
		Integration: o.Integration,
		
		ShouldExportScreenRecordings: o.ShouldExportScreenRecordings,
		Alias:    (*Alias)(o),
	})
}

func (o *Integrationexport) UnmarshalJSON(b []byte) error {
	var IntegrationexportMap map[string]interface{}
	err := json.Unmarshal(b, &IntegrationexportMap)
	if err != nil {
		return err
	}
	
	if Integration, ok := IntegrationexportMap["integration"].(map[string]interface{}); ok {
		IntegrationString, _ := json.Marshal(Integration)
		json.Unmarshal(IntegrationString, &o.Integration)
	}
	
	if ShouldExportScreenRecordings, ok := IntegrationexportMap["shouldExportScreenRecordings"].(bool); ok {
		o.ShouldExportScreenRecordings = &ShouldExportScreenRecordings
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Integrationexport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
