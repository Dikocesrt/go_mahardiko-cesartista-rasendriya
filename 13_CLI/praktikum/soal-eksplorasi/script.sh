#!/bin/bash

echo "Simple API Client"

read -p "Insert endpoint: " endpoint
read -p "Insert HTTP method (GET/POST/PUT/DELETE): " method
read -p "Insert request body: " body
read -p "Insert request body type (e.g., application/json): " content_type

case $method in
    "GET")
        response=$(curl -s -X $method $endpoint)
        ;;
    "POST"|"PUT"|"DELETE")
        response=$(curl -s -X $method -d "$body" -H "Content-Type: $content_type" $endpoint)
        ;;
    *)
        echo "Invalid HTTP method"
        exit 1
        ;;
esac

echo "Output:"
echo $response
