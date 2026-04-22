sudo apt install python3-opencv
pip install opencv-python
yolo download model=yolov8n.pt


import cv2
import requests
import time
from ultralytics import YOLO

# YOLOv8n モデル読み込み
model = YOLO("yolov8n.pt")

GO_SERVER = "http://192.168.0.7:8080/events"  # ← あなたのラズパイIPに合わせる

def send_event(cls, conf, x, y, w, h):
    data = {
        "class": cls,
        "confidence": float(conf),
        "x": int(x),
        "y": int(y),
        "w": int(w),
        "h": int(h),
        "timestamp": int(time.time())
    }
    try:
        requests.post(GO_SERVER, json=data, timeout=0.5)
    except:
        pass

def main():
    cap = cv2.VideoCapture(0)

    while True:
        ret, frame = cap.read()
        if not ret:
            continue

        results = model(frame, imgsz=320)[0]

        for box in results.boxes:
            cls = model.names[int(box.cls)]
            conf = box.conf
            x1, y1, x2, y2 = box.xyxy[0]
            w = x2 - x1
            h = y2 - y1

            send_event(cls, conf, x1, y1, w, h)

        # 軽量化のために少し待つ
        time.sleep(0.05)

if __name__ == "__main__":
    main()
