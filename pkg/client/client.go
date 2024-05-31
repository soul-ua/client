package client

import "github.com/soul-ua/server/pkg/sdk"

type Client struct {
	sdk *sdk.SDK
}

func NewClient(serverURI string) (*Client, error) {
	serverSDK, err := sdk.NewSDK(serverURI)
	if err != nil {
		return nil, err
	}

	return &Client{sdk: serverSDK}, nil
}
