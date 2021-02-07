subinclude("//misc")

sed(
    name = "sed",
    src = "testdata/sedinput",
    out = "sedoutput",
    command = "s/TEST/OK/g",
)

diff_test(
    name = "sed_test",
    actual = ":sed",
    expected = "testdata/sedoutput",
)

sha256sum(
    name = "sha256sum",
    srcs = [
        "testdata/sample1.yaml",
        "testdata/sample2.yaml",
    ],
    out = "sha256sum_checksums.txt",
)

diff_test(
    name = "sha256sum_test",
    actual = ":sha256sum",
    expected = "testdata/sha256sumoutput",
)

build_artifacts(
    name = "artifacts",
    artifacts = [
        "testdata/sample1.yaml",
        "testdata/sample2.yaml",
    ],
)
