#!/bin/bash

arr=(ca69a981a 34a7894deda22)
arrn=(a.com b.com)
per_page=200
email=email@gmail.com
key=512509c91dff829c5b8

for i in {0..2};
do
  echo "===== ${arrn[i]} ${arr[i]} ====="

  path=""

  if [ "${arrn[i]}" = "a.com" ]
  then
    path=/healthcheck
  elif [ "${arrn[i]}" = "b.com" ]
  then
    path=/service/healthcheck
  fi

  echo "===== ${arrn[i]} ${arr[i]} =====" >> result
  curl --silent -X GET "https://api.cloudflare.com/client/v4/zones/${arr[i]}/dns_records?type=A,CNAME&per_page=$per_page" \
     -H "X-Auth-Email: $email" \
     -H "X-Auth-Key: $key" \
     -H "Content-Type: application/json" | jq -r '.result[]|.name' | grep -v '^_' \
     | xargs -n1 -I '{}' curl -o /dev/null -s -w "https://{}\t\thttp_code: %{http_code}\n" https://{}$path >> result

  echo >> result
  echo >> result
done

