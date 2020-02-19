package platformclientv2
import (
	"time"
	"encoding/json"
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


	// RampPeriodStartDate - Date-time the ramp period starts. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	RampPeriodStartDate *time.Time `json:"rampPeriodStartDate,omitempty"`


	// RampPeriodEndDate - Date-time the ramp period ends. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	RampPeriodEndDate *time.Time `json:"rampPeriodEndDate,omitempty"`


	// BillingPeriodStartDate - Date-time the billing period started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	BillingPeriodStartDate *time.Time `json:"billingPeriodStartDate,omitempty"`


	// BillingPeriodEndDate - Date-time the billing period ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	BillingPeriodEndDate *time.Time `json:"billingPeriodEndDate,omitempty"`


	// Usages - Usages for the specified period.
	Usages *[]Subscriptionoverviewusage `json:"usages,omitempty"`


	// ContractAmendmentDate - Date-time the contract was last amended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ContractAmendmentDate *time.Time `json:"contractAmendmentDate,omitempty"`


	// ContractEffectiveDate - Date-time the contract became effective. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ContractEffectiveDate *time.Time `json:"contractEffectiveDate,omitempty"`


	// ContractEndDate - Date-time the contract ends. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ContractEndDate *time.Time `json:"contractEndDate,omitempty"`


	// MinimumMonthlyAmount - Minimum amount that will be charged for the month
	MinimumMonthlyAmount *string `json:"minimumMonthlyAmount,omitempty"`


	// InRampPeriod
	InRampPeriod *bool `json:"inRampPeriod,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trusteebillingoverview) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
