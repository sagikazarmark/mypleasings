subinclude("//k8s")

CONFIG.setdefault("KIND_CLUSTER_NAME", "kind")

kind_kubeconfig(
    name = "kubeconfig",
    cluster_name = CONFIG.KIND_CLUSTER_NAME,
    labels = ["manual"],
    visibility = ["//test/k8s/..."],
)
