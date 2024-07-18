# Use the specified image
FROM localai/localai:latest-aio-gpu-nvidia-cuda-12

# Set the working directory
WORKDIR /app

# Copy the local directories into the container
COPY ./models /models
COPY ./tmp/generated/images /tmp/generated/images
COPY ./configuration /configuration

# Set the environment variable
ENV DEBUG=true

# Expose the specified port
EXPOSE 8090

# Run the command
CMD ["--models-path", "/models", "--context-size", "2048", "--threads", "10", "--localai-config-dir", "/configuration", "worker", "p2p-llama-cpp-rpc", "--token", "${P2PTOKEN}"]
