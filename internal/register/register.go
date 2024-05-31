package register

import (
	"fmt"
	"log"

	"github.com/ProtonMail/gopenpgp/v2/crypto"

	"github.com/soul-ua/client/internal/keychain"
	"github.com/soul-ua/server/pkg/sdk"
)

func Register(username, serverURL string, kc *keychain.Keychain) (*sdk.SDK, error) {

	ecKey, err := crypto.GenerateKey(username, fmt.Sprintf("%s@soul.ua", username), "x25519", 0)
	if err != nil {
		panic(err)
	}

	soul, err := sdk.NewSDK(serverURL, kc, username, ecKey)
	if err != nil {
		panic(err)
	}

	privateKeyArmor, err := ecKey.Armor()
	if err != nil {
		panic(err)
	}

	publicKeyArmor, err := ecKey.GetArmoredPublicKey()
	if err != nil {
		panic(err)
	}

	if err := soul.Register(username, publicKeyArmor); err != nil {
		panic(err)
	}

	if err := kc.SavePrivateKey(username, privateKeyArmor); err != nil {
		panic(err)
	}

	if err := kc.SavePublicKey(username, publicKeyArmor); err != nil {
		panic(err)
	}

	log.Println("success")

	return soul, nil
}
