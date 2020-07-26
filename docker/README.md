# Docker    



## Greylog2
```bash
docker run -dit \
    --log-driver=gelf \
    --log-opt gelf-address=udp://192.168.0.42:12201 \
    alpine sh
```

## License 

MIT

