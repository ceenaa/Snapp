package coding

import (
	"Project/models"
	"bytes"
	"encoding/gob"
)

func Hash(rule models.Rule) []byte {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(rule)
	if err != nil {
		return []byte{}
	}
	return b.Bytes()
}

func UnHash(t string) models.Rule {
	b := []byte(t)
	var rule models.Rule
	err := gob.NewDecoder(bytes.NewBuffer(b)).Decode(&rule)
	if err != nil {
		return models.Rule{}
	}
	return rule
}

func HashRaw(rule models.RawRule) []byte {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(rule)
	if err != nil {
		return []byte{}
	}
	return b.Bytes()
}

func UnHashRaw(t string) models.RawRule {
	b := []byte(t)
	var rule models.RawRule
	err := gob.NewDecoder(bytes.NewBuffer(b)).Decode(&rule)
	if err != nil {
		return models.RawRule{}
	}
	return rule
}
