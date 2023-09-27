/*
Tokopedia API

Testing ShopAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tokopedia

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/dhimas-sirclo/apiclient/tokopedia"
)

func Test_tokopedia_ShopAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ShopAPIService CreateShowcase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ShopAPI.CreateShowcase(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopAPIService DeleteShowcase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ShopAPI.DeleteShowcase(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopAPIService GetShopInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ShopAPI.GetShopInfo(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopAPIService GetShowcase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ShopAPI.GetShowcase(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopAPIService GetShowcaseAllEtalase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ShopAPI.GetShowcaseAllEtalase(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopAPIService UpdateShopStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ShopAPI.UpdateShopStatus(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShopAPIService UpdateShowcase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.ShopAPI.UpdateShowcase(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}