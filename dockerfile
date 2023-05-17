FROM golang:1.18

# create directory app
RUN mkdir /app

# set or make /app our working directory
WORKDIR /app

# copy all files to /app
COPY . .

# Expose port 8000 to the host on
EXPOSE 8000

RUN go build -o portal-api

CMD [ "./portal-api" ]
