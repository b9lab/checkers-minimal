package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	checkersv1 "github.com/alice/checkers/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: nil,
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: checkersv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateGame",
					Use:       "create index black red",
					Short:     "Creates a new checkers game at the index for the black and red players",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"},
						{ProtoField: "black"},
						{ProtoField: "red"},
					},
				},
			},
		},
	}
}
