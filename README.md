# Imager

**Imager** is a powerful CLI tool written in Go for efficiently scaling down and cropping images. Designed for speed and simplicity, Imager is perfect for developers and power users who need to quickly resize or crop images directly from the command line.

### Features:
- **Fast Image Processing**: Leverages the efficiency of Go to quickly scale down and crop images.
- **Simple CLI**: Easy-to-use command-line interface with straightforward commands and options.
- **Customizable Outputs**: Specify dimensions and crop areas to meet your exact needs.

## Getting Started:
1. **Installation**: 
   ```bash
   go install github.com/yourusername/imager@latest
   ```

### Usage
```bash
imager --input path/to/image --width 800 --height 600 --crop 100,100,500,500 --output path/to/output
```

## Contributions
Feel free to open issues or submit pull requests to help improve Imager!
