CONFIG.setdefault("HELM_VERSION", "3.2.4")

genrule(
    name = "helm",
    srcs = [remote_file(
        name = "helm",
        _tag = "download",
        url = f"https://get.helm.sh/helm-v{CONFIG.HELM_VERSION}-{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}.tar.gz",
        out = f"helm-{CONFIG.HELM_VERSION}.tar.gz",
    )],
    outs = ["helm"],
    binary = True,
    cmd = '"$TOOL" x "$SRCS" helm -o tmp && mv $(find tmp -name "helm") "$OUT"',
    tools = [CONFIG.JARCAT_TOOL],
    visibility = ["PUBLIC"],
)
