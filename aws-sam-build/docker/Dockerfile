FROM docker:dind

WORKDIR /app 

RUN apk get update && \
    apk add --no-cache bash python3 py3-pip

RUN pip install aws-sam-cli==1.150.1 --break-system-packages
ADD aws-sam-build/release/linux/amd64/plugin /bin/


ENTRYPOINT ["/bin/plugin"]