subinclude("//misc")

yaml_merge_test_cases = [
    {},
    {
        "overwrite": True,
    },
    {
        "append": True,
    },
    {
        "overwrite": True,
        "append": True,
    },
    {
        "pretty": True,
    },
    {
        "indent": 4,
        "pretty": True,
    },
    {
        "json": True,
    },
    {
        "json": True,
        "pretty": True,
    },
    {
        "indent": 4,
        "json": True,
        "pretty": True,
    },
]

i = 1

for test_case in yaml_merge_test_cases:
    yaml_merge(
        name = f"yaml_merge{i}",
        srcs = ["testdata/sample1.yaml", "testdata/sample2.yaml"],
        out = f"merged{i}.yaml",
        overwrite = test_case.get("overwrite", False),
        append = test_case.get("append", False),
        indent = test_case.get("indent", None),
        json = test_case.get("json", False),
        pretty = test_case.get("pretty", False),
    )

    diff_test(
        name = f"yaml_merge{i}_test",
        expected = f"testdata/merged{i}.yaml",
        actual = f":yaml_merge{i}",
    )

    i += 1

sed(
    name = "sed",
    src = "testdata/sedinput",
    out = "sedoutput",
    command = "s/TEST/OK/g",
)

diff_test(
    name = "sed_test",
    expected = "testdata/sedoutput",
    actual = ":sed",
)

sha256sum(
    name = "sha256sum",
    srcs = ["testdata/sample1.yaml", "testdata/sample2.yaml"],
    out = "checksums.txt",
)

diff_test(
    name = "sha256sum_test",
    expected = "testdata/sha256sumoutput",
    actual = ":sha256sum",
)
