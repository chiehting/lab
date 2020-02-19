# 上傳到GCP storage

## GCP document

[Cloud Storage Client Libraries](https://cloud.google.com/storage/docs/reference/libraries?hl=zh-tw)

[auth-cloud-implicit](https://cloud.google.com/docs/authentication/production?hl=zh-tw#auth-cloud-implicit-java)

[storage-upload-object](https://cloud.google.com/storage/docs/uploading-objects#storage-upload-object-java)


## Example

必須先建立[帳戶金鑰](https://cloud.google.com/iam/docs/creating-managing-service-account-keys)

使用python為範例

```bash
# Install client libraries
$ pip3.7 install --upgrade google-cloud-storage

# Export envrionment
$ export GOOGLE_APPLICATION_CREDENTIALS=./developer-leading-project-256707-afb0320d7a66.json

# Run python program
$ python3 upload_gcp_storage.py
```
