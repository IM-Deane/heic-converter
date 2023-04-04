## HEIC Converter

### Overview

Contains the logic for converting images from .heic to .jpeg

Note: This is currently being used by
[SwiftConvert](https://github.com/IM-Deane/swift-convert)

### Local dev:

Run the following command to start the server on 8080

```bash
go run .
```

Build the project using:

```bash
go build -tags netgo -ldflags '-s -w' -o app
```

### Prod build example:

https://www.callicoder.com/docker-golang-image-container-example/
