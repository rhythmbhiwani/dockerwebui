# dockerwebui

### Docker Compose File:
```
version: '3.5'
services:
  mywebhost:
    image: rhythmbhiwani/dockerwebui
    container_name: mywebhost
    ports:
      - 80:80
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    
```
