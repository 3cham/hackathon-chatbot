Implement a chat bot which clients could register & broadcast message to others

The used communication protocol is HTTP

Server & Client run inside a docker container

Server:
- The central point, which clients could register & send message to

- /api/register -> POST data = { "client-name": "" }
    The information is saved in a file

- /api/send_message
    -> POST data { "client-name": "", "message": "" }
    
- /api/receive_message 
   GET data { "from-timestamp": "" }
   -> [
    { "client-name": "",
      "message": "",
      "timestamp": ""
      },
    { "client-name": "",
      "message": "",
      "timestamp": ""
      }   
      ]

Client:
- When the client is started, the IP address of the server must be provided beforehand
- Call /api/register to server address to register itself with server
- Call /api/send_message to send messages
- Call /api/receive_message every 5s to receive new messages