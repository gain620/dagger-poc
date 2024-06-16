FROM debian:stretch-slim

ENTRYPOINT ["/bin/bash", "-c", "echo 'Hello, World!' && sleep 10 && echo 'Goodbye, World!'"]