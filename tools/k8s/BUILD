subinclude("//k8s:tools")

CONFIG.setdefault("KUBECTL_VERSION", "1.18.6")

kubectl_binary(
    name = "kubectl",
    version = CONFIG.KUBECTL_VERSION,
    visibility = ["PUBLIC"],
)


CONFIG.setdefault("HELM_VERSION", "3.2.4")

helm_binary(
    name = "helm",
    version = CONFIG.HELM_VERSION,
    visibility = ["PUBLIC"],
)
