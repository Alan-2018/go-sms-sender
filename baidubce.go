package go_sms_sender

import (
	"bytes"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/sms"
	"github.com/baidubce/bce-sdk-go/services/sms/api"
)

type BceClient struct {
	sign     string
	template string
	core     *sms.Client
}

func GetBceClient(accessId, accessKey, endpoint, sign, template string) (*BceClient, error) {
	client, err := sms.NewClient(accessId, accessKey, endpoint)
	if err != nil {
		return nil, err
	}

	bceClient := &BceClient{
		sign:     sign,
		template: template,
		core:     client,
	}

	return bceClient, nil
}

func (c *BceClient) SendMessage(param map[string]string, targetPhoneNumber ...string) error {
	code, ok := param["code"]
	if !ok {
		return fmt.Errorf("missing parameter: msg code")
	}

	phoneNumbers := bytes.Buffer{}
	phoneNumbers.WriteString(targetPhoneNumber[0])
	for _, s := range targetPhoneNumber[1:] {
		phoneNumbers.WriteString(",")
		phoneNumbers.WriteString(s)
	}

	contentMap := make(map[string]interface{})
	contentMap["code"] = code
	// contentMap["minute"] = "1"

	sendSmsArgs := &api.SendSmsArgs{
		Mobile:      phoneNumbers.String(),
		SignatureId: c.sign,
		Template:    c.template,
		ContentVar:  contentMap,
	}

	result, err := c.core.SendSms(sendSmsArgs)
	if err != nil {
		fmt.Printf("send sms error, %s", err)
		return err
	}

	fmt.Printf("send sms success. %s", result)
	return nil
}
