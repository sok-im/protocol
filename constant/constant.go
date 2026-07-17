package constant

const (

	///ContentType
	//UserRelated.
	ContentTypeBegin = 100
	Text             = 101
	Picture          = 102
	Voice            = 103
	Video            = 104
	File             = 105
	AtText           = 106
	Merger           = 107
	Card             = 108
	Location         = 109
	Custom           = 110
	Revoke           = 111
	Typing           = 113
	Quote            = 114

	AdvancedText                 = 117
	MarkdownText                 = 118
	CustomNotTriggerConversation = 119
	CustomOnlineOnly             = 120
	ReactionMessageModifier      = 121
	ReactionMessageDeleter       = 122
	Common                       = 200
	GroupMsg                     = 201
	SignalMsg                    = 202
	CustomNotification           = 203

	// SysRelated.
	NotificationBegin = 1000

	FriendApplicationApprovedNotification = 1201 // add_friend_response
	FriendApplicationRejectedNotification = 1202 // add_friend_response
	FriendApplicationNotification         = 1203 // add_friend
	FriendAddedNotification               = 1204
	FriendDeletedNotification             = 1205 // delete_friend
	FriendRemarkSetNotification           = 1206 // set_friend_remark?
	BlackAddedNotification                = 1207 // add_black
	BlackDeletedNotification              = 1208 // remove_black
	FriendInfoUpdatedNotification         = 1209
	FriendsInfoUpdateNotification         = 1210 //update friend info

	ConversationChangeNotification = 1300 // change conversation opt

	UserNotificationBegin         = 1301
	UserInfoUpdatedNotification   = 1303 // SetSelfInfoTip              = 204
	UserStatusChangeNotification  = 1304
	UserCommandAddNotification    = 1305
	UserCommandDeleteNotification = 1306
	UserCommandUpdateNotification = 1307

	UserSubscribeOnlineStatusNotification = 1308

	UserNotificationEnd = 1399
	OANotification      = 1400

	GroupNotificationBegin = 1500

	GroupCreatedNotification                     = 1501
	GroupInfoSetNotification                     = 1502
	JoinGroupApplicationNotification             = 1503
	MemberQuitNotification                       = 1504
	GroupApplicationAcceptedNotification         = 1505
	GroupApplicationRejectedNotification         = 1506
	GroupOwnerTransferredNotification            = 1507
	MemberKickedNotification                     = 1508
	MemberInvitedNotification                    = 1509
	MemberEnterNotification                      = 1510
	GroupDismissedNotification                   = 1511
	GroupMemberMutedNotification                 = 1512
	GroupMemberCancelMutedNotification           = 1513
	GroupMutedNotification                       = 1514
	GroupCancelMutedNotification                 = 1515
	GroupMemberInfoSetNotification               = 1516
	GroupMemberSetToAdminNotification            = 1517
	GroupMemberSetToOrdinaryUserNotification     = 1518
	GroupInfoSetAnnouncementNotification         = 1519
	GroupInfoSetNameNotification                 = 1520
	GroupMessagePinnedNotification               = 1521
	GroupCallStartedNotification                 = 1522
	GroupCallEndedNotification                   = 1523
	GroupCallParticipantCountUpdatedNotification = 1527
	GroupCallParticipantDeclinedNotification     = 1528
	GroupBurnDurationSetNotification             = 1524
	GroupFaceURLSetNotification                  = 1525
	GroupNeedVerificationSetNotification         = 1526
	GroupE2EENotification                        = 1529
	GroupPermissionChangedNotification           = 1530

	//SignalingNotificationBegin = 1600
	//SignalingNotification      = 1601
	//SignalingNotificationEnd   = 1649

	SuperGroupNotificationBegin  = 1650
	SuperGroupUpdateNotification = 1651
	MsgDeleteNotification        = 1652
	SuperGroupNotificationEnd    = 1699

	ConversationPrivateChatNotification = 1701
	ConversationUnreadNotification      = 1702
	ClearConversationNotification       = 1703
	ConversationDeleteNotification      = 1704
	ConversationE2EENotification        = 1705

	// 钱包动作通知（红包/转账）
	RedPacketClaimNotification   = 1801 // 红包领取通知
	TransferReceiveNotification  = 1802 // 转账接收通知
	RedPacketExpiredNotification = 1803 // 红包过期通知
	TransferExpiredNotification  = 1804 // 转账过期通知

	BusinessNotificationBegin = 2000
	BusinessNotification      = 2001
	ServiceNotification       = 2002 // 服务通知
	PaymentNotification       = 2003 // 支付通知
	BusinessNotificationEnd   = 2099

	MsgRevokeNotification  = 2101
	DeleteMsgsNotification = 2102
	// MsgReactionUpdatedNotification 消息表情反应聚合快照更新通知（不计未读、不改会话摘要）
	MsgReactionUpdatedNotification = 2103

	HasReadReceipt = 2200

	NotificationEnd = 5000

	// status.
	MsgNormal  = 1
	MsgDeleted = 4

	// MsgFrom.
	UserMsgType = 100
	SysMsgType  = 200

	// SessionType.
	SingleChatType = 1
	// WriteGroupChatType Not enabled temporarily
	WriteGroupChatType   = 2
	ReadGroupChatType    = 3
	NotificationChatType = 4
	// token.
	NormalToken  = 0
	InValidToken = 1
	KickedToken  = 2
	ExpiredToken = 3

	// MultiTerminalLogin.
	DefalutNotKick = 0
	// Full-end login, but the same end is mutually exclusive.
	AllLoginButSameTermKick = 1
	// New login kicks all existing sessions across all platforms; only one session allowed at any time.
	AllKick = 2
	// The PC side is mutually exclusive, and the mobile side is mutually exclusive, but the web side can be online at
	// the same time.
	AllLoginButSameClassKick = 4
	// The PC terminal can be online at the same time,but other terminal only one of the endpoints can login.
	PCAndOther = 5

	// OnlineStatus  = "online"
	// OfflineStatus = "offline"
	// Registered    = "registered"
	// UnRegistered  = "unregistered"

	Online  = 1
	Offline = 0

	Registered   = 1
	UnRegistered = 0

	// MsgReceiveOpt.
	ReceiveMessage          = 0
	NotReceiveMessage       = 1
	ReceiveNotNotifyMessage = 2

	// OptionsKey.
	IsHistory                  = "history"
	IsPersistent               = "persistent"
	IsOfflinePush              = "offlinePush"
	IsUnreadCount              = "unreadCount"
	IsConversationUpdate       = "conversationUpdate"
	IsSenderSync               = "senderSync"
	IsNotPrivate               = "notPrivate"
	IsSenderConversationUpdate = "senderConversationUpdate"
	IsSenderNotificationPush   = "senderNotificationPush"
	IsReactionFromCache        = "reactionFromCache"
	IsNotNotification          = "isNotNotification"
	IsSendMsg                  = "isSendMsg"

	// GroupStatus.
	GroupOk              = 0
	GroupBanChat         = 1
	GroupStatusDismissed = 2
	GroupStatusMuted     = 3

	// GroupType.
	NormalGroup  = 0
	SuperGroup   = 1
	WorkingGroup = 2

	GroupBaned          = 3
	GroupBanPrivateChat = 4

	// UserJoinGroupSource.
	JoinByAdmin = 1

	JoinByInvitation = 2
	JoinBySearch     = 3
	JoinByQRCode     = 4
	JoinByInviteLink = 5

	// Minio.
	MinioDurationTimes = 3600
	// Aws.
	AwsDurationTimes = 3600

	// callbackCommand.
	CallbackBeforeSendSingleMsgCommand                   = "callbackBeforeSendSingleMsgCommand"
	CallbackAfterSendSingleMsgCommand                    = "callbackAfterSendSingleMsgCommand"
	CallbackBeforeSendGroupMsgCommand                    = "callbackBeforeSendGroupMsgCommand"
	CallbackAfterSendGroupMsgCommand                     = "callbackAfterSendGroupMsgCommand"
	CallbackMsgModifyCommand                             = "callbackMsgModifyCommand"
	CallbackUserOnlineCommand                            = "callbackUserOnlineCommand"
	CallbackUserOfflineCommand                           = "callbackUserOfflineCommand"
	CallbackUserKickOffCommand                           = "callbackUserKickOffCommand"
	CallbackOfflinePushCommand                           = "callbackOfflinePushCommand"
	CallbackOnlinePushCommand                            = "callbackOnlinePushCommand"
	CallbackSuperGroupOnlinePushCommand                  = "callbackSuperGroupOnlinePushCommand"
	CallbackBeforeAddFriendCommand                       = "callbackBeforeAddFriendCommand"
	CallbackBeforeUpdateUserInfoCommand                  = "callbackBeforeUpdateUserInfoCommand"
	CallbackBeforeCreateGroupCommand                     = "callbackBeforeCreateGroupCommand"
	CallbackBeforeMemberJoinGroupCommand                 = "callbackBeforeMemberJoinGroupCommand"
	CallbackBeforeSetGroupMemberInfoCommand              = "CallbackBeforeSetGroupMemberInfoCommand"
	CallbackBeforeSetMessageReactionExtensionCommand     = "callbackBeforeSetMessageReactionExtensionCommand"
	CallbackBeforeDeleteMessageReactionExtensionsCommand = "callbackBeforeDeleteMessageReactionExtensionsCommand"
	CallbackGetMessageListReactionExtensionsCommand      = "callbackGetMessageListReactionExtensionsCommand"
	CallbackAddMessageListReactionExtensionsCommand      = "callbackAddMessageListReactionExtensionsCommand"

	// callback actionCode.
	ActionAllow     = 0
	ActionForbidden = 1
	// callback callbackHandleCode.
	CallbackHandleSuccess = 0
	CallbackHandleFailed  = 1

	// minioUpload.
	OtherType = 1
	VideoType = 2
	ImageType = 3

	// sendMsgStaus.
	MsgStatusNotExist = 0
	MsgIsSending      = 1
	MsgSendSuccessed  = 2
	MsgSendFailed     = 3
)

