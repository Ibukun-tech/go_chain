FROM go

WORKDIR /app

COPY . /app

CMD [ "make run" ]