import threading, os


def send():
    os.system('curl -F "file=@1.mp4" 138.68.102.39:8070/upload')

for i in range(20):
    threading.Thread(target=send).start()