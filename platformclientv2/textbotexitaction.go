package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotexitaction - Settings for a next-action of exiting the bot flow. Any output variables are available in the details.
type Textbotexitaction struct { 
	// Reason - The reason for the exit.
	Reason *string `json:"reason,omitempty"`


	// ReasonExtendedInfo - Extended information related to the reason, if available.
	ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`


	// ActiveIntent - The active intent at the time of the exit.
	ActiveIntent *string `json:"activeIntent,omitempty"`


	// FlowLocation - Describes where in the Bot Flow the user was when the exit occurred.
	FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`


	// OutputData - The output data for the bot flow.
	OutputData *Textbotinputoutputdata `json:"outputData,omitempty"`


	// FlowOutcomes - The list of Flow Outcomes for the bot flow and their details.
	FlowOutcomes *[]Textbotflowoutcome `json:"flowOutcomes,omitempty"`

}

func (o *Textbotexitaction) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotexitaction
	
	return json.Marshal(&struct { 
		Reason *string `json:"reason,omitempty"`
		
		ReasonExtendedInfo *string `json:"reasonExtendedInfo,omitempty"`
		
		ActiveIntent *string `json:"activeIntent,omitempty"`
		
		FlowLocation *Textbotflowlocation `json:"flowLocation,omitempty"`
		
		OutputData *Textbotinputoutputdata `json:"outputData,omitempty"`
		
		FlowOutcomes *[]Textbotflowoutcome `json:"flowOutcomes,omitempty"`
		*Alias
	}{ 
		Reason: o.Reason,
		
		ReasonExtendedInfo: o.ReasonExtendedInfo,
		
		ActiveIntent: o.ActiveIntent,
		
		FlowLocation: o.FlowLocation,
		
		OutputData: o.OutputData,
		
		FlowOutcomes: o.FlowOutcomes,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotexitaction) UnmarshalJSON(b []byte) error {
	var TextbotexitactionMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotexitactionMap)
	if err != nil {
		return err
	}
	
	if Reason, ok := TextbotexitactionMap["reason"].(string); ok {
		o.Reason = &Reason
	}
	
	if ReasonExtendedInfo, ok := TextbotexitactionMap["reasonExtendedInfo"].(string); ok {
		o.ReasonExtendedInfo = &ReasonExtendedInfo
	}
	
	if ActiveIntent, ok := TextbotexitactionMap["activeIntent"].(string); ok {
		o.ActiveIntent = &ActiveIntent
	}
	
	if FlowLocation, ok := TextbotexitactionMap["flowLocation"].(map[string]interface{}); ok {
		FlowLocationString, _ := json.Marshal(FlowLocation)
		json.Unmarshal(FlowLocationString, &o.FlowLocation)
	}
	
	if OutputData, ok := TextbotexitactionMap["outputData"].(map[string]interface{}); ok {
		OutputDataString, _ := json.Marshal(OutputData)
		json.Unmarshal(OutputDataString, &o.OutputData)
	}
	
	if FlowOutcomes, ok := TextbotexitactionMap["flowOutcomes"].([]interface{}); ok {
		FlowOutcomesString, _ := json.Marshal(FlowOutcomes)
		json.Unmarshal(FlowOutcomesString, &o.FlowOutcomes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotexitaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
