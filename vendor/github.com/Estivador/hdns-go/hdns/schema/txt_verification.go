package schema

// TxtVerification defines the schema of TXT Verification object.
type TxtVerification struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

// TxtVerificationResponse defines the schema of the response when
// retrieving TXT Verification object.
type TxtVerificationResponse struct {
	TxtVerification TxtVerification `json:"txt_verification"`
}
