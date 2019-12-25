FROM kbehouse/grpc-python-alpine

COPY pb/*.py /app/pb/
COPY greeter_server.py /app/greeter_server.py
WORKDIR /app
CMD ["python", "greeter_server.py"]