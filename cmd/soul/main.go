package main

import (
	"flag"
	"github.com/soul-ua/client/internal/keychain"
	"github.com/soul-ua/client/internal/register"
	"github.com/soul-ua/client/internal/tui"
	"github.com/soul-ua/server/pkg/sdk"
	"go.etcd.io/bbolt"
	"log"
	"os"
)

var (
	databasePath string
	username     string
	serverUrl    string

	registerAtStartup bool
)

func init() {
	flag.StringVar(&databasePath, "db", "", "optional database path (default: ~/.soul/client.db)")
	flag.StringVar(&username, "username", "", "current user's username")
	flag.StringVar(&serverUrl, "server", "https://soul.ua", "optional server URL (default: https://soul.ua)")
	flag.BoolVar(&registerAtStartup, "register", false, "register given username if not registered yet")
	flag.Parse()
}

func main() {
	if username == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if databasePath == "" {
		databasePath = os.ExpandEnv("$HOME/.soul/client.db")
	}

	db, err := bbolt.Open(databasePath, 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var soul *sdk.SDK

	kc := keychain.NewKeychain(db)
	if privateKey, err := kc.GetPrivateKey(username); err == nil {
		soul, err = sdk.NewSDKArmor(serverUrl, kc, username, privateKey)
		if err != nil {
			panic(err)
		}
	} else {
		if !registerAtStartup {
			log.Println("private key not found for given username")
			os.Exit(1)
		}
		soul, err = register.Register(username, serverUrl, kc)
		if err != nil {
			panic(err)
		}
	}

	serverInfo, err := soul.GetServerInfo()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to:", serverInfo.PublicKey)

	tui.Run()
}

//
//func sendRequest() {
//	kc := keychain.NewKeychain()
//	soul := getSoul(username, kc)
//
//	if err := soul.ContactRequest(recipient); err != nil {
//		panic(err)
//	}
//}
//
//func inbox() {
//	log.Println("get inbox")
//	kc := keychain.NewKeychain()
//	soul := getSoul(recipient, kc)
//
//	newInboxMessages, err := soul.GetInbox("")
//	if err != nil {
//		panic(err)
//	}
//	log.Println("received")
//
//	for _, envelope := range newInboxMessages {
//		log.Printf("[%s] %s>", envelope.PayloadType, envelope.From)
//	}
//
//	log.Println("contact request sent")
//}
