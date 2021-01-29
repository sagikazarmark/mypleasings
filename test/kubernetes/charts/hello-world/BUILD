subinclude("//kubernetes")

helm_package(
    name = "hello-world",
)

helm_lint(
    name = "lint",
    chart = ":hello-world",
)

helm_e2e_test(
    name = "e2e",
    chart = ":hello-world",
    kubeconfig = "//test/kubernetes:kubeconfig",
    labels = ["kind"],
    run_tests = True,
)
