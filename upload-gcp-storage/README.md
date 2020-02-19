# 上傳到GCP storage

## GCP document

[Cloud Storage Client Libraries](https://cloud.google.com/storage/docs/reference/libraries?hl=zh-tw)

[auth-cloud-implicit](https://cloud.google.com/docs/authentication/production?hl=zh-tw#auth-cloud-implicit-java)

[storage-upload-object](https://cloud.google.com/storage/docs/uploading-objects#storage-upload-object-java)


## Example

其中認證 key在 [NAS](http://192.168.160.250:5000/d/f/537679534667903264) 上

使用python為範例

```bash
# Install client libraries
$ pip3.7 install --upgrade google-cloud-storage

# Export envrionment
$ export GOOGLE_APPLICATION_CREDENTIALS=./developer-leading-project-256707-afb0320d7a66.json

# Run python program
$ python3 upload_gcp_storage.py
```
