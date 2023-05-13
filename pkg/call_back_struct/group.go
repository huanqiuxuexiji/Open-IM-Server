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

package call_back_struct

import (
	"Open_IM/pkg/proto/group"
	commonPb "Open_IM/pkg/proto/sdk_ws"
)

type CallbackBeforeCreateGroupReq struct {
	CallbackCommand string `json:"callbackCommand"`
	OperationID     string `json:"operationID"`
	commonPb.GroupInfo
	InitMemberList []*group.GroupAddMemberInfo `json:"initMemberList"`
}

type CallbackBeforeCreateGroupResp struct {
	*CommonCallbackResp
	GroupID           *string `json:"groupID"`
	GroupName         *string `json:"groupName"`
	Notification      *string `json:"notification"`
	Introduction      *string `json:"introduction"`
	FaceURL           *string `json:"faceURL"`
	OwnerUserID       *string `json:"ownerUserID"`
	Ex                *string `json:"ex"`
	Status            *int32  `json:"status"`
	CreatorUserID     *string `json:"creatorUserID"`
	GroupType         *int32  `json:"groupType"`
	NeedVerification  *int32  `json:"needVerification"`
	LookMemberInfo    *int32  `json:"lookMemberInfo"`
	ApplyMemberFriend *int32  `json:"applyMemberFriend"`
}

type CallbackBeforeMemberJoinGroupReq struct {
	CallbackCommand string `json:"callbackCommand"`
	OperationID     string `json:"operationID"`
	GroupID         string `json:"groupID"`
	UserID          string `json:"userID"`
	Ex              string `json:"ex"`
	GroupEx         string `json:"groupEx"`
}

type CallbackBeforeMemberJoinGroupResp struct {
	*CommonCallbackResp
	NickName    *string `json:"nickName"`
	FaceURL     *string `json:"faceURL"`
	RoleLevel   *int32  `json:"roleLevel"`
	MuteEndTime *int64  `json:"muteEndTime"`
	Ex          *string `json:"ex"`
}

type CallbackBeforeSetGroupMemberInfoReq struct {
	CallbackCommand string `json:"callbackCommand"`
	OperationID     string `json:"operationID"`
	GroupID         string `json:"groupID"`
	UserID          string `json:"userID"`
	Nickname        string `json:"nickName"`
	FaceURL         string `json:"faceURL"`
	RoleLevel       int32  `json:"roleLevel"`
	Ex              string `json:"ex"`
}

type CallbackBeforeSetGroupMemberInfoResp struct {
	*CommonCallbackResp
	Ex        *string `json:"ex"`
	Nickname  *string `json:"nickName"`
	FaceURL   *string `json:"faceURL"`
	RoleLevel *int32  `json:"roleLevel"`
}
