package test_utils

//
//type MockOasisAPIService struct {
//	oasisProto.UnimplementedOasisApiServiceServer
//}
//
//type MockChannelApiCallbacks struct {
//	NewMessageEvent func(context.Context, *oasisProto.OasisMessageId) (*oasisProto.OasisEmpty, error)
//}
//
//var oasisCallbacks *MockChannelApiCallbacks
//
//func SetChannelApiCallbacks(c *MockChannelApiCallbacks) {
//	oasisCallbacks = c
//}
//
//func (s MockOasisAPIService) NewMessageEvent(c context.Context, messageId *oasisProto.OasisMessageId) (*oasisProto.OasisEmpty, error) {
//	if oasisCallbacks != nil && oasisCallbacks.NewMessageEvent != nil {
//		return oasisCallbacks.NewMessageEvent(c, messageId)
//	}
//	return nil, status.Errorf(codes.Unimplemented, "method SendMessageEvent not implemented")
//}
