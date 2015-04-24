# Developing Ruby with Docker

For Ruby, the key thing here is to vendor your dependencies with:

```sh
docker run --rm -v "$(pwd)":/app -w /app google/ruby sh -c 'bundle install --standalone'
```

That will install all the required gems you need into the bundle directory. Even if you have
native extensions that are causing problems on your Mac or Windows, doing it in Docker will make it 
no problem. 

Now that we have everything we need in our working directory, let's run it:

```sh
docker run --rm -v "$(pwd)":/app -w /app google/ruby sh -c 'ruby hello.rb'
```

Donezo. 
