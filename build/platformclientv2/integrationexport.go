package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Integrationexport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
