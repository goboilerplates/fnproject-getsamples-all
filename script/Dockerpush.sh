echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push goboilerplates/fnproject-getsamples-allbyname