package config

import (
	"crypto/ed25519"
	"encoding/hex"
	"github.com/google/uuid"
	"log"
	"strings"
)

type PasetoConfig struct {
	Version string `json:"version,omitempty"`
	Used string `json:"used,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

func (that *PasetoConfig) Init() *PasetoConfig {
	var (
		defaultVersion = "v2"
		defaultUsed = "public"
	)
	that.Version = strings.ToLower(that.Version)
	that.Used = strings.ToLower(that.Used)
	if that.Version == "" {
		that.Version = defaultVersion
	} else if that.Version != "v1" && that.Version != "v2" {
		that.Version = defaultVersion
	}
	if that.Used == "" {
		that.Used = defaultUsed
	} else if that.Used != "local" && that.Used != "public" {
		that.Used = defaultUsed
	}
	if that.Used == "local" {
		if that.PrivateKey == "" {
			that.PrivateKey = strings.Replace(uuid.New().String(), "-", "", 4)
		} else if len(that.PrivateKey) < 32 {
			panic("The length of the paseto private key must be 32.")
		} else {
			that.PrivateKey = that.PrivateKey[:32]
		}
		that.PublicKey = ""
	} else if that.Used == "public" {
		if that.PrivateKey == "" && that.PublicKey == "" {
			publicKey, privateKey, err := ed25519.GenerateKey(nil)
			if err != nil {
				log.Fatalln("create paseto public key fatal", err)
			}
			that.PrivateKey = hex.EncodeToString(privateKey)
			that.PublicKey = hex.EncodeToString(publicKey)
		} else if that.PrivateKey != "" && that.PublicKey != "" {
			that.PrivateKey = strings.TrimSpace(that.PrivateKey)
			that.PublicKey = strings.TrimSpace(that.PublicKey)
		} else if that.PrivateKey != "" {
			b, _ := hex.DecodeString(that.PrivateKey)
			privateKey := ed25519.PrivateKey(b)
			public := privateKey.Public()
			publicKey, _ := public.([]byte)
			that.PublicKey = hex.EncodeToString(publicKey)
		} else {
			panic("paseto public key must")
		}
	}
	return that
}