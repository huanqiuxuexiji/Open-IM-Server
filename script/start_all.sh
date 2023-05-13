#!/usr/bin/env bash
# Copyright © 2023 OpenIM. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#fixme This script is the total startup script
#fixme The full name of the shell script that needs to be started is placed in the need_to_start_server_shell array

#fixme Put the shell script name here
need_to_start_server_shell=(
  start_rpc_service.sh
  push_start.sh
  msg_transfer_start.sh
  sdk_svr_start.sh
  msg_gateway_start.sh
  demo_svr_start.sh
  start_cron.sh
)
time=`date +"%Y-%m-%d %H:%M:%S"`
echo "==========================================================">>../logs/openIM.log 2>&1 &
echo "==========================================================">>../logs/openIM.log 2>&1 &
echo "==========================================================">>../logs/openIM.log 2>&1 &
echo "==========server start time:${time}===========">>../logs/openIM.log 2>&1 &
echo "==========================================================">>../logs/openIM.log 2>&1 &
echo "==========================================================">>../logs/openIM.log 2>&1 &
echo "==========================================================">>../logs/openIM.log 2>&1 &

for i in ${need_to_start_server_shell[*]}; do
  chmod +x $i
  ./$i
    if [ $? -ne 0 ]; then
        exit -1
  fi
done
