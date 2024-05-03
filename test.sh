#!/bin/bash

BASE_URL="http://localhost:8080"
USER_ID="user123"
USER_EMAIL="user@example.com"

test_sign_up() {
    echo "Testing Sign-Up..."
    RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" -X POST -d "{\"id\":\"$USER_ID\", \"email\":\"$USER_EMAIL\"}" -H "Content-Type: application/json" "$BASE_URL/signup")
    if [ "$RESPONSE" -eq 200 ]; then
        echo "Sign-Up Test Passed"
    else
        echo "Sign-Up Test Failed: Status Code $RESPONSE"
    fi
}

test_sign_in() {
    echo "Testing Sign-In..."
    RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" -X POST -d "{\"id\":\"$USER_ID\"}" -H "Content-Type: application/json" "$BASE_URL/signin")
    if [ "$RESPONSE" -eq 200 ]; then
        echo "Sign-In Test Passed"
    else
        echo "Sign-In Test Failed: Status Code $RESPONSE"
    fi
}

test_update_points() {
    echo "Testing Points Update..."
    RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" -X POST -d "{\"points\":1000}" -H "Content-Type: application/json" "$BASE_URL/points?id=$USER_ID")
    if [ "$RESPONSE" -eq 200 ]; then
        echo "Points Update Test Passed"
    else
        echo "Points Update Test Failed: Status Code $RESPONSE"
    fi
}

# Function to check user points
test_check_points() {
    echo "Testing Points Checking..."
    RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/points?id=$USER_ID")
    if [ "$RESPONSE" -eq 200 ]; then
        echo "Points Checking Test Passed"
    else
        echo "Points Checking Test Failed: Status Code $RESPONSE"
    fi
}

test_sign_up
test_sign_in
test_update_points
test_check_points
