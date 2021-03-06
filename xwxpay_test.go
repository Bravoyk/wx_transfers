package wx_transfers

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXWXPay_Transfers(t *testing.T) {
	ctx  := context.Background()
	certPemFile := "apiclient_cert.pem"
	keyPemFile := "apiclient_key.pem"

	wxPayClient,err := NewXWXPay("apikey",certPemFile,keyPemFile)
	assert.NoError(t,err)
	param := &ParamPayRequest{
		MchAppid:       "MchAppid",
		Mchid:          "Mchid",
		PartnerTradeNo: "PartnerTradeNo",
		Openid:         "Openid",
		CheckName:      "NO_CHECK",
		Amount:         "100",
		Desc:           "test",
	}
	resp,err := wxPayClient.Transfers(ctx,param)
	assert.NoError(t,err)
	fmt.Printf("%+v",resp)
}

func TestXWXPay_GetTransferInfo(t *testing.T) {
	ctx  := context.Background()
	certPemFile := "apiclient_cert.pem"
	keyPemFile := "apiclient_key.pem"

	wxPayClient,err := NewXWXPay("apikey",certPemFile,keyPemFile)
	assert.NoError(t,err)
	param := &ParamPayResultRequest{
		PartnerTradeNo: "PartnerTradeNo",
		Appid:          "Appid",
		MchId:          "MchId",
	}
	resp,err := wxPayClient.GetTransferInfo(ctx,param)
	assert.NoError(t,err)
	fmt.Printf("%+v",resp)
}

func TestXWXPay_CreateSign(t *testing.T) {
	ctx  := context.Background()
	client,_ := NewXWXPay("9999","","")

	param := &ParamPayResultRequest{
		NonceStr:       "123",
		Sign:           "",
		PartnerTradeNo: "456",
		Appid:          "789",
		MchId:          "111",
	}
	sign,err := client.createSign(ctx,param)
	assert.NoError(t,err)
	fmt.Println(sign)
}