import requests
import json
import re
import csv

# 設定參數
search = 'b.elb.ap-northeast-1.amazonaws.com'
replace = 'a.elb.ap-northeast-1.amazonaws.com'
email = 'example@gmail.com'
key = '512ce'

# get all doname name and domain ID
domainID = {}
cfAPI = 'https://api.cloudflare.com/client/v4/zones'
payload = {'status':'active','page':1,'per_page': 50,'order':'status','direction':'desc','match':'all'}
headers = {'X-Auth-Email': email,"X-Auth-Key":key,"Content-Type":"application/json"}

r = requests.get(cfAPI, headers=headers, params=payload)
responseJson = r.json()

while responseJson['result'] != []:
    print("loadding zones of page " + str(payload['page']))
    for rows in responseJson['result']:
        domainID[rows['id']] = rows['name']
    payload['page'] += 1
    r = requests.get(cfAPI, headers=headers, params=payload)
    responseJson = r.json()

# get all sub domain
record=[]
payload = {'type':'A,CNAME','page':1,'per_page': 50}
for ID in domainID:
    print('loadding content ' + domainID[ID])
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

def change_record(change=False):
    payload = {}
    for row in record:
        if row['content'] == search:
          responseJson = {"success":"false","messages":"not change"}
          if change:
            cfAPI='https://api.cloudflare.com/client/v4/zones/' + row['zone_id'] + '/dns_records/' + row['id']
            payload['type'] = row['type']
            payload['name'] = row['name']
            payload['content'] = replace
            payload['proxied'] = False
            payload['ttl'] = row['ttl']
            r = requests.put(cfAPI, headers=headers, json=payload)
            responseJson = r.json()
          print( row['name'],':', responseJson['success'], responseJson['messages'])


change_record(False)
def yes_or_no(question):
    while "the answer is invalid":
        reply = str(input(question+' (y/n): ')).lower().strip()
        if reply[:1] == 'y':
            return True
        else:
            return False

print("===== change record =====")
change_record(yes_or_no("Do you want change record?"))

