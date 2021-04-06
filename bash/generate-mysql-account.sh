#!/bin/bash

host="%"
maxUserConnections=10
privileges="all privileges"
accountPrefix=""
databasePrefix=""

rows=(
# account,db1 privileges,db2 privileges,...
andarbahar-backend:andarbahar
andarbahar-backend-api:andarbahar,bac
andarbahar-clients-service:andarbahar
baccarat-backend:baccarat
baccarat-backend-api:baccarat,bac
baccarat-clients-service:baccarat
bbin-money:bac
bbserver-env:bac
dragontiger-backend:dragonTiger
dragontiger-backend-api:dragonTiger,bac
dragontiger-clients-service:dragonTiger
game-merchant-backstage-api:member
go-game-detail:member
go-router-bb-to-cc:member
heartz-member-backend:member
heartz-member-clients-service:member
member-backend:member
member-backend-api:member
member-clients-service:member,bac
member-merchant-api:member
sedie-backend:sedie
sedie-backend-api:sedie,bac
sedie-clients-service:sedie
socketserver-fixed-env:bac
teenpatti-backend:teenpatti
teenpatti-backend-api:teenpatti,bac
teenpatti-clients-service:teenpatti
tequila-backstage-api:bac
yuxiaxie-backend:yuxiaxie
yuxiaxie-backend-api:yuxiaxie,bac
yuxiaxie-clients-service:yuxiaxie
)


grant(){
  account=$(echo $1|cut -d':' -f1)
  dbs=$(echo $1|cut -d':' -f2|tr ',' ' ')

  for db in ${dbs};
  do
      echo "GRANT $privileges ON \`${databasePrefix}${db}\`.* TO '${account}'@'${host}';"
  done
  printf "\n"
}

rowsLen=${#rows[@]}
for((i=0; i<$rowsLen; i++))
do
    [[ ! ${rows[i]} =~ ^.*:.*$ ]] && echo "[ FORMAT ERROR ] ${rows[i]}" && exit 1

    account=$(echo ${rows[i]}|cut -d':' -f1)
    password=$(openssl rand -base64 8 |md5 |head -c12)

    echo "CREATE USER '${accountPrefix}${account}'@'${host}' IDENTIFIED BY '${password}';"
    echo "ALTER USER '${accountPrefix}${account}'@'${host}' WITH MAX_USER_CONNECTIONS ${maxUserConnections};"
    grant ${accountPrefix}${rows[i]}
done

echo FLUSH PRIVILEGES;

