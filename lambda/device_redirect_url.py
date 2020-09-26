def lambda_handler(event, context):
    request = event['Records'][0]['cf']['request']
    headers = request['headers']
    # 這邊直接濾掉s3的域名
    origin_url = 'https://' + headers['host'][0]['value'].replace('.s3.amazonaws.com', '')
    origin_uri = request['uri']
    origin_querystring = '?' + request['querystring'] if request['querystring'] != '' else ''
    device = ''

    if 'cloudfront-is-desktop-viewer' in headers and headers['cloudfront-is-desktop-viewer'][0]['value'] == 'true':
        device = 'desktop'
    elif 'cloudfront-is-mobile-viewer' in headers and headers['cloudfront-is-mobile-viewer'][0]['value'] == 'true':
        device = 'mobile'
    elif 'cloudfront-is-tablet-viewer' in headers and headers['cloudfront-is-tablet-viewer'][0]['value'] == 'true':
        device = 'mobile'

    url_is_valid = match_url(device,origin_url)

    if url_is_valid:
        response =  request
    else:
        redirect_url = replace_url(device, origin_url)
        response = {
            'status': '302',
            'statusDescription': 'Found',
            'headers': {
                'location': [{
                    'key': 'Location',
                    'value': redirect_url + origin_uri + origin_querystring
                }]
            }
        }
    return response

def match_url(device, origin_url):
    result = {
    'mobile':lambda origin_url: ("-m" in origin_url or "//m" in origin_url),
    'desktop':lambda origin_url: ("-m" not in origin_url and "//m" not in origin_url),
    '':lambda origin_url: True
    }[device](origin_url)
    return result

def replace_url(device, origin_url):
    new_url = origin_url
    if device == 'mobile':
        if ("dev-" in origin_url):
            new_url = origin_url.replace("dev-", "dev-m");
        elif ("stage-" in origin_url):
            new_url = origin_url.replace("stage-", "stage-m");
        else:
            new_url = origin_url.replace("://", "://m");
    elif device == 'desktop':
        if ("dev-" in origin_url):
            new_url = origin_url.replace("dev-m", "dev-");
        elif ("stage-" in origin_url):
            new_url = origin_url.replace("stage-m", "stage-");
        else:
            new_url = origin_url.replace("://m", "://");
    return new_url

event = {
  "Records": [
    {
      "cf": {
        "config": {
          "distributionDomainName": "d111111abcdef8.cloudfront.net",
          "distributionId": "EDFDVBD6EXAMPLE",
          "eventType": "viewer-request",
          "requestId": "4TyzHTaYWb1GX1qTfsHhEqV6HUDd_BzoBZnwfnvQc_1oF26ClkoUSEQ=="
        },
        "request": {
          "clientIp": "203.0.113.178",
          "headers": {
            "host": [
              {
                "key": "Host",
                "value": "dev-test.com.s3.amazonaws.com"
              }
            ],
            "cloudfront-is-mobile-viewer": [
              {
                "key": "Host",
                "value": "true"
              }
            ],
            "user-agent": [
              {
                "key": "User-Agent",
                "value": "curl/7.66.0"
              }
            ],
            "accept": [
              {
                "key": "accept",
                "value": "*/*"
              }
            ]
          },
          "method": "GET",
          "querystring": "token=test",
          "uri": "/index.html"
        }
      }
    }
  ]
}

print(lambda_handler(event,'test'))
