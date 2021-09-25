A drone config extension to force drone use configuration from your primary branch.  
This can protect your drone from running unwanted ci pipeline edit by unauthorized people from pull request.  
You may find more infomation [here](https://discourse.drone.io/t/drone-protected-build-workaround/3585)  
_Please note this project requires Drone server version 1.4 or higher._

## Note
For I'm living in China where it's hard to connect to github directly, this project use [fastgithub](https://github.com/dotnetcore/FastGithub)  to speed up connection in China.  


## Installation

Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the extension:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=GITHUB_TOKEN=xxxxxxxxxxxxxxxxxxxxxxx \
  --restart=always \
  --name=config registry.cn-hangzhou.aliyuncs.com/pivotstudio/drone-master-config
```

Update your Drone server configuration to include the extension address and the shared secret.

```text
DRONE_YAML_ENDPOINT=http://1.2.3.4:3000
DRONE_YAML_SECRET=bea26a2221fd8090ea38720fc445eca6
```
For k8s usage:  
Change the template value in [deploy files](./deploy) then apply them.  
## Testing

Use the command line tools to test your extension. _This extension uses http-signatures to authorize client access and will reject unverified requests. You will be unable to test this extension using a simple curl command._

```text
export DRONE_YAML_ENDPOINT=http://1.2.3.4:3000
export DRONE_YAML_SECRET=bea26a2221fd8090ea38720fc445eca6

drone plugins config get <repo>
```
