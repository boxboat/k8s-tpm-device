ARG VERSION=master
ARG GO_VERSION=1.20.4

FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION}-alpine as build

RUN apk --no-cache add make ca-certificates
RUN adduser -D tpm
WORKDIR /src
COPY go.mod go.sum /src/
RUN go mod download
COPY . /src/
ARG TARGETOS
ARG TARGETARCH
ARG VERSION
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} VERSION=${VERSION} make docker-build
USER tpm
ENTRYPOINT ["/src/bin/k8s-tpm-device"]

FROM --platform=${TARGETPLATFORM} gcr.io/distroless/static as release

COPY --from=build /etc/passwd /etc/group /etc/
COPY --from=build /src/bin/k8s-tpm-device /bin/k8s-tpm-device
USER tpm
ENTRYPOINT ["/bin/k8s-tpm-device"]