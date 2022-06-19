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

## Demo Video

[![Demo Video](https://img.youtube.com/vi/y3rHGgTeFCY/0.jpg)](https://www.youtube.com/watch?v=y3rHGgTeFCY)
