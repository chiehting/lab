from sshtunnel import SSHTunnelForwarder
import redis

ssh_host = '1.2.1.2'
ssh_port = 6379
ssh_username = 'jenkins'
ssh_key = './.ssh/tmp/id_rsa'

redis_host = '10.2.202.70'
redis_port = 6379
redis_password = ''

# 本地端口
local_port = 6379

# 建立 SSH 隧道
with SSHTunnelForwarder(
    (ssh_host, ssh_port),
    ssh_username=ssh_username,
    ssh_private_key=ssh_key,
    remote_bind_address=(redis_host, redis_port),
    local_bind_address=('127.0.0.1', local_port)
) as tunnel:
    print(f"SSH tunnel created: localhost:{local_port} -> {redis_host}:{redis_port}")

    # 連線 Redis
    redis_conn = redis.StrictRedis(
        host='127.0.0.1',
        port=local_port,
        password=redis_password,
        decode_responses=True
    )

    # 示例1：set、get、delete
    redis_conn.set('example_key', 'example_value')
    result = redis_conn.get('example_key')
    redis_conn.delete('example_key')
    print(f"Result from Redis: {result}")

    result = redis_conn.hgetall('user_token_1')
    print(f"Result from Redis key user_token_1: {result}")

print("SSH tunnel closed.")
