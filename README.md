# Helper

![alt text](https://github.com/JackBekket/hellper/blob/master/img/helper.jpg)

# We are under heavy development!
it may work unexpectedly


   ![alt text](https://github.com/JackBekket/hellper/blob/master/img/local_ai.png)
   


https://github.com/mudler/LocalAI -- original repo of local ai golang node


# How to run?

simply run `docker-compose up --build` (with `-d` for deattach) -- it will create two containers -- local-ai and this bot

if your local-ai node is already running, then you can just build this bot as `docker build -t helper-bot .` and then `docker run -d --name helper-bot --restart=always -p 8085:8085 helper-bot`
add `-p 8080:8080 for web UI



# How to setup bot to work with your models locally with localai?
<details>
<summary>Download models and set-up promt templates</summary>

1. download your models:
```
mkdir models
wget https://huggingface.co/TheBloke/Wizard-Vicuna-13B-Uncensored-GGUF/blob/main/Wizard-Vicuna-13B-Uncensored.Q4_K_M.gguf
wget https://huggingface.co/TheBloke/Wizard-Vicuna-30B-Uncensored-GGUF/blob/main/Wizard-Vicuna-30B-Uncensored.Q4_K_M.gguf
```
I am using wizard-llm-uncensored with 13b and 30b parameters and I download them into local folder
Notes:
1. 30 billion parameters require 22Gb ram minumum, 13b ~= 13Gb RAM min
2. You can download models directly from hugginface
3. You need .gguf format and optimised quntisation choice
4. Wizard Uncensored LLM's are basically the same wizard-vicuna models but they are was trained at edited dataset, in which was removed biased answers
5. I am using stable-diffusion for image generation, for more info see localai stablediffusion
https://localai.io/features/image-generation/
6. If there are necessity of using embedded generations, you should also download bert model
https://gpt4all.io/models/gguf/all-MiniLM-L6-v2-f16.gguf

2. Setup model template
```
# Use a template from the examples local-ai
# https://localai.io/docs/getting-started/customize-model/
# Here is template code for wizard-uncensored 13 billion:
name: wizard-uncensored-13b
f16: false # true to GPU acceleration
cuda: false # true to GPU acceleration
gpu_layers: 10 # this model have max 40 layers, 15-20 is reccomended for half-load at NVIDIA 4060 TiTan (more layers -- more VRAM required), (i guess 0 is no GPU)
parameters:
  model: wizard-uncensored-13b.gguf
#backend: diffusers
template:

  chat: &template |
    Instruct: {{.Input}}
    Output:
  # Modify the prompt template here ^^^ as per your requirements
  completion: *template
```

**Note** you can find templates at original localai repo and edit them to match with your model
**Note** there is currently a bug with GPU's load -- it will try to load all accessible layers to your GPU and fail if not enough memory. In this case you should adjust number of layers to fit your GPU if needed. 
This file should be placed in models directory as yaml

</details>


## How to setup local-ai manually
<details>
   <summary>It may be interesting if you want to change standard docker-compose setup, for example make it CPU instead of GPU</summary>
3. Run localai at localhost:8080, attach models directory, set context-size and CPU threads. it also attach tmp directory for generated images

CPU setup
```
docker run -p 8080:8080 --name local-ai -v $PWD/models:/models -v $PWD/tmp/generated/images:/tmp/generated/images -ti  localai/localai:latest-aio-cpu --models-path /models --context-size 700 --threads 8 
```
NVIDIA GPU setup
```
docker run -p 8080:8080 --gpus all --name local-ai -e DEBUG=true -v $PWD/models:/models -v $PWD/tmp/generated/images:/tmp/generated/images -v $PWD/configuration:/configuration -d  localai/localai:latest-aio-gpu-nvidia-cuda-12 --models-path /models --context-size 1024 --threads 8 
```
you can use `-e DEBUG=true` for debug/verbose mode, `-d` instead of `-ti` for deatached mode, and so on. Also make sure that you have installed CUDA and nvidia-smi for containers, and your docker is installed as `apt-get install docker.io` (not from snap!)

you can also create your own api keys for access and share it to other people. keys should be listed in api_keys.json file under configuration directory

you can also build localai from source.

3.1. Troubleshooting with GPU


4. Now your local ai node is deployed locally and listen to localhost:8080
you can check it work like
```
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
     "model": "wizard-uncensored-13b",
     "messages": [{"role": "user", "content": "How are you?"}],
     "temperature": 0.9
   }'
```

```
curl http://localhost:8080/v1/images/generations -H "Content-Type: application/json" -d '{
  "prompt": "floating hair, portrait, ((loli)), ((one girl)), cute face, hidden hands, asymmetrical bangs, beautiful detailed eyes, eye shadow, hair ornament, ribbons, bowties, buttons, pleated skirt, (((masterpiece))), ((best quality)), colorful|((part of the head)), ((((mutated hands and fingers)))), deformed, blurry, bad anatomy, disfigured, poorly drawn face, mutation, mutated, extra limb, ugly, poorly drawn hands, missing limb, blurry, floating limbs, disconnected limbs, malformed hands, blur, out of focus, long neck, long body, Octane renderer, lowres, bad anatomy, bad hands, text",
  "size": "256x256"
}'
```

Now you need to setup telegram bot to point at localhost.
add to your .env file string
```
AI_ENDPOINT=http://localhost:8080/v1/chat/completions
```

In case if you need to change url/port just change it in .env file

# Build bot
` go build`

</details>


### Commands

Authorize for additional commands: /help -- print this message.  
/restart -- restart session (if you want to switch between local-ai and openai chatGPT).  
/search_doc -- searching documents.  
/rag -- process Retrival-Augmented Generation.   
/instruct -- use system promt template instead of langchain (higher priority, see examples).   
/image -- generate image ....all funcs are experimental so bot can halt and catch fire.  


### DEV
TODO
