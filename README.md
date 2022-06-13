A plugin to Drone plugin to allow users to login and authenticate with container image registries.

# Usage

The following settings changes this plugin's behavior.

* server_address (optional) The container image registry address e.g docker.io. Default to `docker.io`
* user The username to authenticate with the container image registry.
* password The password of the user to authenticate with the container image registry.

Below is an example `.drone.yml` that uses this plugin.

```yaml
kind: pipeline
name: default

steps:
- name: run quay.io/kameshsampath/drone-image-registry-auth plugin
  image: quay.io/kameshsampath/drone-image-registry-auth
  pull: if-not-exists
  settings:
    server_address:
     from_secret: registry_name
    user: 
     from_secret: registry_username
    password:
     from_secret: registry_password
```

# Building

Build the plugin binary:

```text
scripts/build.sh
```

Build the plugin image:

```text
docker build -t quay.io/kameshsampath/drone-image-registry-auth -f docker/Dockerfile .
```

# Testing

Execute the plugin from your current working directory:

```text
docker run --rm -e PLUGIN_SERVER_ADDRESS=$QUAYIO_SERVER \
  -e PLUGIN_USER=$QUAYIO_USERNAME \
  -e PLUGIN_PASSWORD=$QUAYIO_PASSWORD \
  -e DRONE_COMMIT_SHA=8f51ad7884c5eb69c11d260a31da7a745e6b78e2 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_BUILD_NUMBER=43 \
  -e DRONE_BUILD_STATUS=success \
  -w /drone/src \
  -v $(pwd):/drone/src \
  quay.io/kameshsampath/drone-image-registry-auth
```
