import requests
import json
import re
import csv

# 設定參數
file_name='get-cloudflare-all-records.csv'
per_page=50
email = 'example@gmail.com'
key = '512ce'

# get all doname name and domain ID
domainID = {}
cfAPI = 'https://api.cloudflare.com/client/v4/zones'
payload = {'status':'active','page':1,'per_page': per_page,'order':'status','direction':'desc','match':'all'}
headers = {'X-Auth-Email': email,"X-Auth-Key":key,"Content-Type":"application/json"}

r = requests.get(cfAPI, headers=headers, params=payload)
responseJson = r.json()

while responseJson['result'] != []:
    print("load projects of page " + str(payload['page']))
    for rows in responseJson['result']:
        domainID[rows['id']] = rows['name']
    # 下一頁
    payload['page'] += 1
    r = requests.get(cfAPI, headers=headers, params=payload)
    responseJson = r.json()

# get all sub domain
record=[]
payload = {'type':'A,CNAME','page':1,'per_page': per_page}
for ID in domainID:
    print('get content ' + domainID[ID] + ' ...')
    payload['page'] = 1
    cfAPI='https://api.cloudflare.com/client/v4/zones/' + ID + '/dns_records'
    r = requests.get(cfAPI, headers=headers, params=payload)
    responseJson = r.json()
    record = record + responseJson['result']
    while responseJson['result'] != []:
        print('page' + str(payload['page']) + ' ...')
        payload['page'] += 1
        r = requests.get(cfAPI, headers=headers, params=payload)
        responseJson = r.json()
        record = record + responseJson['result']

data=[]
for row in record:
    data.insert(0, [row['zone_name'],row['name'], row['type'], row['content'], row['proxied']])

data.insert(0, ['zone_name','name', 'type', 'content', 'proxied'])

file = open(file_name, 'w+', newline ='')
with file:
    write = csv.writer(file)
    write.writerows(data)
