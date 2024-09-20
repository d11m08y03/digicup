#!/bin/bash

URL="http://localhost:8080/users/login"

EMAIL="joe1@example.com"
PASSWORD="123456"

curl -X POST \
  $URL \
  -H 'Content-Type: application/json' \
  -d '{"email": "'$EMAIL'", "password": "'$PASSWORD'"}'
