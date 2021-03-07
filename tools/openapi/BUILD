subinclude("//openapi:tools")

# https://github.com/openapitools/openapi-generator/releases
CONFIG.setdefault("OPENAPI_GENERATOR_IMAGE_NAME", "openapitools/openapi-generator-cli")
CONFIG.setdefault("OPENAPI_GENERATOR_IMAGE_VERSION", "v5.0.1")

openapi_generator_docker(
    name = "openapi-generator-docker",
    image_name = CONFIG.OPENAPI_GENERATOR_IMAGE_NAME,
    image_version = CONFIG.OPENAPI_GENERATOR_IMAGE_VERSION,
    visibility = ["PUBLIC"],
)

# https://github.com/stoplightio/spectral/releases
CONFIG.setdefault("SPECTRAL_VERSION", "5.8.1")

spectral_binary(
    name = "spectral",
    version = CONFIG.SPECTRAL_VERSION,
    visibility = ["PUBLIC"],
)
