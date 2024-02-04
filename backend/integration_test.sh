#!/bin/bash

function check_http_code(){
    if [[ $1 -ne $2 ]]; then
        echo FAIL
        exit 1
    else
        echo OK
    fi
}

API_ADDR="http://127.0.0.1:3001"

echo "1. test whether server is accessible"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null $API_ADDR)
check_http_code code 200
echo ""

echo "2. cleaning database test.db before tests"
sqlite3 -line test.db "DELETE FROM users; DELETE FROM comments;"
echo OK
echo ""

echo "3. trying to register an user"

echo "3.1. trying to register a user without email provided"
code=$(curl --write-out "%{http_code}\n" -s $API_ADDR/user/registration -d '{"Password":"qwerrtyty"}')
check_http_code code 400
echo ""

echo "3.2. trying to register a user without password provided"
code=$(curl --write-out "%{http_code}\n" -s $API_ADDR/user/registration -d '{"Login":"test@mail.com"}')
check_http_code code 400
echo ""

echo "3.3. trying to register a user - should be CORRECT"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null -X POST $API_ADDR/user/registration -d '{"Login":"test@mail.com", "Password":"weoufnksbef"}')
check_http_code code 201
echo ""

echo "3.4. trying to register a DUPLICATE user"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null -X POST $API_ADDR/user/registration -d '{"Login":"test@mail.com", "Password":"weoufnksbef"}')
check_http_code code 400
echo ""

echo "4.1. trying to create a POST"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null -X POST $API_ADDR/create -d '{"PostText": "My second post"}')
check_http_code code 201
echo ""

echo "4.2. trying to create a POST"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null -X POST $API_ADDR/create -d '{"PostText": "My third post"}')
check_http_code code 201
echo ""

echo "4.3. trying to get post with id"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null $API_ADDR/publication/7)
check_http_code code 200
echo ""

echo "5. trying to add a comment to a post"
code=$(curl --write-out "%{http_code}\n" -s -o /dev/null -X POST $API_ADDR/publication/7/create_comment -d '{"CommentText": "My first comment", "UserId": 1}')
check_http_code code 201
echo ""

echo "7. trying to login"
curl -v "%header{Authorization}\n" $API_ADDR/user/login -d '{"Login":"test@mail.com", "Password":"weoufnksbef"}'
echo ""

