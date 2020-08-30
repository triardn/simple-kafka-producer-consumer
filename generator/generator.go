package generator

func GenerateMessage(flag int) DonationCreated {
	amount := 100000 + (flag * 1000)
	return DonationCreated{
		UserID:             44648,
		IsNonLoginDonation: false,
		PaymentMethodID:    19,
		Platform:           "android",
		SendNotification:   false,
		Campaigns: []CampaignData{
			{
				CampaignID:             19117,
				Amount:                 int64(amount),
				IsKBPlusDonation:       false,
				ShareUserdataAgreement: false,
				CampaignerRelation:     true,
				IsAnonymous:            false,
			},
		},
		XKtbsRequestID: "111",
	}
}
