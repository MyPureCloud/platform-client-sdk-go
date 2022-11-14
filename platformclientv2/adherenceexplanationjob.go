package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherenceexplanationjob
type Adherenceexplanationjob struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// VarType - The type of the adherence explanation job
	VarType *string `json:"type,omitempty"`


	// Status - The status of the adherence explanation job
	Status *string `json:"status,omitempty"`


	// AdherenceExplanation - The adherence explanation added or modified by the job once complete; may be null if status == 'Error'. Used if type is in [ 'AddExplanation', 'UpdateExplanation' ]
	AdherenceExplanation *Adherenceexplanationresponse `json:"adherenceExplanation,omitempty"`


	// DownloadUrl - A URL to fetch results of the job. Only set if status == 'Complete' and type is in [ 'QueryAgentExplanations', 'QueryBuExplanations' ]
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// VarError - Error details if status == 'Error'
	VarError *Errorbody `json:"error,omitempty"`


	// AgentQueryResponseTemplate - Schema template for deserializing data returned from the downloadUrl. Use if type == 'QueryAgentExplanations'
	AgentQueryResponseTemplate *Adherenceexplanationlistingagentqueryresponse `json:"agentQueryResponseTemplate,omitempty"`


	// BuQueryResponseTemplate - Schema template for deserializing data returned from the downloadUrl. Use if type == 'QueryBuExplanations'
	BuQueryResponseTemplate *Adherenceexplanationlistingbuqueryresponse `json:"buQueryResponseTemplate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Adherenceexplanationjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adherenceexplanationjob
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		AdherenceExplanation *Adherenceexplanationresponse `json:"adherenceExplanation,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		VarError *Errorbody `json:"error,omitempty"`
		
		AgentQueryResponseTemplate *Adherenceexplanationlistingagentqueryresponse `json:"agentQueryResponseTemplate,omitempty"`
		
		BuQueryResponseTemplate *Adherenceexplanationlistingbuqueryresponse `json:"buQueryResponseTemplate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		
		Status: o.Status,
		
		AdherenceExplanation: o.AdherenceExplanation,
		
		DownloadUrl: o.DownloadUrl,
		
		VarError: o.VarError,
		
		AgentQueryResponseTemplate: o.AgentQueryResponseTemplate,
		
		BuQueryResponseTemplate: o.BuQueryResponseTemplate,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Adherenceexplanationjob) UnmarshalJSON(b []byte) error {
	var AdherenceexplanationjobMap map[string]interface{}
	err := json.Unmarshal(b, &AdherenceexplanationjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AdherenceexplanationjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if VarType, ok := AdherenceexplanationjobMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Status, ok := AdherenceexplanationjobMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if AdherenceExplanation, ok := AdherenceexplanationjobMap["adherenceExplanation"].(map[string]interface{}); ok {
		AdherenceExplanationString, _ := json.Marshal(AdherenceExplanation)
		json.Unmarshal(AdherenceExplanationString, &o.AdherenceExplanation)
	}
	
	if DownloadUrl, ok := AdherenceexplanationjobMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    
	if VarError, ok := AdherenceexplanationjobMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if AgentQueryResponseTemplate, ok := AdherenceexplanationjobMap["agentQueryResponseTemplate"].(map[string]interface{}); ok {
		AgentQueryResponseTemplateString, _ := json.Marshal(AgentQueryResponseTemplate)
		json.Unmarshal(AgentQueryResponseTemplateString, &o.AgentQueryResponseTemplate)
	}
	
	if BuQueryResponseTemplate, ok := AdherenceexplanationjobMap["buQueryResponseTemplate"].(map[string]interface{}); ok {
		BuQueryResponseTemplateString, _ := json.Marshal(BuQueryResponseTemplate)
		json.Unmarshal(BuQueryResponseTemplateString, &o.BuQueryResponseTemplate)
	}
	
	if SelfUri, ok := AdherenceexplanationjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adherenceexplanationjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
