# golang_shell_encryption
golang_shell_encryption

## usage
if you want to use your own public key, export ENTRYPOINT_PUBLIC_KEY=key,before execute below commands.

```shell
./encryption abusolute_path_of_file
```

```shell
./decryption abusolute_path_of_file other_args_pass_to_entrypoint
```
### example
original Dockerfile:
```dockerfile
FROM daocloud.io/daocloud/rabbitmq:portmapping
ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
CMD ["rabbitmq-server"]
```

1. encrypt /usr/local/bin/docker-entrypoint.sh.
```shell
mykey=$(date +%s | sha256sum | base64 | head -c 32 ; echo)
echo $myket
export ENTRYPOINT_PUBLIC_KEY=$mykey
./encryption /home/rabbitmq/docker-entrypoint.sh
```
2. modify Dockerfile and rebuild
```dockerfile
FROM daocloud.io/daocloud/rabbitmq:portmapping
ENV ENTRYPOINT_PUBLIC_KEY=mykey_value
COPY decryption /usr/bin/decryption
RUN chmod +x /usr/bin/decryption
COPY docker-entrypoint.sh.enc /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["/usr/bin/decryption", "/usr/local/bin/docker-entrypoint.sh", "rabbitmq-server"]
```

## develope
before developing,please cover github.com/Tiked/FileEncryption with FileEncryption.go 

## Todo:
1. Package decryption and sercet file together, like shc
