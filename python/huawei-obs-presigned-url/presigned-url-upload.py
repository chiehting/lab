# Docs:
#   - https://support.huaweicloud.com/sdk-python-devg-obs/obs_22_0400.html
#   - https://support.huaweicloud.com/intl/en-us/sdk-python-devg-obs/obs_22_1301.html
# Install: pip install esdk-obs-python --trusted-host pypi.org

from obs import ObsClient
import os
import traceback
import base64
import requests

# Obtain an AK and SK pair using environment variables or import the AK and SK pair in other ways. Using hard coding may result in leakage.
# Obtain an AK and SK pair on the management console. For details, see https://support.huaweicloud.com/intl/en-us/usermanual-ca/ca_01_0003.html.
ak = os.getenv("AccessKeyID")
sk = os.getenv("SecretAccessKey")
print(ak, sk)

# (Optional) If you use a temporary AK and SK pair and a security token to access OBS, obtain them from environment variables.
# security_token = os.getenv("SecurityToken")
# Set server to the endpoint corresponding to the bucket. CN-Hong Kong is used here as an example. Replace it with the one in use.
server = "https://obs.cn-east-3.myhuaweicloud.com"
bucketName = "obs-bucket-name"
objectKey = "tmp"

# Create an obsClient instance.
# If you use a temporary AK and SK pair and a security token to access OBS, you must specify security_token when creating an instance.
obsClient = ObsClient(access_key_id=ak, secret_access_key=sk, server=server)
try:
    # # Create a signed URL for creating a bucket.
    # res1 = obsClient.createSignedUrl(method='PUT', bucketName=bucketName, expires=3600)
    # print('signedUrl:', res1.signedUrl)
    # print('actualSignedRequestHeaders:', res1.actualSignedRequestHeaders)

    # Create a signed URL for uploading an object.
    res2 = obsClient.createSignedUrl(method='PUT', bucketName=bucketName, objectKey=objectKey, expires=3600,
                                     headers={'Content-Type': 'text/plain'})
    print('signedUrl:', res2.signedUrl)
    print('actualSignedRequestHeaders:', res2.actualSignedRequestHeaders)

    # with open(objectKey, "rb") as f:
    #     data = f.read()

    # headers = res2.actualSignedRequestHeaders
    # response = requests.put(res2.signedUrl, data=data, headers=headers)
    # print("HTTP status:", response.status_code)
    # print("Response:", response.text)

    # # Create a signed URL for setting an object ACL.
    # res3 = obsClient.createSignedUrl(method='PUT', bucketName=bucketName, objectKey=objectKey, specialParam='acl',
    #                                  expires=3600, headers={'x-obs-acl': 'private'})
    # print('signedUrl:', res3.signedUrl)
    # print('actualSignedRequestHeaders:', res3.actualSignedRequestHeaders)

    # # Create a signed URL for downloading an object.
    # res4 = obsClient.createSignedUrl(method='GET', bucketName=bucketName, objectKey=objectKey, expires=3600)
    # print('signedUrl:', res4.signedUrl)
    # print('actualSignedRequestHeaders:', res4.actualSignedRequestHeaders)

    # # Create a signed URL for deleting an object.
    # res5 = obsClient.createSignedUrl(method='DELETE', bucketName=bucketName, objectKey=objectKey, expires=3600)
    # print('signedUrl:', res5.signedUrl)
    # print('actualSignedRequestHeaders:', res5.actualSignedRequestHeaders)

    # # Create a signed URL for deleting a bucket.
    # res6 = obsClient.createSignedUrl(method='DELETE', bucketName=bucketName, expires=3600)
    # print('signedUrl:', res6.signedUrl)
    # print('actualSignedRequestHeaders:', res6.actualSignedRequestHeaders)

    # # Create a signed URL for initiating a multipart task.
    # res7 = obsClient.createSignedUrl(method='POST', bucketName=bucketName, objectKey=objectKey,
    #                                  specialParam='uploads', expires=3600)
    # print('signedUrl:', res7.signedUrl)
    # print('actualSignedRequestHeaders:', res7.actualSignedRequestHeaders)

    # # Create a signed URL for uploading a part.
    # res8 = obsClient.createSignedUrl(method='PUT', bucketName=bucketName, objectKey=objectKey, expires=3600,
    #                                  queryParams={'partNumber': '1', 'uploadId': '00000*****'})
    # print('signedUrl:', res8.signedUrl)
    # print('actualSignedRequestHeaders:', res8.actualSignedRequestHeaders)

    # # Create a signed URL for assembling parts.
    # res9 = obsClient.createSignedUrl(method='POST', bucketName=bucketName, objectKey=objectKey, expires=3600,
    #                                  queryParams={'uploadId': '00000*****'})
    # print('signedUrl:', res9.signedUrl)
    # print('actualSignedRequestHeaders:', res9.actualSignedRequestHeaders)

    # # Create a signed URL for image persistency.
    # # Name of the bucket that stores the source object
    # bucketName = 'originBucketName'; 
    # # Source object name before the processing
    # objectKey = 'test.png';

    # # Name of the object after processing
    # targetObjectName ="save.png"
    # # (Optional) Name of the bucket that stores the new object
    # targetBucketName ="saveBucketName"
    # queryParams={}
    # queryParams["x-image-process"]="image/resize,w_100"
    # queryParams["x-image-save-object"]=base64.b64encode(targetObjectName .encode("utf-8")).decode()
    # # Optional parameter
    # queryParams["x-image-save-bucket"]=base64.b64encode(targetBucketName .encode("utf-8")).decode()

    # res10 = obsClient.createSignedUrl(method='GET', bucketName=bucketName, objectKey=objectKey, queryParams=queryParams, expires=3600)
    # print('signedUrl:', res10.signedUrl)
    # print('actualSignedRequestHeaders:', res10.actualSignedRequestHeaders)
except:
    print(traceback.format_exc())