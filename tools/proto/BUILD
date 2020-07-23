CONFIG.setdefault("PROTOC_VERSION", "3.12.3")

protoc_binary(
    name = "protoc",
    version = CONFIG.PROTOC_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("BUF_VERSION", "0.20.4")

_buf_os = {
    "darwin": "Darwin",
    "linux": "Linux",
}

remote_file(
    name = "buf",
    url = f"https://github.com/bufbuild/buf/releases/download/v%s/buf-%s-x86_64" % (CONFIG.BUF_VERSION, _buf_os[CONFIG.HOSTOS]),
    out = "bin/buf",
    binary = True,
    visibility = ["PUBLIC"],
)
