# Setting Up Redis on EC2

## Launch an EC2 Instance
1. Go to the AWS Management Console and navigate to EC2.
2. Click on "Launch Instance" and select a suitable Amazon Machine Image (AMI), such as Amazon Linux 2.
3. Choose an instance type (e.g., t2.micro for testing).
4. Configure the instance, ensuring it has:
   - A public IP address
   - Security group that allows inbound traffic on port 6379 (the default Redis port) from your local IP address

## Install Redis
1. SSH into your EC2 instance using your key pair.
2. Update the package manager and install Redis with the following commands:
   ```bash
   sudo yum update -y
   sudo yum install -y gcc make
   cd /usr/local/src
   sudo wget http://download.redis.io/redis-stable.tar.gz
   sudo tar xvzf redis-stable.tar.gz
   cd redis-stable
   sudo make
   sudo cp src/redis-server /usr/local/bin/
   sudo cp src/redis-cli /usr/local/bin/
   ```

## Start the Redis Server
```bash
redis-server
```

## Verify Redis Installation
1. Open another terminal session
2. SSH back into your EC2 instance
3. Run:
   ```bash
   redis-cli ping
   ```
   You should receive a response of `PONG`, indicating that Redis is running.