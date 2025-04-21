# build

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Mpacket.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Mpacket.proto=github.com/jito-go \
packet.proto

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Mauth.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Mauth.proto=github.com/jito-go \
auth.proto

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Mshared.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Mshared.proto=github.com/jito-go \
shared.proto

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Mbundle.proto=github.com/jito-go \
--go_opt=Mpacket.proto=github.com/jito-go \
--go_opt=Mshared.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Mbundle.proto=github.com/jito-go \
--go-grpc_opt=Mpacket.proto=github.com/jito-go \
--go-grpc_opt=Mshared.proto=github.com/jito-go \
bundle.proto

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Mblock.proto=github.com/jito-go \
--go_opt=Mshared.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Mblock.proto=github.com/jito-go \
--go-grpc_opt=Mshared.proto=github.com/jito-go \
block.proto

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Mblock_engine.proto=github.com/jito-go \
--go_opt=Mpacket.proto=github.com/jito-go \
--go_opt=Mshared.proto=github.com/jito-go \
--go_opt=Mbundle.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Mblock_engine.proto=github.com/jito-go \
--go-grpc_opt=Mpacket.proto=github.com/jito-go \
--go-grpc_opt=Mshared.proto=github.com/jito-go \
--go-grpc_opt=Mbundle.proto=github.com/jito-go \
block_engine.proto

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Msearcher.proto=github.com/jito-go \
--go_opt=Mblock_engine.proto=github.com/jito-go \
--go_opt=Mpacket.proto=github.com/jito-go \
--go_opt=Mshared.proto=github.com/jito-go \
--go_opt=Mbundle.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Msearcher.proto=github.com/jito-go \
--go-grpc_opt=Mblock_engine.proto=github.com/jito-go \
--go-grpc_opt=Mpacket.proto=github.com/jito-go \
--go-grpc_opt=Mshared.proto=github.com/jito-go \
--go-grpc_opt=Mbundle.proto=github.com/jito-go \
searcher.proto

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Mshredstream.proto=github.com/jito-go \
--go_opt=Mblock_engine.proto=github.com/jito-go \
--go_opt=Mpacket.proto=github.com/jito-go \
--go_opt=Mshared.proto=github.com/jito-go \
--go_opt=Mbundle.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Mshredstream.proto=github.com/jito-go \
--go-grpc_opt=Mblock_engine.proto=github.com/jito-go \
--go-grpc_opt=Mpacket.proto=github.com/jito-go \
--go-grpc_opt=Mshared.proto=github.com/jito-go \
--go-grpc_opt=Mbundle.proto=github.com/jito-go \
shredstream.proto

protoc --proto_path=./mev-protos \
--go_out=. \
--go_opt=Mrelayer.proto=github.com/jito-go \
--go_opt=Mblock_engine.proto=github.com/jito-go \
--go_opt=Mpacket.proto=github.com/jito-go \
--go_opt=Mshared.proto=github.com/jito-go \
--go_opt=Mbundle.proto=github.com/jito-go \
--go-grpc_out=. \
--go-grpc_opt=Mrelayer.proto=github.com/jito-go \
--go-grpc_opt=Mblock_engine.proto=github.com/jito-go \
--go-grpc_opt=Mpacket.proto=github.com/jito-go \
--go-grpc_opt=Mshared.proto=github.com/jito-go \
--go-grpc_opt=Mbundle.proto=github.com/jito-go \
relayer.proto
