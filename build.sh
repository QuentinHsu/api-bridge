USERNAME=quentinhsu
IMAGE=api-bridge

docker build -t $USERNAME/$IMAGE:latest .
# delete dangling images
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)

# print the result
echo "Docker image built and temporary images removed."