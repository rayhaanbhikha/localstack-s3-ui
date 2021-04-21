## Initial set up

This example will start localstack and the localstack-s3-ui in seperate docker container using the following docker-compose.yml file.

```yml
  localstack:
    container_name: localstack
    image: localstack/localstack
    ports:
      - "4566-4599:4566-4599"
      - "8080:8080"
    environment:
      - SERVICES=s3
      - DEBUG=1
      - USE_LIGHT_IMAGE=1
      - DATA_DIR=/tmp/localstack/data
      - PORT_WEB_UI=8080
      - DOCKER_HOST=unix:///var/run/docker.sock
    volumes:
      - "./.localstack:/tmp/localstack"
      - "/var/run/docker.sock:/var/run/docker.sock"

  localstack-s3-ui:
    container_name: localstack-s3-ui
    depends_on: 
      - localstack
    image: gosuper/localstack-s3-ui
    ports:
      - 9000:9000
    environment:
      - API_REQ_FILE_PATH=/tmp/localstack/data/recorded_api_calls.json #Required
      - PORT=9000 # Defaults to 9000
    volumes:
      - "./.localstack:/tmp/localstack"
```

Ensure that you provide the `DATA_DIR` env variable for the `localstack container` so S3 data persists. In the example above, the `.localstack` directory on the host machine is mapped to `/tmp/localstack` on the container and Localstack will store all S3 requests here `/tmp/localstack/data/recorded_api_calls.json`.

The localstack-s3-ui container requires `API_REQ_FILE_PATH` the env variable to parse the Localstack S3 api requests. You can also provide a PORT env for the Web user interface / dashboard (default: 9000).

In the above example both containers also have the same volume mappings to allow both containers to read from the same local data directory.

More information on configuration settings for the Localstack container can be found on the official docs.

## Starting Locally

First start both the containers using docker-compose:

```sh
docker-compose up --build
```

The localstack-s3-ui container will wait for `.localstack/data/recorded_api_calls.json` file to exist, before starting the web server.

Once both containers are running you can run the init script. This will create a bucket and copy the web-app directory into Localstack S3.

```sh
./init.sh
```

You should now be able to view the contents of the S3 bucket on `http://localhost:9000/s3`

The web server will automatically detect changes in S3 and update the UI accordingly.
