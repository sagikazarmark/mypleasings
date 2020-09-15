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

    gentest(
        name = f"yaml_merge{i}_test",
        data = {
            "actual": [f":yaml_merge{i}"],
            "expected": [f"testdata/merged{i}.yaml"],
        },
        test_cmd = 'diff "$DATA_EXPECTED" "$DATA_ACTUAL"',
        no_test_output = True,
    )

    i += 1

sed(
    name = "sed",
    src = "testdata/sedinput",
    out = "sedoutput",
    command = "s/TEST/OK/g",
)

gentest(
    name = "sed_test",
    data = {
        "actual": [":sed"],
        "expected": ["testdata/sedoutput"],
    },
    test_cmd = 'diff "$DATA_EXPECTED" "$DATA_ACTUAL"',
    no_test_output = True,
)
