# Use a base image with Python installed
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install any required packages (if you have a requirements.txt file)
# Uncomment the next line if you have dependencies to install
RUN pip install --no-cache-dir -r requirements.txt

# Make the script executable (optional)
RUN chmod +x enigma_machine.py

# Command to run the script
ENTRYPOINT ["python", "enigma_machine.py"]