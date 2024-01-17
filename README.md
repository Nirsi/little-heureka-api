# Little Heureka API
A little Heureka API is a simple API which aims to simulate a simple Heureka like API for categories, products and offers.

Written in pure Go using [Chi](https://github.com/go-chi/chi)

## Usage
### Docker
```bash
# Build the image
docker build -t little-heureka-api .
```

```bash
# Run the image
docker run -p 8080:8080 little-heureka-api
```

