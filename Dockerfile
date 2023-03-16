# build VUE3 app
FROM node:18-alpine as build
WORKDIR /app
COPY app/package*.json ./
RUN npm install
COPY app /app
RUN npm run build

# build golang app
FROM golang:1.20-alpine AS deploy
WORKDIR $GOPATH/kuberity
COPY --from=build /app/dist $GOPATH/kuberity/app/dist
COPY . .
# Download all the dependencies
RUN go get -d -v ./...
# Install the package
RUN go install -v ./...
CMD ["kuberity"]