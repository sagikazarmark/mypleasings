subinclude("//k8s:tools")

CONFIG.setdefault("KUBECTL_VERSION", "1.18.8")

kubectl_binary(
    name = "kubectl",
    version = CONFIG.KUBECTL_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("KIND_VERSION", "0.8.1")

kind_binary(
    name = "kind",
    version = CONFIG.KIND_VERSION,
    visibility = ["PUBLIC"],
)

CONFIG.setdefault("HELM_VERSION", "3.3.0")

helm_binary(
    name = "helm",
    version = CONFIG.HELM_VERSION,
    visibility = ["PUBLIC"],
)
