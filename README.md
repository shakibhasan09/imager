Here’s a draft for your `README.md`:

````markdown
# Imager

Imager is a Cloudflare Images alternative written in Go. It allows you to set up an API server to upload and optimize images based on variants you define.

## Features

- Upload images via an API.
- Define custom variants for resizing and optimizing images.
- Serve optimized images on demand.

## Table of Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

Follow these instructions to set up and run Imager on your local machine or server.

### Prerequisites

- [Go](https://go.dev/) (1.20 or later)
- [Docker](https://www.docker.com/) (optional, for containerized deployment)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/shakibhasan09/imager.git
   cd imager
   ```
````

2. Build the project:

   ```bash
   go build -o imager
   ```

3. Run the server:
   ```bash
   ./imager
   ```

The server will start on `http://localhost:8080` by default.

### Docker Deployment

To run Imager in a Docker container:

1. Build the Docker image:

   ```bash
   docker build -t imager .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 imager
   ```

## Configuration

Imager supports configuration through environment variables:

- `PORT`: The port the server will run on (default: `8080`).
- `UPLOAD_DIR`: Directory where uploaded images will be stored.
- `VARIANTS`: JSON string defining the image variants (e.g., `{"thumbnail": {"width": 100, "height": 100}}`).

## Usage

1. **Upload an Image**  
   Use the API endpoint `/upload` to upload images:

   ```bash
   curl -X POST -F "file=@your-image.jpg" http://localhost:8080/upload
   ```

2. **Access Optimized Images**  
   Access your image through the variant URL:

   ```
   http://localhost:8080/images/{variant}/{image_name}
   ```

   Replace `{variant}` with your defined variant name and `{image_name}` with the uploaded image's name.

## Contributing

We welcome contributions! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push them to your fork.
4. Open a pull request.

Please ensure your code follows the Go coding standards and includes tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to open an issue or discussion if you have any questions or suggestions.

```

Let me know if you'd like any changes!
```
