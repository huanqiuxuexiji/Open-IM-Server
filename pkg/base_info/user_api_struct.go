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

package base_info

import (
	open_im_sdk "Open_IM/pkg/proto/sdk_ws"
)

type GetUsersInfoReq struct {
	OperationID string   `json:"operationID" binding:"required"`
	UserIDList  []string `json:"userIDList" binding:"required"`
}
type GetUsersInfoResp struct {
	CommResp
	UserInfoList []*open_im_sdk.PublicUserInfo `json:"-"`
	Data         []map[string]interface{}      `json:"data" swaggerignore:"true"`
}

type UpdateSelfUserInfoReq struct {
	ApiUserInfo
	OperationID string `json:"operationID" binding:"required"`
}
type SetGlobalRecvMessageOptReq struct {
	OperationID      string `json:"operationID" binding:"required"`
	GlobalRecvMsgOpt *int32 `json:"globalRecvMsgOpt" binding:"omitempty,oneof=0 1 2"`
}
type SetGlobalRecvMessageOptResp struct {
	CommResp
}
type UpdateUserInfoResp struct {
	CommResp
}

type GetSelfUserInfoReq struct {
	OperationID string `json:"operationID" binding:"required"`
	UserID      string `json:"userID" binding:"required"`
}
type GetSelfUserInfoResp struct {
	CommResp
	UserInfo *open_im_sdk.UserInfo  `json:"-"`
	Data     map[string]interface{} `json:"data" swaggerignore:"true"`
}

type GetFriendIDListFromCacheReq struct {
	OperationID string `json:"operationID" binding:"required"`
}

type GetFriendIDListFromCacheResp struct {
	CommResp
	UserIDList []string `json:"userIDList" binding:"required"`
}

type GetBlackIDListFromCacheReq struct {
	OperationID string `json:"operationID" binding:"required"`
}

type GetBlackIDListFromCacheResp struct {
	CommResp
	UserIDList []string `json:"userIDList" binding:"required"`
}
