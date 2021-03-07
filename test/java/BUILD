subinclude("//java")

http_archive(
    name = "jre-download-darwin_amd64",
    urls = [
        "https://github.com/AdoptOpenJDK/openjdk8-binaries/releases/download/jdk8u282-b08/OpenJDK8U-jre_x64_mac_hotspot_8u282b08.tar.gz",
    ],
    strip_prefix = "jdk8u282-b08-jre/Contents/Home",
    hashes = ["9e7a40d570d5151aae23a2fb017359248f5fb82c547c3ecd860c992770228afb"],
)

genrule(
    name = "jre-darwin_amd64",
    srcs = [":jre-download-darwin_amd64"],
    outs = ["jre"],
    cmd = [
        "rm -rf $SRC/jdk8u282-b08-jre",
        "cp -r $SRC $OUT",
    ],
    visibility = ["PUBLIC"],
)

http_archive(
    name = "jre-linux_amd64",
    urls = [
        "https://github.com/AdoptOpenJDK/openjdk8-binaries/releases/download/jdk8u282-b08/OpenJDK8U-jre_x64_linux_hotspot_8u282b08.tar.gz",
    ],
    strip_prefix = "jdk8u282-b08-jre",
    hashes = ["3b2e2c6ad3ee04a58ffb8d629e3e242b0ae87b38cfd06425e4446b1f9490f521"],
)

remote_file(
    name = "openapi-generator-cli-jar",
    url = "https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.0.1/openapi-generator-cli-5.0.1.jar",
    hashes = ["e4e45d5441283b2f0f4bf988d02186b85425e7b708b4be0b06e3bfd7c7aa52c7"],
    out = "openapi-generator-cli.jar",
)

warp_jar(
    name = "openapi-generator",
    jar = ":openapi-generator-cli-jar",
    jre = ":jre-" + CONFIG.OS + "_" + CONFIG.ARCH,
)
