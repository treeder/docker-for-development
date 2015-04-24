## Quick Example for a Java Worker (3 minutes)

First, let's build it on the right architecture using the actual Docker image it will be running on. The
dependencies are passed into javac. 

```sh
docker run --rm -v "$(pwd)":/app -w /app java sh -c 'javac -cp "json-java.jar:gson-2.2.4.jar" Hello.java'
```

Now run it to test it out:

```sh
docker run --rm -v "$(pwd)":/app -w /app java sh -c 'java -cp gson-2.2.4.jar:json-java.jar:. Hello'
```
