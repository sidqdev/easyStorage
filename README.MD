
# Easy telegram storage

API for save your files on telegram servers


## API Reference

#### Upload file

```
  POST /upload
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `file` | `form file` | **Required**. Your file to upload |

##### Response
    Code: 200, Content: fileID
    Code: 500, Content: error

#### Get file

```
  GET /get/{fileID}
```

##### Response
    Code: 200, Content: file
    Code: 404, Content: error(file not found)
    Code: 500, Content: error

#### Get file info

```
  GET /getInfo/{fileID}
```

##### Response

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `filename` | `string` | Name of file from form |
| `fileID` | `string` | generated file id |
| `hash` | `string` | md5 hesh hex of file |
| `size` | `int` | size of file in bytes |
| `publicURL` | `string` | ngrok link to file |
## Documentation

### Fast run

```bash
git clone https://github.com/sidqdev/easyStorage
cd easyStorage
go run .
```

### Configure
```json
{
    "bots": [
        {
            "token": "",
            "channel_id": 0
        }
    ],
    "storage_directory": "files",
    "split_file_length": 20971520,
    "max_upload_file_length": -1,
    "host": "0.0.0.0:8070",
    "send_frame_delay": 35,
    
    "ngrok_api_key": "",
    "use_ngrok_link": 0
}

```
```
bots - list of bot token and chat\channel id, which bot can send message
storage_directory - directory where saved file data
split_file_length - 20MB limit of telegram file to send
max_upload_file_length - max file length upload to server(-1 - no limit)
host - host and port of server
send_frame_delay - delay between send file to current bot(telegram limit 30 notifications per second)
ngrok_api_key - your ngrok api key
use_ngrok_link - 0 = False, 1 = True- turn off or on ngrok
```
