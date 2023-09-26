/*
Tokopedia API

Testing InteractionAPIService

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

func Test_tokopedia_InteractionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InteractionAPIService GetListMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.InteractionAPI.GetListMessage(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InteractionAPIService GetListReply", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64
		var msgId int64

		resp, httpRes, err := apiClient.InteractionAPI.GetListReply(context.Background(), fsId, msgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InteractionAPIService InitiateChat", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64

		resp, httpRes, err := apiClient.InteractionAPI.InitiateChat(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InteractionAPIService SendReply", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId int64
		var msgId int64

		resp, httpRes, err := apiClient.InteractionAPI.SendReply(context.Background(), fsId, msgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
