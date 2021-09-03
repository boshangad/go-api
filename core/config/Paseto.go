package config

import (
	"crypto/ed25519"
	"encoding/hex"
	"github.com/google/uuid"
	"log"
	"strings"
)

type PasetoConfig struct {
	Version string `json:"version"`
	Used string `json:"used"`
	PrivateKey string `json:"private_key"`
	PublicKey string `json:"public_key"`
}

func (c *PasetoConfig) initDefaultData() {
	var (
		defaultVersion = "v2"
		defaultUsed = "public"
	)
	c.Version = strings.ToLower(c.Version)
	c.Used = strings.ToLower(c.Used)
	if c.Version == "" {
		c.Version = defaultVersion
	} else if c.Version != "v1" && c.Version != "v2" {
		c.Version = defaultVersion
	}
	if c.Used == "" {
		c.Used = defaultUsed
	} else if c.Used != "local" && c.Used != "public" {
		c.Used = defaultUsed
	}
	if c.Used == "local" {
		if c.PrivateKey == "" {
			c.PrivateKey = strings.Replace(uuid.New().String(), "-", "", 4)
		} else if len(c.PrivateKey) < 32 {
			panic("The length of the paseto private key must be 32.")
		} else {
			c.PrivateKey = c.PrivateKey[:32]
		}
		c.PublicKey = ""
	} else if c.Used == "public" {
		if c.PrivateKey == "" && c.PublicKey == "" {
			publicKey, privateKey, err := ed25519.GenerateKey(nil)
			if err != nil {
				log.Fatalln("create paseto public key fatal", err)
			}
			c.PrivateKey = hex.EncodeToString(privateKey)
			c.PublicKey = hex.EncodeToString(publicKey)
		} else if c.PrivateKey != "" && c.PublicKey != "" {
			c.PrivateKey = strings.TrimSpace(c.PrivateKey)
			c.PublicKey = strings.TrimSpace(c.PublicKey)
		} else if c.PrivateKey != "" {
			b, _ := hex.DecodeString(c.PrivateKey)
			privateKey := ed25519.PrivateKey(b)
			public := privateKey.Public()
			publicKey, _ := public.([]byte)
			c.PublicKey = hex.EncodeToString(publicKey)
		} else {
			panic("paseto public key must")
		}
	}

}