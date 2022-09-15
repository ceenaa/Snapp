package coding

import (
	"Project/models"
	"bytes"
	"encoding/gob"
)

func Hash(rule models.Rule) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(rule)
	return b.Bytes()
}

func UnHash(t string) models.Rule {
	b := []byte(t)
	var rule models.Rule
	gob.NewDecoder(bytes.NewBuffer(b)).Decode(&rule)
	return rule
}

func HashRaw(rule models.RawRule) []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(rule)
	return b.Bytes()
}

func UnHashRaw(t string) models.RawRule {
	b := []byte(t)
	var rule models.RawRule
	gob.NewDecoder(bytes.NewBuffer(b)).Decode(&rule)
	return rule
}
