package models

import "encoding/json"

type SecretType string

const (
	SecretTypeAuth SecretType = "auth"
	SecretTypeText SecretType = "text"
	SecretTypeBin  SecretType = "bin"
	SecretTypeBank SecretType = "bank"
)

func (s SecretType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

func (s *SecretType) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = SecretType(str)
	return nil
}
