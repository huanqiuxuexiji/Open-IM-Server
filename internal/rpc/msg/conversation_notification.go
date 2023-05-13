// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/constant"
	"Open_IM/pkg/common/log"
	open_im_sdk "Open_IM/pkg/proto/sdk_ws"
	"Open_IM/pkg/utils"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func SetConversationNotification(operationID, sendID, recvID string, contentType int, m proto.Message, tips open_im_sdk.TipsComm) {
	log.NewInfo(operationID, "args: ", sendID, recvID, contentType, m.String(), tips.String())
	var err error
	tips.Detail, err = proto.Marshal(m)
	if err != nil {
		log.NewError(operationID, "Marshal failed ", err.Error(), m.String())
		return
	}
	marshaler := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: false,
	}
	tips.JsonDetail, _ = marshaler.MarshalToString(m)
	var n NotificationMsg
	n.SendID = sendID
	n.RecvID = recvID
	n.ContentType = int32(contentType)
	n.SessionType = constant.SingleChatType
	n.MsgFrom = constant.SysMsgType
	n.OperationID = operationID
	n.Content, err = proto.Marshal(&tips)
	if err != nil {
		log.Error(operationID, utils.GetSelfFuncName(), "Marshal failed ", err.Error(), tips.String())
		return
	}
	Notification(&n)
}

// SetPrivate调用
func ConversationSetPrivateNotification(operationID, sendID, recvID string, isPrivateChat bool) {
	log.NewInfo(operationID, utils.GetSelfFuncName())
	conversationSetPrivateTips := &open_im_sdk.ConversationSetPrivateTips{
		RecvID:    recvID,
		SendID:    sendID,
		IsPrivate: isPrivateChat,
	}
	var tips open_im_sdk.TipsComm
	var tipsMsg string
	if isPrivateChat == true {
		tipsMsg = config.Config.Notification.ConversationSetPrivate.DefaultTips.OpenTips
	} else {
		tipsMsg = config.Config.Notification.ConversationSetPrivate.DefaultTips.CloseTips
	}
	tips.DefaultTips = tipsMsg
	SetConversationNotification(operationID, sendID, recvID, constant.ConversationPrivateChatNotification, conversationSetPrivateTips, tips)
}

// 会话改变
func ConversationChangeNotification(operationID, userID string) {
	log.NewInfo(operationID, utils.GetSelfFuncName())
	ConversationChangedTips := &open_im_sdk.ConversationUpdateTips{
		UserID: userID,
	}
	var tips open_im_sdk.TipsComm
	tips.DefaultTips = config.Config.Notification.ConversationOptUpdate.DefaultTips.Tips
	SetConversationNotification(operationID, userID, userID, constant.ConversationOptChangeNotification, ConversationChangedTips, tips)
}

//会话未读数同步
func ConversationUnreadChangeNotification(operationID, userID, conversationID string, updateUnreadCountTime int64) {
	log.NewInfo(operationID, utils.GetSelfFuncName())
	ConversationChangedTips := &open_im_sdk.ConversationUpdateTips{
		UserID:                userID,
		ConversationIDList:    []string{conversationID},
		UpdateUnreadCountTime: updateUnreadCountTime,
	}
	var tips open_im_sdk.TipsComm
	tips.DefaultTips = config.Config.Notification.ConversationOptUpdate.DefaultTips.Tips
	SetConversationNotification(operationID, userID, userID, constant.ConversationUnreadNotification, ConversationChangedTips, tips)
}
