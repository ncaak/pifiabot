# pifiabot

This app uses Telegram's [Bot API](https://core.telegram.org/bots/api) to deliver random results based on dice notation. The app sets a webhook as an entry point for Telegram messages.

## Deployment

The app requires two environmental variables to work. It runs as a server receiving requests from Telegram services on the webhook endpoint, and sending callback requests using API methods. For authentication it uses a self-signed certificate in PEM format following API [documentation](https://core.telegram.org/bots/self-signed).

### Env variables

**ENDPOINT**

It is the external IP Address of the machine where the app is running. It is used to setup the webhook.

**BOT_TOKEN**

It is the identifier that Telegram's Bot API sets on bot creation. It is used to communicate with the API.

### Certificate

The server runs under HTTPS protocol and uses a keypair with this format: __private.key__ and __cert.pem__.

### Deployment with Dockerfile

#### Builiding the image

```docker build -t <image_name> --build-arg BOT_TOKEN=<bot_token> .```

#### Running the container

```docker run -td -p 443:443 -e ENDPOINT=<endpoint_ip> <image_name>```

