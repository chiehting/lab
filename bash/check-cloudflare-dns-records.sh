#!/bin/bash

per_page=400
email=team@gmail.com
key=51255e60bf8ce


id=$(curl --silent -X GET "https://api.cloudflare.com/client/v4/zones?status=active&page=1&per_page=${per_page}&order=status&direction=desc&match=all" \
     -H "X-Auth-Email: $email" \
     -H "X-Auth-Key: $key" \
     -H "Content-Type: application/json" | jq -r '.result[].id')

name=$(curl --silent -X GET "https://api.cloudflare.com/client/v4/zones?status=active&page=1&per_page=${per_page}&order=status&direction=desc&match=all" \
     -H "X-Auth-Email: $email" \
     -H "X-Auth-Key: $key" \
     -H "Content-Type: application/json" | jq -r '.result[].name')

num=0
for i in $id
do
  id[$num]=$i
  num=$( expr $num + 1 )
done

num=0
for i in $name
do
  name[$num]=$i
  num=$( expr $num + 1 )
done

for (( i=0; i<$num; i++ ))
do
  echo "===== ${id[i]} ${name[i]} ====="

  path=""

  if [ "${name[i]}" = "a.com" ]
  then
    path=/healthcheck
  elif [ "${name[i]}" = "b.com" ]
  then
    path=/service/healthcheck
  fi

  echo "===== ${id[i]} ${name[i]} =====" >> result
  curl --silent -X GET "https://api.cloudflare.com/client/v4/zones/${id[i]}/dns_records?type=A,CNAME&per_page=$per_page" \
     -H "X-Auth-Email: $email" \
     -H "X-Auth-Key: $key" \
     -H "Content-Type: application/json" | jq -r '.result[]|.name' | grep -v '^_' \
     | xargs -n1 -I '{}' curl -o /dev/null -s -w "https://{}\t\thttp_code: %{http_code}\n" https://{}$path >> result

  echo >> result
  echo >> result
done

