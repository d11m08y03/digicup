#!/bin/bash

TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvZTFAZXhhbXBsZS5jb20iLCJleHAiOjE3Mjk1MzE5MDB9.W0sXUy3zBqEJ8WaR0ZCzaH9PMep8-jg1F96t62eOsLs"

curl -X GET http://localhost:8080/users/getUsers \
-H "Authorization: Bearer $TOKEN"
