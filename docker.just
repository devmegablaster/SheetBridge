up:
  docker-compose up -d

buildapi tag:
  @echo "Building docker image with {{tag}} tag..."
  docker build -t tag . -f Dockerfile.api

dbclean:
  @echo "Cleaning database..."
  docker-compose down -v db
  docker-compose up -d db

kafka:
  @echo "Starting kafka..."
  docker-compose up -d zookeeper kafka
