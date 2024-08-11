#!/bin/bash

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Please provide an argument: docker or local"
    exit 1
fi

# Set the host based on the argument
if [ "$1" = "docker" ]; then
    DB_HOST="postgres"
    RABBIT_HOST="rabbitmq"
elif [ "$1" = "local" ]; then
    DB_HOST="localhost"
    RABBIT_HOST="localhost"
else
    echo "Invalid argument. Please use 'docker' or 'local'"
    exit 1
fi

# Generate the YAML content
cat << EOF
app_config:
  app_name: "technical-test"
  port: 8005
  debug_mode: true
  url: "http://localhost"
  public_dir: "uploads"

database_config:
  host: "${DB_HOST}"
  port: 5432
  user: "technical_test"
  password: "Technical!"
  database: "technical_test"
  usage_sql: "pgsql"
  max_idle_conn: 10
  max_open_conn: 100
  max_life_time_conn: 3600

rabbit_config:
  user: "technical_test"
  password: "Technical!"
  host: "${RABBIT_HOST}"
  port: 5672
EOF