FROM golang:alpine
RUN apk add build-base
RUN apk add --update npm nodejs-current
RUN mkdir /gorvel
ADD . /gorvel
WORKDIR /gorvel

# install missing reflex, build go and svelte
RUN sh cmd/install-reflex.sh
RUN sh cmd/check-binaries.sh
RUN sh cmd/build-go.sh
RUN sh cmd/build-svelte.sh

# expose these ports 8080 for gorvel and 8081 for sveltekit
EXPOSE 8080 8081

# to serve, use this command
CMD ["sh", "./serve", "docker"]

# to use the build, most likely for production
# CMD [".build/gorvel", "-host", "0.0.0.0", "-port", "8334"]
