subinclude("//proto:tools")

CONFIG.setdefault("PROTOC_VERSION", "3.13.0")

protoc_binary(
    name = "protoc",
    version = CONFIG.PROTOC_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("BUF_VERSION", "0.20.5")

buf_binary(
    name = "buf",
    version = CONFIG.BUF_VERSION,
    visibility = ["PUBLIC"],
)
