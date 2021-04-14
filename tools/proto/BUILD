subinclude("//proto:tools")

# https://github.com/protocolbuffers/protobuf/releases
CONFIG.setdefault("PROTOC_VERSION", "3.15.8")

protoc_binary(
    name = "protoc",
    version = CONFIG.PROTOC_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/protocolbuffers/protobuf-go/releases
CONFIG.setdefault("PROTOC_GEN_GO_VERSION", "1.26.0")

protoc_gen_go_binary(
    name = "protoc-gen-go",
    version = CONFIG.PROTOC_GEN_GO_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/grpc/grpc-go/releases
CONFIG.setdefault("PROTOC_GEN_GO_GRPC_VERSION", "1.1.0")

protoc_gen_go_grpc_binary(
    name = "protoc-gen-go-grpc",
    version = CONFIG.PROTOC_GEN_GO_GRPC_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/sagikazarmark/protoc-gen-kit/releases
CONFIG.setdefault("PROTOC_GEN_KIT_VERSION", "0.2.0")

protoc_gen_kit_binary(
    name = "protoc-gen-kit",
    version = CONFIG.PROTOC_GEN_KIT_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/bufbuild/buf/releases
CONFIG.setdefault("BUF_VERSION", "0.41.0")

buf_binary(
    name = "buf",
    version = CONFIG.BUF_VERSION,
    visibility = ["PUBLIC"],
)
