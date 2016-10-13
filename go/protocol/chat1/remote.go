// Auto-generated by avdl-compiler v1.3.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/remote.avdl

package chat1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type MessageBoxed struct {
	ServerHeader     *MessageServerHeader `codec:"serverHeader,omitempty" json:"serverHeader,omitempty"`
	ClientHeader     MessageClientHeader  `codec:"clientHeader" json:"clientHeader"`
	HeaderCiphertext EncryptedData        `codec:"headerCiphertext" json:"headerCiphertext"`
	BodyCiphertext   EncryptedData        `codec:"bodyCiphertext" json:"bodyCiphertext"`
	KeyGeneration    int                  `codec:"keyGeneration" json:"keyGeneration"`
}

type ThreadViewBoxed struct {
	Messages   []MessageBoxed `codec:"messages" json:"messages"`
	Pagination *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetInboxRemoteRes struct {
	Inbox     InboxView  `codec:"inbox" json:"inbox"`
	RateLimit *RateLimit `codec:"rateLimit,omitempty" json:"rateLimit,omitempty"`
}

type GetInboxByTLFIDRemoteRes struct {
	Convs     []Conversation `codec:"convs" json:"convs"`
	RateLimit *RateLimit     `codec:"rateLimit,omitempty" json:"rateLimit,omitempty"`
}

type GetThreadRemoteRes struct {
	Thread    ThreadViewBoxed `codec:"thread" json:"thread"`
	RateLimit *RateLimit      `codec:"rateLimit,omitempty" json:"rateLimit,omitempty"`
}

type GetConversationMetadataRemoteRes struct {
	Conv      Conversation `codec:"conv" json:"conv"`
	RateLimit *RateLimit   `codec:"rateLimit,omitempty" json:"rateLimit,omitempty"`
}

type PostRemoteRes struct {
	MsgHeader MessageServerHeader `codec:"msgHeader" json:"msgHeader"`
	RateLimit *RateLimit          `codec:"rateLimit,omitempty" json:"rateLimit,omitempty"`
}

type NewConversationRemoteRes struct {
	ConvID    ConversationID `codec:"convID" json:"convID"`
	RateLimit *RateLimit     `codec:"rateLimit,omitempty" json:"rateLimit,omitempty"`
}

type GetMessagesRemoteRes struct {
	Msgs      []MessageBoxed `codec:"msgs" json:"msgs"`
	RateLimit *RateLimit     `codec:"rateLimit,omitempty" json:"rateLimit,omitempty"`
}

type MarkAsReadRes struct {
	RateLimit *RateLimit `codec:"rateLimit,omitempty" json:"rateLimit,omitempty"`
}

type S3AttachmentParams struct {
	Endpoint   string `codec:"endpoint" json:"endpoint"`
	Bucket     string `codec:"bucket" json:"bucket"`
	ObjectKey  string `codec:"objectKey" json:"objectKey"`
	Acl        string `codec:"acl" json:"acl"`
	Credential string `codec:"credential" json:"credential"`
	Algorithm  string `codec:"algorithm" json:"algorithm"`
	Date       string `codec:"date" json:"date"`
	Policy     string `codec:"policy" json:"policy"`
	Signature  string `codec:"signature" json:"signature"`
}

type S3Params struct {
	Bucket               string `codec:"bucket" json:"bucket"`
	ObjectKey            string `codec:"objectKey" json:"objectKey"`
	AccessKey            string `codec:"accessKey" json:"accessKey"`
	Acl                  string `codec:"acl" json:"acl"`
	RegionName           string `codec:"regionName" json:"regionName"`
	RegionEndpoint       string `codec:"regionEndpoint" json:"regionEndpoint"`
	RegionBucketEndpoint string `codec:"regionBucketEndpoint" json:"regionBucketEndpoint"`
}

type GetInboxRemoteArg struct {
	Query      *GetInboxQuery `codec:"query,omitempty" json:"query,omitempty"`
	Pagination *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetThreadRemoteArg struct {
	ConversationID ConversationID  `codec:"conversationID" json:"conversationID"`
	Query          *GetThreadQuery `codec:"query,omitempty" json:"query,omitempty"`
	Pagination     *Pagination     `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type PostRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MessageBoxed   MessageBoxed   `codec:"messageBoxed" json:"messageBoxed"`
}

type NewConversationRemoteArg struct {
	IdTriple ConversationIDTriple `codec:"idTriple" json:"idTriple"`
}

type NewConversationRemote2Arg struct {
	IdTriple   ConversationIDTriple `codec:"idTriple" json:"idTriple"`
	TLFMessage MessageBoxed         `codec:"TLFMessage" json:"TLFMessage"`
}

type GetMessagesRemoteArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MessageIDs     []MessageID    `codec:"messageIDs" json:"messageIDs"`
}

type MarkAsReadArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MsgID          MessageID      `codec:"msgID" json:"msgID"`
}

type TlfFinalizeArg struct {
	TlfID TLFID `codec:"tlfID" json:"tlfID"`
}

type GetS3AttachmentParamsArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
}

type GetS3ParamsArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
}

type S3SignArg struct {
	Payload []byte `codec:"payload" json:"payload"`
}

type RemoteInterface interface {
	GetInboxRemote(context.Context, GetInboxRemoteArg) (GetInboxRemoteRes, error)
	GetThreadRemote(context.Context, GetThreadRemoteArg) (GetThreadRemoteRes, error)
	PostRemote(context.Context, PostRemoteArg) (PostRemoteRes, error)
	NewConversationRemote(context.Context, ConversationIDTriple) (NewConversationRemoteRes, error)
	NewConversationRemote2(context.Context, NewConversationRemote2Arg) (NewConversationRemoteRes, error)
	GetMessagesRemote(context.Context, GetMessagesRemoteArg) (GetMessagesRemoteRes, error)
	MarkAsRead(context.Context, MarkAsReadArg) (MarkAsReadRes, error)
	TlfFinalize(context.Context, TLFID) error
	GetS3AttachmentParams(context.Context, ConversationID) (S3AttachmentParams, error)
	GetS3Params(context.Context, ConversationID) (S3Params, error)
	S3Sign(context.Context, []byte) ([]byte, error)
}

func RemoteProtocol(i RemoteInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.remote",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getInboxRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetInboxRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getThreadRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetThreadRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetThreadRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetThreadRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetThreadRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"postRemote": {
				MakeArg: func() interface{} {
					ret := make([]PostRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PostRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]PostRemoteArg)(nil), args)
						return
					}
					ret, err = i.PostRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationRemote": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationRemoteArg)(nil), args)
						return
					}
					ret, err = i.NewConversationRemote(ctx, (*typedArgs)[0].IdTriple)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationRemote2": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationRemote2Arg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationRemote2Arg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationRemote2Arg)(nil), args)
						return
					}
					ret, err = i.NewConversationRemote2(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getMessagesRemote": {
				MakeArg: func() interface{} {
					ret := make([]GetMessagesRemoteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetMessagesRemoteArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetMessagesRemoteArg)(nil), args)
						return
					}
					ret, err = i.GetMessagesRemote(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"markAsRead": {
				MakeArg: func() interface{} {
					ret := make([]MarkAsReadArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]MarkAsReadArg)
					if !ok {
						err = rpc.NewTypeError((*[]MarkAsReadArg)(nil), args)
						return
					}
					ret, err = i.MarkAsRead(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"tlfFinalize": {
				MakeArg: func() interface{} {
					ret := make([]TlfFinalizeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TlfFinalizeArg)
					if !ok {
						err = rpc.NewTypeError((*[]TlfFinalizeArg)(nil), args)
						return
					}
					err = i.TlfFinalize(ctx, (*typedArgs)[0].TlfID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getS3AttachmentParams": {
				MakeArg: func() interface{} {
					ret := make([]GetS3AttachmentParamsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetS3AttachmentParamsArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetS3AttachmentParamsArg)(nil), args)
						return
					}
					ret, err = i.GetS3AttachmentParams(ctx, (*typedArgs)[0].ConversationID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getS3Params": {
				MakeArg: func() interface{} {
					ret := make([]GetS3ParamsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetS3ParamsArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetS3ParamsArg)(nil), args)
						return
					}
					ret, err = i.GetS3Params(ctx, (*typedArgs)[0].ConversationID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"s3Sign": {
				MakeArg: func() interface{} {
					ret := make([]S3SignArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]S3SignArg)
					if !ok {
						err = rpc.NewTypeError((*[]S3SignArg)(nil), args)
						return
					}
					ret, err = i.S3Sign(ctx, (*typedArgs)[0].Payload)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RemoteClient struct {
	Cli rpc.GenericClient
}

func (c RemoteClient) GetInboxRemote(ctx context.Context, __arg GetInboxRemoteArg) (res GetInboxRemoteRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.getInboxRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetThreadRemote(ctx context.Context, __arg GetThreadRemoteArg) (res GetThreadRemoteRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.getThreadRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) PostRemote(ctx context.Context, __arg PostRemoteArg) (res PostRemoteRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.postRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) NewConversationRemote(ctx context.Context, idTriple ConversationIDTriple) (res NewConversationRemoteRes, err error) {
	__arg := NewConversationRemoteArg{IdTriple: idTriple}
	err = c.Cli.Call(ctx, "chat.1.remote.newConversationRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) NewConversationRemote2(ctx context.Context, __arg NewConversationRemote2Arg) (res NewConversationRemoteRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.newConversationRemote2", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetMessagesRemote(ctx context.Context, __arg GetMessagesRemoteArg) (res GetMessagesRemoteRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.getMessagesRemote", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) MarkAsRead(ctx context.Context, __arg MarkAsReadArg) (res MarkAsReadRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.remote.markAsRead", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) TlfFinalize(ctx context.Context, tlfID TLFID) (err error) {
	__arg := TlfFinalizeArg{TlfID: tlfID}
	err = c.Cli.Call(ctx, "chat.1.remote.tlfFinalize", []interface{}{__arg}, nil)
	return
}

func (c RemoteClient) GetS3AttachmentParams(ctx context.Context, conversationID ConversationID) (res S3AttachmentParams, err error) {
	__arg := GetS3AttachmentParamsArg{ConversationID: conversationID}
	err = c.Cli.Call(ctx, "chat.1.remote.getS3AttachmentParams", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) GetS3Params(ctx context.Context, conversationID ConversationID) (res S3Params, err error) {
	__arg := GetS3ParamsArg{ConversationID: conversationID}
	err = c.Cli.Call(ctx, "chat.1.remote.getS3Params", []interface{}{__arg}, &res)
	return
}

func (c RemoteClient) S3Sign(ctx context.Context, payload []byte) (res []byte, err error) {
	__arg := S3SignArg{Payload: payload}
	err = c.Cli.Call(ctx, "chat.1.remote.s3Sign", []interface{}{__arg}, &res)
	return
}
