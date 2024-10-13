#!/bin/bash

# Function to display response
function display_response() {
    echo "===================================="
    echo "Response from API:"
    echo "$response"
    echo "===================================="
}

# Input from user
echo "Masukkan API endpoint (contoh: https://jsonplaceholder.typicode.com/posts):"
read endpoint

echo "Masukkan HTTP method (GET, POST, PUT, DELETE):"
read method

echo "Masukkan request body (jika tidak ada, tekan enter):"
read request_body

echo "Masukkan tipe request body (application/json, application/x-www-form-urlencoded, dll.):"
read content_type

# Check if request body is empty
if [ -z "$request_body" ]; then
    if [ "$method" == "GET" ] || [ "$method" == "DELETE" ]; then
        response=$(curl -s -X "$method" "$endpoint")
    else
        response=$(curl -s -X "$method" -H "Content-Type: $content_type" "$endpoint")
    fi
else
    response=$(curl -s -X "$method" -H "Content-Type: $content_type" -d "$request_body" "$endpoint")
fi

# Display response
display_response
