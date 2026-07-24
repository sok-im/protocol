package sdkws

import (
	"testing"

	"github.com/openimsdk/protocol/constant"
)

func TestMsgDataCheck_WalletActionNotificationSessionTypes(t *testing.T) {
	base := func(sessionType, contentType int32) *MsgData {
		return &MsgData{
			SendID:      "u1",
			Content:     []byte(`{}`),
			SessionType: sessionType,
			ContentType: contentType,
		}
	}

	t.Run("red packet claim in notification session", func(t *testing.T) {
		if err := base(constant.NotificationChatType, constant.RedPacketClaimNotification).Check(); err != nil {
			t.Fatalf("expected ok, got %v", err)
		}
	})

	t.Run("red packet claim in group session", func(t *testing.T) {
		if err := base(constant.ReadGroupChatType, constant.RedPacketClaimNotification).Check(); err != nil {
			t.Fatalf("expected ok for group wallet notification, got %v", err)
		}
	})

	t.Run("transfer receive in group session", func(t *testing.T) {
		if err := base(constant.ReadGroupChatType, constant.TransferReceiveNotification).Check(); err != nil {
			t.Fatalf("expected ok for group wallet notification, got %v", err)
		}
	})

	t.Run("service notification still requires notification session", func(t *testing.T) {
		err := base(constant.ReadGroupChatType, constant.ServiceNotification).Check()
		if err == nil {
			t.Fatal("expected error for ServiceNotification outside NotificationChatType")
		}
	})
}
