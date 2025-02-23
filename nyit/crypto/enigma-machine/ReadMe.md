## Run Docker
To run a Dockerfile, you need to follow a series of steps that involve building a Docker image from the Dockerfile and then running a container based on that image. Here’s a step-by-step guide:

### Step 1: Install Docker

Make sure you have Docker installed on your machine. You can download and install Docker from the [official Docker website](https://www.docker.com/get-started).

### Step 2: Create Your Dockerfile

Ensure you have your `Dockerfile` ready in your project directory. For example, if your `Dockerfile` is named `Dockerfile` and is located in the directory `nyit/crypto/enigma-machine/`, navigate to that directory in your terminal.

### Step 3: Build the Docker Image

Open your terminal and navigate to the directory containing your `Dockerfile`. Use the following command to build the Docker image:

```bash
docker build -t your_image_name .
```

- Replace `your_image_name` with a name you want to give your image (e.g., `enigma-machine`).
- The `.` at the end specifies the current directory as the build context.

### Step 4: Run the Docker Container

After the image is built successfully, you can run a container from that image using the following command:

```bash
docker run --rm your_image_name "Your message here"
```

- Replace `your_image_name` with the name you used in the build step (e.g., `enigma-machine`).
- Replace `"Your message here"` with the actual message you want to encrypt.

### Example Commands

1. **Build the Image:**
   ```bash
   cd nyit/crypto/enigma-machine/
   docker build -t enigma-machine .
   docker build -t enigma . 
   ```

2. **Run the Container:**
   ```bash
   docker run --rm enigma "Hello, World!"
   ```

### Additional Notes

- **List Docker Images:** You can list all Docker images on your machine with:
  ```bash
  docker images
  ```

- **List Running Containers:** To see running containers, use:
  ```bash
  docker ps
  ```

- **Stop a Running Container:** If you need to stop a running container, you can use:
  ```bash
  docker stop container_id
  ```
  Replace `container_id` with the actual ID of the container.

- **Remove an Image:** If you want to remove an image, use:
  ```bash
  docker rmi your_image_name
  ```

By following these steps, you should be able to successfully run your Dockerfile and execute your command-line tool within a Docker container.
