subinclude("//proto:tools")

CONFIG.setdefault("PROTOC_VERSION", "3.13.0")

protoc_binary(
    name = "protoc",
    version = CONFIG.PROTOC_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("PROTOC_GEN_GO_VERSION", "1.25.0")

protoc_gen_go_binary(
    name = "protoc-gen-go",
    version = CONFIG.PROTOC_GEN_GO_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("PROTOC_GEN_KIT_VERSION", "0.2.0")

protoc_gen_kit_binary(
    name = "protoc-gen-kit",
    version = CONFIG.PROTOC_GEN_KIT_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("BUF_VERSION", "0.24.0")

buf_binary(
    name = "buf",
    version = CONFIG.BUF_VERSION,
    visibility = ["PUBLIC"],
)
