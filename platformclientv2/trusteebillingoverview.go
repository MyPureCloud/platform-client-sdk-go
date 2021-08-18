package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trusteebillingoverview
type Trusteebillingoverview struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Organization - Organization
	Organization *Namedentity `json:"organization,omitempty"`


	// Currency - The currency type.
	Currency *string `json:"currency,omitempty"`


	// EnabledProducts - The charge short names for products enabled during the specified period.
	EnabledProducts *[]string `json:"enabledProducts,omitempty"`


	// SubscriptionType - The subscription type.
	SubscriptionType *string `json:"subscriptionType,omitempty"`


	// RampPeriodStartDate - Date-time the ramp period starts. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RampPeriodStartDate *time.Time `json:"rampPeriodStartDate,omitempty"`


	// RampPeriodEndDate - Date-time the ramp period ends. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RampPeriodEndDate *time.Time `json:"rampPeriodEndDate,omitempty"`


	// BillingPeriodStartDate - Date-time the billing period started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	BillingPeriodStartDate *time.Time `json:"billingPeriodStartDate,omitempty"`


	// BillingPeriodEndDate - Date-time the billing period ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	BillingPeriodEndDate *time.Time `json:"billingPeriodEndDate,omitempty"`


	// Usages - Usages for the specified period.
	Usages *[]Subscriptionoverviewusage `json:"usages,omitempty"`


	// ContractAmendmentDate - Date-time the contract was last amended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ContractAmendmentDate *time.Time `json:"contractAmendmentDate,omitempty"`


	// ContractEffectiveDate - Date-time the contract became effective. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ContractEffectiveDate *time.Time `json:"contractEffectiveDate,omitempty"`


	// ContractEndDate - Date-time the contract ends. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ContractEndDate *time.Time `json:"contractEndDate,omitempty"`


	// MinimumMonthlyAmount - Minimum amount that will be charged for the month
	MinimumMonthlyAmount *string `json:"minimumMonthlyAmount,omitempty"`


	// InRampPeriod
	InRampPeriod *bool `json:"inRampPeriod,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Trusteebillingoverview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trusteebillingoverview

	
	RampPeriodStartDate := new(string)
	if u.RampPeriodStartDate != nil {
		
		*RampPeriodStartDate = timeutil.Strftime(u.RampPeriodStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RampPeriodStartDate = nil
	}
	
	RampPeriodEndDate := new(string)
	if u.RampPeriodEndDate != nil {
		
		*RampPeriodEndDate = timeutil.Strftime(u.RampPeriodEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RampPeriodEndDate = nil
	}
	
	BillingPeriodStartDate := new(string)
	if u.BillingPeriodStartDate != nil {
		
		*BillingPeriodStartDate = timeutil.Strftime(u.BillingPeriodStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		BillingPeriodStartDate = nil
	}
	
	BillingPeriodEndDate := new(string)
	if u.BillingPeriodEndDate != nil {
		
		*BillingPeriodEndDate = timeutil.Strftime(u.BillingPeriodEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		BillingPeriodEndDate = nil
	}
	
	ContractAmendmentDate := new(string)
	if u.ContractAmendmentDate != nil {
		
		*ContractAmendmentDate = timeutil.Strftime(u.ContractAmendmentDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ContractAmendmentDate = nil
	}
	
	ContractEffectiveDate := new(string)
	if u.ContractEffectiveDate != nil {
		
		*ContractEffectiveDate = timeutil.Strftime(u.ContractEffectiveDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ContractEffectiveDate = nil
	}
	
	ContractEndDate := new(string)
	if u.ContractEndDate != nil {
		
		*ContractEndDate = timeutil.Strftime(u.ContractEndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ContractEndDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Organization *Namedentity `json:"organization,omitempty"`
		
		Currency *string `json:"currency,omitempty"`
		
		EnabledProducts *[]string `json:"enabledProducts,omitempty"`
		
		SubscriptionType *string `json:"subscriptionType,omitempty"`
		
		RampPeriodStartDate *string `json:"rampPeriodStartDate,omitempty"`
		
		RampPeriodEndDate *string `json:"rampPeriodEndDate,omitempty"`
		
		BillingPeriodStartDate *string `json:"billingPeriodStartDate,omitempty"`
		
		BillingPeriodEndDate *string `json:"billingPeriodEndDate,omitempty"`
		
		Usages *[]Subscriptionoverviewusage `json:"usages,omitempty"`
		
		ContractAmendmentDate *string `json:"contractAmendmentDate,omitempty"`
		
		ContractEffectiveDate *string `json:"contractEffectiveDate,omitempty"`
		
		ContractEndDate *string `json:"contractEndDate,omitempty"`
		
		MinimumMonthlyAmount *string `json:"minimumMonthlyAmount,omitempty"`
		
		InRampPeriod *bool `json:"inRampPeriod,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Organization: u.Organization,
		
		Currency: u.Currency,
		
		EnabledProducts: u.EnabledProducts,
		
		SubscriptionType: u.SubscriptionType,
		
		RampPeriodStartDate: RampPeriodStartDate,
		
		RampPeriodEndDate: RampPeriodEndDate,
		
		BillingPeriodStartDate: BillingPeriodStartDate,
		
		BillingPeriodEndDate: BillingPeriodEndDate,
		
		Usages: u.Usages,
		
		ContractAmendmentDate: ContractAmendmentDate,
		
		ContractEffectiveDate: ContractEffectiveDate,
		
		ContractEndDate: ContractEndDate,
		
		MinimumMonthlyAmount: u.MinimumMonthlyAmount,
		
		InRampPeriod: u.InRampPeriod,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trusteebillingoverview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
