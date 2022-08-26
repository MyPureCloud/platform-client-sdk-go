package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentversion
type Knowledgedocumentversion struct { 
	// Id - Globally unique identifier for the document version.
	Id *string `json:"id,omitempty"`


	// DatePublished - Published date of document version. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`


	// Document - The document which is versioned.
	Document *Knowledgedocumentresponse `json:"document,omitempty"`


	// RestoreFromVersionId - The globally unique identifier for the document version. If the value is provided, the document is restored to the given version. If not, it publishes the draft changes as a new version of the document.
	RestoreFromVersionId *string `json:"restoreFromVersionId,omitempty"`


	// VersionNumber - Version Number of the document.
	VersionNumber *int `json:"versionNumber,omitempty"`


	// DateExpires - Expiry date of document version, applicable only to the 'Archived' version of the document. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpires *time.Time `json:"dateExpires,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgedocumentversion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentversion
	
	DatePublished := new(string)
	if o.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(o.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublished = nil
	}
	
	DateExpires := new(string)
	if o.DateExpires != nil {
		
		*DateExpires = timeutil.Strftime(o.DateExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpires = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		Document *Knowledgedocumentresponse `json:"document,omitempty"`
		
		RestoreFromVersionId *string `json:"restoreFromVersionId,omitempty"`
		
		VersionNumber *int `json:"versionNumber,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		DatePublished: DatePublished,
		
		Document: o.Document,
		
		RestoreFromVersionId: o.RestoreFromVersionId,
		
		VersionNumber: o.VersionNumber,
		
		DateExpires: DateExpires,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentversion) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentversionMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentversionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentversionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if datePublishedString, ok := KnowledgedocumentversionMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if Document, ok := KnowledgedocumentversionMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	
	if RestoreFromVersionId, ok := KnowledgedocumentversionMap["restoreFromVersionId"].(string); ok {
		o.RestoreFromVersionId = &RestoreFromVersionId
	}
    
	if VersionNumber, ok := KnowledgedocumentversionMap["versionNumber"].(float64); ok {
		VersionNumberInt := int(VersionNumber)
		o.VersionNumber = &VersionNumberInt
	}
	
	if dateExpiresString, ok := KnowledgedocumentversionMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if SelfUri, ok := KnowledgedocumentversionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentversion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
