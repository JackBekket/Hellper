# Use the specified image
FROM localai/localai:latest-aio-gpu-nvidia-cuda-12

# Set the working directory
WORKDIR /app


RUN mkdir /models
# Copy the local directories into the container
COPY ./models /models
RUN mkdir tmp && cd tmp
RUN mkdir generated && cd generated
RUN mkdir images && cd images
RUN cd .. && cd .. && cd ..
#COPY ./tmp/generated/images /tmp/generated/images
COPY ./configuration /configuration

# Set the environment variable
ENV DEBUG=true

# Expose the specified port
EXPOSE 8090

# Run the command
CMD ["--models-path", "/models", "--context-size", "2048", "--threads", "10", "--localai-config-dir", "/configuration", "worker", "p2p-llama-cpp-rpc", "--token", "${P2PTOKEN}"]
