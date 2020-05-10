# Localstack S3 UI

At the moment [Localstack](https://github.com/localstack/localstack) does not offer a UI to view content stored in S3. Assuming you're running localstack via docker and have provided the correct envs and volume mappings. Localstack will store the S3 content in this file `/tmp/localstack/data/recorded_api_calls.json`.

This [docker image](https://hub.docker.com/repository/docker/gosuper/localstack-s3-ui) when used with Localstack, will parse the above file and generate a UI.

More information is available [here](./example/README.md) to get set up.