const (
	WriteDiffusion = 0
	ReadDiffusion  = 1
)

const (
	UnreliableNotification    = 1
	ReliableNotificationNoMsg = 2
	ReliableNotificationMsg   = 3
)

const (
	AtAllString       = "AtAllTag"
	AtNormal          = 0
	AtMe              = 1
	AtAll             = 2
	AtAllAtMe         = 3
	GroupNotification = 4
)

var ContentType2PushContent = map[int64]string{
	Picture:   "[PICTURE]",
	Voice:     "[VOICE]",
	Video:     "[VIDEO]",
	File:      "[File]",
	Text:      "[TEXT]",
	AtText:    "[@TEXT]",
	GroupMsg:  "[GROUPMSG]]",
	Common:    "[NEWMSG]",
	SignalMsg: "[SIGNALINVITE]",
}

const (
	FieldRecvMsgOpt    = 1
	FieldIsPinned      = 2
	FieldAttachedInfo  = 3
	FieldIsPrivateChat = 4
	FieldGroupAtType   = 5
	FieldEx            = 7
	FieldUnread        = 8
	FieldBurnDuration  = 9
	FieldHasReadSeq    = 10
)

const (
	IMOrdinaryUser       = 0
	AppOrdinaryUsers     = 1
	AppAdmin             = 2
	AppNotificationAdmin = 3
	AppRobotAdmin        = 4

	GroupOwner         = 100
	GroupAdmin         = 60
	GroupOrdinaryUsers = 20

	GroupResponseAgree  = 1
	GroupResponseRefuse = -1

	FriendResponseNotHandle = 0
	FriendResponseAgree     = 1
	FriendResponseRefuse    = -1

	Male   = 1
	Female = 2
)

