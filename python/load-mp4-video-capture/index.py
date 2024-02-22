import cv2 # pip3 install opencv-python
import os

# 打開文件
video_capture = cv2.VideoCapture('sample-mp4-file-small.mp4')

try:
    # creating a folder named data
    if not os.path.exists('data'):
        os.makedirs('data')
# if not created then raise error
except OSError:
    print ('Error: Creating directory of data')

i = 0
while True:
    # 讀取一幀
    ret, frame = video_capture.read()
    if not ret:
        break

    # 獲取時間戳
    timestamp = video_capture.get(cv2.CAP_PROP_POS_MSEC)
    print("Timestamp (milliseconds):", timestamp)

    # 可以在這裡處理幀
    name = './data/frame' + str(i) + '.jpg'
    print ('Creating...' + name)
    # writing the extracted images
    cv2.imwrite(name, frame)
    i += 1
    print(i)

# 釋放視訊捕捉對象
video_capture.release()

