# Crabe Life
Hello world like service, for educational purposes

## Description

http://crabe.life is a very simple statically compiled, 4MB service in Go.

It's a great candidate to get first experience in application hosting, especially in Cloud native context.

Crabe.life can run as Docker container or as simple linux process with systemd.

### Systemd setup

Download crabe.life binary (releases page)

    wget http://[TODO]
    chmod +x crabe-life_1.0
    mv crabe-life_1.0 /usr/local/bin/crabe-life

Create systemd service

    cat << EOF | sudo tee /lib/systemd/system/crabe-life.service
    [Unit]
    Description=Crabe Life

    [Service]
    ExecStart=/usr/local/bin/crabe-life

    [Install]
    WantedBy=multi-user.target

    EOF

Enable systemd service

    sudo ln -s /lib/systemd/system/crabe-life.service  /etc/systemd/system/crabe-life.service
    sudo systemctl daemon-reload
    sudo systemctl enable crabe-life
    sudo systemctl start crabe-life
    sudo systemctl status crabe-life

### Docker setup

    docker build -t crabe-life:latest .
    docker run -p 8080:8080 crabe-life:latest

## Things you probably want to know

* There is a `/healthz` route that returns 200 OK. You probably want to use it for probing.
* The crab gif and cache headers are configurable (see options `crabe-life -h`)

```
Usage of crabe-life:
-cache-ttl int
        time in second to cache HTTP response
-crab-url string
        url to crab gif (default "https://i.gifer.com/3QZn.gif")
```

## Licence

Public domain. You can copy, modify, distribute and perform the work, even for commercial purposes, all without asking permission