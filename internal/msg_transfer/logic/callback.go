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

package logic

import (
	cbApi "Open_IM/pkg/call_back_struct"
	"Open_IM/pkg/common/callback"
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/constant"
	"Open_IM/pkg/common/http"
	"Open_IM/pkg/common/log"
	pbChat "Open_IM/pkg/proto/msg"
	"Open_IM/pkg/utils"
	http2 "net/http"
)

func callbackAfterConsumeGroupMsg(msg []*pbChat.MsgDataToMQ, triggerID string) cbApi.CommonCallbackResp {
	callbackResp := cbApi.CommonCallbackResp{OperationID: triggerID}
	if !config.Config.Callback.CallbackAfterConsumeGroupMsg.Enable {
		return callbackResp
	}
	for _, v := range msg {
		if v.MsgData.SessionType == constant.SuperGroupChatType || v.MsgData.SessionType == constant.GroupChatType {
			commonCallbackReq := copyCallbackCommonReqStruct(v)
			commonCallbackReq.CallbackCommand = constant.CallbackAfterConsumeGroupMsgCommand
			req := cbApi.CallbackAfterConsumeGroupMsgReq{
				CommonCallbackReq: commonCallbackReq,
				GroupID:           v.MsgData.GroupID,
			}
			resp := &cbApi.CallbackAfterConsumeGroupMsgResp{CommonCallbackResp: &callbackResp}
			defer log.NewDebug(triggerID, utils.GetSelfFuncName(), req, *resp)
			if err := http.CallBackPostReturn(config.Config.Callback.CallbackUrl, constant.CallbackAfterConsumeGroupMsgCommand, req, resp, config.Config.Callback.CallbackAfterConsumeGroupMsg.CallbackTimeOut); err != nil {
				callbackResp.ErrCode = http2.StatusInternalServerError
				callbackResp.ErrMsg = err.Error()
				return callbackResp
			}
		}
	}

	log.NewDebug(triggerID, utils.GetSelfFuncName(), msg)

	return callbackResp
}
func copyCallbackCommonReqStruct(msg *pbChat.MsgDataToMQ) cbApi.CommonCallbackReq {
	req := cbApi.CommonCallbackReq{
		SendID:           msg.MsgData.SendID,
		ServerMsgID:      msg.MsgData.ServerMsgID,
		ClientMsgID:      msg.MsgData.ClientMsgID,
		OperationID:      msg.OperationID,
		SenderPlatformID: msg.MsgData.SenderPlatformID,
		SenderNickname:   msg.MsgData.SenderNickname,
		SessionType:      msg.MsgData.SessionType,
		MsgFrom:          msg.MsgData.MsgFrom,
		ContentType:      msg.MsgData.ContentType,
		Status:           msg.MsgData.Status,
		CreateTime:       msg.MsgData.CreateTime,
		AtUserIDList:     msg.MsgData.AtUserIDList,
		SenderFaceURL:    msg.MsgData.SenderFaceURL,
		Content:          callback.GetContent(msg.MsgData),
		Seq:              msg.MsgData.Seq,
		Ex:               msg.MsgData.Ex,
	}
	return req
}
