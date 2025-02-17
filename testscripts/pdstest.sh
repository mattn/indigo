#!/bin/bash
set -e
set -x 


echo "1. Creating Account"
./gosky --pds-host="http://localhost:4989" newAccount test@foo.com testman.test password > test.auth

echo "2. Some Content"
./gosky --pds-host="http://localhost:4989" --auth-file="test.auth" post "cats are really cool and the best"
./gosky --pds-host="http://localhost:4989" --auth-file="test.auth" post "paul frazee needs to buy a sweater"

echo "3. View That Content"
./gosky --pds-host="http://localhost:4989" --auth-file="test.auth" feed --raw --author=self


echo "4. Make a second account"
./gosky --pds-host="http://localhost:4989" newAccount test2@foo.com friendbot.test password > test2.auth

echo "5. Post on second account"
./gosky --pds-host="http://localhost:4989" --auth-file="test2.auth" post "Im a big fan of the snow"

echo "6. Upvote content"
posturi=$(./gosky --pds-host=http://localhost:4989 --auth-file=test.auth feed --raw --author=self | jq -r .post.uri | head -n1)
./gosky --pds-host="http://localhost:4989" --auth-file="test2.auth" vote $posturi up

echo "7. Check notifications"
./gosky --pds-host="http://localhost:4989" --auth-file="test.auth" notifs

echo "8. Follow"
./gosky --pds-host="http://localhost:4989" --auth-file="test2.auth" follows add $(cat test.auth | jq -r .did)

echo "9. Check notifications"
./gosky --pds-host="http://localhost:4989" --auth-file="test.auth" notifs



echo "Success!"
