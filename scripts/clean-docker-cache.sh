docker rmi --force $(docker images | awk '/^<none>/ {print $3}')
