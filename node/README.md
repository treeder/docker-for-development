## Using Docker to Develop with Node

Install the dependencies to your system:

```sh
docker run --rm -v "$(pwd)":/app -w /app google/nodejs sh -c 'npm install'
```

This vendors the node dependencies to `node_modules` directory in your working directory. 

Now let's run it:

```sh
docker run --rm -v "$(pwd)":/app -w /app google/nodejs sh -c 'node hello.js'
```
