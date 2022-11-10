# maximum-portage
When you need the most ports

## Running

```
docker build -t maximum-portage .
docker run -it --cap-add SYS_ADMIN --cap-add NET_ADMIN --volume /proc:/host/proc:ro maximum-portage
```

## References

- https://github.com/aws/amazon-ecs-agent/blob/5fb8e801e9d11630c3470abb778291913ebb9d7f/ecs-init/docker/docker_config.go#L35
- https://man7.org/linux/man-pages/man5/proc.5.html
- https://www.kernel.org/doc/Documentation/networking/proc_net_tcp.txt