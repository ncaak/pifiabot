# pifiabot

This app uses Telegram's [Bot API](https://core.telegram.org/bots/api) to deliver random results based on dice notation. The app sets a webhook as an entry point for Telegram messages.

## Deployment

The app requires two environmental variables to work. It runs as a server receiving requests from Telegram services on the webhook endpoint, and sending callback requests using API methods. For authentication it uses a self-signed certificate in PEM format following API [documentation](https://core.telegram.org/bots/self-signed).

### Env variables

**ENDPOINT**

It is the external IP Address of the machine where the app is running. It is used to setup the webhook.

**PORT**

It's the port of the service exposed to the webhook callbacks. It is used to set the webhook.

**BOT_TOKEN**

It is the identifier that Telegram's Bot API sets on bot creation. It is used to communicate with the API.

### Certificate

The server runs under HTTPS protocol and uses a keypair with this format: __private.key__ and __cert.pem__.

### Messages

The program also expects a JSON formatted file with the messages to display, using error ids as keys. A dummy JSON file could be provided and the error ids would be displayed instead.

Current format:
```json
{
    "cth_skill_limits": "",
    "cth_wrong_argument": "",
    "dice_number": "",
    "faces_number": "",
    "no_drop" : "",
    "notation_max_length": "",
    "repeat_iter_error": "",
    "repeat_nodice_error": "",
    "unknown_action": "",
    "unknown_error": "",
}
```

### Deployment with Dockerfile

#### Builiding the image

```docker build -t <image_name> --build-arg BOT_TOKEN=<bot_token> .```

#### Running the container

```docker run -td -p <endpoint_port>:<endpoint_port> -e ENDPOINT=<endpoint_ip> -e PORT=<endpoint_port> <image_name>```

#### Required files

Both keys and the messages file should be available at build time in the following path:

* `deploy/cert.pem`
* `deploy/private.key`
* `deploy/messages.json`
