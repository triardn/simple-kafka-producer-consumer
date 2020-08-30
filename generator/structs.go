package generator

import "gopkg.in/guregu/null.v3"

type DonationCreated struct {
	UserID                 uint64         `json:"user_id"`
	IsNonLoginDonation     bool           `json:"is_non_login_donation"`
	PaymentMethodID        int64          `json:"payment_methods_id"`
	Platform               string         `json:"platform"`
	SendNotification       bool           `json:"send_notification"`
	PaymentToken           string         `json:"payment_token,omitempty"`
	RedirectCallback       string         `json:"redirect_callback,omitempty"`
	RedirectCallbackParams string         `json:"redirect_callback_params,omitempty"`
	Cashtag                string         `json:"cashtag,omitempty"`
	PromoCode              string         `json:"promo_code,omitempty"`
	Campaigns              []CampaignData `json:"campaigns"`
	XKtbsRequestID         string         `json:"X-Ktbs-Request-ID"`
}

// CampaignData defines campaign data when creating donation
type CampaignData struct {
	CampaignID             uint64      `json:"campaign_id"`
	Amount                 int64       `json:"amount"`
	IsParentDonation       null.Bool   `json:"is_parent_donation" validate:"required"`
	IsKBPlusDonation       bool        `json:"is_kbplus_donation,omitempty"`
	Comment                null.String `json:"comment"`
	ShareUserdataAgreement bool        `json:"share_userdata_agreement"`
	CampaignerRelation     bool        `json:"campaigner_relation"`
	IsAnonymous            bool        `json:"is_anonymous"`
	DependentID            string      `json:"dependent_id,omitempty"`
	DependentStatus        int         `json:"dependent_status,string,omitempty"`
	UTM                    UTM         `json:"utm"`
}

// UTM defines utm data
type UTM struct {
	Source   null.String `json:"source,omitempty"`
	Medium   null.String `json:"medium,omitempty"`
	Campaign null.String `json:"campaign,omitempty"`
	Content  null.String `json:"content,omitempty"`
}
