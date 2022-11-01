# Encodr
Simple QR Code encoding service


## API Methods
| API Methods                        | Description                                                                                                         |
|------------------------------------|---------------------------------------------------------------------------------------------------------------------|
| GET /health                        | Simple API health check with timestamp/status.                                                                      |
| GET /qrcode/:url?size=<pixelCount> | Generate QR Code (png file) for specified URL. Optional query parameter of size for the QR code. (default is 256px) |


## Dockerfile
Build service from project root
```
docker build . -t encodr
```
Run service
```
docker run -p 8080:80 encodr
```


## Build for Multi-Platform using Docker buildx (used for Raspberry PI)

```
docker buildx build \
--push \
--platform linux/arm/v7,linux/arm64/v8,linux/amd64  --tag alexchomiak/encodr:buildx-latest .
```