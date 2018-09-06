package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	sms1 := NewCommonMessage(new(MessageSMS))
	sms1.SendMessage("hello bob, this SMS1", "bob")

	email1 := NewCommonMessage(new(MessageEmail))
	email1.SendMessage("hello bob, this email1", "bob")

	sms2 := NewUrgencyMessage(new(MessageSMS))
	sms2.SendMessage("hello bob, this SMS2", "bob")

	email2 := NewUrgencyMessage(new(MessageEmail))
	email2.SendMessage("hello bob, this email2", "bob")
}
