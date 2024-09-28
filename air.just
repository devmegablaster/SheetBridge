all: api

api:
  @echo "Running Api..."
  air -c "cmd/api/.air.toml"
