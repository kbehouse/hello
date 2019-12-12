FROM kbehouse/grpc-python-alpine

RUN apk add bash bash-doc bash-completion 

COPY pb/*.py /app/pb/
COPY greeter_server.py /app/greeter_server.py
WORKDIR /app
CMD ["python", "greeter_server.py"]