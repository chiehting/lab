import requests
import json
import re
import csv

# 設定參數
gitalbAPI = 'https://gitlab.example.tw/api/v4/projects'
payload = {'simple': 'true', 'archived': 'false', 'per_page': 50, 'page': 0}
headers = {'PRIVATE-TOKEN': 'UzNy3D'}

# 定義函式
def getFrontendUrl(gitlabId, branch):
    url = gitalbAPI + '/' + str(gitlabId) + '/repository/files/.gitlab-ci.yml/raw?ref=' + branch + '&private_token=' + headers['PRIVATE-TOKEN']
    r = requests.get(url)
    search = re.search('ProdS3:.*', r.text)
    if search:
        return "https://" + search.group()[7:].strip().replace('"','')
    else:
        return 'None'

def getBackendtUrl(projectName):
    url = gitalbAPI + '/6/repository/files/configuration%2Fvalues%2F' + projectName.lower() + '.yaml/raw?ref=master&private_token=' + headers['PRIVATE-TOKEN']
    r = requests.get(url)
    search = re.search('host:.*', r.text)
    if search:
        return "https://" + re.sub("^(dev|stage|uat)-",'',search.group()[5:].strip())
    else:
        if r.status_code != 200:
            return 'Undefined helm deployment file'
        else:
            return 'None'

def getServiceUrl(gitlabId, gitlabGroup, data):
    if gitlabGroup == 'frontend':
        url = getFrontendUrl(gitlabId, data[gitlabId]['branch'])
        data[gitlabId]['url'] = url
    elif gitlabGroup == 'backend':
        url = getBackendtUrl(data[gitlabId]['name'])
        data[gitlabId]['url'] = url
    return

# 取所有專案資料
data = {}
r = requests.get(gitalbAPI, headers=headers, params=payload)
responseJson = r.json()

while responseJson != []:
    print("load projects of page " + str(payload['page']))
    for rows in responseJson:
        data[rows['id']] = {
            'name': rows['name'],
            'description': rows['description'],
            'group': rows['namespace']['name'],
            'topics': rows['tag_list'],
            'branch': rows['default_branch']
        }
    # 下一頁
    payload['page'] += 1
    r = requests.get(gitalbAPI, headers=headers, params=payload)
    responseJson = r.json()

# 將資料寫成檔案做,不要重複撈取
# with open('rowdata.json', 'w') as f:
#     json.dump(data, f)
# # 讀檔, 記得註解 while responseJson != []
# f = open('rowdata.json', 'r')
# data = json.load(f)

# 處理欄位資料
i=0
for index in data:
    if  (i % 20) == 0: print("loadding....")
    if data[index]['description'] is None: data[index]['description'] = ''
    i += 1
    # data[index]['description'] = re.sub(r"http\S+", "", data[index]['description']).rstrip()
    data[index]['topics'] = ', '.join(data[index]['topics'])
    getServiceUrl(index, data[index]['group'], data)

frontend = []
backend = []
for index in data:
    if data[index]['group'] == 'frontend':
        frontend.insert(0, [index, data[index]['name'], data[index]['topics'], data[index]['description'], data[index]['url']])
    elif data[index]['group'] == 'backend':
        backend.insert(0, [index, data[index]['name'], data[index]['topics'], data[index]['description'], data[index]['url']])

frontend.insert(0, ['id', 'name', 'topics', 'description', 'url'])
backend.insert(0, ['id', 'name', 'topics','description', 'url'])
backend.insert(0, [''])

file = open('get-gitlab-project.csv', 'w+', newline ='')
with file:
    write = csv.writer(file)
    write.writerows(frontend)
    write.writerows(backend)
