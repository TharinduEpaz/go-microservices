# Go Microservices with Redis

This is a simple example of how to implement microservices with Redis in Go.

## Features

- Redis storage for the data
- Redis hosted on AWS EC2
- Graceful shutdown using goroutines
- CRUD operations

---
## How to setup Redis on AWS EC2

### Update Your System:

Before installing Redis, ensure that your system is up to date. Run:

```bash
sudo dnf update -y
```

### Install Redis:

The package available for Redis in Amazon Linux 2023 is redis6. Install it using the following command:
```bash
sudo dnf install -y redis6
```

### Configure Redis:

After installation, you may need to edit the configuration file to suit your requirements. The configuration file can be found at /etc/redis6/redis6.conf. Open it using a text editor:

```bash
sudo nano /etc/redis6/redis6.conf
```

You can set parameters such as enabling password authentication and adjusting the binding address.

### Start Redis Service:

To start the Redis server, use:

```bash
sudo systemctl start redis6
```

To enable Redis to start on boot, run:

```bash
sudo systemctl enable redis6
```

### Verify Installation:

Check if Redis is running by using the following command:

```bash
redis6-cli ping
```

If everything is set up correctly, you should receive a response of PONG.
