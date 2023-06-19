protoc --go_out=plugins=grpc:./auth --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/auth auth/auth.proto
protoc --go_out=plugins=grpc:./conversation --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/conversation conversation/conversation.proto
protoc --go_out=plugins=grpc:./errinfo --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/errinfo errinfo/errinfo.proto
protoc --go_out=plugins=grpc:./friend --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/friend friend/friend.proto
protoc --go_out=plugins=grpc:./group --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/group group/group.proto
protoc --go_out=plugins=grpc:./msg --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/msg msg/msg.proto
protoc --go_out=plugins=grpc:./msggateway --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/msggateway msggateway/msggateway.proto
protoc --go_out=plugins=grpc:./push --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/push push/push.proto
protoc --go_out=plugins=grpc:./sdkws --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/sdkws sdkws/sdkws.proto
protoc --go_out=plugins=grpc:./third --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/third third/third.proto
protoc --go_out=plugins=grpc:./user --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/user user/user.proto
protoc --go_out=plugins=grpc:./wrapperspb --go_opt=module=github.com/OpenIMSDK/Open-IM-Server/pkg/proto/wrapperspb wrapperspb/wrapperspb.proto