const (
	OperationID     = "operationID"
	OpUserID        = "opUserID"
	ConnID          = "connID"
	OpUserPlatform  = "platform"
	Token           = "token"
	RpcCustomHeader = "customHeader" // custom ctx parameter used by RPC middleware
	CheckKey        = "CheckKey"
	TriggerID       = "triggerID"
	RemoteAddr      = "remoteAddr"
)

const (
	BecomeFriendByImport = 1 // imported by administrator
	BecomeFriendByApply  = 2 // added via application
)

const (
	// 以下取值仅用于群分享/邀请链接入群（JoinGroupByInviteLink）；主动申请、邀请拉人不再读取 needVerification。
	ApplyNeedVerification = 0 // 分享链接入群需审批
	AllNeedVerification   = 1 // 分享链接入群需审批（与 0 等价）
	Directly              = 2 // 分享链接入群免审，直接入群（开启分享链接时的默认值）
)

const (
	GroupRPCRecvSize = 30
	GroupRPCSendSize = 30
)

const FriendAcceptTip = "You have successfully become friends, so start chatting"

func GroupIsBanChat(status int32) bool {
	if status != GroupStatusMuted {
		return false
	}
	return true
}

func GroupIsBanPrivateChat(status int32) bool {
	if status != GroupBanPrivateChat {
		return false
	}
	return true
}

// IsNotificationSessionContentType 判断是否为通知会话消息类型（需配合 NotificationChatType 使用）。
func IsNotificationSessionContentType(contentType int32) bool {
	switch contentType {
	case OANotification, ServiceNotification, PaymentNotification,
		RedPacketClaimNotification, TransferReceiveNotification,
		RedPacketExpiredNotification, TransferExpiredNotification:
		return true
	default:
		return false
	}
}

const LogFileName = "OpenIM.log"

const LocalHost = "0.0.0.0"

// flag parse.
const (
	FlagPort                  = "port"
	FlagWsPort                = "ws_port"
	FlagTransferProgressIndex = "transferProgressIndex"
	FlagPrometheusPort        = "prometheus_port"
	FlagConf                  = "config_folder_path"
)

const OpenIMCommonConfigKey = "OpenIMServerConfig"

const CallbackCommand = "command"

const BatchNum = 100

// Subscribe to user constants
const (
	SubscriberUser = 1
	Unsubscribe    = 2
)

const (
	GroupSearchPositionHead = 1
	GroupSearchPositionAny  = 2
)

const (
	FirstPageNumber   = 1
	MaxSyncPullNumber = 500
)

const (
	MsgStatusSending     = 1
	MsgStatusSendSuccess = 2
	MsgStatusSendFailed  = 3
	MsgStatusHasDeleted  = 4
	MsgStatusFiltered    = 5
)
