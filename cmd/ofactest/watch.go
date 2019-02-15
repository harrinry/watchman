// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/moov-io/base"
	moov "github.com/moov-io/ofac/client"
)

func addCompanyWatch(ctx context.Context, api *moov.APIClient, id string) error {
	_, resp, err := api.OFACApi.AddCompanyWatch(ctx, id, moov.WatchRequest{
		AuthToken: base.ID(),
		Webhook:   "https://moov.io/ofac",
	}, nil)
	if err != nil {
		return fmt.Errorf("addCompanyWatch: %v", err)
	}
	defer resp.Body.Close()
	return nil
}

func addCustomerWatch(ctx context.Context, api *moov.APIClient, id string) error {
	_, resp, err := api.OFACApi.AddCustomerWatch(ctx, id, moov.WatchRequest{
		AuthToken: base.ID(),
		Webhook:   "https://moov.io/ofac",
	}, nil)
	if err != nil {
		return fmt.Errorf("addCustomerWatch: %v", err)
	}
	defer resp.Body.Close()
	return nil
}