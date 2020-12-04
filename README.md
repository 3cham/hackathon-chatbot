Implement a chat bot which clients could register & broadcast message to others

The used communication protocol is HTTP

Server & Client run inside a docker container

Server:
- The central point, which clients could register & send message to

- /api/register -> POST data = { "clientname": "" }
    The information is saved in a file

- /api/send_message
    -> POST data { "clientname": "", "message": "" }
    
- /api/get_messages?from={STARTTIME}
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

JSON Parsing:
 - key name is case insensitive
 - JSON structure could be different, still be parsed if the keys are correct
 
Client:
- When the client is started, the IP address of the server must be provided beforehand
  Call /api/register to server address to register itself with server
- In an underground go routine, the client calls /api/receive_message every 1s to receive new messages

- otherwise, client wait for input form Stdin, if Enter is pushed, 
it sends the last sentence to server with /api/send_message

# Build & Deploy server

```bash§
# Build docker container
docker build . -t hackathon-chatserver:latest

# start the server 
docker run -d -p 8080:5000 hackathon-chatserver:latest
``` 


# Start client against deployed server

```bash
#build the client
go build -o cli

./cli -type client -name YourName -address http://asprd02.ov.otto.de:8080
```